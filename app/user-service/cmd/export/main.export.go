package serviceexporter

import (
	userdapi "github.com/go-micro-saas/user-service/api/user-service"
	"github.com/go-micro-saas/user-service/app/user-service/internal/conf"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
)

func ExportServiceConfig() []configutil.Option {
	return conf.LoadServiceConfig()
}

func ExportAuthWhitelist() []map[string]middlewareutil.TransportServiceKind {
	return []map[string]middlewareutil.TransportServiceKind{
		userdapi.GetAuthWhiteList(),
	}
}

//func ExportServices(launcherManager setuputil.LauncherManager, serverManager serverutil.ServerManager) (cleanuputil.CleanupManager, error) {
//	hs, err := serverManager.GetHTTPServer()
//	if err != nil {
//		return nil, err
//	}
//	gs, err := serverManager.GetGRPCServer()
//	if err != nil {
//		return nil, err
//	}
//	//return exportServices(launcherManager, hs, gs)
//	return cleanuputil.Merge(exportServices(launcherManager, hs, gs))
//}

//func ExportServicesWithDatabaseMigration(launcherManager setuputil.LauncherManager, serverManager serverutil.ServerManager) (cleanuputil.CleanupManager, error) {
//	settingConfig := launcherManager.GetConfig().GetSetting()
//
//	if settingConfig.GetEnableMigrateDb() {
//		dbConn, err := setuputil.GetRecommendDBConn(launcherManager)
//		if err != nil {
//			return nil, err
//		}
//		logger, err := setuputil.GetLogger(launcherManager)
//		if err != nil {
//			return nil, err
//		}
//		dbmigrate.Run(dbConn, dbutil.WithLogger(logger))
//	}
//	return ExportServices(launcherManager, serverManager)
//}

//func ExportDatabaseMigration() []dbutil.MigrationFunc {
//	return []dbutil.MigrationFunc{
//		dbmigrate.Run,
//	}
//}
