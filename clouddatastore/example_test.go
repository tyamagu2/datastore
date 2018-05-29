package clouddatastore

import (
	"context"
	"fmt"

	"go.mercari.io/datastore"
)

const ProjectID = "datastore-wrapper"

func ExampleFromContext() {
	ctx := context.Background()
	client, err := FromContext(
		ctx,
		datastore.WithProjectID(ProjectID),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	type Data struct {
		Name string
	}

	key := client.IncompleteKey("Data", nil)
	entity := &Data{Name: "mercari"}
	key, err = client.Put(ctx, key, entity)
	if err != nil {
		panic(err)
	}

	entity = &Data{}
	err = client.Get(ctx, key, entity)
	if err != nil {
		panic(err)
	}

	fmt.Println(entity.Name)
	// Output: mercari
}