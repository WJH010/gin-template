# gin-study

一个基于 Gin 框架和 GORM 的分层架构示例项目，旨在为初学者提供清晰的后端开发参考，展示如何规范地实现 CRUD、分页查询、批量操作等常见业务场景。目前还在完善中，后续会添加更多的业务示例和功能。

## 项目特点

- **清晰的分层架构**：严格遵循 Controller-Service-Repository 分层模式，职责边界明确
- **完整的业务示例**：包含基础 CRUD、分页查询、批量操作、软删除等常见功能实现
- **规范的代码风格**：统一的错误处理、响应格式和日志记录，符合生产级开发标准
- **灵活的配置管理**：支持 YAML 配置文件和环境变量注入，方便多环境部署
- **主流技术栈**：基于 Gin + GORM 构建，集成日志、配置解析等常用工具库

## 技术栈

- 框架：Gin v1.10.1
- ORM：GORM v1.30.0
- 数据库驱动：MySQL
- 配置解析：gopkg.in/yaml.v3
- 日志：logrus v1.9.3
- 开发语言：Go 1.24.3

## 项目结构

```
gin-study/
├── cmd/
│   └── main.go               # 程序入口
├── config.yaml               # 配置文件
├── internal/
│   ├── app/
│   │   ├── config/           # 配置相关
│   │   ├── database/         # 数据库连接
│   │   ├── middleware/       # 中间件
│   │   └── routes/           # 路由注册
│   ├── demo/                 # 示例模块
│   │   ├── controller/       # 控制器层（处理HTTP请求）
│   │   ├── service/          # 服务层（业务逻辑）
│   │   ├── repository/       # 数据访问层
│   │   ├── model/            # 数据模型
│   │   └── dto/              # 数据传输对象
│   └── utils/                # 工具函数
├── logs/                     # 日志文件（git忽略）
└── test/                     # 测试相关
```

## 快速开始

### 前置条件

- Go 1.24+ 环境
- MySQL 数据库

### 步骤

1. 克隆仓库

```bash
git clone https://github.com/your-username/gin-study.git
cd gin-study
```

2. 配置数据库

修改 `config.yaml` 文件中的数据库配置：

```yaml
database:
  driver: mysql
  host: localhost
  port: 3306
  username: root
  password: your-password
  dbname: your-dbname
```

也可通过环境变量覆盖配置（如 `DB_HOST`、`DB_PORT` 等）

3. 安装依赖

```bash
go mod tidy
```

4. 启动服务

```bash
go run cmd/main.go
```

服务将在 `8080` 端口启动，可通过 `http://localhost:8080/api/test` 验证是否启动成功

## API 示例

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/api/demo` | 获取所有数据 |
| GET | `/api/demo/page` | 分页查询数据 |
| GET | `/api/demo/:id` | 根据ID获取详情 |
| POST | `/api/demo` | 创建数据 |
| POST | `/api/demo/batch` | 批量创建数据 |
| PUT | `/api/demo/:id` | 更新数据 |
| DELETE | `/api/demo/soft/:id` | 软删除数据 |
| DELETE | `/api/demo/hard/:id` | 物理删除数据 |

## 核心设计说明

### 分层架构

1. **Controller 层**：处理 HTTP 请求，负责参数绑定、调用服务层、统一响应格式
2. **Service 层**：实现核心业务逻辑，处理 DTO 与 Model 的转换
3. **Repository 层**：封装数据库操作，隔离数据访问细节
4. **Model 层**：定义数据模型，映射数据库表结构
5. **DTO 层**：定义数据传输对象，区分 API 输入输出与内部模型

### 错误处理

- 自定义 `BusinessError`（业务错误）和 `SystemError`（系统错误）
- 通过 `utils.HandlerFunc` 统一处理并返回标准化错误响应

### 日志系统

- 基于 logrus 实现，支持不同级别日志染色输出
- 包含时间戳、日志级别、请求信息等关键上下文

## 许可证

[MIT](LICENSE)