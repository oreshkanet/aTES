package database

//TODO: перевезти в отдельный пакет packages

import (
	"context"
	migrate "github.com/rubenv/sql-migrate"
)

type DB interface {
	Select(context.Context, interface{}, string, ...interface{}) error
	Insert(context.Context, string, interface{}) error
	Update(context.Context, string, interface{}) error
	Delete(context.Context, string, interface{}) error
	MigrateUp(migrations *migrate.MemoryMigrationSource) error
}

type DBParam struct {
	Name  string
	Value interface{}
}
