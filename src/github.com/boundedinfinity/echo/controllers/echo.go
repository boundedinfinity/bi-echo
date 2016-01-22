package controllers

import (
    "github.com/astaxie/beego"
)

type EchoController struct {
    beego.Controller
}

func (c *EchoController) Get() {
    c.Data["Channel"] = c.Ctx.Input.Param(":channel")
    c.TplName = "echo.html"
}

