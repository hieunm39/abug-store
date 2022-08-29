package middleware

import (
	"abug-store/common"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find from DB

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
	}
}