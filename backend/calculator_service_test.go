package main

import (
	"context"
	"testing"

	calculatorv1 "calculator/gen/proto"

	"github.com/bufbuild/connect-go"
)

func TestCalculatorServer_Add(t *testing.T) {
	server := &CalculatorServer{}

	testCases := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"正数加法", 5, 3, 8},
		{"负数加法", -5, -3, -8},
		{"零加法", 0, 0, 0},
		{"小数加法", 1.5, 2.5, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := connect.NewRequest(&calculatorv1.AddRequest{
				A: tc.a,
				B: tc.b,
			})

			resp, err := server.Add(context.Background(), req)

			if err != nil {
				t.Fatalf("加法计算出错: %v", err)
			}

			if resp.Msg.Error != "" {
				t.Fatalf("加法返回错误信息: %s", resp.Msg.Error)
			}

			if resp.Msg.Result != tc.expected {
				t.Errorf("加法计算错误, 期望: %f, 实际: %f", tc.expected, resp.Msg.Result)
			}
		})
	}
}

func TestCalculatorServer_Subtract(t *testing.T) {
	server := &CalculatorServer{}

	testCases := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"正数减法", 5, 3, 2},
		{"负数减法", -5, -3, -2},
		{"零减法", 0, 0, 0},
		{"小数减法", 4.5, 2.5, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := connect.NewRequest(&calculatorv1.SubtractRequest{
				A: tc.a,
				B: tc.b,
			})

			resp, err := server.Subtract(context.Background(), req)

			if err != nil {
				t.Fatalf("减法计算出错: %v", err)
			}

			if resp.Msg.Error != "" {
				t.Fatalf("减法返回错误信息: %s", resp.Msg.Error)
			}

			if resp.Msg.Result != tc.expected {
				t.Errorf("减法计算错误, 期望: %f, 实际: %f", tc.expected, resp.Msg.Result)
			}
		})
	}
}

func TestCalculatorServer_Multiply(t *testing.T) {
	server := &CalculatorServer{}

	testCases := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"正数乘法", 5, 3, 15},
		{"负数乘法", -5, -3, 15},
		{"零乘法", 0, 5, 0},
		{"小数乘法", 1.5, 2, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := connect.NewRequest(&calculatorv1.MultiplyRequest{
				A: tc.a,
				B: tc.b,
			})

			resp, err := server.Multiply(context.Background(), req)

			if err != nil {
				t.Fatalf("乘法计算出错: %v", err)
			}

			if resp.Msg.Error != "" {
				t.Fatalf("乘法返回错误信息: %s", resp.Msg.Error)
			}

			if resp.Msg.Result != tc.expected {
				t.Errorf("乘法计算错误, 期望: %f, 实际: %f", tc.expected, resp.Msg.Result)
			}
		})
	}
}

func TestCalculatorServer_Divide(t *testing.T) {
	server := &CalculatorServer{}

	testCases := []struct {
		name        string
		a           float64
		b           float64
		expected    float64
		expectError bool
		errorMsg    string
	}{
		{"正数除法", 6, 3, 2, false, ""},
		{"负数除法", -6, -3, 2, false, ""},
		{"小数除法", 5, 2, 2.5, false, ""},
		{"除以零", 5, 0, 0, true, "除数不能为零"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := connect.NewRequest(&calculatorv1.DivideRequest{
				A: tc.a,
				B: tc.b,
			})

			resp, err := server.Divide(context.Background(), req)

			if err != nil {
				t.Fatalf("除法计算出错: %v", err)
			}

			if tc.expectError {
				if resp.Msg.Error != tc.errorMsg {
					t.Errorf("除法错误消息不匹配, 期望: %s, 实际: %s", tc.errorMsg, resp.Msg.Error)
				}
			} else {
				if resp.Msg.Error != "" {
					t.Fatalf("除法返回错误信息: %s", resp.Msg.Error)
				}

				if resp.Msg.Result != tc.expected {
					t.Errorf("除法计算错误, 期望: %f, 实际: %f", tc.expected, resp.Msg.Result)
				}
			}
		})
	}
}
