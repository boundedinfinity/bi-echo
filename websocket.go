package main

//
// Reference
//
// https://echo.labstack.com/recipes/websocket
//
//

import (
    "github.com/labstack/echo"
    "github.com/gorilla/websocket"
    "container/list"
    "time"
    "encoding/json"
    "net/http"
    "fmt"
    "github.com/labstack/echo/engine/standard"
    log "github.com/Sirupsen/logrus"
)

func InitializeWebsocket(e *echo.Echo) error {

    e.GET("/ws", standard.WrapHandler(http.HandlerFunc(ws())))

    return nil
}

var (
    upgrader = websocket.Upgrader{}
)

func ws() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        c, err := upgrader.Upgrade(w, r, nil)

        if err != nil {
            log.Print("upgrade:", err)
            return
        }
        defer c.Close()

        for {
            // Write
            err := c.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
            if err != nil {
                log.Fatal(err)
            }

            // Read
            _, msg, err := c.ReadMessage()
            if err != nil {
                log.Fatal(err)
            }
            fmt.Printf("%s\n", msg)
        }
    }
}

type EventType int

const (
    EVENT_JOIN = iota
    EVENT_LEAVE
)

type Event struct {
    Type      EventType // JOIN, LEAVE
    Channel   string
    Timestamp int
}

type Message struct {
    Timestamp int
    RestDescriptor RestDescriptor
}

type Subscriber struct {
    Channel string
    Conn *websocket.Conn
}

var (
    subscribe = make(chan Subscriber, 10)
    unsubscribe = make(chan string, 10)
    publish = make(chan RestDescriptor, 10)
    subscribers = list.New()
)

func newEvent(eventType EventType, channel string) Event {
    return Event{
        Type: eventType,
        Channel: channel,
        Timestamp: int(time.Now().Unix()),
    }
}

func Join(channel string, ws *websocket.Conn) {
    subscribe <- Subscriber{Channel: channel, Conn: ws}
}

func Leave(channel string) {
    unsubscribe <- channel
}

func Publish(descriptor RestDescriptor) {
    publish <- descriptor
}

func direct() {
    for {
        select {
        case sub := <-subscribe:
            subscribeToChannel(sub)
        case message := <-publish:
            sendToChannel(message)
        case unsub := <-unsubscribe:
            unsubscribeToChannel(unsub)
        }
    }
}

func cleanup() {
    for {
        for item := subscribers.Front(); item != nil; item = item.Next() {
            sub := item.Value.(Subscriber)

            if sub.Channel == "" || sub.Conn == nil {
                //beego.Warn("Reaping: " + sub.Channel)
                unsubscribe <- sub.Channel
            }
        }

        time.Sleep(1 * time.Second)
    }
}

func init() {
    go direct()
    go cleanup()
}

func subscribeToChannel(subscription Subscriber) {
    if !ChannelExists(subscription.Channel) {
        subscribers.PushBack(subscription)
        //beego.Info("New Channel:", subscription.Channel, ";WebSocket:", subscription.Conn != nil)
    } else {
        //beego.Info("Old Channel:", subscription.Channel, ";WebSocket:", subscription.Conn != nil)
    }
}

func unsubscribeToChannel(channel string) {
    for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
        if sub.Value.(Subscriber).Channel == channel {
            subscribers.Remove(sub)

            ws := sub.Value.(Subscriber).Conn

            if ws != nil {
                ws.Close()
                //beego.Error("WebSocket closed:", channel)
            }

            break
        }
    }
}

func sendToChannel(descriptor RestDescriptor) {
    data, err := json.Marshal(descriptor)

    //beego.Info("Descriptor:", descriptor.Channel, ";Method:", descriptor.Method)

    if err != nil {
        //beego.Error("Fail to marshal descriptor:", err)
        return
    }

    for item := subscribers.Front(); item != nil; item = item.Next() {
        sub := item.Value.(Subscriber)

        if sub.Channel == descriptor.Channel {
            ws := sub.Conn

            if ws != nil {
                if ws.WriteMessage(websocket.TextMessage, data) != nil {
                    unsubscribe <- sub.Channel
                }
            }
        }
    }
}

func sendToBroadcast(descriptor RestDescriptor) {
    data, err := json.Marshal(descriptor)

    //beego.Info("Descriptor:", descriptor.Channel, ";Method:", descriptor.Method)

    if err != nil {
        //beego.Error("Fail to marshal descriptor:", err)
        return
    }

    for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
        ws := sub.Value.(Subscriber).Conn

        if ws != nil {
            if ws.WriteMessage(websocket.TextMessage, data) != nil {
                unsubscribe <- sub.Value.(Subscriber).Channel
            }
        }
    }
}

func ChannelExists(channel string) bool {
    for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
        if sub.Value.(Subscriber).Channel == channel {
            return true
        }
    }
    return false
}

func ChannelList() []string {
    channels := make([]string, 0)

    for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
        channels = append(channels, sub.Value.(Subscriber).Channel)
    }

    return channels
}
