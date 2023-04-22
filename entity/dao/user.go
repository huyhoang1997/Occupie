package dao

import "occupie/entity"

type User struct {
	entity.BaseModel
	Profile Profile
	Role    string
}
