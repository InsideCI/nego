package services

import (
	"errors"
	"os"

	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/repositories"
	"github.com/InsideCI/nego/src/utils/exceptions"
	"golang.org/x/crypto/bcrypt"
)

//UserService represents a service for the User model.
type UserService struct {
	repo *repositories.UserRepository
}

//NewUserService returns a new instance of User service.
func NewUserService() *UserService {
	return &UserService{
		repo: repositories.NewUserRepository(),
	}
}

//Register creates a new user.
func (s *UserService) Register(db *driver.DB, user *models.User) error {
	if user.Token == os.Getenv("REG_KEY") {
		// The user has a valid registraton token.
		if !s.repo.ExistsByEmail(db.Postgres, user.Email) {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
			if err != nil {
				return errors.New("not a valid password: " + err.Error())
			}

			user.Password = string(hashedPassword)

			err = s.repo.Create(db.Postgres, user)
			if err != nil {
				return errors.New("could not create user: " + err.Error())
			}

			//User created.
			//TODO: return JWT.
			return nil
		}
		return errors.New("email already registered")
	}
	return errors.New("not a valid register token")
}

//Login checks if credentials match with registered user.
func (s *UserService) Login(db *driver.DB, user *models.User) *models.NegoError {
	password := user.Password
	err := s.repo.GetByEmail(db.Postgres, user)
	if err != nil {
		return exceptions.NotRegistered
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return exceptions.WrongPassword
	}

	return nil
}
