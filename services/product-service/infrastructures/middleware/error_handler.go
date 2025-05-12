package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zhunismp/nanow4t3r/services/product/core/errors"
)

func validationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", e.Field(), e.Param())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email", e.Field())
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", e.Field(), e.Param())
	default:
		return fmt.Sprintf("%s is invalid", e.Field())
	}
}

func statusFromAppError(t errors.ErrorType) int {
	switch t {
	case errors.Validation:
		return http.StatusBadRequest
	case errors.NotFound:
		return http.StatusNotFound
	case errors.Conflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Handle application errors
		if v, exists := c.Get(errors.AppErrorKey); exists {
			if appErr, ok := v.(*errors.AppError); ok {
				c.JSON(statusFromAppError(appErr.Type), gin.H{"error": appErr.Message})
				return
			}
		}

		// Handle Gin errors
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				switch e.Type {
				case gin.ErrorTypeBind:
					if verrs, ok := e.Err.(validator.ValidationErrors); ok {
						errMap := make(map[string]string)
						for _, fe := range verrs {
							errMap[fe.Field()] = validationErrorToText(fe)
						}
						c.JSON(http.StatusBadRequest, gin.H{"errors": errMap})
						return
					}

					c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
					return

				case gin.ErrorTypePublic:
					c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
					return

				default:
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
					return
				}
			}
		}
	}
}
