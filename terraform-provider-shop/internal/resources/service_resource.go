package resources

import (
	"context"

	"github.com/YashGaykar0309/terraform-provider-shop/internal/client"
	"github.com/YashGaykar0309/terraform-provider-shop/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ServiceResource struct {
	client *client.Client
}

func NewServiceResource() resource.Resource {
	return &ServiceResource{}
}

func (r *ServiceResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = "shop_service"
}

func (r *ServiceResource) Schema(
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
			"price": schema.Float64Attribute{
				Required: true,
			},
		},
	}
}

func (r *ServiceResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *ServiceResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan models.ServiceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	serviceReq := client.Service{
		Name:  plan.Name.ValueString(),
		Price: plan.Price.ValueFloat64(),
	}

	service, err := r.client.CreateService(ctx, serviceReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating service",
			err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(service.ServiceID)

	resp.State.Set(ctx, plan)
}

func (r *ServiceResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state models.ServiceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	service, err := r.client.GetServiceByID(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading service",
			err.Error(),
		)
		return
	}

	// Service deleted outside Terraform
	if service == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	state.Name = types.StringValue(service.Name)
	state.Price = types.Float64Value(service.Price)

	resp.State.Set(ctx, &state)
}

func (r *ServiceResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan models.ServiceModel
	var state models.ServiceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := client.Service{
		ServiceID: state.ID.ValueString(),
		Name:      plan.Name.ValueString(),
		Price:     plan.Price.ValueFloat64(),
	}

	err := r.client.UpdateService(ctx, state.ID.ValueString(), updateReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating service",
			err.Error(),
		)
		return
	}

	plan.ID = state.ID

	resp.State.Set(ctx, plan)
}

func (r *ServiceResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state models.ServiceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteService(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting service",
			err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}
