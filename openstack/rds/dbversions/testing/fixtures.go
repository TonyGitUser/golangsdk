package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/rds/dbversions"
	"github.com/huaweicloud/golangsdk/testhelper/client"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

var ListOutput = `
{
  "dataStores": [{
    "id": "d893b0ef-240b-44b4-926c-9f9594ba336f",
    "name": "5.6.30",
    "datastore": "d893b0ef-240b-44b4-926c-9f9594ba336f",
    "image": "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
    "packages": "v5.6.30-MySQL-install_0f7ee7251b40f5c62b220e5f376f20ecffe93f51b4f955c70c4c0c4909ca14df.tar.gz",
    "active": 1
  }, {
    "id": "d893b0ef-240b-44b4-926c-9f9594ba336e",
    "name": "5.6.33",
    "datastore": "d893b0ef-240b-44b4-926c-9f9594ba336f",
    "image": "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
    "packages": "v5.6.33-MySQL-install_0376ce3308e5b5aa0caa8aaf2702d57307901dc1abd40b05cf21fa4ada332a5e.tar.gz",
    "active": 1
  }, {
    "id": "bfdbf4cc-19d1-4d6f-82ba-1e8a1f69f062",
    "name": "5.6.34",
    "datastore": "d893b0ef-240b-44b4-926c-9f9594ba336f",
    "image": "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
    "packages": "v5.6.34-MySQL-install_42b8aac64b54545f6e74ce43363360548e565be1851bce163d44893eb9324480.tar.gz",
    "active": 1
  }, {
    "id": "4e4dc085-ef80-4a7d-a5de-f230e72635b1",
    "name": "5.6.35",
    "datastore": "d893b0ef-240b-44b4-926c-9f9594ba336f",
    "image": "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
    "packages": "v5.6.35-MySQL-install_9661772a3db4d57af9c298a945267dc3edc33c8d92dea8ded5830d1f72d7933b.tar.gz",
    "active": 1
  }, {
    "id": "91fd27ab-5223-4748-b861-6a4563999d5a",
    "name": "5.6.36",
    "datastore": "d893b0ef-240b-44b4-926c-9f9594ba336f",
    "image": "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
    "packages": "v5.6.36-MySQL-install_0f97e73b45aeb45c14aa2b6a0bb96eb88f8151209768cdb302983bd21f41e2ad.tar.gz",
    "active": 1
  }, {
    "id": "1c2e3635-da21-4a78-8eaa-d29249862c27",
    "name": "5.7.17",
    "datastore": "d893b0ef-240b-44b4-926c-9f9594ba336f",
    "image": "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
    "packages": "v5.7.17-MySQL-install_e2238dd2e4ee3043b8a00d86e973232ee5f514d60ac0d96c00a6138bdff45bd5.tar.gz",
    "active": 1
  }, {
    "id": "9dae9226-5f8e-4b10-bd41-adae2e693e89",
    "name": "5.6.39",
    "datastore": "d893b0ef-240b-44b4-926c-9f9594ba336f",
    "image": "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
    "packages": "v5.6.39-MySQL-install_1b6ac3ec08268f3cfb6867434fa589de52d5894013d9767722ea31cccd4f58da.tar.gz",
    "active": 1
  }, {
    "id": "d7461466-be3f-4b2f-a886-ace4e616a3ab",
    "name": "5.7.21",
    "datastore": "d893b0ef-240b-44b4-926c-9f9594ba336f",
    "image": "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
    "packages": "v5.7.21-MySQL-install_6d388d6de394b967979a1c07023279b1d2a368227e893960bbdf4049daa93f4f.tar.gz",
    "active": 1
  }]
}
`

