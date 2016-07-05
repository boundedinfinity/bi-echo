package main

//
// Reference
//
// https://github.com/Sirupsen/logrus
//

//go:generate go-bindata-assetfs -ignore=\\.gitignore view/... static/...

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "fmt"
    log "github.com/Sirupsen/logrus"
    "os"
    "net/http"
)

type Config struct {
    Port int
}

func main() {
    log.SetFormatter(&log.JSONFormatter{})
    log.SetOutput(os.Stdout)
    log.SetLevel(log.InfoLevel)

    config := Config{
        Port: 8080,
    }

    e := echo.New()

    err := InitializeWebsocket(e)

    if err != nil {
        panic(err)
    }

    err = InitializeRenderers(e)

    if err != nil {
        panic(err)
    }

    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "view/index.html", "")
    })

    log.Info(config)
    e.Run(standard.New(fmt.Sprintf(":%d", config.Port)))
}
