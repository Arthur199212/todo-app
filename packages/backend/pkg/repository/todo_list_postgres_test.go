package repository

import (
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestTodoListPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	r := NewTodoListPostgres(sqlxDB)

	type args struct {
		userId int
		listId int
	}
	type mockBehaviour func(args args)

	testTable := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				userId: 1,
				listId: 5,
			},
			wantErr: false,
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(fmt.Sprintf("delete from %s", todoListTable)).
					WithArgs(args.userId, args.listId).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec(fmt.Sprintf("delete from %s", todoItemsTable)).
					WithArgs(args.listId).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "List delete fail",
			args: args{
				userId: 1,
				listId: 5,
			},
			wantErr: true,
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(fmt.Sprintf("delete from %s", todoListTable)).
					WithArgs(args.userId, args.listId).
					WillReturnError(errors.New("test error"))

				mock.ExpectRollback()
			},
		},
		{
			name: "Items delete fail",
			args: args{
				userId: 1,
				listId: 5,
			},
			wantErr: true,
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(fmt.Sprintf("delete from %s", todoListTable)).
					WithArgs(args.userId, args.listId).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec(fmt.Sprintf("delete from %s", todoItemsTable)).
					WithArgs(args.listId).
					WillReturnError(errors.New("test error"))

				mock.ExpectRollback()
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.args)
			err := r.Delete(testCase.args.userId, testCase.args.listId)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
