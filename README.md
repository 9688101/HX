# 构建前端
cd HX/web/default
npm install
npm run build

# 构建后端
cd ../..
go mod download
go build -ldflags "-s -w" -o HX


HX/
├── config/                 # 配置文件
├── docs/                   # 文档
├── internal/               # 内部代码
│   ├── app/               # 应用程序核心
│   ├── controller/        # HTTP 控制器
│   ├── entity/           # 实体定义
│   ├── middleware/       # 中间件
│   ├── repo/             # 数据仓库层
│   └── usecase/          # 用例层
├── pkg/                   # 公共包
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