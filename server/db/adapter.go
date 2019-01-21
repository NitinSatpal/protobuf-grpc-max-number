// Package adapter contains the interfaces to be implemented by the database adapter
package adapter

import (
	t "max/server/store/types"
)

// Adapter is the interface that must be implemented by a database
type Adapter interface {
	// General

	// Open and configure the adapter
	Open(path string) error
	// Close the adapter
	Close() error
	// IsOpen checks if the adapter is ready for use
	IsOpen() bool
	// GetName returns the name of the adapter
	GetName() string


	// User management

	// UserCreate creates user record
	UserCreate(usr *t.User) error
	// UserGet returns record for a given user ID
	UserGet(id string) (*t.User, error)
	// UserDelete deletes user record
	UserDelete(id string) error
	// UserUpdate updates user record
	UserUpdate(uid string, update map[string]interface{}) error
}
