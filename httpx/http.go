package httpx

import (
    "fmt"
    "net/http"
    "time"
)

type HttpConfig struct {
    Name string `yaml:"name"`
    Host string `yaml:"host"`
    Port int    `yaml:"port"`
}

type HttpServer struct {
    Config *HttpConfig
    Router *Router
}

func NewHttpServer(config *HttpConfig) *HttpServer {
    return &HttpServer{
        Config: config,
    }
}

// StartHttpServer 启动http服务器
func (h *HttpServer)StartHttpServer() {
    addr := fmt.Sprintf("%s:%d", h.Config.Host, h.Config.Port)

    // 创建HTTP服务器
    ser := &http.Server{
        Addr:    addr,
        Handler: h.Router,
    }

    go func() {
        if err := ser.ListenAndServe(); err != nil {
            fmt.Println(fmt.Sprintf("[%s] http服务器启动失败:%s", time.Now().Format("2006-01-02 15:04:05.000000"), addr))
            panic(err)
        }
    }()
    fmt.Println(fmt.Sprintf("[%s] http服务器启动:%s", time.Now().Format("2006-01-02 15:04:05.000000"), addr))
}
