package stat

import (
	"gorm.io/gorm"
	"github.com/xhigher/hzgo/demo/model/db"
)

func DB() *gorm.DB {
	return db.StatDB()
}
