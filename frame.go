package main

import (
    "fmt"
    "github.com/sjingzhao/go-frame/core"
    "github.com/sjingzhao/go-frame/logic"
    "time"
)

func main() {
    fmt.Println(fmt.Sprintf("[%s] 服务启动", time.Now().Format("2006-01-02 15:04:05.000000")))
    server := core.NewServer()
    server.LoadService()

    // 注册路由
    server.SetHttpServer(logic.RegisterRoute())
    // 注册脚本
    //server.SetScriptServer(logic.RegisterScript())

    server.Start()
}
