package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/context"
)

const (
	// ForgeContextKey defines the context key that stores the forge.
	ForgeContextKey = "forge"
)

// Forge gets the forge from the context.
func Forge(c *gin.Context) *model.Forge {
	v, ok := c.Get(ForgeContextKey)

	if !ok {
		return nil
	}

	r, ok := v.(*model.Forge)

	if !ok {
		return nil
	}

	return r
}

// SetForge injects the forge into the context.
func SetForge() gin.HandlerFunc {
	return func(c *gin.Context) {
		record := &model.Forge{}

		res := context.Store(c).Where(
			"forges.id = ?",
			c.Param("forge"),
		).Or(
			"forges.slug = ?",
			c.Param("forge"),
		).First(
			&record,
		)

		if res.RecordNotFound() {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  http.StatusNotFound,
					"message": "Failed to find Forge version",
				},
			)

			c.Abort()
		} else {
			c.Set(ForgeContextKey, record)
			c.Next()
		}
	}
}
