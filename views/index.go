package views

import (
	"github.com/gin-gonic/gin"
	"github.com/ordinary-dev/phoenix/database"
	"gorm.io/gorm"
	"net/http"
)

func ShowMainPage(c *gin.Context, db *gorm.DB) {
	// Get a list of groups with links
	var groups []database.Group
	result := db.
		Model(&database.Group{}).
		Preload("Links").
		Find(&groups)

	if result.Error != nil {
		ShowError(c, result.Error)
		return
	}

	c.HTML(http.StatusOK, "index.html.tmpl", gin.H{
		"groups": groups,
	})
}
