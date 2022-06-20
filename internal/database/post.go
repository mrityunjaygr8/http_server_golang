package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	data, err := c.readDB()
	if err != nil {
		return Post{}, err
	}

	_, ok := data.Users[userEmail]
	if !ok {
		return Post{}, fmt.Errorf("User with email, %s doesn't exists", userEmail)
	}

	id := uuid.New().String()

	post := Post{UserEmail: userEmail, Text: text, ID: id, CreatedAt: time.Now().UTC()}

	data.Posts[id] = post

	err = c.updateDB(data)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	data, err := c.readDB()
	if err != nil {
		return []Post{}, err
	}

	posts := make([]Post, 0)

	for _, value := range data.Posts {
		if value.UserEmail == userEmail {
			posts = append(posts, value)
		}
	}

	return posts, nil
}

func (c Client) DeletePost(id string) error {
	data, err := c.readDB()
	if err != nil {
		return err
	}

	if _, ok := data.Posts[id]; !ok {
		return fmt.Errorf("Post with ID: %s does not exists", id)
	}

	delete(data.Posts, id)
	err = c.updateDB(data)
	if err != nil {
		return err
	}

	return nil
}
