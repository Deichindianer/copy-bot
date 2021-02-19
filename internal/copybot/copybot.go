package copybot

import (
	"fmt"
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gin-gonic/gin"
)

type CopyBot struct {
	mux       *gin.Engine
	converter *md.Converter
}

func New() *CopyBot {
	cp := new(CopyBot)
	cp.mux = gin.New()
	cp.converter = md.NewConverter("", true, nil)
	cp.mux.GET("/convert", cp.convert)
	return cp
}

func (cp *CopyBot) convert(c *gin.Context) {
	url := c.Query("url")
	markdown, err := cp.converter.ConvertURL(url)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to convert html: %s", err))
		return
	}
	c.String(http.StatusOK, markdown)
}

func (cp *CopyBot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cp.mux.ServeHTTP(w, r)
}
