// Code generated by hertz generator.

package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	"github.com/krismil/biz/gomall/app/frontend/infra/rpc"
	"github.com/krismil/biz/gomall/app/frontend/middleware"
	frontendUtils "github.com/krismil/biz/gomall/app/frontend/utils"
	"time"

	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/pprof"
	"github.com/krismil/biz/gomall/app/frontend/biz/router"
	"github.com/krismil/biz/gomall/app/frontend/conf"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	_ = godotenv.Load()
	// init dal
	// dal.Init()

	rpc.Init()
	address := conf.GetConf().Hertz.Address
	h := server.New(server.WithHostPorts(address))

	registerMiddleware(h)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h)
	h.LoadHTMLGlob("template/*")
	h.Static("/static", "./")

	h.GET("/sign-in", func(c context.Context, ctx *app.RequestContext) {
		date := utils.H{
			"Title": "Sign In",
			"Next":  ctx.Query("next"),
		}
		ctx.HTML(consts.StatusOK, "sign-in", date)
	})
	h.GET("/sign-up", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "sign-up", utils.H{"Title": "Sign Up"})
	})
	h.GET("/about", middleware.Auth(), func(c context.Context, ctx *app.RequestContext) {
		userID := frontendUtils.GetUserIdFromCtx(c)
		ctx.HTML(consts.StatusOK, "about", utils.H{
			"Title":   "about",
			"user_id": userID,
		})
	})
	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	store, err := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	if err != nil {
		hlog.Fatal("Failed to initialize Redis store: ", err)
	}
	h.Use(sessions.New("biz", store))
	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		err := asyncWriter.Sync()
		if err != nil {
			return
		}
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())

	middleware.Register(h)
}
