package initialize

import (
	"gsafety/global"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	GS_DB := func() *gorm.DB {
		switch global.GS_CONFIG.System.DbType {
		// case "mysql":
		// 	return GormMysql()
		// case "pgsql":
		// 	return GormPgSql()
		case "cockroach":
			return GormCockroachDB()
		default:
			return GormCockroachDB()
		}
	}()
	global.GS_DB = GS_DB
	global.GS_LOG.Info("GS_DB:", zap.String("db", GS_DB.Name()))
	return GS_DB
}
