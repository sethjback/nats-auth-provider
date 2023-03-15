package natsauth

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nkeys"
)

var _ resource.Resource = &OperatorResource{}
var _ resource.ResourceWithImportState = &OperatorResource{}

func NewOperatorResource() resource.Resource {
	return &OperatorResource{}
}

// OperatorResource defines the resource implementation.
type OperatorResource struct {
}

// OperatorResourceModel describes the resource data model.
type OperatorResourceModel struct {
	Name     types.String `tfsdk:"name"`
	Operator AccountModel `tfsdk:"operator"`
	// SysAccount AccountModel `tfsdk:"system_account"`
}

type AccountModel struct {
	ID  types.String `tfsdk:"id"`
	JWT types.String `tfsdk:"jwt"`
	// RootKey NatsKeyModel `tfsdk:"root_key"`
	// SigningKeys []t `tfsdk:"signing_keys"`
}

type NatsKeyModel struct {
	ID   types.String `tfsdk:"id"`
	Seed types.String `tfsdk:"seed"`
}

func (r *OperatorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_operator"
}

func (r *OperatorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "NATS Operator resource",
		Description:         "NATS Operator resource",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "Operator name",
				Computed:    false,
				Required:    true,
			},
			"operator": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Operator ID",
						Computed:    true,
					},
					"jwt": schema.StringAttribute{
						Description: "JWT for the NATS Operator",
						Computed:    true,
					},
					// "root_key": schema.SingleNestedAttribute{
					// 	Computed:    true,
					// 	Description: "Root key for the operator account",
					// 	Attributes: map[string]schema.Attribute{
					// 		"id": schema.StringAttribute{
					// 			Description: "Key ID",
					// 			Computed:    true,
					// 		},
					// 		"seed": schema.StringAttribute{
					// 			Description: "Key Seed",
					// 			Computed:    true,
					// 			Sensitive:   true,
					// 		},
					// 	},
					// },
					// "signing_keys": schema.ListNestedAttribute{
					// 	Computed:    true,
					// 	Description: "Signing keys for the operator",
					// 	NestedObject: schema.NestedAttributeObject{
					// 		Attributes: map[string]schema.Attribute{
					// 			"id": schema.StringAttribute{
					// 				Description: "Key ID",
					// 				Computed:    true,
					// 			},
					// 			"seed": schema.StringAttribute{
					// 				Description: "Key Seed",
					// 				Computed:    true,
					// 				Sensitive:   true,
					// 			},
					// 		},
					// 	},
					// },
				},
			},
			// "system_account": schema.SingleNestedAttribute{
			// 	Computed: true,
			// 	Attributes: map[string]schema.Attribute{
			// 		"id": schema.StringAttribute{
			// 			Description: "System ID",
			// 			Computed:    true,
			// 		},
			// 		"jwt": schema.StringAttribute{
			// 			Description: "JWT for the NATS SYS account",
			// 			Computed:    true,
			// 		},
			// 		"root_key": schema.SingleNestedAttribute{
			// 			Computed:    true,
			// 			Description: "Root key for the system account",
			// 			Attributes: map[string]schema.Attribute{
			// 				"id": schema.StringAttribute{
			// 					Description: "Key ID",
			// 					Computed:    true,
			// 				},
			// 				"seed": schema.StringAttribute{
			// 					Description: "Key Seed",
			// 					Computed:    true,
			// 					Sensitive:   true,
			// 				},
			// 			},
			// 		},
			// 		// "signing_keys": schema.ListNestedAttribute{
			// 		// 	Computed:    true,
			// 		// 	Description: "Signing keys for the system account",
			// 		// 	NestedObject: schema.NestedAttributeObject{
			// 		// 		Attributes: map[string]schema.Attribute{
			// 		// 			"id": schema.StringAttribute{
			// 		// 				Description: "Key ID",
			// 		// 				Computed:    true,
			// 		// 			},
			// 		// 			"seed": schema.StringAttribute{
			// 		// 				Description: "Key Seed",
			// 		// 				Computed:    true,
			// 		// 				Sensitive:   true,
			// 		// 			},
			// 		// 		},
			// 		// 	},
			// 		// },
			// 	},
			// },
		},
	}
}

func (r *OperatorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// NOOP
}

