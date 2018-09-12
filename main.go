package main

import (
	"math/rand"
	"time"

	"github.com/hashicorp/terraform/plugin"
	"github.com/neodino/terraform-provider-libvirt/libvirt"
)

func main() {
	defer libvirt.CleanupLibvirtConnections()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: libvirt.Provider,
	})
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
