package graphql

import (
	"context"
	util "github.com/eagle7410/go_util/libs"
	"github.com/gorilla/securecookie"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"go_graphql_frame/db"
	"net/http"
)

const contextKey = "store"
const cookUserName = "user"

type (
	graphQl struct {
		schema graphql.Schema
	}
	graphQlContext struct {
		Writer       http.ResponseWriter
		Cookies      []*http.Cookie
		Securecookie *securecookie.SecureCookie
		User         db.User
	}
	Env interface {
		GetCookHashKeyLink() *string
		GetCookHashValueLink() *string
	}
)

func (i *graphQl) GetSchemaLink() *graphql.Schema {
	return &i.schema
}

func (i *graphQl) Init() error {

	Dataloders.Init()

	util.Logf("-- Init dataloader is OK...")

	schema, err := getSchema()

	if err != nil {
		return err
	}

	i.schema = schema

	util.Logf("-- Init schema is OK...")

	return nil
}

var Schema graphQl

func getSchema() (graphql.Schema, error) {
	// Schema

	fields := graphql.Fields{
		"me": &graphql.Field{
			Type:    userType,
			Resolve: ResolveMe,
		},
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
		"users": &graphql.Field{
			Type:    graphql.NewList(userType),
			Resolve: ResolveUsers,
		},
	}

	mutation := graphql.Fields{
		"auth": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"login": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"pass": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: ResolveAuth,
		},
		"userDelete": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: ResolveUserRemove,
		},
		"userCreate": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"login": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"pass": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: ResolveUserCreate,
		},
		"userUpdate": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"login": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pass": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: ResolveUserUpdate,
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	rootMutation := graphql.ObjectConfig{Name: "RootMutation", Fields: mutation}

	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}

	return graphql.NewSchema(schemaConfig)
}

func GetHandlerGraphQl(env Env) http.HandlerFunc {

	graphQLHandler := GetGraphQLHandler()

	hashKey := []byte(*env.GetCookHashKeyLink())
	blockKey := []byte(*env.GetCookHashValueLink())

	sc := securecookie.New(hashKey, blockKey)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUser := db.User{}

		if cookie, err := r.Cookie(cookUserName); err == nil {
			if err := sc.Decode(cookUserName, cookie.Value, &currentUser); err != nil {
				util.Logf("Error decode cook %v", err)
			}
		}

		contextValue := graphQlContext{
			Cookies:      r.Cookies(),
			Writer:       w,
			Securecookie: sc,
			User:         currentUser,
		}

		ctx := context.WithValue(r.Context(), contextKey, contextValue)

		graphQLHandler.ContextHandler(ctx, w, r)
	})
}

func GetGraphQLHandler() *handler.Handler {
	return handler.New(&handler.Config{
		Schema:     Schema.GetSchemaLink(),
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
}
