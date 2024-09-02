package utils

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/go-playground/validator/v10"
)

var Validate = validator.New()

// WriteJSON 响应编码返回
func WriteJSON(w http.ResponseWriter, status int, v any) error {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(status)
    return json.NewEncoder(w).Encode(v)
}

// WriteError 返回 error
func WriteError(w http.ResponseWriter, status int, err error) {
    WriteJSON(w, status, map[string]string{"error": err.Error()})
}

// ParseJSON 解析请求
func ParseJSON(r *http.Request, payload any) error {
    if r.Body == nil {
        return fmt.Errorf("missing request body")
    }
    // 创建一个 JSON 解码器将数据解码到结构体对象上
    return json.NewDecoder(r.Body).Decode(payload)
}
