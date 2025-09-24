package dto

// DemoListRequest demo请求查询参数结构体
type DemoListRequest struct {
	Field1 int    `form:"field1"`
	Field2 string `form:"field2"`
}

// DemoListResponse demo响应结构体
type DemoListResponse struct {
	ID     int    `json:"id"`
	Field1 int    `json:"field1"`
	Field2 string `json:"field2"`
}

// PageQueryRequest 分页查询参数
type PageQueryRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

// DemoPageResponse 分页查询响应
type DemoPageListResponse struct {
	ID     int    `json:"id"`
	Field1 int    `json:"field1"`
	Field2 string `json:"field2"`
}

// DemoIDRequest ID请求参数
type DemoIDRequest struct {
	ID int `uri:"id"`
}

// DemoDetailResponse 详情查询响应
type DemoDetailResponse struct {
	ID     int    `json:"id"`
	Field1 int    `json:"field1"`
	Field2 string `json:"field2"`
}

// DemoCreateRequest demo创建请求参数结构体
type DemoCreateRequest struct {
	Field1 int    `json:"field1"`
	Field2 string `json:"field2"`
}

// DemoCreateResponse demo创建响应结构体
type DemoCreateResponse struct {
	ID int `json:"id"`
}

// DemoUpdateRequest demo更新请求参数结构体
type DemoUpdateRequest struct {
	Field1 *int    `json:"field1"`
	Field2 *string `json:"field2"`
}
