package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
		Users    []User    `json:"users"`
		Profiles []Profile `json:"profiles"`
	}
)

func (i *db) Init(dbStorePath string) error {
	contentByte, err := ioutil.ReadFile(path.Join(dbStorePath, "data.json"))

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

func (i *db) GetUserById(id *int) interface{} {

	for _, user := range i.Users {
		if user.Id == *id {
			return user
		}
	}

	return nil
}

func (i *db) GetUsers () []User  {
	return i.Users
}

var Data db
