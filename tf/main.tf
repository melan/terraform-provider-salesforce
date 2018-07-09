variable "client_id" {}
variable "client_secret" {}
variable "username" {}
variable "password" {}

provider "salesforce" {
  api_version = "v43.0"
  client_id = "${var.client_id}"
  client_secret = "${var.client_secret}"
  username = "${var.username}"
  password = "${var.password}"
  trace_on = false
}

data "sfdx_org" "my_org" {
  provider = "salesforce"
}

resource "sfdx_user" "foo" {
  username = "foo@another.domain.net"
  last_name = "Bar"
  first_name = "Foo"
  email = "melanchenko+foo@gmail.com"
  alias = "Foo"
  is_active = false

  provider = "salesforce"
}

resource "sfdx_user" "boo" {
  username = "boo@another.domain.net"
  last_name = "Bar"
  first_name = "Boo"
  email = "melanchenko+boo@gmail.com"
  alias = "Boo"
  is_active = false

  provider = "salesforce"
}

output "org_id" {
  value = "${data.sfdx_org.my_org.org_id}"
}

output "user_id" {
  value = "${data.sfdx_org.my_org.user_id}"
}

output "foo_uid" {
  value = "${sfdx_user.foo.id}"
}

output "boo_uid" {
  value = "${sfdx_user.boo.id}"
}