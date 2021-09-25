package repositories

import (
	"context"
	"database/sql"
	"sync"

	"github.com/gopackager/micro-based/models"
)

func (r *repository) Update(ctx context.Context, payload models.UsersUpdateInput) (sql.Result, error) {
	mute := sync.Mutex{}
	mute.Lock()
	defer mute.Unlock()
	stmt, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	var args []interface{}
	query := "UPDATE users SET"

	if payload.Fullname != "" {
		query += " fullname = ?"
		args = append(args, payload.Fullname)
	} else if payload.Email != "" {
		query += " email = ?"
		args = append(args, payload.Email)
	} else if payload.Email != "" && payload.Fullname != "" {
		query += " email = ?, fullname = ?"
		args = append(args, payload.Email, payload.Fullname)
	}

	query += " WHERE id = ?"
	args = append(args, payload.RowID)

	st, err := stmt.PrepareContext(ctx, query)
	if err != nil {
		stmt.Rollback()
		return nil, err
	}

	result, err := st.ExecContext(ctx, args...)
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
