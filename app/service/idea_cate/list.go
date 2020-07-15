package idea_cate

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_member/app/model"
)

func (IdeaCate) List()  ([]*model.Cate,error) {
	var ideaCate []*model.Cate
	c:=model.Common{
		ListModel: model.ListModel{
			FindFun: func(db *gorm.DB) error {
				return db.Find(&ideaCate).Error
			},
			OrderFun: func(db *gorm.DB) *gorm.DB {
				return db.Order("id desc")
			},
		},
	}
	return ideaCate,c.ListModel.List().Err
}