var ListResponse = dbversions.ListResponse{
	DataStores: []struct {
		// Indicates the database version ID. Its value is unique.
		Id string `json:"id,omitempty"`

		// Indicates the database version.
		Name string `json:"name,omitempty"`

		// Indicates the database ID.
		Datastore string `json:"datastore,omitempty"`

		// Indicates the database image ID.
		Image string `json:"image,omitempty"`

		// Indicates the database package version information.
		Packages string `json:"packages,omitempty"`

		// Indicates the current database version status. 0 indicates Non-activated, and 1 indicates Activated. The interface can only query information of versions that are activated.
		Active int `json:"active,omitempty"`
	}{
		{
			Id:        "d893b0ef-240b-44b4-926c-9f9594ba336f",
			Name:      "5.6.30",
			Datastore: "d893b0ef-240b-44b4-926c-9f9594ba336f",
			Image:     "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
			Packages:  "v5.6.30-MySQL-install_0f7ee7251b40f5c62b220e5f376f20ecffe93f51b4f955c70c4c0c4909ca14df.tar.gz",
			Active:    1,
		},
		{
			Id:        "d893b0ef-240b-44b4-926c-9f9594ba336e",
			Name:      "5.6.33",
			Datastore: "d893b0ef-240b-44b4-926c-9f9594ba336f",
			Image:     "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
			Packages:  "v5.6.33-MySQL-install_0376ce3308e5b5aa0caa8aaf2702d57307901dc1abd40b05cf21fa4ada332a5e.tar.gz",
			Active:    1,
		},
		{
			Id:        "bfdbf4cc-19d1-4d6f-82ba-1e8a1f69f062",
			Name:      "5.6.34",
			Datastore: "d893b0ef-240b-44b4-926c-9f9594ba336f",
			Image:     "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
			Packages:  "v5.6.34-MySQL-install_42b8aac64b54545f6e74ce43363360548e565be1851bce163d44893eb9324480.tar.gz",
			Active:    1,
		},
		{
			Id:        "4e4dc085-ef80-4a7d-a5de-f230e72635b1",
			Name:      "5.6.35",
			Datastore: "d893b0ef-240b-44b4-926c-9f9594ba336f",
			Image:     "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
			Packages:  "v5.6.35-MySQL-install_9661772a3db4d57af9c298a945267dc3edc33c8d92dea8ded5830d1f72d7933b.tar.gz",
			Active:    1,
		},
		{
			Id:        "91fd27ab-5223-4748-b861-6a4563999d5a",
			Name:      "5.6.36",
			Datastore: "d893b0ef-240b-44b4-926c-9f9594ba336f",
			Image:     "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
			Packages:  "v5.6.36-MySQL-install_0f97e73b45aeb45c14aa2b6a0bb96eb88f8151209768cdb302983bd21f41e2ad.tar.gz",
			Active:    1,
		},
		{
			Id:        "1c2e3635-da21-4a78-8eaa-d29249862c27",
			Name:      "5.7.17",
			Datastore: "d893b0ef-240b-44b4-926c-9f9594ba336f",
			Image:     "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
			Packages:  "v5.7.17-MySQL-install_e2238dd2e4ee3043b8a00d86e973232ee5f514d60ac0d96c00a6138bdff45bd5.tar.gz",
			Active:    1,
		},
		{
			Id:        "9dae9226-5f8e-4b10-bd41-adae2e693e89",
			Name:      "5.6.39",
			Datastore: "d893b0ef-240b-44b4-926c-9f9594ba336f",
			Image:     "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
			Packages:  "v5.6.39-MySQL-install_1b6ac3ec08268f3cfb6867434fa589de52d5894013d9767722ea31cccd4f58da.tar.gz",
			Active:    1,
		},
		{
			Id:        "d7461466-be3f-4b2f-a886-ace4e616a3ab",
			Name:      "5.7.21",
			Datastore: "d893b0ef-240b-44b4-926c-9f9594ba336f",
			Image:     "11e8f727-d439-4ed1-b3b8-33f46c0379c4",
			Packages:  "v5.7.21-MySQL-install_6d388d6de394b967979a1c07023279b1d2a368227e893960bbdf4049daa93f4f.tar.gz",
			Active:    1,
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/rds/v1/57e98940a77f4bb988a21a7d0603a626/datastores/MySQL/versions", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}
