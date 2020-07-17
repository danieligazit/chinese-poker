package utility

import "sync/atomic"
import "fmt"

type UserNameRegistry interface {
	AddUserName(userName string) (id uint32)
	GetUserName(id uint32) (username string, exists bool)
}

type UserNameRegistryMap struct {
	lastId      uint32
	username2id map[string]uint32
	id2username map[uint32]string
}

func NewUserNameRegistryMap() *UserNameRegistryMap {
	return &UserNameRegistryMap{
		username2id: make(map[string]uint32),
		id2username: make(map[uint32]string),
		lastId:      0,
	}
}

func (registry *UserNameRegistryMap) AddUserName(username string) (id uint32) {
	id, exists := registry.username2id[username]
	if exists {
		fmt.Println(id)
		return
	}

	id = registry.nextId()

	registry.id2username[id] = username
	registry.username2id[username] = id
	return
}

func (registry *UserNameRegistryMap) GetUserName(id uint32) (username string, exists bool) {
	username, exists = registry.id2username[id]
	return
}

func (registry *UserNameRegistryMap) nextId() uint32 {
	return atomic.AddUint32(&registry.lastId, 1) - 1
}
