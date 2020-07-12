package model

type Model struct {
	ID         uint32 `gorm:"primary_key" json="id"`
	CreateBy   string `json="created_by"`
	ModifiedBy string `json="modify_by"`
	CreateOn   uint32 `json="created_on"`
	ModifiedOn uint32 `json="modify_on"`
	DeletedOn  uint32 `json="deleted_on"`
	IsDel      uint8  `json="is_del"`
}
