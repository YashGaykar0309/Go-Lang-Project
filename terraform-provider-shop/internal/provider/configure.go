package provider

import (
	"context"
	"strings"

	"github.com/YashGaykar0309/terraform-provider-shop/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type providerModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *ShopProvider) Configure(
	ctx context.Context,
	req provider.ConfigureRequest,
	resp *provider.ConfigureResponse,
) {
	var config providerModel

	// Read provider configuration
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate endpoint
	if config.Endpoint.IsUnknown() || config.Endpoint.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Shop API Endpoint",
			"The provider requires an endpoint to communicate with the Shop Application.",
		)
		return
	}

	endpoint := strings.TrimRight(config.Endpoint.ValueString(), "/")

	// Create API client
	shopClient := client.New(endpoint)

	// Make client available to all resources
	resp.ResourceData = shopClient
	resp.DataSourceData = shopClient
}
