package idea_cate

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_member/app/model"
	"github.com/zngue/go_tool/src/common/request"
)

func (IdeaCate) Detail(request request.IDStatusRequest) (cateOne *model.Cate,err error)  {
	var cate *model.Cate
	c :=model.Common{
		DetailModel: model.DetailModel{
			Request: request,
			First: func(db *gorm.DB) error {
				return db.First(&cate).Error
			},
		},
	}
	err =c.DetailModel.DelTail().Err
	cateOne =cate
	return
}
