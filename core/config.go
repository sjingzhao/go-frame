package core

import (
    "fmt"
    "github.com/sjingzhao/go-frame/httpx"
    "github.com/sjingzhao/go-frame/logx"
    "io/ioutil"

    "gopkg.in/yaml.v2"
)

// SystemConfig 系统配置
type SystemConfig struct {
    Server *ServerConfig `yaml:"Server"`

    Http *httpx.HttpConfig `yaml:"http"`

    Log *logx.LogConfig `yaml:"log"`
}

func NewSystemConfig() *SystemConfig {
    return &SystemConfig{}
}

// LoadSystemConfig 加载系统配置
func LoadSystemConfig() *SystemConfig {
    config := NewSystemConfig()
    parseConfigFile("./etc/frame.yaml", &config)
    return config
}

// LoadCustomConfig 加载自定义配置
// key:  类型的名字
// path: yaml文件路径
// conf: 对应结构引用
func LoadCustomConfig(key string, path string, conf interface{}) interface{} {
    parseConfigFile(path, conf)
    return conf
}

// 解析配置 yaml文件 => 结构
// path: yaml文件路径
// conf: 对应结构引用
func parseConfigFile(path string, conf interface{}) {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        panic(fmt.Sprintf("加载系统配置错误:%s", err.Error()))
    }

    if err = yaml.Unmarshal(content, conf); err != nil {
        panic(fmt.Sprintf("解析系统配置错误:%s", err.Error()))
    }
}
