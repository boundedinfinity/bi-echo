package main

import (
    "github.com/labstack/echo"
    "fmt"
    "net/http"
    "io"
    "html/template"
    "github.com/labstack/echo/middleware"
)

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

func RenderStatic (c echo.Context) error {
    a, err := Asset(fmt.Sprintf("static%s", c.P(0)))

    if err != nil {
        panic(err)
    }

    return c.String(http.StatusOK, string(a))
}

func InitializeRenderers (e *echo.Echo) error {
    //e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.SetRenderer(&EchoRenderer{})

    e.GET("/static*", RenderStatic)

    return nil
}
