package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_member/app/model"
)

func Migration(db *gorm.DB)  {
	db.AutoMigrate(
		new(model.Members),
		new(model.MemberGroup),
		new(model.Idea),
		new(model.Cate),
	)
}
