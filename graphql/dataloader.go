package graphql

import (
	"context"
	"go_graphql_frame/db"
	util "github.com/eagle7410/go_util/libs"
	"strconv"
)
import "github.com/graph-gophers/dataloader"

type appDataLoader struct {
	ProfileLoader *dataloader.Loader
}

func (i *appDataLoader) Init() {
	i.ProfileLoader = dataloader.NewBatchedLoader(profileBatch)
}

func profileBatch(_ context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result

	for _, key := range keys {
		data := dataloader.Result{}
		id64, err := strconv.ParseInt(key.String(), 10, 64)

		if err != nil {
			data.Error = err
		} else {
			id := int(id64)

			util.Logf("Use data loader profileBatch %v", id)

			data.Data = db.Data.GetProfileById(&id)
		}

		results = append(results, &data)

	}

	return results
}

var Dataloders appDataLoader
