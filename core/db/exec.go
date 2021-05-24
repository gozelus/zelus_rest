package db

import (
	"context"
)

type execSQL interface {
	Exec(ctx context.Context, sql string, value ...interface{}) error
}
