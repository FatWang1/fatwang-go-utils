package utils

import "testing"

func TestInterfaces(t *testing.T) {
	// 这些是接口定义，主要用于类型约束
	// 我们可以通过确保有类型实现这些接口来测试它们
	
	// 测试InfoLogger接口
	var _ InfoLogger = (*testInfoLogger)(nil)
	
	// NUMBER接口是类型约束，不能直接测试
	// 但我们已经在其他函数中使用它了，比如Max和Min函数
}

// 实现InfoLogger接口用于测试
type testInfoLogger struct{}

func (l *testInfoLogger) Info(msg string) {
	// 测试实现，不需要实际做任何事情
}