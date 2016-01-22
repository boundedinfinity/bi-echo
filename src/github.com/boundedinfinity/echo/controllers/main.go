package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "http://boundedinfinity.github.io"
	c.Data["Email"] = "brad.babb@boundedinfinity.com"
	c.TplName = "index.html"
}

func (c *MainController) Post() {
    channel := c.Input().Get("channel")

    if channel != "" {
        c.Redirect("/echo/" + channel, 302)
    } else {
        c.TplName = "index.html"
    }
}

func (c *MainController) Channels() {
    c.Data["json"] = ChannelList()
    c.ServeJSON()
}
