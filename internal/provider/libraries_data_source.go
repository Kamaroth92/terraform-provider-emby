package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &LibrariesDataSource{}

type LibrariesDataSource struct {
	data *EmbyProviderData
}

type LibrariesDataSourceModel struct {
	Libraries map[string]LibraryModel `tfsdk:"libraries"`
}

type LibraryModel struct {
	Id             types.String   `tfsdk:"id"`
	Name           types.String   `tfsdk:"name"`
	CollectionType types.String   `tfsdk:"collection_type"`
	Locations      []types.String `tfsdk:"locations"`
}

func NewLibrariesDataSource() datasource.DataSource {
	return &LibrariesDataSource{}
}

func (d *LibrariesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_libraries"
}

func (d *LibrariesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Lists all libraries (virtual folders) on the Emby server.",
		Attributes: map[string]schema.Attribute{
			"libraries": schema.MapNestedAttribute{
				Computed:            true,
				MarkdownDescription: "The libraries keyed by name.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The library ID.",
						},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The library name.",
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
				},
			},
		},
	}
}

func (d *LibrariesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *LibrariesDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	result, _, err := d.data.Client.LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery(d.data.Auth).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to read libraries", err.Error())
		return
	}

	var state LibrariesDataSourceModel
	state.Libraries = make(map[string]LibraryModel)

	for _, folder := range result.GetItems() {
		locations := make([]types.String, 0, len(folder.GetLocations()))
		for _, loc := range folder.GetLocations() {
			locations = append(locations, types.StringValue(loc))
		}

		libraryName := folder.GetName()
		state.Libraries[libraryName] = LibraryModel{
			Id:             types.StringValue(folder.GetId()),
			Name:           types.StringValue(libraryName),
			CollectionType: types.StringValue(folder.GetCollectionType()),
			Locations:      locations,
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
