package database

import "context"

type DB interface {
	Select(context.Context, interface{}, string, ...interface{}) error
	Insert(context.Context, string, interface{}) error
	Update(context.Context, string, interface{}) error
	Delete(context.Context, string, interface{}) error
}
