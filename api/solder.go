package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleister/kleister-api/model/solder"
	"github.com/kleister/kleister-api/router/middleware/session"
	"github.com/kleister/kleister-api/store"
)

// SolderPacks retrieves the packs compatible to Technic Platform.
func SolderPacks(c *gin.Context) {
	records, _ := store.GetSolderPacks(
		c,
	)

	c.JSON(
		http.StatusOK,
		solder.NewPacksFromList(
			records,
			session.Client(c),
			session.Key(c),
			c.Query("include"),
		),
	)
}

// SolderPack retrieves the pack compatible to Technic Platform.
func SolderPack(c *gin.Context) {
	record, err := store.GetSolderPack(
		c,
		c.Param("pack"),
	)

	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "Modpack does not exist",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		solder.NewPackFromModel(
			record,
			session.Client(c),
			session.Key(c),
			c.Query("include"),
		),
	)
}

// SolderBuild retrieves the build compatible to Technic Platform.
func SolderBuild(c *gin.Context) {
	record, err := store.GetSolderBuild(
		c,
		c.Param("pack"),
		c.Param("build"),
	)

	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "Build does not exist",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		solder.NewBuildFromModel(
			record,
			session.Client(c),
			session.Key(c),
			c.Query("include"),
		),
	)
}

// SolderMods retrieves the mods compatible to Technic Platform.
func SolderMods(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"error": "No mod requested",
		},
	)
}

// SolderMod retrieves the mod compatible to Technic Platform.
func SolderMod(c *gin.Context) {
	record, res := store.GetMod(
		c,
		c.Param("mod"),
	)

	if res.RecordNotFound() {
		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "Mod does not exist",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		solder.NewModFromModel(
			record,
			session.Client(c),
			session.Key(c),
			c.Query("include"),
		),
	)
}

// SolderVersion retrieves the version compatible to Technic Platform.
func SolderVersion(c *gin.Context) {
	parent, res := store.GetMod(
		c,
		c.Param("mod"),
	)

	if res.RecordNotFound() {
		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "Mod does not exist",
			},
		)

		return
	}

	record, res := store.GetVersion(
		c,
		parent.ID,
		c.Param("version"),
	)

	if res.RecordNotFound() {
		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "Version does not exist",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		solder.NewVersionFromModel(
			record,
			session.Client(c),
			session.Key(c),
			c.Query("include"),
		),
	)
}
