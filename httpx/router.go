package httpx

import (
    "net/http"
    "strings"
)

// Middleware http中间件
type Middleware func(next http.HandlerFunc) http.HandlerFunc

// Route http路由
type Route struct {
    Method  string
    Path    string
    Handler http.HandlerFunc
}

type Router struct {
    Route map[string]map[string]http.HandlerFunc
}

// MiddlewareRoutes 注册带中间件的路由
func MiddlewareRoutes(middlewares []Middleware, routes []Route) []Route {
    l := len(middlewares)
    for i := l - 1; i >= 0; i-- {
        for x, _ := range routes {
            routes[x] = Route{
                Method:  routes[x].Method,
                Path:    routes[x].Path,
                Handler: middlewares[i](routes[x].Handler),
            }
        }
    }
    return routes
}

func NewRouter() *Router {
    return &Router{}
}

// ServeHTTP 实现Handler接口，匹配方法以及路径
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    if h, ok := r.Route[req.Method][req.URL.Path]; ok {
        h(w, req)
    } else {
        RequestNotFound(w, req)
    }
}

// HandleFunc 根据方法、路径将方法注册到路由
func (r *Router) HandleFunc(method, path string, f http.HandlerFunc) {
    if r.Route == nil {
        r.Route = make(map[string]map[string]http.HandlerFunc)
    }
    if r.Route[method] == nil {
        r.Route[method] = make(map[string]http.HandlerFunc)
    }
    r.Route[method][path] = f
}

// LogFileName url转日志文件名
func LogFileName(str string) string {
    index := strings.Index(str, "/")
    if index == 0 {
        str = str[1:]
    }

    split := strings.Split(str, "/")
    return strings.Join(split, "_")
}