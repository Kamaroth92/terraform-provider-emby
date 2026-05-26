package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccApiKeyResource_basic(t *testing.T) {
	app := acctest.RandomWithPrefix("tf-test-app")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApiKeyConfig(app),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("emby_api_key.test", "app", app),
					resource.TestCheckResourceAttrSet("emby_api_key.test", "access_token"),
				),
			},
		},
	})
}

func testAccApiKeyConfig(app string) string {
	return fmt.Sprintf(`
provider "emby" {}

resource "emby_api_key" "test" {
  app = %[1]q
}
`, app)
}
