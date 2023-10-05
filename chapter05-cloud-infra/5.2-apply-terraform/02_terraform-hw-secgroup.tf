resource "huaweicloud_networking_secgroup" "hw_secgroup" {
  name        = "${local.resource_client}-sg"
  delete_default_rules = true
}

resource "huaweicloud_networking_secgroup_rule" "secgroup_inbound_rule_0" {
  security_group_id = huaweicloud_networking_secgroup.hw_secgroup.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 22
  port_range_max    = 22
  remote_ip_prefix  = "0.0.0.0/0"
  description = "Allow Inbound SSH Traffic"
}

resource "huaweicloud_networking_secgroup_rule" "secgroup_inbound_rule_1" {
  security_group_id = huaweicloud_networking_secgroup.hw_secgroup.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 80
  port_range_max    = 80
  remote_ip_prefix  = "0.0.0.0/0"
  description = "Allow Inbound Http port"
}

resource "huaweicloud_networking_secgroup_rule" "secgroup_inbound_rule_2" {
  security_group_id = huaweicloud_networking_secgroup.hw_secgroup.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 443
  port_range_max    = 443
  remote_ip_prefix  = "0.0.0.0/0"
  description = "Allow Inbound Https port"
}

resource "huaweicloud_networking_secgroup_rule" "secgroup_outbound_rule_0" {
  security_group_id = huaweicloud_networking_secgroup.hw_secgroup.id
  direction         = "egress"
  ethertype         = "IPv4"
  protocol          = "0"
  port_range_min    = 0
  port_range_max    = 0
  remote_ip_prefix  = "0.0.0.0/0"
  description = "Explicitly Allow All Outbound Traffic"
}
