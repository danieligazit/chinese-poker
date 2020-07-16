package utils

import "sync/atomic"

type UserNamesRegistry struct {
	userNames map[uint64]string
	lastId       uint64
}

func NewUserNameRegistry() *UserNamesRegistry {
	return &UserNamesRegistry{
		userNames: make(map[uint64]string),
	    lastId: 0,
	}
}

func (registry *UserNamesRegistry) AddUserName(userName string) (id uint64){
    id = atomic.AddUint64(registry.lastId, 1)
	registry.userNames[id] = userName
	return
}

func (registry *UserNamesRegistry) GetUserName(id uint64) string {
	return registry.userNames[id]
}