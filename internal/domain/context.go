package domain

import "time"

type UserContext struct {
	user        *User
	sessCreated time.Time
	data        map[string]interface{}
}

func NewUserContext(user *User) *UserContext {
	return &UserContext{
		user:        user,
		sessCreated: time.Now(),
		data:        make(map[string]interface{}),
	}
}

func (c *UserContext) User() *User {
	return c.user
}

func (c *UserContext) CreatedAt() time.Time {
	return c.sessCreated
}

func (c *UserContext) Get(key string) interface{} {
	return c.data[key]
}

func (c *UserContext) Add(key string, val interface{}) {
	c.data[key] = val
}
