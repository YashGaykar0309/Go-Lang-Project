package models

import "github.com/hashicorp/terraform-plugin-framework/types"

type ServiceModel struct {
	ID    types.String  `tfsdk:"id"`
	Name  types.String  `tfsdk:"name"`
	Price types.Float64 `tfsdk:"price"`
}
