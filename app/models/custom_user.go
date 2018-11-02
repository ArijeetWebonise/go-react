package models

import "github.com/graphql-go/graphql"

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

//UserSchema schema for user table
var UserSchema = graphql.NewObject(graphql.ObjectConfig{
	Name:        "user",
	Description: "User Table",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"lastName": &graphql.Field{
			Type: graphql.String,
		},
		"modifiedAt": &graphql.Field{
			Type: graphql.DateTime,
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
