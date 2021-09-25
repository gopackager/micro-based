package repositories

import (
	"context"
	"database/sql"
	"sync"
)

func (r *repository) Delete(ctx context.Context, rowID int) (sql.Result, error) {
	mute := sync.Mutex{}
	mute.Lock()
	defer mute.Unlock()
	stmt, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	st, err := stmt.PrepareContext(ctx, "DELETE FROM users WHERE id = ?")
	if err != nil {
		stmt.Rollback()
		return nil, err
	}

	result, err := st.ExecContext(ctx, rowID)
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
