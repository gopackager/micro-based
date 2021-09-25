package repositories

import (
	"context"
	"database/sql"
	"sync"

	"github.com/gopackager/micro-based/models"
)

func (r *repository) Create(ctx context.Context, payload models.UsersCreateInput) (sql.Result, error) {
	mute := sync.Mutex{}
	mute.Lock()
	defer mute.Unlock()
	stmt, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	st, err := stmt.PrepareContext(ctx, "INSERT INTO users SET fullname = ?, password = ?, email = ?")
	if err != nil {
		stmt.Rollback()
		return nil, err
	}

	result, err := st.ExecContext(ctx, payload.Fullname, payload.ConfirmPassword, payload.Email)
	if err != nil {
		stmt.Rollback()
		return nil, err
	}
	defer st.Close()

	if err := stmt.Commit(); err != nil {
		stmt.Rollback()
		return nil, err
	}

	return result, nil
}
