package database

import (
	"errors"
	"os"
)

type Client struct {
	Path string
}

func NewClient(db_path string) Client {
	client := &Client{
		Path: db_path,
	}

	return *client
}

func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.Path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return c.createDB()
		}
		return err
	}

	return nil
}
