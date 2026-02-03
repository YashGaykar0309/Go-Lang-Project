package models

import "github.com/hashicorp/terraform-plugin-framework/types"

type CustomerModel struct {
	ID           types.String `tfsdk:"id"`
	FirstName    types.String `tfsdk:"first_name"`
	LastName     types.String `tfsdk:"last_name"`
	EmailAddress types.String `tfsdk:"email_address"`
	PhoneNumber  types.String `tfsdk:"phone_number"`
	Address      types.String `tfsdk:"address"`
}
