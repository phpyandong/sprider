package model

import (
	"sprider/config"

	"github.com/jinzhu/gorm"
	"github.com/hicsgo/ging"
	"fmt"
	"sprider/basic"
	"github.com/pkg/errors"
)

// CREATE TABLE `yue_comment` (
//   `id` int(11) NOT NULL AUTO_INCREMENT,
//   `from_uid` int(11) NOT NULL DEFAULT '0' COMMENT '评论人用户id',
//   `to_uid` int(11) NOT NULL DEFAULT '0' COMMENT '被评论用户id',
//   `comment` text NOT NULL COMMENT '评论内容',
//   `create_time` int(10) NOT NULL DEFAULT '0' COMMENT '评论时间',
//   `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '评论类型，1文本评论 2语音评论',
//   `voice_url` varchar(255) NOT NULL DEFAULT '' COMMENT '音频地址',
//   `voice_time` smallint(6) NOT NULL DEFAULT '0' COMMENT '播放时间',
//   `source_os_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '评论来源系统类型 0: 未知, 1: iOS 2: Android 3: H5 4: Mini',
//   PRIMARY KEY (`id`),
//   KEY `idx1` (`create_time`,`from_uid`,`to_uid`)
// ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='评论表';

type Comment struct {
	BaseModel
	ID           int64  `gorm:"primary_key;column:id"` //主键
	FromUID      int64  `gorm:"column:from_uid"`       //
	ToUID        int64  `gorm:"column:to_uid"`         //
	Comment      string `grom:"column:comment"`        //评论内容
	CreationDate int64  `gorm:"column:create_time"`    //创建时间
	CType        int64  `gorm:"column:type"`           //评论类型，1文本评论 2语音评论
	VoiceUrl     string `gorm:"column:voice_url"`      //音频地址
	VoiceTime    int64  `gorm:"column:voice_time"`     //
	GroupID      int64  `gorm:"column:grp_id"`         //组
	OsType       int64  `gorm:"column:source_os_type"` //评论来源系统类型 0: 未知, 1: iOS 2: Android 3: H5 4: Mini
	SourceFrom   int64  `gorm:"column:source_from"` //评论来源语言 0: go, 1: php
}

func (comment *Comment) GetProjectName() string {
	return config.ProjectName
}
func (comment *Comment) TableName() string {
	return "yue_comment"
}

type CommentCondition struct {
	Paging        *ging.Paging
	ID            *int64
	GroupID       *int64
	MinCreateTime *int64
	MaxCreateTime *int64
	LastMaxID     *int64
	LastMinID     *int64
	Status        *int64
	SourceFrom    *int64

}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断数据是否存在（多个自定义复杂条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (comment *Comment) IsExists(
	query interface{},
	args ...interface{}) (bool, error) {

	if count, err := comment.GetCount(query, args...); err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return true, err
	} else {
		if count > 0 {
			return true, nil
		}
	}

	return false, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取数据记录数（多个自定义复杂条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (comment *Comment) GetCount(query interface{}, args ...interface{}) (int64, error) {
	comment.ReadModelDbMap(comment)

	paging := ging.Paging{
		PagingIndex: 1,
		PagingSize:  1,
	}

	if err := comment.DbMap.Model(comment).Where(
		query, args...).Count(&paging.TotalRecord).Error; err != nil {
		return 0, err
	}

	return paging.TotalRecord, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取单条数据（单个简单条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (comment *Comment) Select(fieldName string, fieldValue interface{}) error {
	comment.ReadModelDbMap(comment)

	query := map[string]interface{}{}
	query[fieldName] = fieldValue

	if err := comment.DbMap.Find(comment, query).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return basic.NotFoundError
		}
		return err
	}
	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取单条数据（多个自定义复杂条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (comment *Comment) SelectQuery(query interface{}, args ...interface{}) error {
	comment.ReadModelDbMap(comment)

	if err := comment.DbMap.Where(
		query, args...).Find(comment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return basic.NotFoundError
		}
		return err
	}
	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取单条数据（多个自定义复杂条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (comment *Comment) SelectOrderQuery(query interface{}, sortorder string, args ...interface{}) error {
	comment.ReadModelDbMap(comment)

	if err := comment.DbMap.Where(
		query, args...).Order(sortorder).Limit(1).Find(comment).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return basic.NotFoundError
		}
		return err
	}
	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取单条数据（主键标识简单条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (comment *Comment) SelectById(id interface{}) error {

	comment.ReadModelDbMap(comment)

	query := map[string]interface{}{
		"did": id,
	}

	if err := comment.DbMap.Find(comment, query).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.Wrapf(basic.NotFoundError,
				fmt.Sprintf("sql: %s error: %v", "sql语句，orm 没有合适的获得sql的方式",
					err))
		}
		return errors.Wrapf(err,
			fmt.Sprintf("sql: %s error: %v", "sqlxxxxx语句，orm 没有合适的获得sql的方式",
				err))	}

	return nil
}

