package stack

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"testing"
)

func TestCallers(t *testing.T) {
	// 测试直接调用
	result := Callers(2)
	if len(result) == 0 {
		t.Error("Expected non-empty call stack")
	}

	// 打印调用栈信息
	t.Log("直接调用的调用栈信息:")
	printStack(t, result)

	// 测试嵌套调用
	nestedCall(t)
}

func nestedCall(t *testing.T) {
	// 在嵌套函数中获取调用栈
	result := Callers(2)
	if len(result) <= 1 {
		t.Error("Expected nested call stack")
	}

	t.Log("嵌套调用的调用栈信息:")
	printStack(t, result)
}

func printStack(t *testing.T, pcs []uintptr) {
	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()
		t.Logf("函数: %s\n文件: %s:%d", frame.Function, frame.File, frame.Line)
		t.Log("==")
		t.Log(path.Dir(frame.Func.Name()))
		t.Log("------")
		if !more {
			break
		}
	}
}

func TestPath(t *testing.T) {
	var sb strings.Builder
	err := errors.New("test error")
	sb.WriteString("Error: ")
	sb.WriteString(err.Error())
	fmt.Println(sb.String())
}
