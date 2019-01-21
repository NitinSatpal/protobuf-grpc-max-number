package store

import (
	"time"

	"max/server/db"
	"max/server/store/types"
)

var adp adapter.Adapter

func Open(path string) error {
	return adp.Open(path)
}

func Close() error {
	if adp.IsOpen() {
		return adp.Close()
	}

	return nil
}

func IsOpen() bool {
	if adp != nil {
		return adp.IsOpen()
	}

	return false
}

func GetAdapterName() string {
	if adp != nil {
		return adp.GetName()
	}

	return ""
}

func RegisterAdapter(name string, a adapter.Adapter) {
	if a == nil {
		panic("store: Register adapter is nil")
	}

	if adp != nil {
		panic("store: adapter '" + adp.GetName() + "' is already registered")
	}

	adp = a
}

// UsersObjMapper is a users struct to hold methods for persistence mapping for the User object.
type UsersObjMapper struct{}

// Users is the ancor for storing/retrieving User objects
var Users UsersObjMapper

// Create inserts User object into a database, updates creation time and assigns UID
func (UsersObjMapper) Create(user *types.User) (*types.User, error) {
	user.InitTimes()

	err := adp.UserCreate(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (UsersObjMapper) Get(uid string) (*types.User, error) {
	return adp.UserGet(uid)
}

func (UsersObjMapper) Delete(uid string) error {
	return adp.UserDelete(uid)
}

func (UsersObjMapper) Update(uid string, update map[string]interface{}) error {
	update["updatedAt"] = time.Now()
	return adp.UserUpdate(uid, update)
}
