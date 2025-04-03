# 晖雄 AI

## 项目简介

晖雄 AI 是一个基于 Go 语言开发的智能对话系统，支持多种 AI 模型接入，提供友好的 Web 界面。

## 构建前端

cd HX/web/default
npm install
npm run build

## 构建后端

cd ../..
go mod download
go build -ldflags "-s -w" -o HX

## 目录结构

```text
HX/
├── cmd/                    # 命令行入口
├── config/                 # 配置文件
│   ├── config.go          # 配置核心
│   ├── db.go              # 数据库配置
│   ├── defaults.go        # 默认配置
│   ├── loader.go          # 配置加载器
│   ├── models.go          # 配置模型
│   ├── utils.go           # 工具函数
│   └── validator.go       # 配置验证器
├── docs/                   # 文档
├── internal/               # 内部代码
│   ├── app/               # 应用程序核心
│   │   └── app.go         # 应用初始化和启动
│   ├── controller/        # HTTP 控制器
│   │   └── http/         # HTTP 相关控制器
│   ├── entity/           # 实体定义
│   │   ├── message.go    # 消息实体
│   │   ├── option.go     # 选项实体
│   │   └── user.go       # 用户实体
│   ├── middleware/       # 中间件
│   │   ├── auth.go       # 认证中间件
│   │   ├── cors.go       # CORS 中间件
│   │   ├── language.go   # 语言中间件
│   │   ├── logger.go     # 日志中间件
│   │   └── requestid.go  # 请求ID中间件
│   ├── repo/             # 数据仓库层
│   │   ├── option.go     # 选项仓库
│   │   └── user.go       # 用户仓库
│   └── usecase/          # 用例层
│       ├── option.go     # 选项用例
│       └── user.go       # 用户用例
├── pkg/                   # 公共包
│   ├── db/               # 数据库工具
│   ├── i18n/             # 国际化
│   └── logger/           # 日志工具
├── web/                   # 前端代码
│   ├── air/              # Air 主题
│   ├── berry/            # Berry 主题
│   ├── build/            # 构建后的静态文件
│   ├── default/          # 默认主题
│   └── build.sh          # 构建脚本
├── .env.example          # 环境变量示例
├── go.mod                # Go 模块文件
├── go.sum                # Go 依赖版本锁定
├── main.go               # 主程序入口
└── hx.db                 # SQLite 数据库文件
```

## 主要功能

- 多 AI 模型支持
- 用户认证和授权
- 配额管理
- 多语言支持
- 主题切换
- 实时对话
- 消息历史记录
- 系统监控

## 技术栈

- 后端：Go + Gin + GORM
- 前端：React + TypeScript
- 数据库：SQLite
- 缓存：内存缓存
- 国际化：i18n
- 日志：自定义日志系统

## 配置说明

项目支持多种配置方式：

- 配置文件（config.yaml）
- 环境变量
- 命令行参数
- 默认值

## 开发说明

1. 克隆项目
2. 安装依赖
3. 配置环境变量
4. 构建前端
5. 构建后端
6. 运行项目

## 许可证

MIT License
