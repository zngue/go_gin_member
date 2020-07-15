package idea

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_member/app/model"
)

func (Idea) Edit(idea model.Idea) (int , error) {

	c :=model.Common{
		AddModel: model.AddModel{
			AddFun: func(db *gorm.DB) error {
				return db.Create(&idea).Error
			},
		},
		UpdateModel: model.UpdateModel{
			Model: model.Idea{},
			Data: idea,
			WhereFun: func(db *gorm.DB) *gorm.DB {
				return db.Where("id = ?",idea.ID)
			},
		},
	}
	if idea.ID>0 {
		return  idea.ID,c.UpdateModel.Update().Err
	}else{
		return  idea.ID,c.AddModel.Add().Err
	}
}
