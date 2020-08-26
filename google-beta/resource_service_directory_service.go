// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceDirectoryService() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDirectoryServiceCreate,
		Read:   resourceServiceDirectoryServiceRead,
		Update: resourceServiceDirectoryServiceUpdate,
		Delete: resourceServiceDirectoryServiceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceServiceDirectoryServiceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The resource name of the namespace this service will belong to.`,
			},
			"service_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRFC1035Name(2, 63),
				Description: `The Resource ID must be 1-63 characters long, including digits,
lowercase letters or the hyphen character.`,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Metadata for the service. This data can be consumed
by service clients. The entire metadata dictionary may contain
up to 2000 characters, spread across all key-value pairs.
Metadata that goes beyond any these limits will be rejected.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name for the service in the
format 'projects/*/locations/*/namespaces/*/services/*'.`,
			},
		},
	}
}

func resourceServiceDirectoryServiceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	metadataProp, err := expandServiceDirectoryServiceMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}

	url, err := replaceVars(d, config, "{{ServiceDirectoryBasePath}}{{namespace}}/services?serviceId={{service_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Service: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Service: %s", err)
	}
	if err := d.Set("name", flattenServiceDirectoryServiceName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Service %q: %#v", d.Id(), res)

	return resourceServiceDirectoryServiceRead(d, meta)
}

func resourceServiceDirectoryServiceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ServiceDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ServiceDirectoryService %q", d.Id()))
	}

	if err := d.Set("name", flattenServiceDirectoryServiceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}
	if err := d.Set("metadata", flattenServiceDirectoryServiceMetadata(res["metadata"], d, config)); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}

	return nil
}

func resourceServiceDirectoryServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	obj := make(map[string]interface{})
	metadataProp, err := expandServiceDirectoryServiceMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}

	url, err := replaceVars(d, config, "{{ServiceDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Service %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("metadata") {
		updateMask = append(updateMask, "metadata")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Service %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Service %q: %#v", d.Id(), res)
	}

	return resourceServiceDirectoryServiceRead(d, meta)
}

func resourceServiceDirectoryServiceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	url, err := replaceVars(d, config, "{{ServiceDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Service %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Service")
	}

	log.Printf("[DEBUG] Finished deleting Service %q: %#v", d.Id(), res)
	return nil
}

func resourceServiceDirectoryServiceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	nameParts := strings.Split(d.Get("name").(string), "/")
	if len(nameParts) == 8 {
		// `projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}`
		d.Set("namespace", fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", nameParts[1], nameParts[3], nameParts[5]))
		d.Set("service_id", nameParts[7])
	} else if len(nameParts) == 4 {
		// `{{project}}/{{location}}/{{namespace_id}}/{{service_id}}`
		d.Set("namespace", fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", nameParts[0], nameParts[1], nameParts[2]))
		d.Set("service_id", nameParts[3])
		id := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", nameParts[0], nameParts[1], nameParts[2], nameParts[3])
		d.Set("name", id)
		d.SetId(id)
	} else if len(nameParts) == 3 {
		// `{{location}}/{{namespace_id}}/{{service_id}}`
		project, err := getProject(d, config)
		if err != nil {
			return nil, err
		}
		d.Set("namespace", fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", project, nameParts[0], nameParts[1]))
		d.Set("service_id", nameParts[2])
		id := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", project, nameParts[0], nameParts[1], nameParts[2])
		d.Set("name", id)
		d.SetId(id)
	} else {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s, %s or %s",
			d.Get("name"),
			"projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}",
			"{{project}}/{{location}}/{{namespace_id}}/{{service_id}}",
			"{{location}}/{{namespace_id}}/{{service_id}}")
	}
	return []*schema.ResourceData{d}, nil
}

func flattenServiceDirectoryServiceName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenServiceDirectoryServiceMetadata(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandServiceDirectoryServiceMetadata(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
