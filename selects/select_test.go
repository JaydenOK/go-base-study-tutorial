package selects

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

//for select 阻塞，监听信号，重启任务或查询任务信息
func TestA(t *testing.T) {
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	for {
		select {
		case <-exitChan:
			checkStatus()
		}
	}
}

func checkStatus() {
	fmt.Println("aaa")
}
