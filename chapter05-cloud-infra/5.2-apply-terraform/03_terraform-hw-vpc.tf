resource "huaweicloud_vpc" "hw_vpc" {
  name = "${local.resource_client}-vpc"
  cidr = "10.0.0.0/16"
}

resource "huaweicloud_vpc_subnet" "hw_subnet" {
  name       = "${local.resource_client}-subnet"
  cidr       = "10.0.65.0/24"
  gateway_ip = "10.0.65.1"
  vpc_id     = huaweicloud_vpc.hw_vpc.id
  dns_list   = ["100.125.1.250", "100.125.1.251"]
}