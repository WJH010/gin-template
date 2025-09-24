package repository

import (
	"context"
	"errors"
	"fmt"
	"gin-template/internal/demo/dto"
	"gin-template/internal/demo/model"
	"gin-template/internal/utils"

	"gorm.io/gorm"
)

// DemoRepository 数据访问接口，定义数据访问的方法集。规范 Demo 数据访问的方法。定义“做什么”
type DemoRepository interface {
	// ListDemo 获取demo数据
	ListDemo(ctx context.Context, req *dto.DemoListRequest) ([]*model.Demo, error)
	// ListDemoPage 分页查询demo数据
	ListDemoPage(ctx context.Context, page, pageSize int) ([]*model.Demo, int64, error)
	// GetDemoByID 根据ID获取demo数据
	GetDemoByID(ctx context.Context, id int) (*model.Demo, error)
	// CreateDemo 创建demo数据
	CreateDemo(ctx context.Context, demo *model.Demo) (int, error)
	// BatchCreateDemo 批量创建demo数据
	BatchCreateDemo(ctx context.Context, demos []*model.Demo) error
	// UpdateDemo 更新demo数据
	UpdateDemo(ctx context.Context, id int, updateFields map[string]interface{}) error
	// DeleteDemo 删除demo数据
	DeleteDemo(ctx context.Context, id int) error
}

// DemoRepositoryImpl 实现接口的具体结构体。定义“怎么做”
type DemoRepositoryImpl struct {
	db *gorm.DB
}

// NewDemoRepository 创建数据访问实例。用于创建DemoRepository接口的实例，接收一个*gorm.DB（数据库连接）参数，注入到DemoRepositoryImpl结构体中。
func NewDemoRepository(db *gorm.DB) DemoRepository {
	return &DemoRepositoryImpl{db: db}
}

// ListDemo 获取demo数据
func (repo *DemoRepositoryImpl) ListDemo(ctx context.Context, req *dto.DemoListRequest) ([]*model.Demo, error) {
	// 声明一个Demo指针切片，用于存储查询结果
	var demo []*model.Demo

	// 使用GORM进行查询，WithContext将上下文与数据库操作关联，支持超时和取消
	query := repo.db.WithContext(ctx)
	// 拼接查询条件
	if req.Field1 != 0 {
		query = query.Where("field1 = ?", req.Field1)
	}
	if req.Field2 != "" {
		query = query.Where("field2 = ?", req.Field2)
	}
	// Find方法将查询结果存储到demo指针切片中
	result := query.Find(&demo)
	err := result.Error

	// 异常处理
	if err != nil {
		return nil, utils.NewSystemError(fmt.Errorf("数据库查询失败: %v", err))
	}

	return demo, nil
}

// ListDemoPage 分页查询demo数据
func (repo *DemoRepositoryImpl) ListDemoPage(ctx context.Context, page, pageSize int) ([]*model.Demo, int64, error) {

	offset := (page - 1) * pageSize // 计算偏移量

	// 声明一个Demo指针切片，用于存储查询结果
	var demo []*model.Demo

	// 构建基础查询
	query := repo.db.WithContext(ctx).Model(&model.Demo{})

	// 计算总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, utils.NewSystemError(fmt.Errorf("计算总数时数据库查询失败: %v", err))
	}

	// 查询数据
	if err := query.Offset(offset).Limit(pageSize).Find(&demo).Error; err != nil {
		return nil, 0, utils.NewSystemError(fmt.Errorf("数据库查询失败: %v", err))
	}

	return demo, total, nil
}

// GetDemoByID 根据ID获取demo数据
func (repo *DemoRepositoryImpl) GetDemoByID(ctx context.Context, id int) (*model.Demo, error) {
	// 声明一个Demo指针，用于存储查询结果
	var demo *model.Demo

	// 查询数据
	// First 方法：查询符合条件的第一条记录（默认按主键升序排序）
	// 若要按其他字段查询第一条记录，写法：repo.db.First(&demo, "field1 = ?", value)
	result := repo.db.WithContext(ctx).First(&demo, id)
	err := result.Error

	// 异常处理
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewBusinessError(utils.ErrCodeResourceNotFound, "demo数据不存在")
		}
		return nil, utils.NewSystemError(fmt.Errorf("数据库查询失败: %w", err))
	}

	return demo, nil
}

// CreateDemo 创建demo数据
func (repo *DemoRepositoryImpl) CreateDemo(ctx context.Context, demo *model.Demo) (int, error) {
	// 插入数据
	result := repo.db.WithContext(ctx).Create(demo)
	err := result.Error

	// 异常处理
	if err != nil {
		// 检查是否是重复键错误（需在数据库层配置唯一约束）
		if exist, fieldName, value := utils.IsUniqueConstraintError(err); exist {
			if fieldName == "field1" {
				return 0, utils.NewBusinessError(utils.ErrCodeDuplicateKey, fmt.Sprintf("字段一('%s')已存在，不能重复创建", value))
			}
		}
		return 0, utils.NewSystemError(fmt.Errorf("数据库插入失败: %w", err))
	}

	return demo.ID, nil
}

// BatchCreateDemo 批量创建demo数据
func (repo *DemoRepositoryImpl) BatchCreateDemo(ctx context.Context, demos []*model.Demo) error {
	// 批量插入数据
	result := repo.db.WithContext(ctx).Create(demos)
	err := result.Error

	// 异常处理
	if err != nil {
		// 检查是否是重复键错误（需在数据库层配置唯一约束）
		if exist, fieldName, value := utils.IsUniqueConstraintError(err); exist {
			if fieldName == "field1" {
				return utils.NewBusinessError(utils.ErrCodeDuplicateKey, fmt.Sprintf("字段一('%s')已存在，不能重复创建", value))
			}
		}
		return utils.NewSystemError(fmt.Errorf("数据库批量插入失败: %w", err))
	}

	return nil
}

// UpdateDemo 更新demo数据
func (repo *DemoRepositoryImpl) UpdateDemo(ctx context.Context, id int, updateFields map[string]interface{}) error {
	// 更新数据
	result := repo.db.WithContext(ctx).
		Model(&model.Demo{}).
		Where("id = ?", id).
		Updates(updateFields)
	err := result.Error

	// 异常处理
	if err != nil {
		return utils.NewSystemError(fmt.Errorf("更新数据失败: %w", err))
	}
	if result.RowsAffected == 0 {
		return utils.NewBusinessError(utils.ErrCodeResourceNotFound, "数据不存在或已被删除，请刷新页面后重试")
	}

	return nil
}

// DeleteDemo 删除demo数据
func (repo *DemoRepositoryImpl) DeleteDemo(ctx context.Context, id int) error {
	// 删除数据
	result := repo.db.WithContext(ctx).
		Model(&model.Demo{}).
		Where("id = ?", id).
		Delete(&model.Demo{})
	err := result.Error

	// 异常处理
	if err != nil {
		return utils.NewSystemError(fmt.Errorf("删除数据失败: %w", err))
	}
	if result.RowsAffected == 0 {
		return utils.NewBusinessError(utils.ErrCodeResourceNotFound, "数据不存在或已被删除，请刷新页面后重试")
	}

	return nil
}
