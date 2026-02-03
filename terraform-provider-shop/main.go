package main

import (
	"context"
	"log"

	"github.com/YashGaykar0309/terraform-provider-shop/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	err := providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/yashgaykar/shop",
	})

	if err != nil {
		log.Fatal(err)
	}
}
