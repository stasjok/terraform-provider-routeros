# routeros_interface_macvlan (Resource)


## Example Usage
```terraform
resource "routeros_interface_macvlan" "test" {
  interface    = "ether1"
  name         = "macvlan1"
  disabled     = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface` (String) Name of the interface.
- `name` (String) Changing the name of this resource will force it to be recreated.
	> The links of other configuration properties to this resource may be lost!
	> Changing the name of the resource outside of a Terraform will result in a loss of control integrity for that resource!

### Optional

- `arp` (String) Address Resolution Protocol mode:
	* disabled - the interface will not use ARP
	* enabled - the interface will use ARP
	* local-proxy-arp - the router performs proxy ARP on the interface and sends replies to the same interface
	* proxy-arp - the router performs proxy ARP on the interface and sends replies to other interfaces
	* reply-only - the interface will only reply to requests originated from matching IP address/MAC address combinations which are entered as static entries in the ARP table. No dynamic entries will be automatically stored in the ARP table. Therefore for communications to be successful, a valid static entry must already exist.
- `arp_timeout` (String) ARP timeout is time how long ARP record is kept in ARP table after no packets are received from IP. Value auto equals to the value of arp-timeout in IP/Settings, default is 30s. Can use postfix ms, s, M, h, d for milliseconds, seconds, minutes, hours or days. If no postfix is set then seconds (s) is used.
- `comment` (String)
- `disabled` (Boolean)
- `loop_protect` (String)
- `loop_protect_disable_time` (String)
- `loop_protect_send_interval` (String)
- `mac_address` (String) Static MAC address of the interface. A randomly generated MAC address will be assigned when not specified.
- `mode` (String) Sets MACVLAN interface mode:
	private - does not allow communication between MACVLAN instances on the same parent interface.
	bridge - allows communication between MACVLAN instances on the same parent interface.

### Read-Only

- `id` (String) The ID of this resource.
- `l2mtu` (Number) Layer2 Maximum transmission unit. [See](https://wiki.mikrotik.com/wiki/Maximum_Transmission_Unit_on_RouterBoards).

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/macvlan get [print show-ids]]
terraform import routeros_interface_macvlan.test "*0"
```