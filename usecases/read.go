package usecases

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/models"
)

func (uc *usecases) Read(ctx *gin.Context) (interface{}, error) {
	cx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	// business logic here

	// sent to repository for throw to database or third party services
	rows, err := uc.repo.Read(cx)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := make([]models.UsersReadOutputList, 0)
	for rows.Next() {
		v := models.UsersReadOutputList{}
		err = rows.Scan(&v.RowID, &v.Fullname, &v.Email)
		if err != nil {
			panic(err)
		}
		data = append(data, v)
	}

	// sent back to transport for client
	return models.UsersReadOutput{
		Users: data,
	}, nil
}
