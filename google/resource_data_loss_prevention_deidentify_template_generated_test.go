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

func TestAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckDataLossPreventionDeidentifyTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateBasicExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_deidentify_template.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDataLossPreventionDeidentifyTemplate_dlpDeidentifyTemplateBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_deidentify_template" "basic" {
	parent = "projects/%{project}"
	description = "Description"
	display_name = "Displayname"

	deidentify_config {
		info_type_transformations {
			transformations {
				info_types {
					name = "PHONE_NUMBER"
				}
				info_types {
					name = "CREDIT_CARD_NUMBER"
				}

				primitive_transformation {
					replace_config {
						new_value {
							integer_value = 9
						}
					}
				}
			}

			transformations {
				info_types {
					name = "EMAIL_ADDRESS"
				}
				info_types {
					name = "LAST_NAME"
				}

				primitive_transformation {
					character_mask_config {
						masking_character = "X"
						number_to_mask = 4
						reverse_order = true
						characters_to_ignore {
							common_characters_to_ignore = "PUNCTUATION"
						}
					}
				}
			}

			transformations {
				info_types {
					name = "DATE_OF_BIRTH"
				}

				primitive_transformation {
					replace_config {
						new_value {
							date_value {
								year  = 2020
								month = 1
								day   = 1
							}
						}
					}
				}
			}
		}
	}
}
`, context)
}

func testAccCheckDataLossPreventionDeidentifyTemplateDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_loss_prevention_deidentify_template" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DataLossPreventionBasePath}}{{parent}}/deidentifyTemplates/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("DataLossPreventionDeidentifyTemplate still exists at %s", url)
			}
		}

		return nil
	}
}
