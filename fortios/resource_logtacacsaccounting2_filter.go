// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Settings for TACACS+ accounting events filter.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogTacacsAccounting2Filter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogTacacsAccounting2FilterUpdate,
		Read:   resourceLogTacacsAccounting2FilterRead,
		Update: resourceLogTacacsAccounting2FilterUpdate,
		Delete: resourceLogTacacsAccounting2FilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"login_audit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config_change_audit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cli_cmd_audit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogTacacsAccounting2FilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogTacacsAccounting2Filter(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogTacacsAccounting2Filter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogTacacsAccounting2Filter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogTacacsAccounting2Filter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogTacacsAccounting2Filter")
	}

	return resourceLogTacacsAccounting2FilterRead(d, m)
}

func resourceLogTacacsAccounting2FilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogTacacsAccounting2Filter(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogTacacsAccounting2Filter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogTacacsAccounting2Filter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogTacacsAccounting2Filter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogTacacsAccounting2FilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogTacacsAccounting2Filter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogTacacsAccounting2Filter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogTacacsAccounting2Filter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogTacacsAccounting2Filter resource from API: %v", err)
	}
	return nil
}

func flattenLogTacacsAccounting2FilterLoginAudit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogTacacsAccounting2FilterConfigChangeAudit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogTacacsAccounting2FilterCliCmdAudit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogTacacsAccounting2Filter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("login_audit", flattenLogTacacsAccounting2FilterLoginAudit(o["login-audit"], d, "login_audit", sv)); err != nil {
		if !fortiAPIPatch(o["login-audit"]) {
			return fmt.Errorf("Error reading login_audit: %v", err)
		}
	}

	if err = d.Set("config_change_audit", flattenLogTacacsAccounting2FilterConfigChangeAudit(o["config-change-audit"], d, "config_change_audit", sv)); err != nil {
		if !fortiAPIPatch(o["config-change-audit"]) {
			return fmt.Errorf("Error reading config_change_audit: %v", err)
		}
	}

	if err = d.Set("cli_cmd_audit", flattenLogTacacsAccounting2FilterCliCmdAudit(o["cli-cmd-audit"], d, "cli_cmd_audit", sv)); err != nil {
		if !fortiAPIPatch(o["cli-cmd-audit"]) {
			return fmt.Errorf("Error reading cli_cmd_audit: %v", err)
		}
	}

	return nil
}

func flattenLogTacacsAccounting2FilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogTacacsAccounting2FilterLoginAudit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogTacacsAccounting2FilterConfigChangeAudit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogTacacsAccounting2FilterCliCmdAudit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogTacacsAccounting2Filter(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("login_audit"); ok {
		if setArgNil {
			obj["login-audit"] = nil
		} else {

			t, err := expandLogTacacsAccounting2FilterLoginAudit(d, v, "login_audit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["login-audit"] = t
			}
		}
	}

	if v, ok := d.GetOk("config_change_audit"); ok {
		if setArgNil {
			obj["config-change-audit"] = nil
		} else {

			t, err := expandLogTacacsAccounting2FilterConfigChangeAudit(d, v, "config_change_audit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["config-change-audit"] = t
			}
		}
	}

	if v, ok := d.GetOk("cli_cmd_audit"); ok {
		if setArgNil {
			obj["cli-cmd-audit"] = nil
		} else {

			t, err := expandLogTacacsAccounting2FilterCliCmdAudit(d, v, "cli_cmd_audit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cli-cmd-audit"] = t
			}
		}
	}

	return &obj, nil
}
