package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	embyclient "github.com/Kamaroth92/terraform-provider-emby/client"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
					// Verify the library actually exists on the server via direct API call.
					checkLibraryExistsOnServer(name, "movies"),
				),
			},
			// Update name only.
			{
				Config: testAccLibraryConfig(updatedName, "movies", []string{dir1}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("emby_library.test", "name", updatedName),
					checkLibraryExistsOnServer(updatedName, "movies"),
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

// checkLibraryExistsOnServer queries the Emby API directly (not through the
// provider) to verify a library exists with the expected name and collection type.
func checkLibraryExistsOnServer(expectedName, expectedType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, auth := newTestClient()
		result, _, err := client.LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery(auth).Execute()
		if err != nil {
			return fmt.Errorf("direct API check failed: unable to query libraries: %w", err)
		}
		for _, folder := range result.GetItems() {
			if folder.GetName() == expectedName {
				if folder.GetCollectionType() != expectedType {
					return fmt.Errorf("direct API check: library %q has collection_type %q, expected %q",
						expectedName, folder.GetCollectionType(), expectedType)
				}
				if len(folder.GetLocations()) == 0 {
					return fmt.Errorf("direct API check: library %q has no locations", expectedName)
				}
				return nil
			}
		}
		return fmt.Errorf("direct API check: library %q not found on server", expectedName)
	}
}

// newTestClient creates an Emby API client using the same env vars as the provider.
func newTestClient() (*embyclient.APIClient, context.Context) {
	hostname := os.Getenv("EMBY_HOSTNAME")
	apiKey := os.Getenv("EMBY_APIKEY")
	cfg := embyclient.NewConfiguration()
	cfg.Servers = embyclient.OAPIServerConfigs{
		{URL: hostname + "/emby"},
	}
	client := embyclient.NewAPIClient(cfg)
	auth := context.WithValue(context.Background(), embyclient.ContextAPIKeys, map[string]embyclient.APIKey{
		"apikeyauth": {Key: apiKey},
	})
	return client, auth
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
