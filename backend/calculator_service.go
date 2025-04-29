package main

import (
	"context"

	calculatorv1 "calculator/gen/proto"

	"github.com/bufbuild/connect-go"
)

// CalculatorServer 实现计算器服务接口
type CalculatorServer struct{}

// Add 实现加法计算
func (s *CalculatorServer) Add(ctx context.Context, req *connect.Request[calculatorv1.AddRequest]) (*connect.Response[calculatorv1.CalculateResponse], error) {
	return &connect.Response[calculatorv1.CalculateResponse]{
		Msg: &calculatorv1.CalculateResponse{
			Result: req.Msg.A + req.Msg.B,
		},
	}, nil
}

// Subtract 实现减法计算
func (s *CalculatorServer) Subtract(ctx context.Context, req *connect.Request[calculatorv1.SubtractRequest]) (*connect.Response[calculatorv1.CalculateResponse], error) {
	return &connect.Response[calculatorv1.CalculateResponse]{
		Msg: &calculatorv1.CalculateResponse{
			Result: req.Msg.A - req.Msg.B,
		},
	}, nil
}

// Multiply 实现乘法计算
func (s *CalculatorServer) Multiply(ctx context.Context, req *connect.Request[calculatorv1.MultiplyRequest]) (*connect.Response[calculatorv1.CalculateResponse], error) {
	return &connect.Response[calculatorv1.CalculateResponse]{
		Msg: &calculatorv1.CalculateResponse{
			Result: req.Msg.A * req.Msg.B,
		},
	}, nil
}

// Divide 实现除法计算
func (s *CalculatorServer) Divide(ctx context.Context, req *connect.Request[calculatorv1.DivideRequest]) (*connect.Response[calculatorv1.CalculateResponse], error) {
	if req.Msg.B == 0 {
		return &connect.Response[calculatorv1.CalculateResponse]{
			Msg: &calculatorv1.CalculateResponse{
				Error: "除数不能为零",
			},
		}, nil
	}
	return &connect.Response[calculatorv1.CalculateResponse]{
		Msg: &calculatorv1.CalculateResponse{
			Result: req.Msg.A / req.Msg.B,
		},
	}, nil
}
