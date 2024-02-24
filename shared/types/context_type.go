package types

import (
	"strconv"
)

type ManagerContext struct {
	User *UserContext
}

type UserContext struct {
	Id   uint64
	Name string
}

type JwtUserData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewManagerContext() *ManagerContext {
	return &ManagerContext{}
}

func NewUserContext() *UserContext {
	return &UserContext{
		Id:   0,
		Name: "",
	}
}

func (uc *UserContext) LoadFromJwt(data *JwtUserData) *UserContext {
	uc.Id, _ = strconv.ParseUint(data.Id, 10, 0)
	uc.Name = data.Name
	return uc
}

func (pc *ManagerContext) GetUser() *UserContext {
	return pc.User
}
