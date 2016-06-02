package main

//go:generate go-bindata-assetfs -ignore=\\.gitignore view/... static/...

import (
    "github.com/labstack/echo"
    "net/http"
    "github.com/labstack/echo/engine/standard"
    "fmt"
    "github.com/labstack/echo/middleware"
    "html/template"
    "io"
    "github.com/labstack/gommon/log"
)

type Config struct {
    Port int
}

type EchoRenderer struct {

}

func(r *EchoRenderer) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
    a, err := Asset(name)

    if err != nil {
        panic(err)
    }

    t, err := template.New(name).Parse(string(a))

    if err != nil {
        panic(err)
    }

    return t.ExecuteTemplate(w, name, data)
}

func main() {
    config := Config{
        Port: 8080,
    }

    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.SetRenderer(&EchoRenderer{})
    //e.Static("/static", "")

    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "view/index.html", "")
    })

    e.GET("/test", func(c echo.Context) error {
        return c.Render(http.StatusOK, "view/test.html", "")
    })

    e.GET("/static*", func(c echo.Context) error {
        log.Info("=====================")
        log.Info(c.P(0))
        log.Info("=====================")

        a, err := Asset(fmt.Sprintf("static%s", c.P(0)))

        if err != nil {
            panic(err)
        }

        return c.String(http.StatusOK, string(a))
    })

    e.Run(standard.New(fmt.Sprintf(":%d", config.Port)))
}
