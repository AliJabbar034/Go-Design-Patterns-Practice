package models

import "database/sql"

type Repository interface {
	AllDogsBreed() ([]*DogBreed, error)
}

type mySqlRepository struct {
	DB *sql.DB
}

func NewMySqlRepository(db *sql.DB) Repository {
	return &mySqlRepository{
		DB: db,
	}
}

type testRepository struct {
	DB *sql.DB
}

func NewTestRepository(db *sql.DB) Repository {
	return &testRepository{
		DB: nil,
	}
}
