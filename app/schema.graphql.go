package app

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/Arijeet-webonise/go-react/app/models"
	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

//GraphQlService encapsulates GraphQl functions
type GraphQlService interface {
	GetSchema() (graphql.Schema, error)
}

// GraphQlServiceImpl implement GraphQlService
type GraphQlServiceImpl struct {
	DB models.XODB
}

//GetSchema return Schema for GraphQl
func (serviceImpl *GraphQlServiceImpl) GetSchema() (graphql.Schema, error) {

	fields := graphql.Fields{
		"ping": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "pong", nil
			},
		},
		"user": &graphql.Field{
			Type:        models.UserSchema,
			Description: "Fetch User by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, err := strconv.Atoi(p.Args["id"].(string))
				if err != nil {
					return nil, err
				}
				user, err := models.UserByID(serviceImpl.DB, id)
				if err != nil {
					return nil, err
				}
				return &struct {
					ID         int
					Email      string
					FirstName  string
					LastName   string
					ModifiedAt time.Time
					CreatedAt  time.Time
				}{user.ID, user.Email, user.FirstName.String, user.LastName.String, user.ModifiedAt, user.CreatedAt}, nil
			},
		},
		"login": &graphql.Field{
			Type:        models.UserSchema,
			Description: "Check if User Logged in",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(param graphql.ResolveParams) (interface{}, error) {
				user, err := models.UserByEmail(serviceImpl.DB, param.Args["email"].(string))
				if err != nil {
					return nil, err
				}
				if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Args["password"].(string))); err != nil {
					return nil, err
				}
				return &struct {
					ID         int
					Email      string
					FirstName  string
					LastName   string
					ModifiedAt time.Time
					CreatedAt  time.Time
				}{user.ID, user.Email, user.FirstName.String, user.LastName.String, user.ModifiedAt, user.CreatedAt}, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: serviceImpl.GetMutation(),
	}
	return graphql.NewSchema(schemaConfig)
}

//GetMutation return all mutations
func (serviceImpl *GraphQlServiceImpl) GetMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "RootMutaion",
		Description: "Manupulate data in DB",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: models.UserSchema,
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"firstName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"lastName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					firstName := params.Args["firstName"].(string)
					lastName := params.Args["lastName"].(string)
					user := &models.User{
						Email:      params.Args["email"].(string),
						FirstName:  sql.NullString{String: firstName, Valid: (firstName == "")},
						LastName:   sql.NullString{String: lastName, Valid: (lastName == "")},
						ModifiedAt: time.Now(),
						CreatedAt:  time.Now(),
					}

					userService := models.UserServiceImpl{DB: serviceImpl.DB}
					if err := userService.InsertUser(user); err != nil {
						return nil, err
					}

					return &struct {
						ID         int
						Email      string
						FirstName  string
						LastName   string
						ModifiedAt time.Time
						CreatedAt  time.Time
					}{user.ID, user.Email, user.FirstName.String, user.LastName.String, user.ModifiedAt, user.CreatedAt}, nil

				},
			},
		},
	})
}
