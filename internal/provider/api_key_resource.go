// Copyright (c) HashiCorp, Inc.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &ApiKeyResource{}

type ApiKeyResource struct {
	data *EmbyProviderData
}

type ApiKeyResourceModel struct {
	App         types.String `tfsdk:"app"`
	AccessToken types.String `tfsdk:"access_token"`
}

// authKeyItem mirrors the JSON shape returned by GET /Auth/Keys.
type authKeyItem struct {
	Id          int    `json:"Id"`
	AppName     string `json:"AppName"`
	AccessToken string `json:"AccessToken"`
	DateCreated string `json:"DateCreated"`
}

type authKeysResponse struct {
	Items            []authKeyItem `json:"Items"`
	TotalRecordCount int           `json:"TotalRecordCount"`
}

func NewApiKeyResource() resource.Resource {
	return &ApiKeyResource{}
}

func (r *ApiKeyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_key"
}

func (r *ApiKeyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Creates an Emby API key for a named application. Destroying this resource deletes the key.",
		Attributes: map[string]schema.Attribute{
			"app": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The application name to associate with the API key.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"access_token": schema.StringAttribute{
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "The generated API access token.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ApiKeyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	data, ok := req.ProviderData.(*EmbyProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *EmbyProviderData, got: %T.", req.ProviderData),
		)
		return
	}
	r.data = data
}

func (r *ApiKeyResource) listKeys(_ context.Context) ([]authKeyItem, error) {
	httpResp, err := r.data.Client.SessionsServiceAPI.GetAuthKeys(r.data.Auth).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	var result authKeysResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse auth keys response: %w", err)
	}
	return result.Items, nil
}

func (r *ApiKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApiKeyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Snapshot existing tokens before creation so we can identify the new one.
	before, err := r.listKeys(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to list API keys", err.Error())
		return
	}
	existingTokens := make(map[string]struct{}, len(before))
	for _, k := range before {
		existingTokens[k.AccessToken] = struct{}{}
	}

	_, err = r.data.Client.SessionsServiceAPI.PostAuthKeys(r.data.Auth).App(plan.App.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to create API key", err.Error())
		return
	}

	after, err := r.listKeys(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to list API keys after creation", err.Error())
		return
	}

	for _, k := range after {
		if _, exists := existingTokens[k.AccessToken]; !exists && k.AppName == plan.App.ValueString() {
			plan.AccessToken = types.StringValue(k.AccessToken)
			resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
			return
		}
	}

	resp.Diagnostics.AddError("Unable to find newly created API key", fmt.Sprintf("No new key found for app %q", plan.App.ValueString()))
}

func (r *ApiKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApiKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	keys, err := r.listKeys(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to list API keys", err.Error())
		return
	}

	for _, k := range keys {
		if k.AccessToken == state.AccessToken.ValueString() {
			state.App = types.StringValue(k.AppName)
			resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
			return
		}
	}

	// Key no longer exists; remove from state.
	resp.State.RemoveResource(ctx)
}

func (r *ApiKeyResource) Update(_ context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
	// All attributes require replacement; update is never called.
}

func (r *ApiKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApiKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.data.Client.SessionsServiceAPI.DeleteAuthKeysByKey(r.data.Auth, state.AccessToken.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to delete API key", err.Error())
		return
	}

	// Verify the key is actually gone before confirming deletion.
	keys, listErr := r.listKeys(ctx)
	if listErr == nil {
		for _, k := range keys {
			if k.AccessToken == state.AccessToken.ValueString() {
				resp.Diagnostics.AddError(
					"API key still exists after delete",
					"The Emby API returned success but the key was found on verification. Retry on next apply.",
				)
				return
			}
		}
	}
}
