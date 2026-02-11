package utils

// APIResponse API统一响应结构体
// 定义了所有API接口的响应格式
type APIResponse struct {
	Code    int         `json:"code"`    // 响应码：0表示成功，非0表示失败
	Message string      `json:"message"` // 响应消息：描述操作结果
	Data    interface{} `json:"data,omitempty"` // 响应数据：可选，包含操作结果数据
}

// Success 创建成功响应
// 参数：
//   - data: 响应数据
//   - message: 响应消息
// 返回：
//   - APIResponse: 成功响应对象
func Success(data interface{}, message string) APIResponse {
	return APIResponse{
		Code:    0,    // 成功响应码为0
		Message: message,
		Data:    data,
	}
}

// Error 创建错误响应
// 参数：
//   - code: 错误码（通常使用HTTP状态码）
//   - message: 错误消息
// 返回：
//   - APIResponse: 错误响应对象
func Error(code int, message string) APIResponse {
	return APIResponse{
		Code:    code,    // 错误码
		Message: message, // 错误消息
		// 错误响应不包含data字段
	}
}
