package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUserLibraryAccessResource_basic(t *testing.T) {
	userName := acctest.RandomWithPrefix("tf-test-ula-user")
	libName := acctest.RandomWithPrefix("tf-test-ula-lib")
	dir := t.TempDir()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Setup: create a user and a library.
			{
				Config: testAccUserLibraryAccessSetup(userName, libName, dir),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("emby_user.ula_user", "name", userName),
					resource.TestCheckResourceAttr("emby_library.ula_lib", "name", libName),
				),
			},
			// Create: restrict user to only the created library.
			{
				Config: testAccUserLibraryAccessConfig(userName, libName, dir),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("emby_user_library_access.test", "user_id"),
					resource.TestCheckResourceAttr("emby_user_library_access.test", "library_ids.#", "1"),
					// User's enable_all_folders should now be false.
					resource.TestCheckResourceAttr("emby_user.ula_user", "enable_all_folders", "false"),
				),
			},
		},
	})
}

func testAccUserLibraryAccessSetup(userName, libName, path string) string {
	return fmt.Sprintf(`
provider "emby" {}

resource "emby_user" "ula_user" {
  name = %[1]q
}

resource "emby_library" "ula_lib" {
  name            = %[2]q
  collection_type = "movies"
  paths           = [%[3]q]
}
`, userName, libName, path)
}

func testAccUserLibraryAccessConfig(userName, libName, path string) string {
	return fmt.Sprintf(`
provider "emby" {}

resource "emby_user" "ula_user" {
  name = %[1]q
}

resource "emby_library" "ula_lib" {
  name            = %[2]q
  collection_type = "movies"
  paths           = [%[3]q]
}

data "emby_libraries" "all" {}

resource "emby_user_library_access" "test" {
  user_id     = emby_user.ula_user.id
  library_ids = [emby_library.ula_lib.id]
}
`, userName, libName, path)
}
