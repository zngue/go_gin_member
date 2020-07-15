package idea

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_member/app/model"
	"github.com/zngue/go_gin_member/app/request/idea"
)

type IdeaList struct {
	Count int
	Item []*model.Idea
}
func (Idea) List(request idea.ListRequest)  ( list *IdeaList,err error ) {
	var ideaList IdeaList
	c :=model.Common{
		ListModel: model.ListModel{
			Request: &request.Page,
			FindFun: func(db *gorm.DB) error {
				if request.CateId!=0 {
					db = db.Where("cate_id = ?",request.CateId)
				}
				db = db.Preload("Cate").Find(&ideaList.Item)
				return db.Error
			},
			NumFun: func(db *gorm.DB) error {
				return db.Model(&model.Idea{}).Count(&ideaList.Count).Error
			},
			OrderFun: func(db *gorm.DB) *gorm.DB {
				return db.Order("id desc")
			},
		},
	}
	if err =c.ListModel.List().Err; err != nil {
		return
	}else{
		list = &ideaList
		return
	}

}
