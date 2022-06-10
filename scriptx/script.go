package scriptx

import (
    "fmt"
    "time"
)

// Script 脚本
type Script struct {
    Name    string
    Handler func()
}

type ScriptServer struct {
    scripts []Script
}

func NewScriptService() *ScriptServer {
    return &ScriptServer{
        scripts: make([]Script, 0),
    }
}

// AddScripts 添加脚本
func (s *ScriptServer) AddScripts(script []Script) {
    s.scripts = append(s.scripts, script...)
}

// GetScripts 获取脚本
func (s *ScriptServer) GetScripts() []Script {
    return s.scripts
}

// StartScriptServer 启动脚本
func (s *ScriptServer) StartScriptServer() {
    for _, script := range s.scripts {
        go script.Handler()
        fmt.Println(fmt.Sprintf("[%s] script服务器: %s", time.Now().Format("2006-01-02 15:04:05.000000"), script.Name))
    }
    fmt.Println(fmt.Sprintf("[%s] script服务器启动", time.Now().Format("2006-01-02 15:04:05.000000")))
}
