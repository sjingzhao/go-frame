package httpx

import (
    "github.com/mitchellh/mapstructure"
    "net/http"
)

// ParamsAnalysis get和post参数解析
func ParamsAnalysis(r *http.Request, rq interface{}) {
    GetParamsAnalysis(r, rq)
    PostParamsAnalysis(r, rq)
}

// GetParamsAnalysis get参数解析
func GetParamsAnalysis(r *http.Request, rq interface{}) {
    params := make(map[string]interface{}, len(r.Form))
    for name := range r.URL.Query() {
        v := r.URL.Query().Get(name)
        if len(v) > 0 {
            params[name] = v
        }
    }
    _ = mapstructure.WeakDecode(params, rq)
}

// PostParamsAnalysis post参数解析
func PostParamsAnalysis(r *http.Request, rq interface{}) {
    if err := r.ParseForm(); err != nil {
        return
    }

    params := make(map[string]interface{}, len(r.Form))
    for name := range r.Form {
        v := r.Form.Get(name)
        if len(v) > 0 {
            params[name] = v
        }
    }
    _ = mapstructure.WeakDecode(params, rq)
}

func PathParamsAnalysis() {

}
