package model

import (
	"github.com/yunshu2009/blog-service/pkg/setting"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json="id"`
	CreateBy   string `json="created_by"`
	ModifiedBy string `json="modify_by"`
	CreateOn   uint32 `json="created_on"`
	ModifiedOn uint32 `json="modify_on"`
	DeletedOn  uint32 `json="deleted_on"`
	IsDel      uint8  `json="is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, error = gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", 
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime
	))
	// other code
}
