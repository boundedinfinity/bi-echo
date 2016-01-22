package controllers

import (
    "github.com/astaxie/beego"
    "github.com/gorilla/websocket"
    "net/http"
)

type WebsocketController struct {
    beego.Controller
}

func (c *WebsocketController) Get() {
    channel := c.Ctx.Input.Param(":channel")

    beego.Info("ws connection: channel: " + channel)

    if len(channel) == 0 {
        c.Redirect("/", 302)
        return
    }

    ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)

    if _, ok := err.(websocket.HandshakeError); ok {
        http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
        return
    } else if err != nil {
        beego.Error("Cannot setup WebSocket connection:", err)
        return
    }

    Join(channel, ws)
    defer Leave(channel)
}
