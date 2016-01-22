package models

type RestDescriptor struct {
    Channel string `json:"channel"`
    Referer string `json:"referer"`
    Method string `json:"method"`
    Body string `json:"body"`
    Timestamp int `json:"timestamp"`
}

type RestResponse struct {
    Channel string `json:"channel"`
    Message string `json:"message"`
}
