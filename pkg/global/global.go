package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// ćšć±ćé
var (
	DB  *gorm.DB
	Rdb *redis.Client
)
