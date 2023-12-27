package spreadsheet

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func NewSpreadsheetService(ctx context.Context) (*sheets.Service, error) {
	appName := helper.GetAppName(ctx)
	// get user home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// read credentials from file at $HOME/.credentials/${appName}_credential.json
	filePath := fmt.Sprintf("%s/%s_credential.json", home, appName)
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read client secret file")
	}

	// If modifying these scopes, delete your previously saved token.json.
	// config with readwrite access
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse client secret file to config")
	}
	client, err := getClient(ctx, config)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to get client")
	}

	return sheets.NewService(ctx, option.WithHTTPClient(client))
}

func getClient(ctx context.Context, config *oauth2.Config) (*http.Client, error) {
	appName := helper.GetAppName(ctx)
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	tokFile := fmt.Sprintf("%s/%s_token.json", home, appName)
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(ctx, tok), nil
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
