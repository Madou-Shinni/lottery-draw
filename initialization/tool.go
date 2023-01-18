package initialization

import (
	"github.com/Madou-Shinni/lottery-draw/internal/conf"
	"github.com/Madou-Shinni/lottery-draw/pkg/tools/snowflake"
)

func init() {
	// 初始化雪花算法工具
	snowflake.SnowflakeInit(conf.Conf.MachineID)
}
