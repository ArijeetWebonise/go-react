package models

//CustomUserService encapsulates custom user function
type CustomUserService interface {
	GetUser(email string) (*User, error)
}

//CustomUserServiceImpl implements CustomUserService
type CustomUserServiceImpl struct {
	DB XODB
}

//GetUser return user from email id
func (serviceImpl *CustomUserServiceImpl) GetUser(email string) (*User, error) {
	return UserByEmail(serviceImpl.DB, email)
}
