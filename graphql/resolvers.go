package graphql

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"go_graphql_frame/db"
	"log"
)

type KeyId struct {
	Val int
}

func (i *KeyId) String () string {
	return fmt.Sprintf("%v", i.Val)
}

func (i *KeyId) Raw () interface{} {
	return *i;
}

func ResolveUsers(_ graphql.ResolveParams) (interface{}, error) {
	return db.Data.GetUsers(), nil
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

	//profile := db.Data.GetProfileById(&user.Id)
	loader := Dataloders.ProfileLoader

	key := dataloader.StringKey(fmt.Sprintf("%d", user.Id))

	thunk := loader.Load(context.Background(), key) // StringKey is a convenience method that make wraps string to implement `Key` interface

	profile, err := thunk()

	log.Printf("value: %#v", profile)
	return profile, err
}
