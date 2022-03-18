package zia

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/willguibr/terraform-provider-zia/gozscaler/usermanagement"
)

func dataSourceDepartmentManagement() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDepartmentManagementRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Department ID",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Department name",
			},
			"idp_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Identity provider (IdP) ID",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "dditional information about this department",
			},
			"deleted": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceDepartmentManagementRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	var resp *usermanagement.Department
	id, ok := getIntFromResourceData(d, "id")
	if ok {
		log.Printf("[INFO] Getting data for department id: %d\n", id)
		res, err := zClient.usermanagement.GetDepartments(id)
		if err != nil {
			return err
		}
		resp = res
	}
	name, _ := d.Get("name").(string)
	if resp == nil && name != "" {
		log.Printf("[INFO] Getting data for department : %s\n", name)
		res, err := zClient.usermanagement.GetDepartmentsByName(name)
		if err != nil {
			return err
		}
		resp = res
	}

	if resp != nil {
		d.SetId(fmt.Sprintf("%d", resp.ID))
		_ = d.Set("name", resp.Name)
		_ = d.Set("idp_id", resp.IdpID)
		_ = d.Set("comments", resp.Comments)
		_ = d.Set("deleted", resp.Deleted)

	} else {
		return fmt.Errorf("couldn't find any department with name '%s' or id '%d'", name, id)
	}

	return nil
}
