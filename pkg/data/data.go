package data

import "github.com/ShardNguyen/GolangCounter/pkg/entity"

var UserTestData map[int]entity.User

func init() {
	UserTestData[1] = *entity.NewUser(1, "Loc")
	UserTestData[2] = *entity.NewUser(2, "Anh")
	UserTestData[3] = *entity.NewUser(3, "Sang")
	UserTestData[4] = *entity.NewUser(4, "Nhan")
	UserTestData[5] = *entity.NewUser(5, "Thinh (Bo truong bo thuc pham)")
}
