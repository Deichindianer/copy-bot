package copybot

import (
	"fmt"
	"io/ioutil"
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gin-gonic/gin"
)

type CopyBot struct {
	mux        *gin.Engine
	httpClient *http.Client
	converter  *md.Converter
}

func NewCopyBot() *CopyBot {
	cp := new(CopyBot)
	cp.mux = gin.New()
	cp.httpClient = http.DefaultClient
	cp.converter = md.NewConverter("", true, nil)
	cp.mux.GET("/convert", cp.convert)
	return cp
}

func (cp *CopyBot) convert(c *gin.Context) {
	url := c.Query("url")
	resp, err := cp.httpClient.Get(url)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to get url: %s", err))
	}
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to read html: %s", err))
	}
	markdown, err := cp.converter.ConvertBytes(html)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to convert html: %s", err))
	}
	c.String(http.StatusOK, string(markdown))
	return
}

func (cp *CopyBot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cp.mux.ServeHTTP(w, r)
}
