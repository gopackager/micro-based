package repositories

import (
	"context"
	"database/sql"

	"github.com/gopackager/micro-based/helper/asynq/client"
	"github.com/gopackager/micro-based/models"
)

type Repository interface {
	Create(ctx context.Context, payload models.UsersCreateInput) (sql.Result, error)
	CreateEvent(ctx context.Context, payload models.UsersCreateInput) (sql.Result, error)
	Read(ctx context.Context) (*sql.Rows, error)
	DetailByID(ctx context.Context, rowID int) (*sql.Row, error)
	Update(ctx context.Context, payload models.UsersUpdateInput) (sql.Result, error)
	Delete(ctx context.Context, rowID int) (sql.Result, error)
	DetailByEmail(ctx context.Context, payload string) *sql.Row
}

type repository struct {
	DB     *sql.DB
	Queuer client.AsynQueueClient
}

func New(db *sql.DB, q client.AsynQueueClient) Repository {
	return &repository{
		DB:     db,
		Queuer: q,
	}
}
