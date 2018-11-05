package models

import "github.com/graphql-go/graphql"

// NullString null string
var NullString = graphql.NewObject(graphql.ObjectConfig{
	Name: "NullString",
	Fields: graphql.Fields{
		"String": &graphql.Field{
			Type: graphql.String,
		},
		"Valid": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

// NullInteger null integer
var NullInteger = graphql.NewObject(graphql.ObjectConfig{
	Name: "NullInteger",
	Fields: graphql.Fields{
		"Int64": &graphql.Field{
			Type: graphql.String,
		},
		"Valid": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

// NullFloat null float
var NullFloat = graphql.NewObject(graphql.ObjectConfig{
	Name: "NullFloat",
	Fields: graphql.Fields{
		"Float64": &graphql.Field{
			Type: graphql.String,
		},
		"Valid": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

// NullBool null bool
var NullBool = graphql.NewObject(graphql.ObjectConfig{
	Name: "NullBool",
	Fields: graphql.Fields{
		"Bool": &graphql.Field{
			Type: graphql.String,
		},
		"Valid": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
