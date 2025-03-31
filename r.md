# 构建前端
cd HX/web/default
npm install
npm run build

# 构建后端
cd ../..
go mod download
go build -ldflags "-s -w" -o HX