package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Render func gets context, data, template name
// and calls JSON, XML or HTML context methods with data parameter.
// The method called in context depends on the Accept value in the context request header.
func Render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
