package graphql

import (
	"context"
	"errors"
	"fmt"
	util "github.com/eagle7410/go_util/libs"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"go_graphql_frame/db"
	"net/http"
)

func ResolveUsers(_ graphql.ResolveParams) (interface{}, error) {
	return db.Data.GetUsers(), nil
}

func ResolveMe(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context.Value(contextKey).(graphQlContext)

	return ctx.User, nil
}

func ResolveAuth(p graphql.ResolveParams) (interface{}, error) {

	login, _ := p.Args["login"].(string)
	pass, _ := p.Args["pass"].(string)

	user, isFind := db.Data.FindByLoginPass(&login, &pass)

	var err error

	if !isFind {
		err = errors.New("User not found")
	} else {
		ctx := p.Context.Value(contextKey).(graphQlContext)
		encoded, err := ctx.Securecookie.Encode(cookUserName, user)

		if err == nil {
			cookie := &http.Cookie{
				Name:     cookUserName,
				Value:    encoded,
				Path:     "/",
				HttpOnly: true,
			}

			http.SetCookie(ctx.Writer, cookie)
		} else {
			util.LogEF("Error in ResolveAuth, encode cookie: %v", err)
		}
	}

	return user, err
}

func ResolveUserCreate(p graphql.ResolveParams) (interface{}, error) {
	login, _ := p.Args["login"].(string)
	pass, _ := p.Args["pass"].(string)

	user := db.User{
		Login: login,
		Pass:  pass,
	}

	err := db.Data.CreateUser(&user)

	return user, err
}

func ResolveUserRemove(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)

	return db.Data.RemoveUser(&id)
}

func ResolveUserUpdate(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)
	login, _ := p.Args["login"].(string)
	pass, _ := p.Args["pass"].(string)

	user := db.User{
		Id:    id,
		Login: login,
		Pass:  pass,
	}
	return user, db.Data.UpdateUser(&user)
}

func ResolveUser(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)

	user := db.Data.GetUserById(&id)

	return user, nil
}

func ResolveProfile(p graphql.ResolveParams) (interface{}, error) {
	user, _ := p.Source.(db.User)

	loader := Dataloders.ProfileLoader

	key := dataloader.StringKey(fmt.Sprintf("%d", user.Id))

	thunk := loader.Load(context.Background(), key)

	return thunk()
}

func ResolveHello(_ graphql.ResolveParams) (interface{}, error) {
	return "world", nil
}
