# 分布式计算器项目

这是一个基于Go和Next.js构建的分布式计算器应用，使用Connect协议进行前后端通信。

## 项目结构

- `backend/`: Go语言后端服务
  - `proto/`: 协议定义文件
  - `gen/`: 生成的代码
  - `calculator_service.go`: 计算器业务逻辑
  - `connect_service.go`: Connect服务配置
  - `main.go`: 主程序入口

- `frontend/calculator-client/`: Next.js前端应用
  - `pages/`: Next.js页面
  - `src/`: 源代码目录
    - `gen/`: 生成的协议代码
    - `client.ts`: Connect客户端配置

## 技术栈

- **后端**：
  - Go 1.24.x
  - ConnectRPC
  - Protocol Buffers

- **前端**：
  - Next.js 13.x
  - React 18.x
  - TypeScript
  - Tailwind CSS
  - ConnectRPC Web

## 如何启动

### 环境要求

- Go 1.24.x 或更高版本
- Node.js 16.x 或更高版本 (推荐 18.x+)
- npm 或 yarn

### 后端服务

1. 打开终端，进入后端目录：

```bash
cd backend
```

2. 启动Go服务器：

```bash
go run main.go
# 或者使用已编译的二进制文件
./calculator.exe
```

后端服务将在 http://localhost:8080 上运行。

### 前端应用

1. 打开**新的**终端窗口，进入前端目录：

```bash
cd frontend/calculator-client
```

2. 安装依赖：

```bash
npm install
```

3. 启动开发服务器：

```bash
npm run dev
```

前端应用将在 http://localhost:3000 上运行。

## 如何使用

1. 在浏览器中打开 http://localhost:3000
2. 在第一个输入框中输入第一个数字
3. 选择要执行的运算（加法、减法、乘法或除法）
4. 在第二个输入框中输入第二个数字
5. 点击"计算"按钮
6. 查看计算结果

## 特性

- 完全类型安全的API通信
- 使用Connect协议进行前后端通信
- 优雅的错误处理
- 响应式UI设计

## 运行测试

### 后端测试

```bash
cd backend
go test ./...
```

### 前端测试

```bash
cd frontend/calculator-client
npm test
```

## 开发指南

### 修改协议定义

如果需要修改协议定义：

1. 编辑 `backend/proto/calculator.proto` 文件
2. 在后端目录中生成新的代码：
   ```bash
   cd backend
   buf generate
   ```
3. 将更新后的proto文件复制到前端：
   ```bash
   cp backend/proto/calculator.proto frontend/calculator-client/proto/
   ```
4. 在前端目录中生成新的客户端代码：
   ```bash
   cd frontend/calculator-client
   npx buf generate
   ```
5. 如果生成的代码中有`.js`后缀的导入路径，将它们更改为不带后缀的导入路径

## 部署指南

### 后端部署

1. 构建Go二进制文件：
   ```bash
   cd backend
   go build -o calculator-server
   ```

2. 运行编译好的二进制文件：
   ```bash
   ./calculator-server
   ```

### 前端部署

1. 构建前端应用：
   ```bash
   cd frontend/calculator-client
   npm run build
   ```

2. 部署生成的`out`目录到静态Web服务器或CDN

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建Pull Request

## 许可证

本项目基于 MIT 许可证 - 详情请查看 LICENSE 文件。

## 单元测试

### 后端测试

后端使用Go的内置测试框架进行单元测试。测试文件为`backend/calculator_service_test.go`。

运行后端测试：

```bash
cd backend
go test -v
```

### 前端测试

前端使用Jest和React Testing Library进行单元测试。测试文件位于`frontend/calculator-client/__tests__/`目录中。

运行前端测试：

```bash
cd frontend/calculator-client
npm test
``` 