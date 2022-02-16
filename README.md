# hcloud-go-getopts

Get a complete, up-to-date overview of various configuration options for the hcloud-go API

## Some context

The [hcloud-go package](https://github.com/hetznercloud/hcloud-go) developed by [Hetzner](https://www.hetzner.com) provides a simple golang API to interact with Hetzner projects. However, no listings of the supported configuration options are available in their documentation. Developers are thus required to request this information themselves whenever needed. This is an annoyance that is trivial to solve, but it's stupid to make people solve these problems multiple times. So I decided to put this little repo together for anyone who has use for it. 

NOTE: if you find that the json files in this repo have become outdated, ping me and I'll update the files as soon as possible. Or you can just fork the repo, and run the update yourself. Hopefully it's helpful! If not (yet), any contributions are welcome.

## How to use this project

### Quick lookup

> Just view the json files directly here on Github, or in any text editor. 

Currently included config options
- images
- locations
- datacenters
- servertypes

### Outdated listing(s), or you need more specific information? 

> Run a local build, and customize as needed. 

1. Clone this git repo to a local directory 
2. Provide the API token for a valid Hetzner project in main.go
3. Feel free to customize the code to fit your needs (PRs welcome)
4. `cd` into the working directory, then run `go run main.go`
5. The json files under `pwd/opts/` should now be updated
