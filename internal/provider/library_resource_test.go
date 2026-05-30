package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLibraryResource_basic(t *testing.T) {
	name := acctest.RandomWithPrefix("tf-test-library")
	updatedName := name + "-updated"
	dir1 := "/tmp"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create.
			{
				Config: testAccLibraryConfig(name, "movies", []string{dir1}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("emby_library.test", "name", name),
					resource.TestCheckResourceAttr("emby_library.test", "collection_type", "movies"),
					resource.TestCheckResourceAttrSet("emby_library.test", "id"),
					resource.TestCheckResourceAttrSet("emby_library.test", "item_id"),
					resource.TestCheckResourceAttrSet("emby_library.test", "guid"),
					resource.TestCheckResourceAttr("emby_library.test", "locations.#", "1"),
				),
			},
			// Update name only.
			{
				Config: testAccLibraryConfig(updatedName, "movies", []string{dir1}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("emby_library.test", "name", updatedName),
				),
			},
			// Import.
			{
				ResourceName:      "emby_library.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccLibraryConfig(name, collectionType string, paths []string) string {
	pathList := ""
	for _, p := range paths {
		pathList += fmt.Sprintf("%q,\n    ", p)
	}
	return fmt.Sprintf(`
provider "emby" {}

resource "emby_library" "test" {
  name            = %[1]q
  collection_type = %[2]q
  paths           = [
    %[3]s
  ]
}
`, name, collectionType, pathList)
}
