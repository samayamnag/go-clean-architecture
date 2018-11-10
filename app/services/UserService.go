package services

import (
	"strings"
	"github.com/samayamnag/boilerplate/app/models/mongo"
	"github.com/samayamnag/boilerplate/app/repositories"
)

//Service service interface
type UserService struct {
	repo repositories.UserRepoInterface
}

//NewService create new service
func NewUserService(r repositories.UserRepoInterface) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) Store(u mongo.User) (mongo.User, error) {
	u.ID = mongo.NewID()
	u.SetEmailField()
	u.SetFullNameField()
	u.SetPasswordField()
	u.BeforeInsert()
	return s.repo.Store(u)
}

//Count all users
func (s *UserService) Count() (int, error) {
	return s.repo.Count()
}

//Find user
func (s *UserService) Find(id mongo.ID) (*mongo.User, error) {
	return s.repo.Find(id)
}

//FindAll users
func (s *UserService) FindAll() ([]*mongo.User, error) {
	return s.repo.FindAll()
}

//Search users by email
func (s *UserService) FindByEmail(email string) (*mongo.User, error) {
	return s.repo.FindByEmail(email)
}

//Search users
func (s *UserService) Search(query string) ([]*mongo.User, error) {
	return s.repo.Search(strings.ToLower(query))
}

func (s *UserService) Update(u *mongo.User) (*mongo.User, error) {
	u.SetEmailField()
	u.SetFullNameField()
	u.SetPasswordField()
	u.BeforeUpdate()
	return s.repo.Update(u)
}

//Delete user
func (s *UserService) Delete(id mongo.ID) error {
	b, err := s.Find(id)
	if err != nil {
		return err
	}
	if b.Admin {
		return mongo.ErrCannotBeDeleted
	}
	return s.repo.Delete(id)
}
