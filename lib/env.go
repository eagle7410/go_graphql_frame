package lib

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path"
	"reflect"
	"strings"
)

type env struct {
	WorkDir,
	TimeZone,
	CookHashKey,
	CookHashValue,
	Place string
	IsDev bool
}

func (i *env) GetCookHashKeyLink() *string {
	return &i.CookHashKey
}

func (i *env) GetCookHashValueLink() *string {
	return &i.CookHashValue
}

func (i *env) GetIsDev() bool {
	return i.IsDev
}

func (i *env) Init() error {

	pwd, err := os.Getwd()

	fmt.Println(pwd)

	if err != nil {
		return err
	}

	i.WorkDir = pwd

	envPath := path.Join(pwd, ".env")

	if _, err := os.Stat(envPath); err == nil {
		fmt.Println("Env load from file")
		err := godotenv.Load(envPath)

		if err != nil {
			return err
		}
	}

	props := map[string]bool{
		"CookHashValue": true,
		"CookHashKey":   true,
		"TimeZone":      false,
		"Place":         true,
	}

	for prop, isRequired := range props {

		v := os.Getenv(prop)

		if isRequired == true && v == "" {
			return errors.New("Bad " + prop)
		}

		reflect.ValueOf(i).Elem().FieldByName(prop).SetString(v)
	}

	if strings.ToLower(os.Getenv("isDev")) == "true" {
		i.IsDev = true
	}

	if len(i.CookHashValue) < 16 {
		return errors.New("CookHashValue must be 16 char")
	}

	if len(i.CookHashKey) < 16 {
		return errors.New("CookHashKey must be 16 char")
	}

	if i.TimeZone == "" {
		i.TimeZone = "Europe/London"
	}

	os.Setenv("TZ", i.TimeZone)

	return nil
}

var ENV env
