package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 创建一个通道，用于传送信号
	stop := make(chan os.Signal, 1)

	// 将SIGINT信号转发到通道stop
	signal.Notify(stop, os.Interrupt)

	// 创建一个带上下文的group
	group, ctx := errgroup.WithContext(context.Background())

	// 若收到外部终止信号，则优雅地终止所有服务
	group.Go(func() error {
		return listenSignal(ctx, stop)
	})

	// 启动Web服务
	group.Go(func() error {
		return startServer(ctx, "A", ":8080")
	})
	group.Go(func() error {
		//return startServer(ctx, "B" ,":8080") // 故障模拟
		return startServer(ctx, "B", ":8081")
	})

	// 等待所有通过Go方法调用的函数返回
	if err := group.Wait(); err != nil {
		// 打印导致服务终止的错误
		log.Printf("all service shutdown because: %v", err)
	}
}

func listenSignal(ctx context.Context, stop chan os.Signal) error {
	select {
	case s := <-stop:
		// 接收到外部终止信号
		log.Printf("stopped with signal: %v", s)
		return errors.New("interrupted manually")
	case <-ctx.Done():
		// 同组的其他服务出现错误导致context取消
		return nil
	}
}

func startServer(ctx context.Context, name string, addr string) error {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello World From Server " + name))
	})

	server := &http.Server{Addr: addr, Handler: handler}
	done := make(chan error)
	go func() {
		log.Println("starting server " + name + " at " + addr)

		// 故障模拟，A服务延迟启动
		//if name == "A" {
		//	time.Sleep(20 * time.Second)
		//}
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			done <- fmt.Errorf("server %s start err: %v", name, err)
		}
	}()

	select {
	case err := <-done:
		// 启动出错
		return err

	case <-ctx.Done():
		// 同组的其他服务出现错误导致context取消
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel() // 若在倒计时前服务关闭，则提前释放资源
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("server "+name+" shutdown err: %v", err)
			return err
		}
		log.Println("shutdown server " + name + " gracefully!")
		return nil
	}
}
