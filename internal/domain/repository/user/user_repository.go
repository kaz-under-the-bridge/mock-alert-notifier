package user

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/pkg/errors"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/spreadsheet"
)

const (
	spreadsheetId = "xxxxx"
	readRange     = "UserMaster!A3:G990" // ここを変更するときはvalidateRowDataのoffset値も変更すること
)

// ThreadUnsafeな挙動になるのでgoroutineなどで並列処理を行う場合は注意
var sharedUsers *model.Users

type RepositoryInterface interface {
	GetUsers() (*model.Users, error)
}

type Repository struct {
	ctx context.Context
	ds  spreadsheet.SpreadsheetDatastoreInterface
}

func NewUserRepository(ctx context.Context, ds spreadsheet.SpreadsheetDatastoreInterface) *Repository {
	return &Repository{
		ctx: ctx,
		ds:  ds,
	}
}

// データ型のvalidationのみ行う（それ以上のvalidationはmodel側で行う）
func validateRowDataType(rowIndex int, data []interface{}) error {
	rowOffset := 3 // readRangeの開始行が3行目なのでoffsetは3
	colOffset := 0 // readRangeの開始列がA列なのでoffsetは0

	col := helper.GetColSlice()

	invalidDataError := errors.New("Invalid Data")

	if _, ok := data[0].(int); !ok {
		cell := fmt.Sprintf("%s%d", col[colOffset+0], rowIndex+rowOffset)
		helper.Logger.Error("field should be Integer", slog.String("Cell", cell), slog.String("Value", fmt.Sprintf("%v", data)))
		return errors.Wrapf(invalidDataError, "Cell: %s", cell)
	}

	for c, v := range data[1:5] {
		if _, ok := v.(string); !ok {
			cell := fmt.Sprintf("%s%d", col[colOffset+c+1], rowIndex+rowOffset)
			helper.Logger.Error("filed should be String", slog.String("Cell", cell), slog.String("Value", fmt.Sprintf("%v", data)))
			return errors.Wrapf(invalidDataError, "Cell: %s", cell)
		}
	}

	if _, ok := data[5].(int); !ok {
		cell := fmt.Sprintf("%s%d", col[colOffset+5], rowIndex+rowOffset)
		helper.Logger.Error("field should be Integer", slog.String("Cell", cell), slog.String("Value", fmt.Sprintf("%v", data)))
		return errors.Wrapf(invalidDataError, "Cell: %s", cell)
	}

	return nil
}

func (r Repository) getUsersFromSpreadsheet() (*model.Users, error) {
	resp, err := r.ds.Values(r.ctx, spreadsheetId, readRange)
	if err != nil {
		helper.Logger.Error("spreadsheet service values got error", slog.String("SheetID", spreadsheetId), slog.String("ReadRange", readRange))
		return nil, err
	}
	var users model.Users
	for r, row := range resp.Values {
		if err := validateRowDataType(r, row); err != nil {
			return nil, err
		}

		user := model.NewUser(
			row[0].(int),
			row[1].(string),
			row[2].(string),
			row[3].(string),
			row[4].(string),
			row[5].(int),
		)
		users.Push(user)
	}
	sharedUsers = &users
	return &users, nil
}

func (r *Repository) GetUsers() (*model.Users, error) {
	if sharedUsers != nil {
		return sharedUsers, nil
	}

	return r.getUsersFromSpreadsheet()
}
