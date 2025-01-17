package zia

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zscaler/zscaler-sdk-go/zia/services/adminuserrolemgmt"
)

func dataSourceAdminRoles() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAdminRolesRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Admin role Id.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the admin role.",
			},
			"rank": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Admin rank of this admin role. This is applicable only when admin rank is enabled in the advanced settings. Default value is 7 (the lowest rank). The assigned admin rank determines the roles or admin users this user can manage, and which rule orders this admin can access.",
			},
			"policy_access": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Policy access permission.",
			},
			"dashboard_access": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Dashboard access permission.",
			},
			"report_access": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Report access permission.",
			},
			"analysis_access": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Insights logs access permission.",
			},
			"username_access": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Username access permission. When set to NONE, the username will be obfuscated.",
			},
			"admin_acct_access": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Admin and role management access permission.",
			},
			"is_auditor": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this is an auditor role.",
			},
			"permissions": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of functional areas to which this role has access. This attribute is subject to change.",
			},
			"is_non_editable": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether or not this admin user is editable/deletable.",
			},
			"logs_limit": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Log range limit.",
			},
			"role_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The admin role type. ()This attribute is subject to change.)",
			},
		},
	}
}

func dataSourceAdminRolesRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	var resp *adminuserrolemgmt.AdminRoles
	id, ok := getIntFromResourceData(d, "id")
	if ok {
		log.Printf("[INFO] Getting data for admin role id: %d\n", id)
		res, err := zClient.adminuserrolemgmt.Get(id)
		if err != nil {
			return err
		}
		resp = res
	}
	name, _ := d.Get("name").(string)
	if resp == nil && name != "" {
		log.Printf("[INFO] Getting data for admin role name: %s\n", name)
		res, err := zClient.adminuserrolemgmt.GetByName(name)
		if err != nil {
			return err
		}
		resp = res
	}

	if resp != nil {
		d.SetId(fmt.Sprintf("%d", resp.ID))
		_ = d.Set("rank", resp.Rank)
		_ = d.Set("name", resp.Name)
		_ = d.Set("policy_access", resp.PolicyAccess)
		_ = d.Set("dashboard_access", resp.DashboardAccess)
		_ = d.Set("report_access", resp.ReportAccess)
		_ = d.Set("is_auditor", resp.IsAuditor)
		_ = d.Set("analysis_access", resp.AnalysisAccess)
		_ = d.Set("username_access", resp.UsernameAccess)
		_ = d.Set("admin_acct_access", resp.AdminAcctAccess)
		_ = d.Set("is_auditor", resp.IsAuditor)
		_ = d.Set("permissions", resp.Permissions)
		_ = d.Set("is_non_editable", resp.IsNonEditable)
		_ = d.Set("logs_limit", resp.LogsLimit)
		_ = d.Set("role_type", resp.RoleType)

	} else {
		return fmt.Errorf("couldn't find any admin role name '%s' or id '%d'", name, id)
	}

	return nil
}