/* ================================================================================
 * query: []interface{} || map[string]interface{} || string
 * args: if string: interface{}...
 * ================================================================================ */
func (comment *Comment) SelectAll(paging *ging.Paging, query interface{}, args ...interface{}) ([]*Comment, error) {
	comment.ReadModelDbMap(comment)

	var miniCommentModelList = make([]*Comment, 0)
	var err error = nil

	if paging != nil {
		isTotalRecord := true
		if paging.IsTotalOnce {
			if paging.PagingIndex > 1 {
				isTotalRecord = false
			}
		}

		if isTotalRecord && paging.PagingSize > 0 {
			if len(paging.Group) == 0 {
				err = comment.DbMap.Model(comment).
					Where(query, args...).
					Order(paging.Sortorder).
					Offset(paging.Offset()).
					Limit(paging.PagingSize).
					Find(&miniCommentModelList).Error

				// err = comment.DbMap.Model(comment).
				// 	Where(query, args...).
				// 	Count(&paging.TotalRecord).Error
			} else {
				err = comment.DbMap.Model(comment).
					Where(query, args...).
					Group(paging.Group).
					Order(paging.Sortorder).
					Offset(paging.Offset()).
					Limit(paging.PagingSize).
					Find(&miniCommentModelList).Error

				// err = comment.DbMap.Model(comment).
				// 	Where(query, args...).
				// 	Count(&paging.TotalRecord).Error
			}

			//paging.SetTotalRecord(paging.TotalRecord)
		} else {
			if len(paging.Group) == 0 {
				err = comment.DbMap.Model(comment).
					Where(query, args...).
					Order(paging.Sortorder).
					Find(&miniCommentModelList).Error
			} else {
				err = comment.DbMap.Model(comment).
					Where(query, args...).
					Group(paging.Group).
					Order(paging.Sortorder).
					Find(&miniCommentModelList).Error
			}
		}
	} else {
		err = comment.DbMap.Where(query, args...).Find(&miniCommentModelList).Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = basic.NotFoundError
		}
	}

	return miniCommentModelList, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 插入数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (comment *Comment) Insert() error {

	if comment.DbMap == nil {
		comment.WriteModelDbMap(comment)

	}

	if err := comment.DbMap.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

/* ================================================================================
 * data type:
 * Model{"fieldName":"value"...}
 * map[string]interface{}
 * key1,value1,key2,value2
 * ================================================================================ */
func (comment *Comment) Update(data ...interface{}) (int64, error) {

	if comment.ID == 0 || len(data) == 0 {
		return 0, basic.NotFoundError
	}

	if comment.DbMap == nil {
		comment.WriteModelDbMap(comment)

	}

	dbContext := comment.DbMap.Model(comment).UpdateColumns(data)
	rowsAffected, err := dbContext.RowsAffected, dbContext.Error

	return rowsAffected, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 删除数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (comment *Comment) Delete() (int64, error) {
	if comment.ID == 0 {
		return 0, basic.NotFoundError
	}

	if comment.DbMap == nil {
		comment.WriteModelDbMap(comment)

	}

	dbContext := comment.DbMap.Delete(comment)
	rowsAffected, err := dbContext.RowsAffected, dbContext.Error

	return rowsAffected, err
}
