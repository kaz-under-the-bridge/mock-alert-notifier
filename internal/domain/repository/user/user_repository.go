package user

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/spreadsheet"
)

type Repository struct {
	ds spreadsheet.SpreadsheetDatastoreInterface
}

func NewUserRepository(ds spreadsheet.SpreadsheetDatastoreInterface) *Repository {
	return &Repository{
		ds: ds,
	}
}

func (r *Repository) GetUsers(ctx context.Context, spreadsheetId, readRange string) ([]*model.User, error) {
	resp, err := r.ds.Values(ctx, spreadsheetId, readRange)
	if err != nil {
		return nil, err
	}
	var users []*model.User
	for r, row := range resp.Values {
		// check id column(row[0]) is not int type
		if _, ok := row[0].(int); !ok {
			return nil, errors.Errorf("ID column(Row: %d) is not Integer", r)
		}
		// check if row[1:] is not string type
		for c, v := range row[1:] {
			if _, ok := v.(string); !ok {
				return nil, errors.Errorf("Column(%d), Row(%d) is not String", r, c+1)
			}
		}
		user := model.NewUser(
			row[0].(int),
			row[1].(string),
			row[2].(string),
			row[3].(string),
			row[4].(string),
			row[5].(string),
		)
		users = append(users, user)
	}
	return users, nil
}
