---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "zia_location_management Data Source - terraform-provider-zia"
subcategory: ""
description: |-
  
---

# zia_location_management (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **name** (String)

### Read-Only

- **aup_block_internet_until_accepted** (Boolean)
- **aup_enabled** (Boolean)
- **aup_force_ssl_inspection** (Boolean)
- **aup_timeout_in_days** (Number)
- **auth_required** (Boolean)
- **caution_enabled** (Boolean)
- **country** (String)
- **description** (String)
- **display_time_unit** (String)
- **dn_bandwidth** (Number)
- **id** (Number) The ID of this resource.
- **idle_time_in_minutes** (Number)
- **ip_addresses** (List of String)
- **ips_control** (Boolean)
- **ofw_enabled** (Boolean)
- **parent_id** (Number)
- **ports** (String)
- **profile** (String)
- **ssl_scan_enabled** (Boolean)
- **surrogate_ip** (Boolean)
- **surrogate_ip_enforced_for_known_browsers** (Boolean)
- **surrogate_refresh_time_in_minutes** (Number)
- **surrogate_refresh_time_unit** (String)
- **tz** (String)
- **up_bandwidth** (Number)
- **vpn_credentials** (List of Object) (see [below for nested schema](#nestedatt--vpn_credentials))
- **xff_forward_enabled** (Boolean)
- **zapp_ssl_scan_enabled** (Boolean)

<a id="nestedatt--vpn_credentials"></a>
### Nested Schema for `vpn_credentials`

Read-Only:

- **comments** (String)
- **fqdn** (String)
- **id** (Number)
- **location** (List of Object) (see [below for nested schema](#nestedobjatt--vpn_credentials--location))
- **managed_by** (List of Object) (see [below for nested schema](#nestedobjatt--vpn_credentials--managed_by))
- **pre_shared_key** (String)
- **type** (String)

<a id="nestedobjatt--vpn_credentials--location"></a>
### Nested Schema for `vpn_credentials.location`

Read-Only:

- **extensions** (Map of String)
- **id** (Number)
- **name** (String)


<a id="nestedobjatt--vpn_credentials--managed_by"></a>
### Nested Schema for `vpn_credentials.managed_by`

Read-Only:

- **extensions** (Map of String)
- **id** (Number)
- **name** (String)

