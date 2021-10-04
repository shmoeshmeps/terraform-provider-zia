package zia

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/willguibr/terraform-provider-zia/gozscaler/client"
	"github.com/willguibr/terraform-provider-zia/gozscaler/trafficforwarding/staticips"
)

func resourceTrafficForwardingStaticIP() *schema.Resource {
	return &schema.Resource{
		Create:   resourceTrafficForwardingStaticIPCreate,
		Read:     resourceTrafficForwardingStaticIPRead,
		Update:   resourceTrafficForwardingStaticIPUpdate,
		Delete:   resourceTrafficForwardingStaticIPDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"static_ip_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the Static IP.",
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"geo_override": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"latitude": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"longitude": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"routable_ip": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"managed_by": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"extensions": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"last_modified_by": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"extensions": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceTrafficForwardingStaticIPCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandTrafficForwardingStaticIP(d)
	log.Printf("[INFO] Creating zia static ip\n%+v\n", req)

	resp, _, err := zClient.staticips.Create(&req)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Created zia static ip request. ID: %v\n", resp)
	d.SetId(strconv.Itoa(resp.ID))
	_ = d.Set("static_ip_id", resp.ID)

	return resourceTrafficForwardingStaticIPRead(d, m)
}

func resourceTrafficForwardingStaticIPRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id, ok := getIntFromResourceData(d, "static_ip_id")
	if !ok {
		return fmt.Errorf("no Traffic Forwarding zia static ip id is set")
	}
	resp, err := zClient.staticips.Get(id)

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing static ip %s from state because it no longer exists in ZIA", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting static ip:\n%+v\n", resp)
	d.SetId(fmt.Sprintf("%d", resp.ID))
	_ = d.Set("static_ip_id", resp.ID)
	_ = d.Set("ip_address", resp.IpAddress)
	_ = d.Set("geo_override", resp.GeoOverride)
	_ = d.Set("latitude", resp.Latitude)
	_ = d.Set("longitude", resp.Longitude)
	_ = d.Set("routable_ip", resp.RoutableIP)

	if err := d.Set("managed_by", flattenStaticManagedBy(resp.ManagedBy)); err != nil {
		return err
	}

	if err := d.Set("last_modified_by", flattenStaticLastModifiedBy(resp.LastModifiedBy)); err != nil {
		return err
	}

	return nil
}

func resourceTrafficForwardingStaticIPUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id, ok := getIntFromResourceData(d, "static_ip_id")
	if !ok {
		log.Printf("[ERROR] static ip ID not set: %v\n", id)
	}

	log.Printf("[INFO] Updating static ip ID: %v\n", id)
	req := expandTrafficForwardingStaticIP(d)

	if _, _, err := zClient.staticips.Update(id, &req); err != nil {
		return err
	}

	return resourceTrafficForwardingStaticIPRead(d, m)
}

func resourceTrafficForwardingStaticIPDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id, ok := getIntFromResourceData(d, "static_ip_id")
	if !ok {
		log.Printf("[ERROR] static ip ID not set: %v\n", id)
	}
	log.Printf("[INFO] Deleting static ip ID: %v\n", (d.Id()))

	if _, err := zClient.staticips.Delete(id); err != nil {
		return err
	}
	d.SetId("")
	log.Printf("[INFO] static ip deleted")
	return nil
}

func expandTrafficForwardingStaticIP(d *schema.ResourceData) staticips.StaticIP {
	id, _ := getIntFromResourceData(d, "static_ip_id")
	result := staticips.StaticIP{
		ID:          id,
		IpAddress:   d.Get("ip_address").(string),
		GeoOverride: d.Get("geo_override").(bool),
		Latitude:    d.Get("latitude").(float64),
		Longitude:   d.Get("longitude").(float64),
		RoutableIP:  d.Get("routable_ip").(bool),
		Comment:     d.Get("comment").(string),
	}
	managedBy := expandStaticIPManagedBy(d)
	if managedBy != nil {
		result.ManagedBy = managedBy
	}

	lastModifiedBy := expandStaticIPLastModifiedBy(d)
	if lastModifiedBy != nil {
		result.LastModifiedBy = lastModifiedBy
	}
	return result
}

func expandStaticIPManagedBy(d *schema.ResourceData) *staticips.ManagedBy {
	managedByObj, ok := d.GetOk("managed_by")
	if !ok {
		return nil
	}
	managed, ok := managedByObj.(*schema.Set)
	if !ok {
		return nil
	}
	if len(managed.List()) > 0 {
		managedObj := managed.List()[0]
		managed, ok := managedObj.(map[string]interface{})
		if !ok {
			return nil
		}
		return &staticips.ManagedBy{
			ID:         managed["id"].(int),
			Name:       managed["name"].(string),
			Extensions: managed["extensions"].(map[string]interface{}),
		}
	}
	return nil
}

func expandStaticIPLastModifiedBy(d *schema.ResourceData) *staticips.LastModifiedBy {
	lastModiedByObj, ok := d.GetOk("last_modified_by")
	if !ok {
		return nil
	}
	modified, ok := lastModiedByObj.(*schema.Set)
	if !ok {
		return nil
	}
	if len(modified.List()) > 0 {
		lastModiedByObj := modified.List()[0]
		modified, ok := lastModiedByObj.(map[string]interface{})
		if !ok {
			return nil
		}
		return &staticips.LastModifiedBy{
			ID:         modified["id"].(int),
			Name:       modified["name"].(string),
			Extensions: modified["extensions"].(map[string]interface{}),
		}
	}
	return nil
}