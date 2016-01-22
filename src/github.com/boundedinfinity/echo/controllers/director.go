package controllers

import (
    "github.com/boundedinfinity/echo/models"
    "time"
    "github.com/gorilla/websocket"
    "container/list"
    "github.com/astaxie/beego"
    "encoding/json"
)

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
    RestDescriptor models.RestDescriptor
}

type Subscriber struct {
    Channel string
    Conn *websocket.Conn
}

var (
    subscribe = make(chan Subscriber, 10)
    unsubscribe = make(chan string, 10)
    publish = make(chan models.RestDescriptor, 10)
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

func Publish(descriptor models.RestDescriptor) {
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
                beego.Warn("Reaping: " + sub.Channel)
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
        beego.Info("New Channel:", subscription.Channel, ";WebSocket:", subscription.Conn != nil)
    } else {
        beego.Info("Old Channel:", subscription.Channel, ";WebSocket:", subscription.Conn != nil)
    }
}

func unsubscribeToChannel(channel string) {
    for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
        if sub.Value.(Subscriber).Channel == channel {
            subscribers.Remove(sub)

            ws := sub.Value.(Subscriber).Conn

            if ws != nil {
                ws.Close()
                beego.Error("WebSocket closed:", channel)
            }

            break
        }
    }
}

func sendToChannel(descriptor models.RestDescriptor) {
    data, err := json.Marshal(descriptor)

    beego.Info("Descriptor:", descriptor.Channel, ";Method:", descriptor.Method)

    if err != nil {
        beego.Error("Fail to marshal descriptor:", err)
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

func sendToBroadcast(descriptor models.RestDescriptor) {
    data, err := json.Marshal(descriptor)

    beego.Info("Descriptor:", descriptor.Channel, ";Method:", descriptor.Method)

    if err != nil {
        beego.Error("Fail to marshal descriptor:", err)
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
