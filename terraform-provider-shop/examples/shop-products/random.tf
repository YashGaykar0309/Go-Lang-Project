resource "random_string" "customer_suffix" {
  for_each = local.customers

  length  = 6
  upper   = false
  special = false
}

resource "random_string" "customer_last" {
  for_each = local.customers

  length  = 6
  upper   = false
  special = false
}

resource "random_string" "customer_phone" {
  for_each = local.customers

  length  = 10
  upper   = false
  lower   = false
  numeric = true
  special = false
}

resource "random_integer" "customer_city" {
  for_each = local.customers

  min = 1
  max = 999
}

resource "random_string" "vendor_suffix" {
  for_each = local.vendors

  length  = 6
  upper   = false
  special = false
}

resource "random_string" "vendor_contact" {
  for_each = local.vendors

  length  = 6
  upper   = false
  special = false
}

resource "random_string" "vendor_phone" {
  for_each = local.vendors

  length  = 10
  upper   = false
  lower   = false
  numeric = true
  special = false
}

resource "random_integer" "vendor_city" {
  for_each = local.vendors

  min = 1
  max = 999
}

resource "random_string" "service_suffix" {
  for_each = local.services

  length  = 6
  upper   = false
  special = false
}

resource "random_integer" "service_price" {
  for_each = local.services

  min = 100
  max = 10000
}

resource "random_string" "product_suffix" {
  for_each = local.products

  length  = 6
  upper   = false
  special = false
}

resource "random_integer" "product_price" {
  for_each = local.products

  min = 100
  max = 10000
}
