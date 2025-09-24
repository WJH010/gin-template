package controller

import (
	"fmt"
	"gin-template/internal/demo/dto"
	"gin-template/internal/demo/service"
	"gin-template/internal/utils"

	"github.com/gin-gonic/gin"
)

// DemoController 控制器，持有服务层接口实例
type DemoController struct {
	service service.DemoService
}

// NewDemoController 创建控制器实例
func NewDemoController(demoService service.DemoService) *DemoController {
	return &DemoController{
		// 注入服务层实例
		service: demoService,
	}
}

func (ctr *DemoController) ListDemo(ctx *gin.Context) {
	fmt.Println("查询参数：", ctx.Request.URL.Query())
	// 初始化参数结构体并绑定查询参数
	var req dto.DemoListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		utils.HandlerFunc(ctx, utils.NewSystemError(fmt.Errorf("参数绑定失败: %w", err)))
		return
	}
	// 调用服务层
	demo, err := ctr.service.ListDemo(ctx, &req)

	// 处理服务层返回的错误
	if err != nil {
		utils.HandlerFunc(ctx, err)
		return
	}

	// 返回数据
	utils.Success(ctx, "获取成功", demo)
}

// ListDemoPage 分页查询demo数据
func (ctr *DemoController) ListDemoPage(ctx *gin.Context) {
	// 从上下文获取分页参数
	var pageQuery dto.PageQueryRequest
	// 从 URL 查询参数中提取数据并绑定到结构体
	if err := ctx.ShouldBindQuery(&pageQuery); err != nil {
		utils.HandlerFunc(ctx, utils.NewSystemError(fmt.Errorf("分页参数绑定失败: %w", err)))
		return
	}

	// 处理分页参数默认值
	var page, pageSize int
	page = pageQuery.Page
	pageSize = pageQuery.PageSize
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// 调用服务层
	demo, total, err := ctr.service.ListDemoPage(ctx, page, pageSize)

	// 处理服务层返回的错误
	if err != nil {
		utils.HandlerFunc(ctx, err)
		return
	}

	// 返回数据
	utils.SuccessPage(ctx, "获取成功", total, page, pageSize, demo)
}

// GetDemoByID 根据ID获取demo数据
func (ctr *DemoController) GetDemoByID(ctx *gin.Context) {
	// 从 URL 参数中提取 ID
	idReq := &dto.DemoIDRequest{}
	if err := ctx.ShouldBindUri(idReq); err != nil {
		utils.HandlerFunc(ctx, utils.NewSystemError(fmt.Errorf("ID 绑定失败: %w", err)))
		return
	}

	// 调用服务层
	demo, err := ctr.service.GetDemoByID(ctx, idReq.ID)

	// 处理服务层返回的错误
	if err != nil {
		utils.HandlerFunc(ctx, err)
		return
	}

	// 返回数据
	utils.Success(ctx, "获取成功", demo)
}

// CreateDemo 创建demo数据
func (ctr *DemoController) CreateDemo(ctx *gin.Context) {
	// 初始化参数结构体并绑定请求体
	var req dto.DemoCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandlerFunc(ctx, utils.NewSystemError(fmt.Errorf("参数绑定失败: %w", err)))
		return
	}
	// 调用服务层
	resp, err := ctr.service.CreateDemo(ctx, &req)
	if err != nil {
		utils.HandlerFunc(ctx, err)
		return
	}
	// 返回数据
	utils.Success(ctx, "创建成功", resp)
}

// BatchCreateDemo 批量创建demo数据
func (ctr *DemoController) BatchCreateDemo(ctx *gin.Context) {
	// 初始化参数结构体并绑定请求体
	var req []*dto.DemoCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandlerFunc(ctx, utils.NewSystemError(fmt.Errorf("参数绑定失败: %w", err)))
		return
	}
	// 调用服务层
	err := ctr.service.BatchCreateDemo(ctx, req)
	if err != nil {
		utils.HandlerFunc(ctx, err)
		return
	}
	// 返回数据
	utils.Success(ctx, "批量创建成功", nil)
}

// UpdateDemo 更新demo数据
func (ctr *DemoController) UpdateDemo(ctx *gin.Context) {
	// 从 URL 参数中提取 ID
	var idReq dto.DemoIDRequest
	if err := ctx.ShouldBindUri(&idReq); err != nil {
		utils.HandlerFunc(ctx, utils.NewSystemError(fmt.Errorf("ID 绑定失败: %w", err)))
		return
	}
	// 初始化参数结构体并绑定请求体
	var req dto.DemoUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandlerFunc(ctx, utils.NewSystemError(fmt.Errorf("参数绑定失败: %w", err)))
		return
	}
	// 调用服务层
	err := ctr.service.UpdateDemo(ctx, idReq.ID, req)
	if err != nil {
		utils.HandlerFunc(ctx, err)
		return
	}
	// 返回数据
	utils.Success(ctx, "更新成功", nil)
}

// SoftDeleteDemo 软删除demo数据
func (ctr *DemoController) SoftDeleteDemo(ctx *gin.Context) {
	// 从 URL 参数中提取 ID
	var idReq dto.DemoIDRequest
	if err := ctx.ShouldBindUri(&idReq); err != nil {
		utils.HandlerFunc(ctx, utils.NewSystemError(fmt.Errorf("ID 绑定失败: %w", err)))
		return
	}
	// 调用服务层
	err := ctr.service.SoftDeleteDemo(ctx, idReq.ID)
	if err != nil {
		utils.HandlerFunc(ctx, err)
		return
	}
	// 返回数据
	utils.Success(ctx, "删除成功", nil)
}

// DeleteDemo 删除demo数据
func (ctr *DemoController) DeleteDemo(ctx *gin.Context) {
	// 从 URL 参数中提取 ID
	var idReq dto.DemoIDRequest
	if err := ctx.ShouldBindUri(&idReq); err != nil {
		utils.HandlerFunc(ctx, utils.NewSystemError(fmt.Errorf("ID 绑定失败: %w", err)))
		return
	}
	// 调用服务层
	err := ctr.service.DeleteDemo(ctx, idReq.ID)
	if err != nil {
		utils.HandlerFunc(ctx, err)
		return
	}
	// 返回数据
	utils.Success(ctx, "删除成功", nil)
}
