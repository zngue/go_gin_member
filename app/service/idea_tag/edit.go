package idea_tag

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_member/app/model"
)

func (IdeaTag)Add(tag model.IdeaTag) error {

	c:=model.Common{
		AddModel: model.AddModel{
			AddFun: func(db *gorm.DB) error {
				return db.Create(&tag).Error
			},
		},
	}
	return c.AddModel.Add().Err

}