func (r *OperatorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *OperatorResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Operator Keys
	tflog.Info(ctx, "creating operator")

	okr, _ := nkeys.CreateOperator()
	okp, _ := okr.PublicKey()
	// oks, _ := okr.Seed()

	// data.Operator.ID = types.StringValue(okp)
	// data.Operator.RootKey.ID = types.StringValue(okp)
	// data.Operator.RootKey.Seed = types.StringValue(string(oks))

	var oksig []nkeys.KeyPair
	tflog.Info(ctx, "creating operator signing keys")
	for i := 0; i < 2; i++ {
		// k, _ := nkeys.CreateOperator()
		// kp, _ := k.PublicKey()
		// ks, _ := k.Seed()
		// oksig = append(oksig, k)
		// data.Operator.SigningKeys = append(data.Operator.SigningKeys, NatsKeyModel{ID: types.StringValue(kp), Seed: types.StringValue(string(ks))})
	}

	// System Keys
	tflog.Info(ctx, "creating sys account")

	// sysr, _ := nkeys.CreateAccount()
	// sysp, _ := sysr.PublicKey()
	// syss, _ := sysr.Seed()

	// data.SysAccount.ID = types.StringValue(sysp)
	// data.SysAccount.RootKey.ID = types.StringValue(sysp)
	// data.SysAccount.RootKey.Seed = types.StringValue(string(syss))
	// var syssig []nkeys.KeyPair
	// tflog.Info(ctx, "creating sys account signing keys")
	// for i := 0; i < 2; i++ {
	// 	// k, _ := nkeys.CreateAccount()
	// 	// kp, _ := k.PublicKey()
	// 	// ks, _ := k.Seed()
	// 	// syssig = append(syssig, k)
	// 	// data.SysAccount.SigningKeys = append(data.SysAccount.SigningKeys, NatsKeyModel{ID: types.StringValue(kp), Seed: types.StringValue(string(ks))})
	// }

	// Claims

	// Operator
	tflog.Info(ctx, "creating operator jwt")

	oc := jwt.NewOperatorClaims(okp)
	oc.Name = data.Name.String()
	oc.OperatorServiceURLs.Add("nats://0.0.0.0:4222")
	// oc.SystemAccount = sysp
	oc.Operator.StrictSigningKeyUsage = true
	for _, sigKey := range oksig {
		kp, _ := sigKey.PublicKey()
		oc.SigningKeys.Add(kp)
	}
	ojwt, _ := oc.Encode(okr)
	data.Operator.JWT = types.StringValue(ojwt)

	// SYS Account
	// tflog.Info(ctx, "creating SYS account jwt")

	// sysc := jwt.NewAccountClaims(sysp)
	// sysc.Name = "SYS"
	// sysc.Exports = jwt.Exports{&jwt.Export{
	// 	Name:                 "account-monitoring-services",
	// 	Subject:              "$SYS.REQ.ACCOUNT.*.*",
	// 	Type:                 jwt.Service,
	// 	ResponseType:         jwt.ResponseTypeStream,
	// 	AccountTokenPosition: 4,
	// 	Info: jwt.Info{
	// 		Description: `Request account specific monitoring services for: SUBSZ, CONNZ, LEAFZ, JSZ and INFO`,
	// 		InfoURL:     "https://docs.nats.io/nats-server/configuration/sys_accounts",
	// 	},
	// }, &jwt.Export{
	// 	Name:                 "account-monitoring-streams",
	// 	Subject:              "$SYS.ACCOUNT.*.>",
	// 	Type:                 jwt.Stream,
	// 	AccountTokenPosition: 3,
	// 	Info: jwt.Info{
	// 		Description: `Account specific monitoring stream`,
	// 		InfoURL:     "https://docs.nats.io/nats-server/configuration/sys_accounts",
	// 	},
	// }}
	// for _, sigKey := range syssig {
	// 	kp, _ := sigKey.PublicKey()
	// 	sysc.SigningKeys.Add(kp)
	// }
	// sysjwt, _ := sysc.Encode(oksig[0])
	// data.SysAccount.JWT = types.StringValue(sysjwt)

	tflog.Trace(ctx, "created operator")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OperatorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *OperatorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OperatorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *OperatorResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OperatorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *OperatorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *OperatorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
