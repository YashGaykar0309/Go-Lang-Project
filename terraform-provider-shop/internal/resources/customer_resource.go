package resources

import (
	"context"

	"github.com/YashGaykar0309/terraform-provider-shop/internal/client"
	"github.com/YashGaykar0309/terraform-provider-shop/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomerResource struct {
	client *client.Client
}

func NewCustomerResource() resource.Resource {
	return &CustomerResource{}
}

func (r *CustomerResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = "shop_customer"
}

func (r *CustomerResource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"first_name": schema.StringAttribute{
				Required: true,
			},
			"last_name": schema.StringAttribute{
				Required: true,
			},
			"email_address": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"phone_number": schema.StringAttribute{
				Required: true,
			},
			"address": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *CustomerResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *CustomerResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan models.CustomerModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	customerReq := client.Customer{
		FirstName:    plan.FirstName.ValueString(),
		LastName:     plan.LastName.ValueString(),
		EmailAddress: plan.EmailAddress.ValueString(),
		PhoneNumber:  plan.PhoneNumber.ValueString(),
		Address:      plan.Address.ValueString(),
	}

	customer, err := r.client.CreateCustomer(ctx, customerReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating customer",
			err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(customer.CustomerID)

	resp.State.Set(ctx, plan)
}

func (r *CustomerResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state models.CustomerModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	customer, err := r.client.GetCustomerByID(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading customer",
			err.Error(),
		)
		return
	}

	// Customer deleted outside Terraform
	if customer == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	state.FirstName = types.StringValue(customer.FirstName)
	state.LastName = types.StringValue(customer.LastName)
	state.EmailAddress = types.StringValue(customer.EmailAddress)
	state.PhoneNumber = types.StringValue(customer.PhoneNumber)
	state.Address = types.StringValue(customer.Address)

	resp.State.Set(ctx, &state)
}

func (r *CustomerResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan models.CustomerModel
	var state models.CustomerModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := client.Customer{
		CustomerID:   state.ID.ValueString(),
		FirstName:    plan.FirstName.ValueString(),
		LastName:     plan.LastName.ValueString(),
		EmailAddress: plan.EmailAddress.ValueString(),
		PhoneNumber:  plan.PhoneNumber.ValueString(),
		Address:      plan.Address.ValueString(),
	}

	err := r.client.UpdateCustomer(ctx, state.ID.ValueString(), updateReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating customer",
			err.Error(),
		)
		return
	}

	plan.ID = state.ID

	resp.State.Set(ctx, plan)
}

func (r *CustomerResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state models.CustomerModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteCustomer(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting customer",
			err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}
