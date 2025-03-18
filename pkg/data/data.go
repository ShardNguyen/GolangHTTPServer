package data

import "github.com/ShardNguyen/GolangCounter/pkg/entity"

var UserTestData []entity.User

func init() {
	UserTestData = []entity.User{
		{Id: 1, Name: "Loc"},
		{Id: 2, Name: "Anh"},
		{Id: 3, Name: "Sang"},
		{Id: 4, Name: "Nhan"},
	}
}
