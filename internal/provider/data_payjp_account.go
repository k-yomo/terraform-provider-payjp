package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	payjpgo "github.com/payjp/payjp-go/v1"
	"time"
)

func dataPayJPAccount() *schema.Resource {
	return &schema.Resource{
		Description: "A configuration for an Account",
		ReadContext: dataPayJPPlanRead,
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type: schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"merchant": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"merchant_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bank_enabled": {
							Type: schema.TypeBool,
							Computed: true,
						},
						"brands_accepted": {
							Type: schema.TypeString,
							Computed: true,
						},
						"currencies_supported": {
							Type: schema.TypeString,
							Computed: true,
						},
						"default_currency": {
							Type: schema.TypeString,
							Computed: true,
						},
						"business_type": {
							Type: schema.TypeString,
							Computed: true,
						},
						"contact_phone": {
							Type: schema.TypeString,
							Computed: true,
						},
						"country": {
							Type: schema.TypeString,
							Computed: true,
						},
						"charge_type": {
							Type: schema.TypeString,
							Computed: true,
						},
						"product_detail": {
							Type: schema.TypeString,
							Computed: true,
						},
						"product_name": {
							Type: schema.TypeString,
							Computed: true,
						},
						"product_type": {
							Type: schema.TypeString,
							Computed: true,
						},
						"details_submitted": {
							Type: schema.TypeBool,
							Computed: true,
						},
						"live_mode_enabled": {
							Type: schema.TypeBool,
							Computed: true,
						},
						"live_mode_activated_at": {
							Type: schema.TypeString,
							Computed: true,
						},
						"site_published": {
							Type: schema.TypeBool,
							Computed: true,
						},
						"url": {
							Type: schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataPayJPPlanRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	payjpClient := m.(*payjpgo.Service)
	accountRes, err := payjpClient.Account.Retrieve()
	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	}

	d.SetId(accountRes.ID)
	d.Set("account_id", accountRes.ID)
	d.Set("email", accountRes.Email)
	d.Set("created_at", accountRes.CreatedAt.Format(time.RFC3339))
	d.Set("merchant", map[string]interface{}{
		"bank_enabled":           accountRes.Merchant.BankEnabled,
		"brands_accepted":        accountRes.Merchant.BrandsAccepted,
		"currencies_supported":   accountRes.Merchant.CurrenciesSupported,
		"default_currency":       accountRes.Merchant.DefaultCurrency,
		"business_type":          accountRes.Merchant.BusinessType,
		"contact_phone":          accountRes.Merchant.ContactPhone,
		"country":                accountRes.Merchant.Country,
		"charge_type":            accountRes.Merchant.ChargeType,
		"product_detail":         accountRes.Merchant.ProductDetail,
		"product_name":           accountRes.Merchant.ProductName,
		"product_type":           accountRes.Merchant.ProductType,
		"details_submitted":      accountRes.Merchant.DetailsSubmitted,
		"live_mode_enabled":      accountRes.Merchant.LiveModeEnabled,
		"live_mode_activated_at": accountRes.Merchant.LiveModeActivatedAt,
		"site_published":         accountRes.Merchant.SitePublished,
		"url":                    accountRes.Merchant.URL,
	})
	return nil
}
