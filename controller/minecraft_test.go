package controller

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"

	. "github.com/franela/goblin"
)

func TestMinecraft(t *testing.T) {
	gin.SetMode(gin.TestMode)
	store := *model.Test()

	g := Goblin(t)
	g.Describe("GetMinecraft", func() {
		var minecrafts model.Minecrafts

		g.BeforeEach(func() {
			minecrafts = model.Minecrafts{
				&model.Minecraft{
					Name: "1.4.0",
					Type: "snapshot",
				},
				&model.Minecraft{
					Name: "1.10.4",
					Type: "release",
				},
				&model.Minecraft{
					Name: "1.8.0",
					Type: "release",
				},
			}

			for _, record := range minecrafts {
				store.Create(record)
			}
		})

		g.AfterEach(func() {
			store.Delete(&model.Minecraft{})
		})

		g.It("should respond with json content type", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetMinecraft(ctx)

			g.Assert(rw.Code).Equal(200)
			g.Assert(rw.HeaderMap.Get("Content-Type")).Equal("application/json; charset=utf-8")
		})

		g.It("should serve a collection", func() {
			ctx, rw, _ := gin.CreateTestContext()
			ctx.Set("store", store)

			GetMinecraft(ctx)

			out := model.Minecrafts{}
			json.NewDecoder(rw.Body).Decode(&out)

			g.Assert(len(out)).Equal(len(minecrafts))
			g.Assert(out[0]).Equal(minecrafts[2])
		})
	})
}