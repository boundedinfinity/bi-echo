package controllers

import (
    "github.com/astaxie/beego"
    "github.com/boundedinfinity/echo/models"
    "net/http"
    "time"
    "io/ioutil"
)

type RestController struct {
    beego.Controller
}

func handle(c *RestController, method string) {
    channel := c.Ctx.Input.Param(":channel")

    if channel == "" {
        c.Data["json"] = models.RestResponse{
            Channel: "",
            Message: "invalid input",
        }

        c.Ctx.Output.SetStatus(http.StatusForbidden)
    } else {
        if(ChannelExists(channel)) {
            body, err := ioutil.ReadAll(c.Ctx.Request.Body)

            if err != nil {
                c.Ctx.Output.SetStatus(http.StatusInternalServerError)

                c.Data["json"] = models.RestResponse {
                    Channel: channel,
                    Message: err.Error(),
                }
            } else {
                Publish(models.RestDescriptor {
                    Channel: channel,
                    Referer: c.Ctx.Request.Referer(),
                    Method: method,
                    Timestamp: int(time.Now().Unix()),
                    Body: string(body[:]),
                })

                c.Data["json"] = models.RestResponse {
                    Channel: channel,
                    Message: "sent",
                }
            }
        } else {
            c.Data["json"] = models.RestResponse {
                Channel: channel,
                Message: "no subscribers",
            }
        }
    }

    c.ServeJSON()
}

func (c *RestController) Get() {
    handle(c, "GET");
}

func (c *RestController) Post() {
    handle(c, "POST");
}

