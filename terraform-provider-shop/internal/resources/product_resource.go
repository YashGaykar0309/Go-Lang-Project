package resources

import (
	"context"

	"github.com/YashGaykar0309/terraform-provider-shop/internal/client"
	"github.com/YashGaykar0309/terraform-provider-shop/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProductResource struct {
	client *client.Client
}

func NewProductResource() resource.Resource {
	return &ProductResource{}
}

func (r *ProductResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = "shop_product"
}

func (r *ProductResource) Schema(
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
			"vendor_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *ProductResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *ProductResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan models.ProductModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	productReq := client.Product{
		Name:     plan.Name.ValueString(),
		Price:    plan.Price.ValueFloat64(),
		VendorID: plan.VendorID.ValueString(),
	}

	product, err := r.client.CreateProduct(ctx, productReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating product",
			err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(product.ProductID)

	resp.State.Set(ctx, plan)
}

func (r *ProductResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state models.ProductModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	product, err := r.client.GetProductByID(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading product",
			err.Error(),
		)
		return
	}

	// Product deleted outside Terraform
	if product == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	state.Name = types.StringValue(product.Name)
	state.Price = types.Float64Value(product.Price)
	state.VendorID = types.StringValue(product.VendorID)

	resp.State.Set(ctx, &state)
}

func (r *ProductResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan models.ProductModel
	var state models.ProductModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	updateReq := client.Product{
		ProductID: state.ID.ValueString(),
		Name:      plan.Name.ValueString(),
		Price:     plan.Price.ValueFloat64(),
		VendorID:  plan.VendorID.ValueString(),
	}

	err := r.client.UpdateProduct(ctx, state.ID.ValueString(), updateReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating product",
			err.Error(),
		)
		return
	}

	plan.ID = state.ID

	resp.State.Set(ctx, plan)
}

func (r *ProductResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state models.ProductModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteProduct(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting product",
			err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}
