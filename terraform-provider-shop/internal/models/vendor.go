package models

import "github.com/hashicorp/terraform-plugin-framework/types"

type VendorModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	EmailAddress types.String `tfsdk:"email_address"`
	PhoneNumber  types.String `tfsdk:"phone_number"`
	Address      types.String `tfsdk:"address"`
	Contact      types.String `tfsdk:"contact"`
}
