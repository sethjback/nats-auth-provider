package natsauth

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &natsAuthProvider{}

func New() provider.Provider {
	return &natsAuthProvider{}
}

type natsAuthProvider struct{}

func (p *natsAuthProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "natsauth"
}

func (p *natsAuthProvider) Schema(_ context.Context, _ provider.SchemaRequest, _ *provider.SchemaResponse) {
}

func (p *natsAuthProvider) Configure(_ context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *natsAuthProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (p *natsAuthProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewOperatorResource,
	}
}
