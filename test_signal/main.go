package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal)

	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	//signal.Stop(sigs) // 停止接受信号后 exit... 无法输出
	fmt.Println("signal wait....")
	<-done
	fmt.Println("exit...")
}
