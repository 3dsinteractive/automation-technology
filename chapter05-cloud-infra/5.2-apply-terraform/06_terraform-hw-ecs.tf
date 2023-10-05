data "huaweicloud_availability_zones" "hw_az" {
}

data "huaweicloud_compute_flavors" "hw_compute_ecs1" {
  availability_zone = data.huaweicloud_availability_zones.hw_az.names[0]
  performance_type  = "normal"
  cpu_core_count    = 2
  memory_size       = 4
}

data "huaweicloud_compute_flavors" "hw_compute_ecs2" {
  availability_zone = data.huaweicloud_availability_zones.hw_az.names[0]
  performance_type  = "normal"
  cpu_core_count    = 2
  memory_size       = 8
}

data "huaweicloud_images_image" "hw_image_ecs" {
  name        = "Ubuntu 20.04 server 64bit"
  most_recent = true
}

resource "huaweicloud_compute_instance_v2" "hw_ecs1" {
  name              = "${local.resource_client}-ecs1"
  image_id          = data.huaweicloud_images_image.hw_image_ecs.id
  flavor_id         = data.huaweicloud_compute_flavors.hw_compute_ecs1.ids[0]
  key_pair          = local.key_name
  security_group_ids= [huaweicloud_networking_secgroup.hw_secgroup.id]
  availability_zone = data.huaweicloud_availability_zones.hw_az.names[0]

  network {
    uuid = huaweicloud_vpc_subnet.hw_subnet.id
  }
}

resource "huaweicloud_compute_eip_associate" "hw_associated1" {
  public_ip   = huaweicloud_vpc_eip.hw_eip1.address
  instance_id = huaweicloud_compute_instance_v2.hw_ecs1.id
}

resource "huaweicloud_compute_instance_v2" "hw_ecs2" {
  name              = "${local.resource_client}-ecs2"
  image_id          = data.huaweicloud_images_image.hw_image_ecs.id
  flavor_id         = data.huaweicloud_compute_flavors.hw_compute_ecs2.ids[0]
  key_pair          = local.key_name
  security_group_ids= [huaweicloud_networking_secgroup.hw_secgroup.id]
  availability_zone = data.huaweicloud_availability_zones.hw_az.names[0]

  network {
    uuid = huaweicloud_vpc_subnet.hw_subnet.id
  }
}

resource "huaweicloud_compute_eip_associate" "hw_associated2" {
  public_ip   = huaweicloud_vpc_eip.hw_eip2.address
  instance_id = huaweicloud_compute_instance_v2.hw_ecs2.id
}

output "ecs1_piblic_ip" {
  value = huaweicloud_vpc_eip.hw_eip1.address
  sensitive = true
}

output "ecs2_piblic_ip" {
  value = huaweicloud_vpc_eip.hw_eip2.address
  sensitive = true
}

