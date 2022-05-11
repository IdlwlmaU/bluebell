package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"bluebell/settings"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// @title bluebell项目接口文档
// @version 1.0
// @description bluebell社区
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8888
// @BasePath /api/v1
func main() {
	// 1.加载配置
	err := settings.Init()
	if err != nil {
		fmt.Println("init settings failed, err:", err)
		return
	}
	// 2.初始化日志
	err = logger.Init(settings.Conf.LogConfig, settings.Conf.Mode)
	if err != nil {
		fmt.Println("init logger failed, err:", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	// 3.初始化MySQL连接
	err = mysql.Init(settings.Conf.MySQLConfig)
	if err != nil {
		fmt.Println("init mysql failed, err:", err)
		return
	}
	defer mysql.Close()
	// 4.初始化Redis连接
	err = redis.Init(settings.Conf.RedisConfig)
	if err != nil {
		fmt.Println("init redis failed, err:", err)
		return
	}
	defer redis.Close()
	// snowflake
	err = snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID)
	if err != nil {
		fmt.Println("init snowflake failed, err:", err)
		return
	}
	// 初始化gin框架内置的校验器使用的翻译器
	if err = controller.InitTrans("zh"); err != nil {
		fmt.Println("init trans failed, err:", err)
		return
	}
	// 5.注册路由
	r := router.Setup()
	// 6.启动服务额（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}
	fmt.Printf("[GIN-debug] Listening and serving HTTP on :%d\n", settings.Conf.Port)

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Info("listen: %s\n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
