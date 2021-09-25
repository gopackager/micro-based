package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gopackager/micro-based/models"
)

func (r *repository) CreateEvent(ctx context.Context, payload models.UsersCreateInput) (sql.Result, error) {
	payloads, _ := json.Marshal(payload)
	r.Queuer.Delay(10*time.Second, "users:create_new", payloads)
	r.Queuer.Flush()
	return nil, nil
}
