package v1

// BaseResponse 定义了所有回复的公共字段
type BaseResponse struct {
	Success bool   `json:"success"` // 表示操作是否成功
	Message string `json:"message"` // 返回的提示信息
}

// Response 作为统一的回复结构体，可以用于返回任意复杂数据
type Response struct {
	BaseResponse
	Data interface{} `json:"data,omitempty"` // 数据部分，根据需要可以是 map、结构体、数组等
}
