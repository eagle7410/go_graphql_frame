package graphql

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"go_graphql_frame/db"
)

func ResolveUsers(_ graphql.ResolveParams) (interface{}, error) {
	return db.Data.GetUsers(), nil
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

	Logf("Remove user with id %v", id)

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

	Logf("Update user with id %v", id)

	return user, db.Data.UpdateUser(&user)
}

func ResolveUser(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)

	Logf("Get user with id %v", id)

	user := db.Data.GetUserById(&id)

	return user, nil
}

func ResolveProfile(p graphql.ResolveParams) (interface{}, error) {
	user, _ := p.Source.(db.User)

	Logf("Find profile with id %v", user.Id)

	loader := Dataloders.ProfileLoader

	key := dataloader.StringKey(fmt.Sprintf("%d", user.Id))

	thunk := loader.Load(context.Background(), key)

	return thunk()
}
