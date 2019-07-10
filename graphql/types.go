package graphql

import "github.com/graphql-go/graphql"

var profileType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Profile",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name_first": &graphql.Field{
				Type: graphql.String,
			},
			"name_last": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"login": &graphql.Field{
				Type: graphql.String,
			},
			"pass": &graphql.Field{
				Type: graphql.String,
			},
			"profile": &graphql.Field{
				Type: profileType,
				Resolve: ResolveProfile,
			},
		},
	},
)
