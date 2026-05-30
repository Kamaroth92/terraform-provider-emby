package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUserResource_basic(t *testing.T) {
	name := acctest.RandomWithPrefix("tf-test-user")
	updatedName := name + "-updated"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create — verifies no "unknown after apply" errors.
			{
				Config: testAccUserConfig(name, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("emby_user.test", "name", name),
					resource.TestCheckResourceAttr("emby_user.test", "is_administrator", "false"),
					resource.TestCheckResourceAttrSet("emby_user.test", "id"),
					// Spot-check computed fields are no longer unknown.
					resource.TestCheckResourceAttr("emby_user.test", "audio_language_preference", ""),
					resource.TestCheckResourceAttr("emby_user.test", "enable_all_folders", "true"),
				),
			},
			// Update — verifies computed fields stay populated.
			{
				Config: testAccUserConfig(updatedName, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("emby_user.test", "name", updatedName),
					resource.TestCheckResourceAttr("emby_user.test", "is_administrator", "true"),
				),
			},
			// Import.
			{
				ResourceName:      "emby_user.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccUserConfig(name string, admin bool) string {
	return fmt.Sprintf(`
provider "emby" {}

resource "emby_user" "test" {
  name                       = %[1]q
  is_administrator           = %[2]t
  audio_language_preference  = ""
  enable_all_folders         = true
}
`, name, admin)
}
