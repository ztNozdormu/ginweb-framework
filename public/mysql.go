package public

import (
	"github.com/ztNozdormu/gweb-common/lib"
	"github.com/e421083458/gorm"
)
/**
 * 数据库初始化
 */
var (
	GormPool *gorm.DB
)

func InitMysql() error {
	dbpool, err := lib.GetGormPool("default")
	if err != nil {
		return err
	}
	GormPool = dbpool
	return nil
}
