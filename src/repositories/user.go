package repositories

import (
	"github.com/InsideCI/nego/src/models"
	"github.com/jinzhu/gorm"
)

//UserRepository represents a repository for User model.
type UserRepository struct {
	GenericRepository
}

//NewUserRepository returns a new instance of user repository.
func NewUserRepository() *UserRepository {
	return &UserRepository{
		struct{ Type interface{} }{Type: models.User{}},
	}
}

//ExistsByEmail checks if there's an already registered user with that email.
func (r *UserRepository) ExistsByEmail(db *gorm.DB, email string) bool {
	var count int
	user := models.User{}
	db.Model(&user).Where("email = ?", email).Count(&count)
	return count != 0
}

//GetByEmail returns a User by it's Email.
func (r *UserRepository) GetByEmail(db *gorm.DB, user *models.User) error {
	return db.Where("email = ?", user.Email).Find(&user).Error
}
