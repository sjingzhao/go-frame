package core

import (
    "context"
    "fmt"
    "go-frame/crontabx"
    "go-frame/httpx"
    "go-frame/logx"
    "go-frame/scriptx"
    "sync"
    "time"
)

type ServerConfig struct {
}

type Server struct {
    name string
    host string
    port int

    wg sync.WaitGroup

    ctx        context.Context
    cancelFunc context.CancelFunc

    log *logx.Logger

    serverConfig *ServerConfig          // 配置
    customConfig map[string]interface{} // 自定义配置

    scriptServer  *scriptx.ScriptServer   // 脚本服务器
    crontabServer *crontabx.CrontabServer // 定时任务服务器
    rpcServer     string                  // rpc服务器
    httpServer    *httpx.HttpServer       // http服务器
}

func NewServer() *Server {
    ctx, cancelFunc := context.WithCancel(context.Background())

    return &Server{
        ctx:          ctx,
        cancelFunc:   cancelFunc,
        customConfig: make(map[string]interface{}),
    }
}

// SetServiceConfig 设置系统配置
func (s *Server) SetServiceConfig(config *ServerConfig) {
    s.serverConfig = config
}

// GetSystemConfig 设置系统配置
func (s *Server) GetSystemConfig() *ServerConfig {
    return s.serverConfig
}

// SetCustomConfig 设置系统配置
// Server.SetCustomConfig("system", core.LoadCustomConfig("system", "/Users/cltx/p/go-frame/etc/frame.yaml", core.NewSystemConfig()))
// fmt.Println(Server.GetCustomConfig("system").(*core.SystemConfig).Name)
func (s *Server) SetCustomConfig(key string, config interface{}) {
    s.customConfig[key] = config
}

// GetCustomConfig 获取系统配置
// Server.GetCustomConfig("mysql").(*core.SystemConfig).Name
// *core.SystemConfig 为断言的类型
func (s *Server) GetCustomConfig(key string) interface{} {
    return s.customConfig[key]
}

// SetScriptServer 设置脚本服务器
func (s *Server) SetScriptServer(scripts []scriptx.Script) {
    s.scriptServer.AddScripts(scripts)
}

// GetScriptServer 获取脚本服务器
func (s *Server) GetScriptServer() *scriptx.ScriptServer {
    return s.scriptServer
}

// SetCrontabServer 设置定时任务服务器
func (s *Server) SetCrontabServer() {

}

// SetRpcServer 设置rpc服务器
func (s *Server) SetRpcServer() {

}

// SetHttpServer 设置http服务器
func (s *Server) SetHttpServer(routes []httpx.Route) {
    fmt.Println(fmt.Sprintf("[%s] http服务器注册路由", time.Now().Format("2006-01-02 15:04:05.000000")))
    // 注册路由, 用于监听http请求转发
    router := httpx.NewRouter()
    for _, route := range routes {
        router.HandleFunc(route.Method, route.Path, route.Handler)
    }
    s.httpServer.Router = router
}

// GetHttpServer 获取http服务器
func (s *Server) GetHttpServer() *httpx.HttpServer {
    return s.httpServer
}

func (s *Server) LoadService() {
    // 项目配置
    fmt.Println(fmt.Sprintf("[%s] 加载系统配置", time.Now().Format("2006-01-02 15:04:05.000000")))
    systemConfig := LoadSystemConfig()
    s.SetServiceConfig(systemConfig.Server)

    // 自定义配置
    fmt.Println(fmt.Sprintf("[%s] 加载自定义配置", time.Now().Format("2006-01-02 15:04:05.000000")))
    s.SetCustomConfig("system", LoadCustomConfig("system", "./etc/frame.yaml", NewSystemConfig()))

    // 系统日志
    s.log = logx.NewLogger("").GetLogger("system")

    // 注册脚本
    fmt.Println(fmt.Sprintf("[%s] 脚本启动", time.Now().Format("2006-01-02 15:04:05.000000")))
    s.scriptServer = scriptx.NewScriptService()

    //fmt.Println(fmt.Sprintf("[%s] 定时任务启动", time.Now().Format("2006-01-02 15:04:05.000000")))
    //fmt.Println(fmt.Sprintf("[%s] rpc启动", time.Now().Format("2006-01-02 15:04:05.000000")))

    fmt.Println(fmt.Sprintf("[%s] 加载http服务器", time.Now().Format("2006-01-02 15:04:05.000000")))
    s.httpServer = httpx.NewHttpServer(systemConfig.Http)
}

func (s *Server) Start() {
    // 启动脚本服务
    if len(s.scriptServer.GetScripts()) > 0 {
        s.wg.Add(1)
        s.scriptServer.StartScriptServer()
    }

    // 启动http服务
    if s.httpServer.Router != nil {
        s.wg.Add(1)
        s.httpServer.StartHttpServer()
    }

    s.wg.Wait()
}
