package middleware_mock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type AuthMiddlewareMock struct {
	mock.Mock
}

func (a *AuthMiddlewareMock) RequireToken(roles ...string) gin.HandlerFunc {
	return func(context *gin.Context) {
		// yg penting sebuah handlerfunc
	}
}
