package rest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/repositories"
	"github.com/gopackager/micro-based/usecases"
	"github.com/stretchr/testify/assert"
)

func Test_handler_Create(t *testing.T) {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // mock sql.DB
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.New(db, nil)
	uc := usecases.New(repo)
	tp := New(uc)

	tests := []struct {
		name         string
		url          string
		wantErr      bool
		wantRespCode int
		body         string
	}{
		// TODO: Add test cases.
		{
			name:         "should be success",
			url:          "http://localhost:8088/v1/users",
			wantErr:      false,
			wantRespCode: http.StatusOK,
			body:         `{"fullname": "suwondo","new_password": "suwondo","confirm_password": "suwondo","email": "popo@mail.com"}`,
		},
	}
	for _, tt := range tests {
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("v1/users", tp.Create)

		req, err := http.NewRequest(http.MethodPost, tt.url, strings.NewReader(tt.body))
		assert.NoError(t, err)
		req.Header.Add("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		assert.Equal(t, resp.Code, tt.wantRespCode)
	}
}
