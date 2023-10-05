resource "huaweicloud_vpc_eip" "hw_eip1" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    share_type  = "PER"
    name        = "${local.resource_client}-eip1"
    size        = 50
    charge_mode = "traffic"
  }
}

resource "huaweicloud_vpc_eip" "hw_eip2" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    share_type  = "PER"
    name        = "${local.resource_client}-eip2"
    size        = 50
    charge_mode = "traffic"
  }
}