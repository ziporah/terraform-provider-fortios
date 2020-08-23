// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPasswordPolicyUpdate,
		Read:   resourceSystemPasswordPolicyRead,
		Update: resourceSystemPasswordPolicyUpdate,
		Delete: resourceSystemPasswordPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apply_to": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"minimum_length": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(8, 128),
				Optional: true,
				Computed: true,
			},
			"min_lower_case_letter": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional: true,
				Computed: true,
			},
			"min_upper_case_letter": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional: true,
				Computed: true,
			},
			"min_non_alphanumeric": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional: true,
				Computed: true,
			},
			"min_number": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional: true,
				Computed: true,
			},
			"change_4_characters": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expire_status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expire_day": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 999),
				Optional: true,
				Computed: true,
			},
			"reuse_password": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}


func resourceSystemPasswordPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemPasswordPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemPasswordPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemPasswordPolicy(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemPasswordPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPasswordPolicy")
	}

	return resourceSystemPasswordPolicyRead(d, m)
}

func resourceSystemPasswordPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemPasswordPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPasswordPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPasswordPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemPasswordPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemPasswordPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPasswordPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemPasswordPolicy resource from API: %v", err)
	}
	return nil
}


func flattenSystemPasswordPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyApplyTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMinimumLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMinLowerCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMinUpperCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMinNonAlphanumeric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMinNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyChange4Characters(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyExpireStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyExpireDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyReusePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemPasswordPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("status", flattenSystemPasswordPolicyStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("apply_to", flattenSystemPasswordPolicyApplyTo(o["apply-to"], d, "apply_to")); err != nil {
		if !fortiAPIPatch(o["apply-to"]) {
			return fmt.Errorf("Error reading apply_to: %v", err)
		}
	}

	if err = d.Set("minimum_length", flattenSystemPasswordPolicyMinimumLength(o["minimum-length"], d, "minimum_length")); err != nil {
		if !fortiAPIPatch(o["minimum-length"]) {
			return fmt.Errorf("Error reading minimum_length: %v", err)
		}
	}

	if err = d.Set("min_lower_case_letter", flattenSystemPasswordPolicyMinLowerCaseLetter(o["min-lower-case-letter"], d, "min_lower_case_letter")); err != nil {
		if !fortiAPIPatch(o["min-lower-case-letter"]) {
			return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
		}
	}

	if err = d.Set("min_upper_case_letter", flattenSystemPasswordPolicyMinUpperCaseLetter(o["min-upper-case-letter"], d, "min_upper_case_letter")); err != nil {
		if !fortiAPIPatch(o["min-upper-case-letter"]) {
			return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
		}
	}

	if err = d.Set("min_non_alphanumeric", flattenSystemPasswordPolicyMinNonAlphanumeric(o["min-non-alphanumeric"], d, "min_non_alphanumeric")); err != nil {
		if !fortiAPIPatch(o["min-non-alphanumeric"]) {
			return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
		}
	}

	if err = d.Set("min_number", flattenSystemPasswordPolicyMinNumber(o["min-number"], d, "min_number")); err != nil {
		if !fortiAPIPatch(o["min-number"]) {
			return fmt.Errorf("Error reading min_number: %v", err)
		}
	}

	if err = d.Set("change_4_characters", flattenSystemPasswordPolicyChange4Characters(o["change-4-characters"], d, "change_4_characters")); err != nil {
		if !fortiAPIPatch(o["change-4-characters"]) {
			return fmt.Errorf("Error reading change_4_characters: %v", err)
		}
	}

	if err = d.Set("expire_status", flattenSystemPasswordPolicyExpireStatus(o["expire-status"], d, "expire_status")); err != nil {
		if !fortiAPIPatch(o["expire-status"]) {
			return fmt.Errorf("Error reading expire_status: %v", err)
		}
	}

	if err = d.Set("expire_day", flattenSystemPasswordPolicyExpireDay(o["expire-day"], d, "expire_day")); err != nil {
		if !fortiAPIPatch(o["expire-day"]) {
			return fmt.Errorf("Error reading expire_day: %v", err)
		}
	}

	if err = d.Set("reuse_password", flattenSystemPasswordPolicyReusePassword(o["reuse-password"], d, "reuse_password")); err != nil {
		if !fortiAPIPatch(o["reuse-password"]) {
			return fmt.Errorf("Error reading reuse_password: %v", err)
		}
	}


	return nil
}

func flattenSystemPasswordPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemPasswordPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyApplyTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinimumLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinLowerCaseLetter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinUpperCaseLetter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinNonAlphanumeric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyChange4Characters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyExpireStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyExpireDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyReusePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemPasswordPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemPasswordPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("apply_to"); ok {
		t, err := expandSystemPasswordPolicyApplyTo(d, v, "apply_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apply-to"] = t
		}
	}

	if v, ok := d.GetOk("minimum_length"); ok {
		t, err := expandSystemPasswordPolicyMinimumLength(d, v, "minimum_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-length"] = t
		}
	}

	if v, ok := d.GetOk("min_lower_case_letter"); ok {
		t, err := expandSystemPasswordPolicyMinLowerCaseLetter(d, v, "min_lower_case_letter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-lower-case-letter"] = t
		}
	}

	if v, ok := d.GetOk("min_upper_case_letter"); ok {
		t, err := expandSystemPasswordPolicyMinUpperCaseLetter(d, v, "min_upper_case_letter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-upper-case-letter"] = t
		}
	}

	if v, ok := d.GetOk("min_non_alphanumeric"); ok {
		t, err := expandSystemPasswordPolicyMinNonAlphanumeric(d, v, "min_non_alphanumeric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-non-alphanumeric"] = t
		}
	}

	if v, ok := d.GetOk("min_number"); ok {
		t, err := expandSystemPasswordPolicyMinNumber(d, v, "min_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-number"] = t
		}
	}

	if v, ok := d.GetOk("change_4_characters"); ok {
		t, err := expandSystemPasswordPolicyChange4Characters(d, v, "change_4_characters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["change-4-characters"] = t
		}
	}

	if v, ok := d.GetOk("expire_status"); ok {
		t, err := expandSystemPasswordPolicyExpireStatus(d, v, "expire_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-status"] = t
		}
	}

	if v, ok := d.GetOk("expire_day"); ok {
		t, err := expandSystemPasswordPolicyExpireDay(d, v, "expire_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-day"] = t
		}
	}

	if v, ok := d.GetOk("reuse_password"); ok {
		t, err := expandSystemPasswordPolicyReusePassword(d, v, "reuse_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-password"] = t
		}
	}


	return &obj, nil
}

