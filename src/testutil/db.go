package testutil

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ezio1119/fishapp-chat/conf"
	"github.com/jinzhu/gorm"
)

func NewGormMock(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, error) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gdb, err := gorm.Open(conf.C.Db.Dbms, db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub gorm database connection", err)
	}
	gdb.LogMode(true)
	return gdb, mock, nil
}
