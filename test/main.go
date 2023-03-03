package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	b := make(chan int, 2)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for i := 0; i <= 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("获取到关闭信息 <-ctx.Done()")
				return

			case b <- i:
				fmt.Println(fmt.Sprintf("输入进去%d", i))
				time.Sleep(time.Second * 1)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("获取信息也能拿到关闭状态 <-ctx.Done()")
				return

			case res, ok := <-b:
				fmt.Println(fmt.Sprintf("获取到信息%d:%+v", res, ok))
				time.Sleep(time.Second * 1)

			}

		}
	}()

	time.Sleep(time.Second * 3)

	//c <- 1
	cancel()
	close(b)
	time.Sleep(time.Second * 1000)

}
