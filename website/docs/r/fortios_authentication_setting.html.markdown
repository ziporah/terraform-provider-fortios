---
subcategory: "FortiGate Authentication"
layout: "fortios"
page_title: "FortiOS: fortios_authentication_setting"
description: |-
  Configure authentication setting.
---

# fortios_authentication_setting
Configure authentication setting.

## Example Usage

```hcl
resource "fortios_authentication_setting" "trname" {
  auth_https              = "enable"
  captive_portal_ip       = "0.0.0.0"
  captive_portal_ip6      = "::"
  captive_portal_port     = 7830
  captive_portal_ssl_port = 7831
  captive_portal_type     = "fqdn"
}
```

## Argument Reference

The following arguments are supported:

* `active_auth_scheme` - Active authentication method (scheme name).
* `sso_auth_scheme` - Single-Sign-On authentication method (scheme name).
* `captive_portal_type` - Captive portal type. Valid values: `fqdn`, `ip`.
* `captive_portal_ip` - Captive portal IP address.
* `captive_portal_ip6` - Captive portal IPv6 address.
* `captive_portal` - Captive portal host name.
* `captive_portal6` - IPv6 captive portal host name.
* `cert_auth` - Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable`, `disable`.
* `cert_captive_portal` - Certificate captive portal host name.
* `cert_captive_portal_ip` - Certificate captive portal IP address.
* `cert_captive_portal_port` - Certificate captive portal port number (1 - 65535, default = 7832).
* `captive_portal_port` - Captive portal port number (1 - 65535, default = 7830).
* `auth_https` - Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
* `captive_portal_ssl_port` - Captive portal SSL port number (1 - 65535, default = 7831).
* `user_cert_ca` - CA certificate used for client certificate verification. The structure of `user_cert_ca` block is documented below.
* `dev_range` - Address range for the IP based device query. The structure of `dev_range` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `user_cert_ca` block supports:

* `name` - CA certificate list.

The `dev_range` block supports:

* `name` - Address name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Authentication Setting can be imported using any of these accepted formats:
```
$ terraform import fortios_authentication_setting.labelname AuthenticationSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_authentication_setting.labelname AuthenticationSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
