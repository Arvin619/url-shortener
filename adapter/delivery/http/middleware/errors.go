package middleware

import (
	"errors"
	http1 "github.com/Arvin619/url-shortener/adapter/delivery/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) != 0 {
			err := ctx.Errors[0].Err
			var err1 *http1.Error
			switch {
			case errors.As(err, &err1):
			default:
				err1 = http1.NewError(http.StatusInternalServerError, "internal server error")
			}
			ctx.JSON(err1.StatusCode, err1)
		}
	}
}
