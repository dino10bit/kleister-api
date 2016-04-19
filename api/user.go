package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/session"
	"github.com/solderapp/solder-api/store"
)

// GetUsers retrieves all available users.
func GetUsers(c *gin.Context) {
	records, err := store.GetUsers(
		c,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch users",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// GetUser retrieves a specific user.
func GetUser(c *gin.Context) {
	record := session.User(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// DeleteUser removes a specific user.
func DeleteUser(c *gin.Context) {
	record := session.User(c)

	err := store.DeleteUser(
		c,
		record,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted user",
		},
	)
}

// PatchUser updates an existing user.
func PatchUser(c *gin.Context) {
	record := session.User(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind user data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user data",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateUser(
		c,
		record,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// PostUser creates a new user.
func PostUser(c *gin.Context) {
	record := &model.User{
		Permission: &model.Permission{},
	}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind user data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind user data",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateUser(
		c,
		record,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// GetUserMods retrieves all mods related to a user.
func GetUserMods(c *gin.Context) {
	user := session.User(c)

	records, err := store.GetUserMods(
		c,
		user.ID,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch mods",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// PatchUserMod appends a mod to a user.
func PatchUserMod(c *gin.Context) {
	user := session.User(c)
	mod := session.Mod(c)

	assigned := store.GetUserHasMod(
		c,
		user.ID,
		mod.ID,
	)

	if assigned == true {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Mod is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateUserMod(
		c,
		user.ID,
		mod.ID,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append mod",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended mod",
		},
	)
}

// DeleteUserMod deleted a mod from a user
func DeleteUserMod(c *gin.Context) {
	user := session.User(c)
	mod := session.Mod(c)

	assigned := store.GetUserHasMod(
		c,
		user.ID,
		mod.ID,
	)

	if assigned == false {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Mod is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeleteUserMod(
		c,
		user.ID,
		mod.ID,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink mod",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked mod",
		},
	)
}
