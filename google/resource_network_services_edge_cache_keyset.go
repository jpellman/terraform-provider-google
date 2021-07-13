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

func resourceNetworkServicesEdgeCacheKeyset() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkServicesEdgeCacheKeysetCreate,
		Read:   resourceNetworkServicesEdgeCacheKeysetRead,
		Update: resourceNetworkServicesEdgeCacheKeysetUpdate,
		Delete: resourceNetworkServicesEdgeCacheKeysetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkServicesEdgeCacheKeysetImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource; provided by the client when the resource is created.
The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
and all following characters must be a dash, underscore, letter or digit.`,
			},
			"public_key": {
				Type:     schema.TypeList,
				Required: true,
				Description: `An ordered list of Ed25519 public keys to use for validating signed requests.
You must specify at least one (1) key, and may have up to three (3) keys.

Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your corresponding private key.
You should ensure that the private key is kept secret, and that only authorized users can add public keys to a keyset.`,
				MinItems: 1,
				MaxItems: 3,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The ID of the public key. The ID must be 1-63 characters long, and comply with RFC1035.
The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]*
which means the first character must be a letter, and all following characters must be a dash, underscore, letter or digit.`,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The base64-encoded value of the Ed25519 public key. The base64 encoding can be padded (44 bytes) or unpadded (43 bytes).
Representations or encodings of the public key other than this will be rejected with an error.`,
							Sensitive: true,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A human-readable description of the resource.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Set of label tags associated with the EdgeCache resource.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceNetworkServicesEdgeCacheKeysetCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesEdgeCacheKeysetDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkServicesEdgeCacheKeysetLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	publicKeysProp, err := expandNetworkServicesEdgeCacheKeysetPublicKey(d.Get("public_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("public_key"); !isEmptyValue(reflect.ValueOf(publicKeysProp)) && (ok || !reflect.DeepEqual(v, publicKeysProp)) {
		obj["publicKeys"] = publicKeysProp
	}

	url, err := replaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/edgeCacheKeysets?edgeCacheKeysetId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EdgeCacheKeyset: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EdgeCacheKeyset: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating EdgeCacheKeyset: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = networkServicesOperationWaitTime(
		config, res, project, "Creating EdgeCacheKeyset", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create EdgeCacheKeyset: %s", err)
	}

	log.Printf("[DEBUG] Finished creating EdgeCacheKeyset %q: %#v", d.Id(), res)

	return resourceNetworkServicesEdgeCacheKeysetRead(d, meta)
}

func resourceNetworkServicesEdgeCacheKeysetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EdgeCacheKeyset: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("NetworkServicesEdgeCacheKeyset %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading EdgeCacheKeyset: %s", err)
	}

	if err := d.Set("description", flattenNetworkServicesEdgeCacheKeysetDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading EdgeCacheKeyset: %s", err)
	}
	if err := d.Set("labels", flattenNetworkServicesEdgeCacheKeysetLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading EdgeCacheKeyset: %s", err)
	}
	if err := d.Set("public_key", flattenNetworkServicesEdgeCacheKeysetPublicKey(res["publicKeys"], d, config)); err != nil {
		return fmt.Errorf("Error reading EdgeCacheKeyset: %s", err)
	}

	return nil
}

func resourceNetworkServicesEdgeCacheKeysetUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EdgeCacheKeyset: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesEdgeCacheKeysetDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkServicesEdgeCacheKeysetLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	publicKeysProp, err := expandNetworkServicesEdgeCacheKeysetPublicKey(d.Get("public_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("public_key"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, publicKeysProp)) {
		obj["publicKeys"] = publicKeysProp
	}

	url, err := replaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating EdgeCacheKeyset %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("public_key") {
		updateMask = append(updateMask, "publicKeys")
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

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating EdgeCacheKeyset %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating EdgeCacheKeyset %q: %#v", d.Id(), res)
	}

	err = networkServicesOperationWaitTime(
		config, res, project, "Updating EdgeCacheKeyset", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceNetworkServicesEdgeCacheKeysetRead(d, meta)
}

func resourceNetworkServicesEdgeCacheKeysetDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EdgeCacheKeyset: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting EdgeCacheKeyset %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "EdgeCacheKeyset")
	}

	err = networkServicesOperationWaitTime(
		config, res, project, "Deleting EdgeCacheKeyset", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting EdgeCacheKeyset %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkServicesEdgeCacheKeysetImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/edgeCacheKeysets/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkServicesEdgeCacheKeysetDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNetworkServicesEdgeCacheKeysetLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNetworkServicesEdgeCacheKeysetPublicKey(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"id":    flattenNetworkServicesEdgeCacheKeysetPublicKeyId(original["id"], d, config),
			"value": flattenNetworkServicesEdgeCacheKeysetPublicKeyValue(original["value"], d, config),
		})
	}
	return transformed
}
func flattenNetworkServicesEdgeCacheKeysetPublicKeyId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNetworkServicesEdgeCacheKeysetPublicKeyValue(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandNetworkServicesEdgeCacheKeysetDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheKeysetLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNetworkServicesEdgeCacheKeysetPublicKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedId, err := expandNetworkServicesEdgeCacheKeysetPublicKeyId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedValue, err := expandNetworkServicesEdgeCacheKeysetPublicKeyValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesEdgeCacheKeysetPublicKeyId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheKeysetPublicKeyValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}