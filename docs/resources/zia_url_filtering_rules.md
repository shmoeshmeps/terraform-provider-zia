---
subcategory: "URL Filtering Rule"
layout: "zscaler"
page_title: "ZIA:  url_filtering_rules"
description: |-
       Creates and manages a URL Filtering Policy rule.
---
# Resource: zia_url_filtering_rules

The **zia_url_filtering_rules** resource creates and manages a URL filtering rules within the Zscaler Internet Access cloud.

## Example Usage

```hcl
# URL filtering rules
resource "zia_url_filtering_rules" "block_streaming" {
    name                = "Block Streaming"
    description         = "Block Video Streaming."
    state               = "ENABLED"
    action              = "BLOCK"
    order               = 2
    url_categories      = ["ANY"]
    device_trust_levels = ["UNKNOWN_DEVICETRUSTLEVEL", "LOW_TRUST", "MEDIUM_TRUST", "HIGH_TRUST"]
    protocols           = ["ANY_RULE"]
    request_methods     = [ "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "OTHER", "POST", "PUT", "TRACE"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the Firewall Filtering policy rule
* `order` - (Required) Order of execution of rule with respect to other URL Filtering rules
* `protocols` - (Required) Protocol criteria. Supported values: `SMRULEF_ZPA_BROKERS_RULE`, `ANY_RULE`, `TCP_RULE`, `UDP_RULE`, `DOHTTPS_RULE`, `TUNNELSSL_RULE`, `HTTP_PROXY`, `FOHTTP_RULE`, `FTP_RULE`, `HTTPS_RULE`, `HTTP_RULE`, `SSL_RULE`, `TUNNEL_RULE`.

### Optional

* `request_methods` - (Optional) Request method for which the rule must be applied. If not set, rule will be applied to all methods
* `rank` - (Optional) Admin rank of the admin who creates this rule
* `state` - (Optional) Rule State
* `end_user_notification_url` - (Optional) URL of end user notification page to be displayed when the rule is matched. Not applicable if either 'overrideUsers' or 'overrideGroups' is specified.
* `block_override` - (Optional) When set to true, a `BLOCK` action triggered by the rule could be overridden. If true and both overrideGroup and overrideUsers are not set, the `BLOCK` triggered by this rule could be overridden for any users. If block)Override is not set, `BLOCK` action cannot be overridden.
* `time_quota` - (Optional) Time quota in minutes, after which the URL Filtering rule is applied. If not set, no quota is enforced. If a policy rule action is set to `BLOCK`, this field is not applicable.
* `size_quota` - (Optional) Size quota in KB beyond which the URL Filtering rule is applied. If not set, no quota is enforced. If a policy rule action is set to `BLOCK`, this field is not applicable.
* `description` - (Optional) Additional information about the rule
* `validity_start_time` - (Optional) If enforceTimeValidity is set to true, the URL Filtering rule will be valid starting on this date and time.
* `validity_end_time` - (Optional) If `enforceTimeValidity` is set to true, the URL Filtering rule will cease to be valid on this end date and time.
* `validity_time_zone_id` - (Optional) If `enforceTimeValidity` is set to true, the URL Filtering rule date and time will be valid based on this time zone ID.
* `last_modified_time` - (Optional) When the rule was last modified
* `enforce_time_validity` - (Optional) Enforce a set a validity time period for the URL Filtering rule.
* `action` - (Optional) Action taken when traffic matches rule criteria. Supported values: `ANY`, `NONE`, `BLOCK`, `CAUTION`, `ALLOW`, `ICAP_RESPONSE`
* `device_trust_levels` - (Optional) List of device trust levels for which the rule must be applied. This field is applicable for devices that are managed using Zscaler Client Connector. The trust levels are assigned to the devices based on your posture configurations in the Zscaler Client Connector Portal. If no value is set, this field is ignored during the policy evaluation. Supported values: `ANY`, `UNKNOWN_DEVICETRUSTLEVEL`, `LOW_TRUST`, `MEDIUM_TRUST`, `HIGH_TRUST`
* `cipa_rule` - (Optional) If set to true, the CIPA Compliance rule is enabled
* `url_categories` - (Optional) List of URL categories for which rule must be applied

* `locations` - (List of Object) The locations to which the Firewall Filtering policy rule applies
  * `id` - (Optional) Identifier that uniquely identifies an entity

* `groups` - (List of Object) The groups to which the Firewall Filtering policy rule applies
  * `id` - (Optional) Identifier that uniquely identifies an entity

* `departments` - (List of Object) The departments to which the Firewall Filtering policy rule applies
  * `id` - (Optional) Identifier that uniquely identifies an entity

* `users` - (List of Object) The users to which the Firewall Filtering policy rule applies
  * `id` - (Number) Identifier that uniquely identifies an entity

* `time_windows` - (List of Object) The time interval in which the Firewall Filtering policy rule applies
  * `id` - (Optional) Identifier that uniquely identifies an entity

* `override_users` - (List of Object) Name-ID pairs of users for which this rule can be overridden. Applicable only if blockOverride is set to `true`, action is `BLOCK` and overrideGroups is not set.If this overrideUsers is not set, `BLOCK` action can be overridden for any user.
  * `id` - (Optional) Identifier that uniquely identifies an entity

* `override_groups` - (List of Object) Name-ID pairs of users for which this rule can be overridden. Applicable only if blockOverride is set to `true`, action is `BLOCK` and overrideGroups is not set.If this overrideUsers is not set, `BLOCK` action can be overridden for any group.
  * `id` - (Optional) Identifier that uniquely identifies an entity

* `location_groups` - (List of Object) The location groups to which the Firewall Filtering policy rule applies
  * `id` - (Optional) Identifier that uniquely identifies an entity

* `labels`
  * `id` - (Optional) Identifier that uniquely identifies an entity
