package memeber

import "github.com/zngue/go_tool/src/common/request"

type ListRequest struct {
	IDString string `json:"id_string" form:"id_string"`
	IDArr []int
	Name string `form:"name"`
	Mobile int `form:"mobile"`
	request.Page
}
