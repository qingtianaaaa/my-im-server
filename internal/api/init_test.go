package api

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestSignalHandling(t *testing.T) {
	// 创建信号通道
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGCONT)

	// 在goroutine中等待信号
	sigReceived := make(chan bool)
	go func() {
		<-sigs
		sigReceived <- true
	}()

	// 发送SIGTERM信号
	process, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("无法找到当前进程: %v", err)
	}
	t.Log(process)
	t.Log("------")
	err = process.Signal(syscall.SIGCONT) 
	if err != nil {
		t.Fatalf("发送信号失败: %v", err)
	}

	// 等待信号处理
	select {
	case <-sigReceived:
		// 成功接收到信号
	case <-time.After(2 * time.Second):
		t.Log("超时：未收到SIGTERM信号")
	}
}
