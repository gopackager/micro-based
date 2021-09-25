package usecases

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/models"
)

func (uc *usecases) Detail(ctx *gin.Context, data int) (interface{}, error) {
	cx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	// business logic here

	// sent to repository for throw to database or third party services
	row, err := uc.repo.DetailByID(cx, data)
	if err != nil {
		return nil, err
	}

	var result models.UsersReadOutputList
	err = row.Scan(&result.RowID, &result.Fullname, &result.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("no records found")
	}

	// sent back to transport for client
	return result, nil
}
