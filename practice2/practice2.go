package practice2

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

/*
日期：
2021/11/7

问题：
1.基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
*/

// StartHttpServer 启动 HTTP server
func StartHttpServer(srv *http.Server) error {
	http.HandleFunc("/hello", HelloServer2)
	fmt.Println("http server start")
	err := srv.ListenAndServe()
	return err
}

// HelloServer2 增加一个 HTTP hanlder
func HelloServer2(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctx)
	srv := &http.Server{Addr: ":9090"}
	group.Go(func() error {
		return StartHttpServer(srv)
	})
	group.Go(func() error {
		<-errCtx.Done()
		fmt.Println("http server stop")
		return srv.Shutdown(errCtx)
	})
	chanel := make(chan os.Signal, 1)
	signal.Notify(chanel)
	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-chanel:
				cancel()
			}
		}
		return nil
	})
	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")
}