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

type VendorResource struct {
	client *client.Client
}

func NewVendorResource() resource.Resource {
	return &VendorResource{}
}

func (r *VendorResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = "shop_vendor"
}

func (r *VendorResource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
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
			"contact": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *VendorResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *VendorResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan models.VendorModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	vendorReq := client.Vendor{
		Name:         plan.Name.ValueString(),
		EmailAddress: plan.EmailAddress.ValueString(),
		PhoneNumber:  plan.PhoneNumber.ValueString(),
		Address:      plan.Address.ValueString(),
		Contact:      plan.Contact.ValueString(),
	}

	vendor, err := r.client.CreateVendor(ctx, vendorReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating vendor",
			err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(vendor.VendorID)

	resp.State.Set(ctx, plan)
}

func (r *VendorResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state models.VendorModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	vendor, err := r.client.GetVendorByID(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading vendor",
			err.Error(),
		)
		return
	}

	// Vendor deleted outside Terraform
	if vendor == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	state.Name = types.StringValue(vendor.Name)
	state.EmailAddress = types.StringValue(vendor.EmailAddress)
	state.PhoneNumber = types.StringValue(vendor.PhoneNumber)
	state.Address = types.StringValue(vendor.Address)
	state.Contact = types.StringValue(vendor.Contact)

	resp.State.Set(ctx, &state)
}

func (r *VendorResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan models.VendorModel
	var state models.VendorModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := client.Vendor{
		VendorID:     state.ID.ValueString(),
		Name:         plan.Name.ValueString(),
		EmailAddress: plan.EmailAddress.ValueString(),
		PhoneNumber:  plan.PhoneNumber.ValueString(),
		Address:      plan.Address.ValueString(),
		Contact:      plan.Contact.ValueString(),
	}

	err := r.client.UpdateVendor(ctx, state.ID.ValueString(), updateReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating vendor",
			err.Error(),
		)
		return
	}

	plan.ID = state.ID

	resp.State.Set(ctx, plan)
}

func (r *VendorResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state models.VendorModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteVendor(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting vendor",
			err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}
