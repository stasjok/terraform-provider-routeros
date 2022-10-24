package routeros

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gnewbury1/terraform-provider-routeros/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testSystemIdentityName = "routeros_system_identity.myself"

func TestAccSystemIdentityTest_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSystemIdentityDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemIdentityConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemIdentityExists(testSystemIdentityName),
					resource.TestCheckResourceAttr(testSystemIdentityName, "name", "myself"),
				),
			},
		},
	})
}

func testAccCheckSystemIdentityExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no id is set")
		}

		return nil
	}
}

func testAccSystemIdentityConfig() string {
	return `

provider "routeros" {
  insecure = true
}

resource "routeros_system_identity" "myself" {
  name   = "myself"
}

`
}

func testAccCheckSystemIdentityDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "routeros_system_identity" {
			continue
		}
		id := rs.Primary.ID
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/ip/pool/%s", c.HostURL, id), nil)
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(c.Username, c.Password)

		res, err := c.HTTPClient.Do(req)
		if err != nil {
			return nil
		}

		if res.StatusCode != 404 {
			return fmt.Errorf("dhcp client %s has been found", id)
		}

		return nil
	}

	return nil
}
