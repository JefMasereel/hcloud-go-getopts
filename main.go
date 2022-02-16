package main

import (
	"context"
	"log"
	"os"

	"encoding/json"
	"io/ioutil"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

func main() {

	ctx := context.TODO()

	// Read a private API token from a local file. This token must correspond to
	// a valid Hetzner project. Create new projects via console.hetzner.cloud.
	APIToken, err := os.ReadFile("../localsecrets/apitoken.txt")
	if err != nil {
		log.Fatalf("failed to read apitoken from source, error: %s", err)
	}

	// Alternatively, enter a private API token here directly.
	// APIToken := "sometokenstring"

	// Initialize a client with this token.
	Client := hcloud.NewClient(hcloud.WithToken(string(APIToken)), hcloud.WithApplication("hcloud-go-getopts", ""))

	// Query all supported options from the API, then dump the resulting list of
	// structs into appropriately named json files for manual lookups.
	ImageOptsList, err := Client.Image.All(ctx)
	if err != nil {
		log.Fatalf("failed to collect Image optslist, error: %s", err)
	}
	file, _ := json.MarshalIndent(ImageOptsList, "", " ")
	_ = ioutil.WriteFile("opts/images.json", file, 0644)

	// Repeat for all relevant API parameters. Note that all API information can
	// be found on https://pkg.go.dev/github.com/hetznercloud/hcloud-go.

	// Datacenter
	DatacenterOptsList, err := Client.Datacenter.All(ctx)
	if err != nil {
		log.Fatalf("failed to collect Datacenter optslist, error: %s", err)
	}
	file, _ = json.MarshalIndent(DatacenterOptsList, "", " ")
	_ = ioutil.WriteFile("opts/datacenters.json", file, 0644)

	// Location
	LocationOptsList, err := Client.Location.All(ctx)
	if err != nil {
		log.Fatalf("failed to collect Location optslist, error: %s", err)
	}
	file, _ = json.MarshalIndent(LocationOptsList, "", " ")
	_ = ioutil.WriteFile("opts/locations.json", file, 0644)

	// ServerType
	ServerTypeOptsList, err := Client.ServerType.All(ctx)
	if err != nil {
		log.Fatalf("failed to collect ServerType optslist, error: %s", err)
	}
	file, _ = json.MarshalIndent(ServerTypeOptsList, "", " ")
	_ = ioutil.WriteFile("opts/servertypes.json", file, 0644)
}
