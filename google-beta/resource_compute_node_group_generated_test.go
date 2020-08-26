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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComputeNodeGroup_nodeGroupBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeNodeGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNodeGroup_nodeGroupBasicExample(context),
			},
			{
				ResourceName:            "google_compute_node_group.nodes",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"node_template", "zone"},
			},
		},
	})
}

func testAccComputeNodeGroup_nodeGroupBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_node_template" "soletenant-tmpl" {
  name      = "tf-test-soletenant-tmpl%{random_suffix}"
  region    = "us-central1"
  node_type = "n1-node-96-624"
}

resource "google_compute_node_group" "nodes" {
  name        = "tf-test-soletenant-group%{random_suffix}"
  zone        = "us-central1-a"
  description = "example google_compute_node_group for Terraform Google Provider"

  size          = 1
  node_template = google_compute_node_template.soletenant-tmpl.id
}
`, context)
}

func TestAccComputeNodeGroup_nodeGroupAutoscalingPolicyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeNodeGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNodeGroup_nodeGroupAutoscalingPolicyExample(context),
			},
		},
	})
}

func testAccComputeNodeGroup_nodeGroupAutoscalingPolicyExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_node_template" "soletenant-tmpl" {
  provider = google-beta
  name      = "tf-test-soletenant-tmpl%{random_suffix}"
  region    = "us-central1"
  node_type = "n1-node-96-624"
}

resource "google_compute_node_group" "nodes" {
  provider = google-beta
  name        = "tf-test-soletenant-group%{random_suffix}"
  zone        = "us-central1-a"
  description = "example google_compute_node_group for Terraform Google Provider"

  size          = 1
  node_template = google_compute_node_template.soletenant-tmpl.id
  autoscaling_policy {
    mode = "ON"
    min_nodes = 1
    max_nodes = 10
  }
}
`, context)
}

func testAccCheckComputeNodeGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_node_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("ComputeNodeGroup still exists at %s", url)
			}
		}

		return nil
	}
}
