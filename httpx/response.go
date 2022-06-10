package httpx

import (
    "encoding/json"
    "fmt"
    "net/http"
)

const (
    // ContentType means Content-Type.
    ContentType = "Content-Type"
    // ApplicationJson means application/json.
    ApplicationJson = "application/json"
)

type rsBody struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data,omitempty"`
}

// RequestNotFound 未知请求
func RequestNotFound(w http.ResponseWriter, r *http.Request) {
    OutputJson(w, 404, "not fount", nil)
}

func OutputJson(w http.ResponseWriter, code int, msg string, data interface{}) {
    w.Header().Set(ContentType, ApplicationJson)

    j, err := json.Marshal(rsBody{
        Code: code,
        Msg:  msg,
        Data: data,
    })
    if err != nil {
        fmt.Println("json", err.Error())
    }

    _, err = fmt.Fprintln(w, string(j))
    if err != nil {
        fmt.Println("write", err.Error())
    }
}
