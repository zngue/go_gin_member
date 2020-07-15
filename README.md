#导入使用包

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_member/app/middleware"
	"github.com/zngue/go_gin_member/migration"
	"github.com/zngue/go_gin_member/router"
	usermigration "github.com/zngue/go_gin_user/migration"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/gin_run"
)

func main() {
	gin_run.GinRun(func(group *gin.RouterGroup) {
		group.Use(middleware.Cors())
		//UserRouter.Router(group)//路由加载自己加载的路由
		//group.Use(middleware.JWTAuth())
		router.Router(group)

	}, func(db *gorm.DB) {
		usermigration.AutoMigrate(db)//数据库迁移
		migration.Migration(db)
	})
	defer db.ConnClose()
}
```