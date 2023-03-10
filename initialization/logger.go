package initialization

import (
	"github.com/Madou-Shinni/go-logger"
	"github.com/Madou-Shinni/lottery-draw/internal/conf"
)

// 自定义日志初始化配置
func init() {
	var err error

	env := conf.Conf.Env
	file := "/log/github.com/Madou-Shinni/lottery-draw.log"

	if env == "dev" {
		// 开发环境
		_, err = logger.NewJSONLogger(
			// 日志等级
			logger.WithDebugLevel(),
			// 时间格式化
			logger.WithTimeLayout("2006-01-02 15:04:05"),
		)
	} else if env == "prod" {
		// 生产环境
		_, err = logger.NewJSONLogger(
			// 日志等级
			logger.WithDebugLevel(),
			// 写出的文件
			logger.WithFileRotationP(file),
			// 不在控制台打印
			logger.WithDisableConsole(),
			// 时间格式化
			logger.WithTimeLayout("2006-01-02 15:04:05"),
		)
	}

	if err != nil {
		panic(err)
	}
}
