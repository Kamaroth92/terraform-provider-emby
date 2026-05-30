// Copyright (c) HashiCorp, Inc.

package provider

import (
	"context"
	"fmt"

	embyclient "github.com/Kamaroth92/terraform-provider-emby/client"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &LibraryResource{}
	_ resource.ResourceWithImportState = &LibraryResource{}
)

type LibraryResource struct {
	data *EmbyProviderData
}

type LibraryResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	CollectionType types.String `tfsdk:"collection_type"`
	Paths          types.List   `tfsdk:"paths"`
	ItemId         types.String `tfsdk:"item_id"`
	Guid           types.String `tfsdk:"guid"`
	Locations      types.List   `tfsdk:"locations"`
}

func NewLibraryResource() resource.Resource {
	return &LibraryResource{}
}

func (r *LibraryResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_library"
}

func (r *LibraryResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Emby library (virtual folder).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The library ID.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The library name.",
			},
			"collection_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The collection type (e.g. movies, tvshows, music, books, homevideos, musicvideos, mixed). Changing this forces the library to be recreated.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"paths": schema.ListAttribute{
				Required:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "The file system paths for the library. At least one path is required.",
			},
			"item_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The internal item ID used for policy bindings (e.g. EnabledFolders).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"guid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The library GUID.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"locations": schema.ListAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "The file system paths for the library, as reported by the server.",
			},
		},
	}
}

func (r *LibraryResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *LibraryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan LibraryResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var paths []string
	resp.Diagnostics.Append(plan.Paths.ElementsAs(ctx, &paths, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := embyclient.NewLibraryAddVirtualFolder()
	createReq.SetName(plan.Name.ValueString())
	createReq.SetCollectionType(plan.CollectionType.ValueString())
	createReq.SetPaths(paths)
	createReq.SetRefreshLibrary(true)

	httpResp, err := r.data.Client.LibraryStructureServiceAPI.PostLibraryVirtualfolders(r.data.Auth).
		LibraryAddVirtualFolder(*createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to create library", err.Error())
		return
	}
	if httpResp != nil && httpResp.StatusCode >= 300 {
		resp.Diagnostics.AddError("Unexpected status creating library", fmt.Sprintf("HTTP %d", httpResp.StatusCode))
		return
	}

	// Find the created library by name to get its computed fields.
	found := r.findLibraryByName(plan.Name.ValueString(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if found == nil {
		resp.Diagnostics.AddError("Library not found after creation", fmt.Sprintf("Library %q was created but could not be found in the library list.", plan.Name.ValueString()))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, found)...)
}

func (r *LibraryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state LibraryResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	found := r.findLibraryById(state.Id.ValueString(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if found == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, found)...)
}

func (r *LibraryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan LibraryResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state LibraryResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = state.Id

	// Handle name change.
	if plan.Name.ValueString() != state.Name.ValueString() {
		renameReq := embyclient.NewLibraryRenameVirtualFolder()
		renameReq.SetId(plan.Id.ValueString())
		renameReq.SetNewName(plan.Name.ValueString())

		_, err := r.data.Client.LibraryStructureServiceAPI.PostLibraryVirtualfoldersName(r.data.Auth).
			LibraryRenameVirtualFolder(*renameReq).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Unable to rename library", err.Error())
			return
		}
	}

	// Handle path changes.
	var planPaths, statePaths []string
	resp.Diagnostics.Append(plan.Paths.ElementsAs(ctx, &planPaths, false)...)
	resp.Diagnostics.Append(state.Paths.ElementsAs(ctx, &statePaths, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	statePathSet := make(map[string]bool, len(statePaths))
	for _, p := range statePaths {
		statePathSet[p] = true
	}
	planPathSet := make(map[string]bool, len(planPaths))
	for _, p := range planPaths {
		planPathSet[p] = true
	}

	// Add paths in plan but not in state.
	for _, p := range planPaths {
		if !statePathSet[p] {
			addReq := embyclient.NewLibraryAddMediaPath()
			addReq.SetId(plan.Id.ValueString())
			addReq.SetPath(p)
			addReq.SetRefreshLibrary(true)

			_, err := r.data.Client.LibraryStructureServiceAPI.PostLibraryVirtualfoldersPaths(r.data.Auth).
				LibraryAddMediaPath(*addReq).Execute()
			if err != nil {
				resp.Diagnostics.AddError("Unable to add library path", err.Error())
				return
			}
		}
	}

	// Remove paths in state but not in plan.
	for _, p := range statePaths {
		if !planPathSet[p] {
			removeReq := embyclient.NewLibraryRemoveMediaPath()
			removeReq.SetId(plan.Id.ValueString())
			removeReq.SetPath(p)
			removeReq.SetRefreshLibrary(true)

			_, err := r.data.Client.LibraryStructureServiceAPI.PostLibraryVirtualfoldersPathsDelete(r.data.Auth).
				LibraryRemoveMediaPath(*removeReq).Execute()
			if err != nil {
				resp.Diagnostics.AddError("Unable to remove library path", err.Error())
				return
			}
		}
	}

	// Read back the updated state.
	found := r.findLibraryById(plan.Id.ValueString(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if found == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, found)...)
}

func (r *LibraryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state LibraryResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	deleteReq := embyclient.NewLibraryRemoveVirtualFolder()
	deleteReq.SetId(state.Id.ValueString())

	_, err := r.data.Client.LibraryStructureServiceAPI.PostLibraryVirtualfoldersDelete(r.data.Auth).
		LibraryRemoveVirtualFolder(*deleteReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to delete library", err.Error())
		return
	}
}

func (r *LibraryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LibraryResource) findLibraryByName(name string, diags *diag.Diagnostics) *LibraryResourceModel {
	result, _, err := r.data.Client.LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery(r.data.Auth).Execute()
	if err != nil {
		diags.AddError("Unable to read libraries", err.Error())
		return nil
	}

	for _, folder := range result.GetItems() {
		if folder.GetName() == name {
			return mapVirtualFolderToState(&folder)
		}
	}
	return nil
}

func (r *LibraryResource) findLibraryById(id string, diags *diag.Diagnostics) *LibraryResourceModel {
	result, _, err := r.data.Client.LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery(r.data.Auth).Execute()
	if err != nil {
		diags.AddError("Unable to read libraries", err.Error())
		return nil
	}

	for _, folder := range result.GetItems() {
		if folder.GetId() == id {
			return mapVirtualFolderToState(&folder)
		}
	}
	return nil
}

func mapVirtualFolderToState(folder *embyclient.VirtualFolderInfo) *LibraryResourceModel {
	locations := make([]types.String, 0, len(folder.GetLocations()))
	paths := make([]types.String, 0, len(folder.GetLocations()))
	for _, loc := range folder.GetLocations() {
		locations = append(locations, types.StringValue(loc))
		paths = append(paths, types.StringValue(loc))
	}

	return &LibraryResourceModel{
		Id:             types.StringValue(folder.GetId()),
		Name:           types.StringValue(folder.GetName()),
		CollectionType: types.StringValue(folder.GetCollectionType()),
		Paths:          types.ListValueMust(types.StringType, toAttrValues(paths)),
		ItemId:         types.StringValue(folder.GetItemId()),
		Guid:           types.StringValue(folder.GetGuid()),
		Locations:      types.ListValueMust(types.StringType, toAttrValues(locations)),
	}
}

func toAttrValues(strs []types.String) []attr.Value {
	vals := make([]attr.Value, len(strs))
	for i, s := range strs {
		vals[i] = s
	}
	return vals
}
