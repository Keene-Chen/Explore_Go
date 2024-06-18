package service

import (
	"Explore_Go/exercise/demo_gin/restful_api/app/model"
)

// 列表方法
// gorm的Unscoped方法设置tx.Statement.Unscoped为true；针对软删除会追加SoftDeleteDeleteClause，即设置deleted_at为指定的时间戳；而callbacks的Delete方法在db.Statement.Unscoped为false的时候才追加db.Statement.Schema.DeleteClauses，而Unscoped则执行的是物理删除。
func UserList(wheres map[string]interface{}) (list []model.User) {
	var users []model.User
	// model.DB().Debug().Where(wheres).Unscoped().Order("id desc").Limit(5).Offset(8).Find(&users)
	model.DB().Debug().Where(wheres).Unscoped().Find(&users)
	return users
}

// 分页方法
func UserPage(wheres map[string]interface{}, pageIndex int, pageSize int) (list []model.User, total int64) {
	var user []model.User
	model.DB().Debug().Where(wheres).Find(&user).Count(&total)
	offset := pageSize * (pageIndex - 1)
	model.DB().Debug().Where(wheres).Unscoped().Order("id desc").Limit(pageSize).Offset(offset).Find(&list)
	return list, total
}

// 取得总行数
func UserTotal(wheres map[string]interface{}) int64 {
	var user []model.User
	var total int64
	model.DB().Debug().Where(wheres).Find(&user).Count(&total)
	return total
}
