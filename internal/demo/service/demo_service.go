package service

import (
	"context"
	"gin-template/internal/demo/dto"
	"gin-template/internal/demo/model"
	"gin-template/internal/demo/repository"
	"gin-template/internal/utils"
)

// DemoService 服务接口，定义服务应该提供的功能
type DemoService interface {
	// ListDemo 获取demo数据
	ListDemo(ctx context.Context, req *dto.DemoListRequest) ([]*dto.DemoListResponse, error)
	// ListDemoPage 分页查询demo数据
	ListDemoPage(ctx context.Context, page, pageSize int) ([]*dto.DemoPageListResponse, int64, error)
	// GetDemoByID 根据ID获取demo数据
	GetDemoByID(ctx context.Context, id int) (*dto.DemoDetailResponse, error)
	// CreateDemo 创建demo数据
	CreateDemo(ctx context.Context, demo *dto.DemoCreateRequest) (*dto.DemoCreateResponse, error)
	// BatchCreateDemo 批量创建demo数据
	BatchCreateDemo(ctx context.Context, demos []*dto.DemoCreateRequest) error
	// UpdateDemo 更新demo数据
	UpdateDemo(ctx context.Context, id int, req dto.DemoUpdateRequest) error
	// SoftDeleteDemo 软删除demo数据
	SoftDeleteDemo(ctx context.Context, id int) error
	// DeleteDemo 删除demo数据
	DeleteDemo(ctx context.Context, id int) error
}

// DemoServiceImpl 实现接口的具体结构体，持有数据访问层接口 Repository 的实例
type DemoServiceImpl struct {
	demoRepo repository.DemoRepository
}

// NewDemoService 创建服务实例
func NewDemoService(demoRepo repository.DemoRepository) DemoService {
	return &DemoServiceImpl{demoRepo: demoRepo}
}

// ListDemo 获取demo数据
func (svc *DemoServiceImpl) ListDemo(ctx context.Context, req *dto.DemoListRequest) ([]*dto.DemoListResponse, error) {
	// 调用数据访问层方法获取数据
	demo, err := svc.demoRepo.ListDemo(ctx, req)
	if err != nil {
		return nil, err
	}
	// 转换为dto（领域模型 -> 数据传输对象）
	var demoResp []*dto.DemoListResponse
	for _, v := range demo {
		demoResp = append(demoResp, &dto.DemoListResponse{
			ID:     v.ID,
			Field1: v.Field1,
			Field2: v.Field2,
		})
	}
	// 返回数据传输对象
	return demoResp, nil
}

// ListDemoPage 分页查询demo数据
func (svc *DemoServiceImpl) ListDemoPage(ctx context.Context, page, pageSize int) ([]*dto.DemoPageListResponse, int64, error) {
	// 调用数据访问层方法获取数据
	demo, total, err := svc.demoRepo.ListDemoPage(ctx, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	// 转换为dto（领域模型 -> 数据传输对象）
	var demoResp []*dto.DemoPageListResponse
	for _, v := range demo {
		demoResp = append(demoResp, &dto.DemoPageListResponse{
			ID:     v.ID,
			Field1: v.Field1,
			Field2: v.Field2,
		})
	}
	// 返回数据传输对象
	return demoResp, total, nil
}

// GetDemoByID 根据ID获取demo数据
func (svc *DemoServiceImpl) GetDemoByID(ctx context.Context, id int) (*dto.DemoDetailResponse, error) {
	// 调用数据访问层方法获取数据
	demo, err := svc.demoRepo.GetDemoByID(ctx, id)
	if err != nil {
		return nil, err
	}
	// 转换为dto（领域模型 -> 数据传输对象）
	demoResp := &dto.DemoDetailResponse{
		ID:     demo.ID,
		Field1: demo.Field1,
		Field2: demo.Field2,
	}
	// 返回数据传输对象
	return demoResp, nil
}

// CreateDemo 创建demo数据
func (svc *DemoServiceImpl) CreateDemo(ctx context.Context, demo *dto.DemoCreateRequest) (*dto.DemoCreateResponse, error) {
	// 此处可编写业务逻辑，如参数校验、转换等

	// 转换为数据模型（数据传输对象 -> 数据模型）
	demoModel := &model.Demo{
		Field1: demo.Field1,
		Field2: demo.Field2,
	}
	// 调用数据访问层方法创建数据
	id, err := svc.demoRepo.CreateDemo(ctx, demoModel)
	if err != nil {
		return nil, err
	}
	// 返回数据转换为dto（领域模型 -> 数据传输对象）
	resp := &dto.DemoCreateResponse{
		ID: id,
	}
	return resp, nil
}

// BatchCreateDemo 批量创建demo数据
func (svc *DemoServiceImpl) BatchCreateDemo(ctx context.Context, demos []*dto.DemoCreateRequest) error {
	// 转换为数据模型（数据传输对象 -> 数据模型）
	demoModels := make([]*model.Demo, 0, len(demos))
	for _, demo := range demos {
		// 1.业务逻辑

		// 2.转换为数据模型
		demoModels = append(demoModels, &model.Demo{
			Field1: demo.Field1,
			Field2: demo.Field2,
		})
	}
	// 调用数据访问层方法批量创建数据
	err := svc.demoRepo.BatchCreateDemo(ctx, demoModels)
	if err != nil {
		return err
	}
	return nil
}

// UpdateDemo 更新demo数据
func (svc *DemoServiceImpl) UpdateDemo(ctx context.Context, id int, req dto.DemoUpdateRequest) error {
	// 构建更新字段
	updateFields := make(map[string]interface{})
	if req.Field1 != nil {
		updateFields["field1"] = *req.Field1
	}
	if req.Field2 != nil {
		updateFields["field2"] = *req.Field2
	}
	// 调用数据访问层方法更新数据
	if len(updateFields) > 0 {
		err := svc.demoRepo.UpdateDemo(ctx, id, updateFields)
		if err != nil {
			return err
		}
	} else {
		return utils.NewBusinessError(utils.ErrCodeResourceNotFound, "无更新数据")
	}
	return nil
}

// SoftDeleteDemo 软删除demo数据
func (svc *DemoServiceImpl) SoftDeleteDemo(ctx context.Context, id int) error {
	// 检查数据是否存在
	demo, err := svc.demoRepo.GetDemoByID(ctx, id)
	if err != nil {
		return err
	}
	if demo == nil {
		return utils.NewBusinessError(utils.ErrCodeResourceNotFound, "数据不存在或已被删除，请刷新后重试")
	}

	// 构建更新数据，复用UpdateDemo方法
	updateFields := map[string]interface{}{
		"is_deleted": "Y",
	}

	// 调用数据访问层方法软删除数据
	err = svc.demoRepo.UpdateDemo(ctx, id, updateFields)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDemo 删除demo数据
func (svc *DemoServiceImpl) DeleteDemo(ctx context.Context, id int) error {
	// 检查数据是否存在
	demo, err := svc.demoRepo.GetDemoByID(ctx, id)
	if err != nil {
		return err
	}
	if demo == nil {
		return utils.NewBusinessError(utils.ErrCodeResourceNotFound, "数据不存在或已被删除，请刷新后重试")
	}

	// 调用数据访问层方法删除数据
	err = svc.demoRepo.DeleteDemo(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
