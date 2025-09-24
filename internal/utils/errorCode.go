package utils

// 错误码定义（按业务模块分类）
const (
	// 通用错误
	ErrCodeSuccess         = 0     // 成功
	ErrCodeParamInvalid    = 10001 // 参数验证失败
	ErrCodeParamBind       = 10002 // 参数绑定失败
	ErrCodeParamTypeError  = 10003 // 参数类型错误
	ErrCodeParamOutOfRange = 10004 // 参数值超出合法范围
	ErrCodeDataFormatError = 10005 // 数据格式错误（如 JSON/XML 格式解析失败）

	// 用户/权限相关
	ErrCodePermissionDenied = 20001 // 权限不足（无访问该资源的权限）

	// 资源相关
	ErrCodeResourceNotFound = 30001 // 资源不存在（如查询的用户 ID / 订单 ID 不存在）
	ErrCodeDuplicateKey     = 30002 // 重复键错误（如创建重复的用户名、订单号等）

	// 服务器/系统相关
	ErrCodeServerInternalError = 50001 // 服务器内部错误（如代码异常、未捕获的异常）
)
