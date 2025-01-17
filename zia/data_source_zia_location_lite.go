package zia

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zscaler/zscaler-sdk-go/zia/services/locationmanagement/locationlite"
)

func dataSourceLocationLite() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLocationLiteRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tz": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kerberos_auth": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"digest_auth_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"xff_forward_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"aup_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"caution_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"aup_block_internet_until_accepted": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"aup_force_ssl_inspection": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"surrogate_ip": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"zapp_ssl_scan_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"surrogate_ip_enforced_for_known_browsers": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"other_sub_location": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"other6_sub_location": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ofw_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ips_control": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ipv6_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ec_location": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceLocationLiteRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	var resp *locationlite.LocationLite
	id, ok := getIntFromResourceData(d, "id")
	if ok {
		log.Printf("[INFO] Getting data for location id: %d\n", id)
		res, err := zClient.locationlite.GetLocationLiteID(id)
		if err != nil {
			return err
		}
		resp = res
	}
	name, _ := d.Get("name").(string)
	if resp == nil && name != "" {
		log.Printf("[INFO] Getting data for location name: %s\n", name)
		res, err := zClient.locationlite.GetLocationLiteByName(name)
		if err != nil {
			return err
		}
		resp = res
	}

	if resp != nil {
		d.SetId(fmt.Sprintf("%d", resp.ID))
		_ = d.Set("name", resp.Name)
		_ = d.Set("kerberos_auth", resp.KerberosAuth)
		_ = d.Set("digest_auth_enabled", resp.DigestAuthEnabled)
		_ = d.Set("parent_id", resp.ParentID)
		_ = d.Set("tz", resp.TZ)
		_ = d.Set("zapp_ssl_scan_enabled", resp.ZappSSLScanEnabled)
		_ = d.Set("xff_forward_enabled", resp.XFFForwardEnabled)
		_ = d.Set("surrogate_ip", resp.SurrogateIP)
		_ = d.Set("surrogate_ip_enforced_for_known_browsers", resp.SurrogateIPEnforcedForKnownBrowsers)
		_ = d.Set("ofw_enabled", resp.OFWEnabled)
		_ = d.Set("ips_control", resp.IPSControl)
		_ = d.Set("aup_enabled", resp.AUPEnabled)
		_ = d.Set("caution_enabled", resp.CautionEnabled)
		_ = d.Set("aup_block_internet_until_accepted", resp.AUPBlockInternetUntilAccepted)
		_ = d.Set("aup_force_ssl_inspection", resp.AUPForceSSLInspection)
		_ = d.Set("ec_location", resp.ECLocation)
		_ = d.Set("other_sub_location", resp.OtherSubLocation)
		_ = d.Set("other6_sub_location", resp.Other6SubLocation)
		_ = d.Set("ipv6_enabled", resp.IPv6Enabled)

	} else {
		return fmt.Errorf("couldn't find any location with name '%s'", name)
	}

	return nil
}
