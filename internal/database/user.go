package database

import (
	"fmt"
	"time"
)

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	data, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	if _, ok := data.Users[email]; ok {
		return User{}, fmt.Errorf("User with email, %s already exists", email)
	}

	user := User{Email: email, Password: password, Name: name, Age: age, CreatedAt: time.Now().UTC()}

	data.Users[email] = user

	err = c.updateDB(data)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	data, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	user, ok := data.Users[email]
	if !ok {
		return User{}, fmt.Errorf("User with email, %s doesn't exists", email)
	}

	newUser := User{Email: email, Password: password, Name: name, Age: age, CreatedAt: user.CreatedAt}

	data.Users[email] = newUser

	err = c.updateDB(data)
	if err != nil {
		return User{}, err
	}
	return newUser, nil
}

func (c Client) GetUser(email string) (User, error) {
	data, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	user, ok := data.Users[email]
	if !ok {
		return User{}, fmt.Errorf("User with email, %s doesn't exists", email)
	}

	return user, nil
}

func (c Client) DeleteUser(email string) error {
	data, err := c.readDB()
	if err != nil {
		return nil
	}

	_, ok := data.Users[email]
	if !ok {
		return fmt.Errorf("User with email, %s doesn't exists", email)
	}

	delete(data.Users, email)
	err = c.updateDB(data)
	if err != nil {
		return err
	}
	return nil
}
