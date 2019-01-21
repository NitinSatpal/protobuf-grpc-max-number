package boltdb

import (
	"errors"
	"time"

	"github.com/asdine/storm/q"

	"github.com/asdine/storm"

	"max/server/store"
	t "max/server/store/types"
)

type adapter struct {
	db      *storm.DB
	dsn     string
	dbName  string
	version int
}

const (
	defaultDatabase = "my.db"

	adapterName = "boltdb"
)

func (a *adapter) Open(path string) error {
	if a.db != nil {
		return errors.New("adapter is already connected")
	}

	a.dbName = path
	if a.dbName == "" {
		a.dbName = defaultDatabase
	}

	var err error
	a.db, err = storm.Open(a.dbName)
	if err != nil {
		return err
	}

	a.db.Init(&t.User{})

	return err
}

// Close closes the underlying database connection
func (a *adapter) Close() error {
	var err error
	if a.db != nil {
		err = a.db.Close()
		a.db = nil
		a.version = -1
	}

	return err
}

func (a *adapter) IsOpen() bool {
	return a.db != nil
}

// GetName returns string that adapter uses to register itself with store.
func (a *adapter) GetName() string {
	return adapterName
}

// UserCreate creates a new user. Returns error and true if error is due to duplicate user name,
// false for any other error
func (a *adapter) UserCreate(user *t.User) error {
	err := a.db.Save(user)
	if err != nil {
		return err
	}

	return nil
}

// UserGet fetches a single user by user id. If user is not found it returns (nil, nil)
func (a *adapter) UserGet(uid string) (*t.User, error) {
	var user t.User
	err := a.db.One("Identity", uid, &user)
	if err == nil {
		return &user, nil
	}

	return nil, err
}

func (a *adapter) UserDelete(uid string) error {
	var err error
	err = a.db.Select(q.Eq("Identity", uid)).Delete(&t.User{})

	return err
}

// UserUpdate updates user object.
func (a *adapter) UserUpdate(uid string, update map[string]interface{}) error {
	return a.db.Select(q.Eq("Identity", uid)).Each(&t.User{}, func(u interface{}) error {
		user := u.(*t.User)

		user.UpdatedAt = update["updatedAt"].(time.Time)
		user.CurrentMaxNumber = update["currentMaxNumber"].(int)

		return a.db.Save(user)
	})
}

func init() {
	store.RegisterAdapter(adapterName, &adapter{})
}
