package model

type IdeaTag struct {
	CateID int 		`gorm:"column:cate_id;type:int(10)" json:"cate_id"`
	IdeaID int 		`gorm:"column:idea_id;type:int(10)" json:"idea_id"`
}

func (IdeaTag) TableName()  string {
	return "idea_tag"
}