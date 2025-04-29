# 计算器后端服务

这是计算器应用的后端服务，基于Go和ConnectRPC实现。

## 功能

- 提供基本的数学运算API：加法、减法、乘法、除法
- 使用Connect协议进行通信
- 支持跨域资源共享(CORS)

## 技术栈

- Go 1.24.x
- buf/connect-go
- Protocol Buffers

## 目录结构

- `proto/`: 协议定义文件
- `gen/`: 生成的代码
- `calculator_service.go`: 实现计算器业务逻辑
- `connect_service.go`: 设置Connect服务
- `main.go`: 应用程序入口

## 如何运行

1. 确保安装了Go 1.24或更高版本
2. 运行服务器:

```bash
go run main.go
```

或者，你可以构建并运行二进制文件:

```bash
go build -o calculator-server
./calculator-server
```

服务器将在 http://localhost:8080 上运行。

## API端点

- POST `/calculator.v1.CalculatorService/Add` - 执行加法运算
- POST `/calculator.v1.CalculatorService/Subtract` - 执行减法运算
- POST `/calculator.v1.CalculatorService/Multiply` - 执行乘法运算
- POST `/calculator.v1.CalculatorService/Divide` - 执行除法运算

## 示例请求

### 加法

```bash
curl -X POST \
  http://localhost:8080/calculator.v1.CalculatorService/Add \
  -H 'Content-Type: application/json' \
  -d '{"a": 5, "b": 3}'
```

响应:

```json
{
  "result": 8
}
```

## 如何修改协议

1. 编辑 `proto/calculator.proto` 文件
2. 使用buf生成新的代码:

```bash
buf generate
```

## 测试

运行单元测试:

```bash
go test ./...
``` 