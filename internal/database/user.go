package database

import (
	"fmt"
	"time"
)

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	data, err := c.readDB()
	if err != nil {
		return User{}, nil
	}

	if _, ok := data.Users[email]; ok {
		return User{}, fmt.Errorf("User with email, %s already exists", email)
	}

	user := User{Email: email, Password: password, Name: name, Age: age, CreatedAt: time.Now().UTC()}

	data.Users[email] = user

	c.updateDB(data)
	return user, nil
}
