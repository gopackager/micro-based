package usecases

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/models"
)

func (uc *usecases) Delete(ctx *gin.Context, data int) (interface{}, error) {
	cx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	// business logic here
	row, err := uc.repo.DetailByID(ctx, data)
	if err != nil {
		return nil, err
	}

	var value models.UsersReadOutputList
	err = row.Scan(&value.RowID, &value.Fullname, &value.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("no records found")
	}

	// sent to repository for throw to database or third party services
	_, err = uc.repo.Delete(cx, data)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
