package main

import (
	"github.com/rajshreepardeshi/terraform/restapi"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return restapi.Provider()
		},
	})
}
