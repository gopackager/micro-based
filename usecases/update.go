package usecases

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/models"
)

func (uc *usecases) Update(ctx *gin.Context, data interface{}) (interface{}, error) {
	cx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	// business logic here
	payload := data.(models.UsersUpdateInput)

	row, err := uc.repo.DetailByID(ctx, payload.RowID)
	if err != nil {
		return nil, err
	}

	var value models.UsersReadOutputList
	err = row.Scan(&value.RowID)
	if err == sql.ErrNoRows {
		return nil, errors.New("no records found")
	}

	// sent to repository for throw to database or third party services
	result, err := uc.repo.Update(cx, payload)
	if err != nil {
		return nil, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return nil, err
	}

	r, _ := uc.repo.DetailByID(ctx, payload.RowID)
	var val models.UsersReadOutputList
	r.Scan(&val.RowID, &val.Fullname, &val.Email)

	// encapsulation response and sent back to transport for client
	return val, nil
}
