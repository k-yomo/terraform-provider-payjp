package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	payjpgo "github.com/payjp/payjp-go/v1"
)

func resourcePayJPPlan() *schema.Resource {
	return &schema.Resource{
		Description: "A configuration for a plan",
		CreateContext: resourcePayJPPlanCreate,
		ReadContext:   resourcePayJPPlanRead,
		UpdateContext: resourcePayJPPlanUpdate,
		DeleteContext: resourcePayJPPlanDelete,
		Schema: map[string]*schema.Schema{
			"plan_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"amount": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"currency": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"interval": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"trial_days": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"billing_day": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"metadata": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}
}

func resourcePayJPPlanCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	payjpClient := m.(*payjpgo.Service)

	plan := payjpgo.Plan{
		Amount:   d.Get("amount").(int),
		Currency: d.Get("currency").(string),
		Interval: d.Get("interval").(string),
		Metadata: expandMetadata(d),
	}

	if id, ok := d.GetOk("plan_id"); ok {
		plan.ID = id.(string)
	}
	planName := d.Get("name")
	if _, ok := d.GetOk("name"); ok {
		plan.Name = planName.(string)
	}
	if trialDays, ok := d.GetOk("trial_days"); ok {
		plan.TrialDays = trialDays.(int)
	}
	if bilingDays, ok := d.GetOk("billing_day"); ok {
		plan.BillingDay = bilingDays.(int)
	}

	planRes, err := payjpClient.Plan.Create(plan)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(planRes.ID)

	return nil
}

func resourcePayJPPlanRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	payjpClient := m.(*payjpgo.Service)
	planRes, err := payjpClient.Plan.Retrieve(d.Id())
	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	} else {
		d.Set("plan_id", planRes.ID)
		d.Set("name", planRes.Name)
		d.Set("amount", planRes.Amount)
		d.Set("currency", planRes.Currency)
		d.Set("interval", planRes.Interval)
		d.Set("trial_days", planRes.TrialDays)
		d.Set("billing_day", planRes.BillingDay)
		d.Set("metadata", planRes.Metadata)
	}
	return nil
}

func resourcePayJPPlanUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	payjpClient := m.(*payjpgo.Service)

	if d.HasChange("name") {
		_, err := payjpClient.Plan.Update(d.Id(), d.Get("name").(string))
		return diag.FromErr(err)
	}
	return nil
}

func resourcePayJPPlanDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	payjpClient := m.(*payjpgo.Service)
	if err := payjpClient.Plan.Delete(d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
