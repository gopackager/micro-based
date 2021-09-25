package usecases

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/helper"
	"github.com/gopackager/micro-based/models"
)

func (uc *usecases) Create(ctx *gin.Context, data interface{}) (interface{}, error) {
	cx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	// business logic here
	payload := data.(models.UsersCreateInput)

	// validate is email already registered
	if err := uc.checkEmailIsAlready(cx, payload.Email); err != nil {
		return nil, err
	}

	payload.ConfirmPassword = string(helper.GenerateFromPassword([]byte(payload.ConfirmPassword)))

	// sent to repository for connecting to database or third party services
	// result, err := uc.repo.Create(cx, payload)
	// if err != nil {
	// 	return nil, err
	// }

	// id, _ := result.LastInsertId()
	uc.repo.CreateEvent(cx, payload)

	// sent back to transport for client
	return models.UsersCreateOutput{
		// RowID:    id,
		Email:    payload.Email,
		Fullname: payload.Fullname,
	}, nil
}

func (uc *usecases) checkEmailIsAlready(ctx context.Context, email string) error {
	row := uc.repo.DetailByEmail(ctx, email)
	if row.Err() != nil {
		return row.Err()
	}

	var id int64
	err := row.Scan(&id)
	if err != sql.ErrNoRows {
		return errors.New("this email already registered")
	}
	return nil
}
