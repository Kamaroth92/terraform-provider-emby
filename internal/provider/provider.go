// Copyright (c) HashiCorp, Inc.

package provider

import (
	"context"
	"os"

	embyclient "github.com/Kamaroth92/terraform-provider-emby/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &EmbyProvider{}

type EmbyProvider struct {
	version string
}

type EmbyProviderModel struct {
	Hostname types.String `tfsdk:"hostname"`
	ApiKey   types.String `tfsdk:"api_key"`
}

// EmbyProviderData is passed to data sources and resources via Configure.
type EmbyProviderData struct {
	Client *embyclient.APIClient
	Auth   context.Context
}

func (p *EmbyProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "emby"
	resp.Version = p.version
}

func (p *EmbyProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The Emby provider allows you to interact with your Emby media server.",
		Attributes: map[string]schema.Attribute{
			"hostname": schema.StringAttribute{
				MarkdownDescription: "The base URL of your Emby server (e.g. `http://emby.example.com:8096`).",
				Optional:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "The API key used to authenticate with Emby.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *EmbyProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data EmbyProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Hostname.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("hostname"),
			"Unknown Emby API Host",
			"The provider cannot create the Emby API client as there is an unknown configuration value for the Emby API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the EMBY_HOST environment variable.",
		)
	}

	if data.ApiKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown Emby API Key",
			"The provider cannot create the Emby API client as there is an unknown configuration value for the Emby API key. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the EMBY_API_KEY environment variable.",
		)
	}

	hostname := os.Getenv("EMBY_HOSTNAME")

	apiKey := os.Getenv("EMBY_APIKEY")

	if !data.Hostname.IsNull() {
		hostname = data.Hostname.ValueString()
	}

	if !data.ApiKey.IsNull() {
		apiKey = data.ApiKey.ValueString()
	}

	cfg := embyclient.NewConfiguration()
	cfg.Servers = embyclient.OAPIServerConfigs{
		{URL: hostname + "/emby"},
	}

	client := embyclient.NewAPIClient(cfg)

	authCtx := context.WithValue(context.Background(), embyclient.ContextAPIKeys, map[string]embyclient.APIKey{
		"apikeyauth": {Key: apiKey},
	})

	providerData := &EmbyProviderData{
		Client: client,
		Auth:   authCtx,
	}

	resp.DataSourceData = providerData
	resp.ResourceData = providerData
}

func (p *EmbyProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewUserResource,
		NewUserLibraryAccessResource,
		NewApiKeyResource,
		NewLibraryResource,
	}
}

func (p *EmbyProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewLibrariesDataSource,
		NewLibraryDataSource,
		NewUserDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &EmbyProvider{
			version: version,
		}
	}
}
