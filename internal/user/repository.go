package user

import (
	"sanguo-server/pkg/database"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUser(dbs ...*gorm.DB) *Repository {
	db := database.GetDB()
	if len(dbs) > 0 {
		db = dbs[0]
	}

	return &Repository{
		db: db,
	}
}

func (r *Repository) TableName() string {
	return "user"
}

func (r *Repository) Create(m *User) error {
	panic("implement me")
}

func (r *Repository) Update(m *User) error {
	panic("implement me")
}

func (r *Repository) Delete(id int) error {
	panic("implement me")
}

func (r *Repository) Get(id int) (*User, error) {
	panic("implement me")
}

func (r *Repository) Query(...any) ([]*User, error) {
	panic("implement me")
}
