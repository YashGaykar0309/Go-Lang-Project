package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type ShopProvider struct {
	version string
}

func New() provider.Provider {
	return &ShopProvider{
		version: "0.1.0",
	}
}

func (p *ShopProvider) Metadata(
	ctx context.Context,
	req provider.MetadataRequest,
	resp *provider.MetadataResponse,
) {
	resp.TypeName = "shop"
	resp.Version = p.version
}

func (p *ShopProvider) Schema(
	ctx context.Context,
	req provider.SchemaRequest,
	resp *provider.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				Required:    true,
				Description: "Shop Application API endpoint (e.g. http://localhost:8080)",
			},
		},
	}
}

// func (p *ShopProvider) Configure(
// 	ctx context.Context,
// 	req provider.ConfigureRequest,
// 	resp *provider.ConfigureResponse,
// ) {
// 	// NEXT STEP: wire client here
// }

func (p *ShopProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}

func (p *ShopProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}
