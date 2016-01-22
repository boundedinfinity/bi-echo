package routers

import (
	"github.com/boundedinfinity/echo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/channels", &controllers.MainController{}, "get:Channels")
    beego.Router("/echo/?:channel", &controllers.EchoController{})
    beego.Router("/ws/?:channel", &controllers.WebsocketController{})
    beego.Router("/rest/?:channel", &controllers.RestController{})
}
