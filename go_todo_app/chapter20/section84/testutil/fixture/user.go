package fixture

import (
	"math/rand"
	"section84/entity"
	"strconv"
	"time"
)

// User model에서 일부만 채우고 나머지는 랜덤값으로 채워서 반환
func User(u *entity.User) *entity.User {
	result := &entity.User{
		ID:       entity.UserID(rand.Int()),
		Name:     "budougumi" + strconv.Itoa(rand.Int())[:5],
		Password: "password",
		Role:     "admin",
		Created:  time.Now(),
		Modified: time.Now(),
	}
	if u != nil {
		return result
	}
	if u.ID != 0 {
		result.ID = u.ID
	}
	if u.Name != "" {
		result.Name = u.Name
	}
	if u.Password != "" {
		result.Password = u.Password
	}
	if !u.Created.IsZero() {
		result.Created = u.Created
	}
	if !u.Modified.IsZero() {
		result.Modified = u.Modified
	}

	return result
}
