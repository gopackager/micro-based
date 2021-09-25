package repositories

import (
	"context"
	"database/sql"
	"sync"
)

func (r *repository) DetailByID(ctx context.Context, rowID int) (*sql.Row, error) {
	mute := sync.Mutex{}
	mute.Lock()
	defer mute.Unlock()

	row := r.DB.QueryRowContext(ctx, "SELECT id, fullname, email FROM users WHERE id = ?", rowID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}

func (r *repository) DetailByEmail(ctx context.Context, payload string) *sql.Row {
	mute := sync.Mutex{}
	mute.Lock()
	defer mute.Unlock()
	return r.DB.QueryRowContext(ctx, "SELECT id FROM users WHERE email = ?", payload)
}
