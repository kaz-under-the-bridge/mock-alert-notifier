package spreadsheet

import (
	"context"

	"google.golang.org/api/sheets/v4"
)

type ValueResponse struct {
	Range          string
	MajorDimension string
	Values         [][]interface{}
}

type SpreadsheetDatastoreInterface interface {
	Values(ctx context.Context, spreadsheetId string, readRange string) (*ValueResponse, error)
}

type SpreadsheetDatastore struct {
	srv sheets.Service
}

func NewSpreadsheetDatastore(ctx context.Context) (SpreadsheetDatastoreInterface, error) {
	srv, err := NewSpreadsheetService(ctx)
	if err != nil {
		return nil, err
	}
	return &SpreadsheetDatastore{
		srv: *srv,
	}, nil
}

func (d *SpreadsheetDatastore) Values(ctx context.Context, spreadsheetId string, readRange string) (*ValueResponse, error) {
	resp, err := d.srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return nil, err
	}
	return &ValueResponse{
		Range:          resp.Range,
		MajorDimension: resp.MajorDimension,
		Values:         resp.Values,
	}, nil
}
