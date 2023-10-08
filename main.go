package main

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	ginutil "github.com/jtagcat/util/gin"
)

func main() {
	ctx := context.Background()
	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)

	router := gin.Default()

	router.GET("/*redirectURL", func(c *gin.Context) {
		escaped := c.Param("redirectURL")[1:] // [1:]: strip leading slash
		if escaped == "/" {
			c.Data(http.StatusBadRequest, "text/plain", []byte("Must redirect elsewhere"))
			return
		}

		unsecaped, err := url.PathUnescape(escaped)
		if err != nil {
			c.Data(http.StatusBadRequest, "text/plain", []byte(err.Error()))
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, unsecaped)
	})

	ginutil.RunWithContext(ctx, router)
}
