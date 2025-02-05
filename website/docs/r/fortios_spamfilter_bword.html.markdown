---
subcategory: "FortiGate SpamFilter"
layout: "fortios"
page_title: "FortiOS: fortios_spamfilter_bword"
description: |-
  Configure AntiSpam banned word list.
---

# fortios_spamfilter_bword
Configure AntiSpam banned word list. Applies to FortiOS Version `<= 6.2.0`.

## Example Usage

```hcl
resource "fortios_spamfilter_bword" "trname" {
  comment = "test"
  fosid   = 1
  name    = "s1"

  entries {
    action       = "clear"
    language     = "western"
    pattern      = "test*patten"
    pattern_type = "wildcard"
    score        = 10
    status       = "enable"
    where        = "subject"
  }
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - (Required) Name of table.
* `comment` - Optional comments.
* `entries` - Spam filter banned word. The structure of `entries` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entries` block supports:

* `status` - Enable/disable status. Valid values: `enable`, `disable`.
* `id` - Banned word entry ID.
* `pattern` - Pattern for the banned word.
* `pattern_type` - Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.
* `action` - Mark spam or good. Valid values: `spam`, `clear`.
* `where` - Component of the email to be scanned. Valid values: `subject`, `body`, `all`.
* `language` - Language for the banned word. Valid values: `western`, `simch`, `trach`, `japanese`, `korean`, `french`, `thai`, `spanish`.
* `score` - Score value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Spamfilter Bword can be imported using any of these accepted formats:
```
$ terraform import fortios_spamfilter_bword.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_spamfilter_bword.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
