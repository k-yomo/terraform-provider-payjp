package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/k-yomo/terraform-provider-payjp/pkg/httputil"
	payjpgo "github.com/payjp/payjp-go/v1"
	"net/http"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"api_key": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("PAYJP_API_KEY", nil),
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"payjp_plan": resourcePayJPPlan(),
			},
			DataSourcesMap: map[string]*schema.Resource{
				"payjp_account": dataPayJPAccount(),
			},
		}
		p.ConfigureContextFunc = configure(version, p)
		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		userAgent := p.UserAgent("terraform-provider-payjp", version)
		httpClient := &http.Client{Transport: httputil.NewAddHeaderTransport(nil, map[string]string{"User-Agent": userAgent})}
		payjpClient := payjpgo.New(d.Get("api_key").(string), httpClient)
		return payjpClient, nil
	}
}

type AddHeaderTransport struct {
	T http.RoundTripper
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "go")
	return adt.T.RoundTrip(req)
}
