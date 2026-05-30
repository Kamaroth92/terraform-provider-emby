package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccLibrariesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `provider "emby" {}
data "emby_libraries" "all" {}`,
				Check: func(s *terraform.State) error {
					attrs := s.RootModule().Resources["data.emby_libraries.all"].Primary.Attributes
					for k := range attrs {
						if strings.HasPrefix(k, "libraries.") {
							return nil
						}
					}
					return fmt.Errorf("expected at least one library entry in 'libraries' map")
				},
			},
		},
	})
}

func TestAccUserDataSource_byName(t *testing.T) {
	name := acctest.RandomWithPrefix("tf-test-ds-user")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create a user so we can look it up.
			{
				Config: testAccUserDataSourceConfig(name, false),
				Check:  resource.TestCheckResourceAttr("emby_user.ds_source", "name", name),
			},
			// Look up by name.
			{
				Config: testAccUserDataSourceConfig(name, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.emby_user.test", "name", name),
					resource.TestCheckResourceAttrSet("data.emby_user.test", "id"),
				),
			},
		},
	})
}

func testAccUserDataSourceConfig(name string, useDataSource bool) string {
	if useDataSource {
		return fmt.Sprintf(`
provider "emby" {}

resource "emby_user" "ds_source" {
  name = %[1]q
}

data "emby_user" "test" {
  name = emby_user.ds_source.name
}
`, name)
	}
	return fmt.Sprintf(`
provider "emby" {}

resource "emby_user" "ds_source" {
  name = %[1]q
}
`, name)
}
