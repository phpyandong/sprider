package config

import (
"fmt"
"time"
)

/* ================================================================================
 * 数据设置模块
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * App taipa数据配置
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
var AppSetting *Setting

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func init() {
	fmt.Printf("%v App settings init\n", time.Now())
	AppSetting = &Setting{
		DatabaseConfig:  getDatabaseConfig(),
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据库连接字符串初始化配置
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func getDatabaseConfig() *DatabaseConfig {
	//env := os.Getenv("ENV")

	////本地开发环境
	//if env == ENVIRONMENT_VARIABLE_DEV {
	//	return getLocalTencentDBConfig()
	//}
	////测试环境
	//if env == ENVIRONMENT_VARIABLE_TEST {
	//	return getTestTencentDBConfig()
	//}
	////生产环境
	//if env == ENVIRONMENT_VARIABLE_PRO {
	//	return getProTencentDBConfig()
	//}

	return getLocalTencentDBConfig()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * (本地开发)数据库连接字符串初始化配置 tencent
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func getLocalTencentDBConfig() *DatabaseConfig {
	//本地开发地址
	databaseHost := "bj-cdb-83ypzrxe.sql.tencentcdb.com:62171"
	databaseName := "yue"

	dbConfig := &DatabaseConfig{
		DatabaseOptions: []*DatabaseOption{
			&DatabaseOption{
				ProjectName: ProjectName,
				ReadDBConns: []*DatabaseConnectionOption{{
					Key:      "love_read_db_1",
					Username: "root",
					Password: "oM#d1Jux*R$ggJ39",
					Host:     databaseHost,
					Database: databaseName,
					Dialect:  "mysql",
					IsLog:    true,
					Weight:   50,
				}, {
					Key:      "love_read_db_2",
					Username: "root",
					Password: "oM#d1Jux*R$ggJ39",
					Host:     databaseHost,
					Database: databaseName,
					Dialect:  "mysql",
					IsLog:    true,
					Weight:   50,
				}},
				WirteDBConns: []*DatabaseConnectionOption{{
					Key:      "love_write_db_1",
					Username: "root",
					Password: "oM#d1Jux*R$ggJ39",
					Host:     databaseHost,
					Database: databaseName,
					Dialect:  "mysql",
					IsLog:    true,
					Weight:   50,
				}, {
					Key:      "love_write_db_2",
					Username: "root",
					Password: "oM#d1Jux*R$ggJ39",
					Host:     databaseHost,
					Database: databaseName,
					Dialect:  "mysql",
					IsLog:    true,
					Weight:   50,
				}},
			},
		},
	}

	return dbConfig
}
