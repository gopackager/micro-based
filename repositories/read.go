package repositories

import (
	"context"
	"database/sql"
	"sync"
)

func (r *repository) Read(ctx context.Context) (*sql.Rows, error) {
	mute := sync.Mutex{}
	mute.Lock()
	defer mute.Unlock()
	return r.DB.QueryContext(ctx, "SELECT id, fullname, email FROM users")
}
