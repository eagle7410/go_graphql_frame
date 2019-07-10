package graphql

import (
	"github.com/graphql-go/graphql"
)

type (
	graphQl struct {
		schema graphql.Schema
	}
)

func (i *graphQl) GetSchemaLink() *graphql.Schema {
	return &i.schema
}

func (i *graphQl) Init() error {

	Dataloders.Init()

	Logf("-- Init dataloader is OK...")

	schema, err := getSchema()

	if err != nil {
		return err
	}

	i.schema = schema

	Logf("-- Init schema is OK...")

	return nil
}

var Schema graphQl

func getSchema() (graphql.Schema, error) {
	// Schema

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type:    graphql.String,
			Resolve: ResolveHello,
		},
		"user": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: ResolveUser,
		},
		"users" : &graphql.Field{
			Type: graphql.NewList(userType),
			Resolve:ResolveUsers,
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	return graphql.NewSchema(schemaConfig)
}

func ResolveHello(_ graphql.ResolveParams) (interface{}, error) {
	return "world", nil
}
