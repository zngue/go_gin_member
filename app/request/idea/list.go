package idea
import "github.com/zngue/go_tool/src/common/request"
type ListRequest struct {
	request.Page
	CateId int `form:"cate_id"`
}