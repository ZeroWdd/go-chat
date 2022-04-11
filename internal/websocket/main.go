package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"go-chat/internal/pkg/im"
	"go-chat/internal/provider"
	"golang.org/x/sync/errgroup"
)

func main() {
	cmd := cli.NewApp()

	cmd.Name = "Websocket Log"
	cmd.Usage = "GoChat 即时聊天应用"

	// 设置参数
	cmd.Flags = []cli.Flag{
		// 配置文件参数
		&cli.StringFlag{Name: "config", Aliases: []string{"c"}, Value: "./config.yaml", Usage: "配置文件路径", DefaultText: "./config.yaml"},

		&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 9504, Usage: "设置端口号", DefaultText: "9504"},
	}

	cmd.Action = func(tx *cli.Context) error {
		ctx, cancel := context.WithCancel(tx.Context)
		defer cancel()

		eg, groupCtx := errgroup.WithContext(ctx)

		// 初始化 IM 渠道配置
		im.Initialize(groupCtx, eg)

		// 读取配置文件
		config := provider.ReadConfig(tx.String("config"))

		// 设置服务端口号
		config.SetPort(tx.Int("port"))

		if !config.Debug() {
			writer, _ := os.OpenFile(fmt.Sprintf("%s/logs/ws-access.log", config.Log.Dir), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

			gin.DefaultWriter = writer
			gin.SetMode(gin.ReleaseMode)
		}

		providers := Initialize(ctx, config)

		c := make(chan os.Signal, 1)

		signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

		// 延时启动守护协程
		time.AfterFunc(3*time.Second, func() {
			providers.Process.Start(eg, groupCtx)
		})

		log.Printf("Websocket Server ID   :%s", config.ServerId())
		log.Printf("Websocket Listen Port :%d", config.App.Port)
		log.Printf("Websocket Server Pid  :%d", os.Getpid())

		return start(c, eg, groupCtx, cancel, providers.WsServer)
	}

	_ = cmd.Run(os.Args)
}

func start(c chan os.Signal, eg *errgroup.Group, ctx context.Context, cancel context.CancelFunc, server *http.Server) error {
	// 启动 http 服务
	eg.Go(func() error {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Websocket Listen Err: %s", err)
		}

		return nil
	})

	eg.Go(func() error {
		defer func() {
			cancel()

			// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
			timeCtx, timeCancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer timeCancel()

			if err := server.Shutdown(timeCtx); err != nil {
				log.Printf("Websocket Shutdown Err: %s \n", err)
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c:
			return nil
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("eg error: %s", err)
		return err
	}

	fmt.Println()
	log.Fatal("Websocket Shutdown")

	return nil
}
