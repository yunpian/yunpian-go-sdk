package yunpian

import (
	"fmt"
)

// ErrorResponse is the return format when the response fails
type ErrorResponse struct {
	HTTPStatusCode int    `json:"http_status_code"`
	Code           int    `json:"code"`
	Message        string `json:"msg"`
	Detail         string `json:"detail"`
}

func (resp ErrorResponse) Error() string {
	return fmt.Sprintf("请求失败，StatusCode: %d, Code: %d, Message: %s, Detail: %s", resp.HTTPStatusCode, resp.Code, resp.Message, resp.Detail)
}
