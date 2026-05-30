// Copyright (c) HashiCorp, Inc.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &LibraryDataSource{}

type LibraryDataSource struct {
	data *EmbyProviderData
}

type LibraryDataSourceModel struct {
	Name           types.String   `tfsdk:"name"`
	Id             types.String   `tfsdk:"id"`
	ItemId         types.String   `tfsdk:"item_id"`
	Guid           types.String   `tfsdk:"guid"`
	CollectionType types.String   `tfsdk:"collection_type"`
	Locations      []types.String `tfsdk:"locations"`
}

func NewLibraryDataSource() datasource.DataSource {
	return &LibraryDataSource{}
}

func (d *LibraryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_library"
}

func (d *LibraryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Gets a specific library (virtual folder) by name from the Emby server.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The name of the library to retrieve.",
			},
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The library ID.",
			},
			"item_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The internal item ID used for policy bindings (e.g. EnabledFolders).",
			},
			"guid": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The library GUID.",
			},
			"collection_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The collection type (e.g. movies, tvshows, music).",
			},
			"locations": schema.ListAttribute{
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "The file system paths for the library.",
			},
		},
	}
}

func (d *LibraryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(*EmbyProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *EmbyProviderData, got: %T.", req.ProviderData),
		)
		return
	}

	d.data = data
}

func (d *LibraryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LibraryDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, _, err := d.data.Client.LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery(d.data.Auth).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to read libraries", err.Error())
		return
	}

	var libraryName, libraryId string
	if !config.Name.IsNull() && !config.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Multiple identifiers specified",
			"Only one of name or id should be specified to identify the library.",
		)
		return
	}
	if !config.Name.IsNull() {
		libraryName = config.Name.ValueString()
	} else if !config.Id.IsNull() {
		libraryId = config.Id.ValueString()
	}
	var foundLibrary *LibraryDataSourceModel

	for _, folder := range result.GetItems() {
		nameMatches := libraryName != "" && folder.GetName() == libraryName
		idMatches := libraryId != "" && folder.GetId() == libraryId
		if nameMatches || idMatches {
			locations := make([]types.String, 0, len(folder.GetLocations()))
			for _, loc := range folder.GetLocations() {
				locations = append(locations, types.StringValue(loc))
			}

			foundLibrary = &LibraryDataSourceModel{
				Name:           types.StringValue(folder.GetName()),
				Id:             types.StringValue(folder.GetId()),
				ItemId:         types.StringValue(folder.GetItemId()),
				Guid:           types.StringValue(folder.GetGuid()),
				CollectionType: types.StringValue(folder.GetCollectionType()),
				Locations:      locations,
			}
			break
		}
	}

	if foundLibrary == nil {
		var errMsg string
		if libraryName != "" {
			errMsg = fmt.Sprintf("No library found with name %q", libraryName)
		} else if libraryId != "" {
			errMsg = fmt.Sprintf("No library found with ID %q", libraryId)
		} else {
			resp.Diagnostics.AddError(
				"Missing library name or id",
				"Either name or id must be specified",
			)
			return
		}

		resp.Diagnostics.AddError(
			"Library not found",
			errMsg,
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, foundLibrary)...)
}
