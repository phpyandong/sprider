package model

import (
"log"
_ "github.com/go-sql-driver/mysql"
"github.com/hicsgo/glib"
"github.com/jinzhu/gorm"
	"sprider/config"
)

/* ================================================================================
 * 数据模型相关信息
 * email   : golang123@outlook.com
 * author  : hicsgo
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 基础数据模型(Id应该直接放入basemodel，业务限制现在不使用)
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type BaseModel struct {
	//Id       string   `gorm:"primary_key;column:id"`
	DbMap       *gorm.DB `msgpack:"-" sql:"-" json:"-"`
	DBKeyName   string   `msgpack:"-" sql:"-" json:"-"`
	TableName   string   `msgpack:"-" sql:"-" json:"-"`
	ProjectName string   `msgpack:"-" sql:"-" json:"-"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 实例化ReadDbMap
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (baseModel *BaseModel) ReadModelDbMap(s ISharing) {

	baseModel.TableName = s.TableName()
	baseModel.ProjectName = s.GetProjectName()

	dbMap := config.ReadDB
	baseModel.DbMap = dbMap

}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 实例化WirteDbMap
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (baseModel *BaseModel) WriteModelDbMap(s ISharing) {

	baseModel.TableName = s.TableName()
	baseModel.ProjectName = s.GetProjectName()

	dbMap := config.WriteDB
	baseModel.DbMap = dbMap
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 注册回调钩子
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
/*
func (baseModel *BaseModel) RegisterCallback() {
	gorm.DefaultCallback.Update().Replace("gorm:after_update", baseModel.AfterUpdate)
}
*/

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * AfterCreate钩子
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
/*
func (baseModel *BaseModel) AfterCreate(scope *gorm.Scope) {
	tableName := scope.TableName()
	log.Printf("AfterCreate Hook: %s, id: %s", tableName, baseModel.Id)
}
*/

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * AfterUpdate钩子
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
/*
func (baseModel *BaseModel) AfterUpdate(scope *gorm.Scope) {
	tableName := scope.TableName()
	log.Printf("AfterUpdate Hook: %s, id: %s", tableName, baseModel.Id)

	RemoveFromCache(baseModel.Id)
}
*/

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * AfterDelete钩子
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
/*
func (baseModel *BaseModel) AfterDelete(scope *gorm.Scope) {
	tableName := scope.TableName()
	log.Printf("AfterDelete Hook: %s, id: %s", tableName, baseModel.Id)

	RemoveFromCache(baseModel.Id)
}
*/

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 事务(只支持写库同一个库事务)
 * fun: 回调函数，接受事务DbMap
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (baseModel *BaseModel) Transactions(s ISharing, fun func(dbMap *gorm.DB)) error {

	baseModel.TableName = s.TableName()
	baseModel.ProjectName = s.GetProjectName()

	var tranDbMap *gorm.DB = nil
	err := glib.Capture2(
		func() {
			log.Printf("Trans Begin")
			tranDbMap = config.WriteDB.Begin()

			baseModel.DbMap = tranDbMap

			fun(tranDbMap)

			tranDbMap.Commit()
			log.Printf("Trans Commit")
		}, func(e interface{}) {
			tranDbMap.Rollback()
			log.Printf("Trans Rollback %v", e)
		})
	return err
}
