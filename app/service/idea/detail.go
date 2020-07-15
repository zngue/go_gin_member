package idea

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_member/app/model"
	"github.com/zngue/go_tool/src/common/request"
)

func ( Idea ) Detail( request request.IDStatusRequest ) (ideas model.Idea,err error)  {
	var idea model.Idea
	c :=model.Common{
		DetailModel: model.DetailModel{
			Request: request,
			First: func(db *gorm.DB) error {
				return db.First(&idea).Error
			},
			WhereFun: func(db *gorm.DB) *gorm.DB {
				return db.Preload("Cate").Where("id = ?" ,request.ID)
			},
		},
	}
	c.DetailModel.DelTail()
	err =c.DetailModel.Err
	ideas = idea
	return
}
