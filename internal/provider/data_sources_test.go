package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLibrariesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `provider "emby" {}
data "emby_libraries" "all" {}`,
				Check: resource.TestCheckResourceAttrSet("data.emby_libraries.all", "libraries"),
			},
		},
	})
}

func TestAccUserDataSource_byName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
provider "emby" {}

data "emby_user" "test" {
  name = "admin-user"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.emby_user.test", "name", "admin-user"),
					resource.TestCheckResourceAttrSet("data.emby_user.test", "id"),
				),
			},
		},
	})
}
