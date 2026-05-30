package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLibraryDataSource_byName(t *testing.T) {
	name := acctest.RandomWithPrefix("tf-test-ds-lib")
	dir := "/tmp"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create a library so we can look it up.
			{
				Config: testAccLibraryDataSourceSetup(name, dir, false),
				Check:  resource.TestCheckResourceAttr("emby_library.ds_source", "name", name),
			},
			// Look up by name.
			{
				Config: testAccLibraryDataSourceSetup(name, dir, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.emby_library.test", "name", name),
					resource.TestCheckResourceAttrSet("data.emby_library.test", "id"),
					resource.TestCheckResourceAttrSet("data.emby_library.test", "item_id"),
					resource.TestCheckResourceAttrSet("data.emby_library.test", "guid"),
				),
			},
		},
	})
}

func TestAccLibraryDataSource_byId(t *testing.T) {
	name := acctest.RandomWithPrefix("tf-test-ds-lib-id")
	dir := "/tmp"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create a library so we can look it up.
			{
				Config: testAccLibraryDataSourceSetup(name, dir, false),
				Check:  resource.TestCheckResourceAttr("emby_library.ds_source", "name", name),
			},
			// Look up by ID.
			{
				Config: testAccLibraryDataSourceById(name, dir),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.emby_library.test", "name", name),
					resource.TestCheckResourceAttrSet("data.emby_library.test", "id"),
					resource.TestCheckResourceAttrSet("data.emby_library.test", "item_id"),
					resource.TestCheckResourceAttrSet("data.emby_library.test", "guid"),
				),
			},
		},
	})
}

func testAccLibraryDataSourceSetup(name, path string, useDataSource bool) string {
	if useDataSource {
		return fmt.Sprintf(`
provider "emby" {}

resource "emby_library" "ds_source" {
  name            = %[1]q
  collection_type = "movies"
  paths           = [%[2]q]
}

data "emby_library" "test" {
  name = emby_library.ds_source.name
}
`, name, path)
	}
	return fmt.Sprintf(`
provider "emby" {}

resource "emby_library" "ds_source" {
  name            = %[1]q
  collection_type = "movies"
  paths           = [%[2]q]
}
`, name, path)
}

func testAccLibraryDataSourceById(name, path string) string {
	return fmt.Sprintf(`
provider "emby" {}

resource "emby_library" "ds_source" {
  name            = %[1]q
  collection_type = "movies"
  paths           = [%[2]q]
}

data "emby_library" "test" {
  id = emby_library.ds_source.id
}
`, name, path)
}
