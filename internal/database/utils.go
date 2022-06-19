package database

import (
	"encoding/json"
	"os"
)

func writeToFile(file_path string, db_schema databaseSchema) error {
	data, _ := json.Marshal(db_schema)
	err := os.WriteFile(file_path, data, 0666)

	return err
}

func (c Client) createDB() error {
	users := make(map[string]User, 0)
	posts := make(map[string]Post, 0)
	err := writeToFile(c.Path, databaseSchema{Users: users, Posts: posts})

	return err
}

func (c Client) readDB() (databaseSchema, error) {
	raw, err := os.ReadFile(c.Path)
	if err != nil {
		return databaseSchema{}, err
	}

	var data databaseSchema
	err = json.Unmarshal(raw, &data)
	if err != nil {
		return databaseSchema{}, err
	}

	return data, nil
}

func (c Client) updateDB(db_schema databaseSchema) error {
	err := writeToFile(c.Path, db_schema)
	return err
}
