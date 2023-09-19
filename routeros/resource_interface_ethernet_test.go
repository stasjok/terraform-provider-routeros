package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testInterfaceEthernetAddress = "routeros_interface_ethernet.test"

func TestAccInterfaceEthernetTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceEthernetConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceEthernetExists(testInterfaceEthernetAddress),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "name", "terraform"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "mtu", "9000"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "advertise", "10000M-full"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "arp", "disabled"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "auto_negotiation", "false"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "tx_flow_control", "auto"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "rx_flow_control", "auto"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "full_duplex", "true"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "mdix_enable", "false"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "sfp_shutdown_temperature", "60"),
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "speed", "100Mbps"),

							// read only properties. #slave and #switch are not returned from the virtual switch
							// so we can add assertions.
							resource.TestCheckResourceAttr(testInterfaceEthernetAddress, "running", "true"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckInterfaceEthernetExists(name string) resource.TestCheckFunc {
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

func testAccInterfaceEthernetConfig() string {
	return providerConfig + `

resource "routeros_interface_ethernet" "test" {
  factory_name              = "ether2"
  name                      = "terraform"
  mtu                       = "9000"
  advertise                 = "10000M-full"
  arp                       = "disabled"
  auto_negotiation          = false
  tx_flow_control           = "auto"
  rx_flow_control           = "auto"
  full_duplex  	            = "true"
  mdix_enable               = false
  sfp_shutdown_temperature  = 60
  speed                     = "100Mbps"
}
`
}