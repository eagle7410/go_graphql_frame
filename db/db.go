package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type (
	User struct {
		Id    int    `json:"id"`
		Login string `json:"login"`
		Pass  string `json:"pass"`
	}
	Profile struct {
		Id        int    `json:"id"`
		NameFirst string `json:"name_first"`
		NameLast  string `json:"name_last"`
	}
	db struct {
		Users     []User    `json:"users"`
		Profiles  []Profile `json:"profiles"`
		StoreFile string    `json:"-"`
	}
)

func (i *db) Init(dbStorePath string) error {
	i.StoreFile = path.Join(dbStorePath, "data.json")

	contentByte, err := ioutil.ReadFile(i.StoreFile)

	if err != nil {
		return errors.New(fmt.Sprintf("Error read data file: %v", err))
	}

	if err := json.Unmarshal(contentByte, i); err != nil {
		return errors.New(fmt.Sprintf("Error parse data: %v", err))
	}

	return nil
}

func (i *db) GetProfileByBatch(id *int) interface{} {
	for _, profile := range i.Profiles {
		if profile.Id == *id {
			return profile
		}
	}

	return nil
}

func (i *db) GetProfileById(id *int) interface{} {

	Logf("GetProfileById profile with id %v", *id)

	for _, profile := range i.Profiles {
		if profile.Id == *id {
			return profile
		}
	}

	return nil
}
func (i *db) Save() error {
	bytes, err := json.MarshalIndent(i, "", "\t")

	if err != nil {
		return err
	}

	return ioutil.WriteFile(i.StoreFile, bytes, os.FileMode(os.O_WRONLY))

}
func (i *db) CreateUser(user *User) error {
	max := 0

	for _, dbUser := range i.Users {
		if max < dbUser.Id {
			max = dbUser.Id
		}
	}

	user.Id = max + 1

	i.Users = append(i.Users, *user)

	return i.Save()
}

func (i *db) RemoveUser(id *int) (bool, error) {
	for inx, dbUser := range i.Users {
		if dbUser.Id == *id {
			i.Users = append(i.Users[:inx], i.Users[inx+1:]...)

			return true, i.Save()
		}
	}

	return false, nil
}

func (i *db) UpdateUser(user *User) error {
	for inx, dbUser := range i.Users {
		if dbUser.Id == user.Id {
			i.Users[inx].Login = user.Login
			i.Users[inx].Pass = user.Pass

			return i.Save()
		}
	}

	return nil
}

func (i *db) GetUserById(id *int) interface{} {

	for _, user := range i.Users {
		if user.Id == *id {
			return user
		}
	}

	return nil
}

func (i *db) GetUsers() []User {
	return i.Users
}

var Data db
