package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"test/models"
	"time"
)

var lg *zap.Logger

// Init 自己初始化的日志
func Init(logconf *models.Logconf) (err error){
	encoder := getEncoder()
	writersyncer := getsyncer(
		//viper.GetInt("max_size"),
		//viper.GetInt("max_backups"),
		//viper.GetInt("max_age"),

		logconf.Maxsize,
		logconf.Maxbackups,
		logconf.Maxage,
		)

	var l = new(zapcore.Level)
	//通过反序列化把里面的东西拿出来
	err = l.UnmarshalText([]byte(viper.GetString("log.level")))

	if err != nil {
		return errors.Wrap(err,"zap init err")
	}

	core := zapcore.NewCore(encoder,writersyncer,l)
	lg = zap.New(core,zap.AddCaller())

	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(lg)
	return
}

func getEncoder()zapcore.Encoder{
	encoder := zap.NewProductionEncoderConfig()
	//化成可看的时间
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	//化为全大写字符串
	encoder.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoder)
}

func getsyncer(maxSize, maxBackup, maxAge int)zapcore.WriteSyncer{
	lumberjack := &lumberjack.Logger{
		Filename: "./test.log",
		MaxSize: maxSize,
		MaxAge: maxAge,
		MaxBackups: maxBackup,
		Compress:   false,
	}
	return zapcore.AddSync(lumberjack)
}


// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}


