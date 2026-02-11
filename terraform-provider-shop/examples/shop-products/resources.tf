locals {
  customers = toset([for i in range(var.customer_count) : tostring(i)])
  vendors   = toset([for i in range(var.vendor_count) : tostring(i)])
  services  = toset([for i in range(var.service_count) : tostring(i)])
  products  = toset([for i in range(var.product_count) : tostring(i)])

  vendor_ids = [for k in sort(keys(shop_vendor.vendor)) : shop_vendor.vendor[k].id]
}

resource "shop_customer" "customer" {
  for_each = local.customers

  first_name    = "Customer-${each.key}-${random_string.customer_suffix[each.key].result}"
  last_name     = "Last-${random_string.customer_last[each.key].result}"
  email_address = "customer.${each.key}.${random_string.customer_suffix[each.key].result}@example.com"
  phone_number  = random_string.customer_phone[each.key].result
  address       = "City-${random_integer.customer_city[each.key].result}"
}

resource "shop_vendor" "vendor" {
  for_each = local.vendors

  name          = "Vendor-${each.key}-${random_string.vendor_suffix[each.key].result}"
  email_address = "vendor.${each.key}.${random_string.vendor_suffix[each.key].result}@example.com"
  phone_number  = random_string.vendor_phone[each.key].result
  address       = "City-${random_integer.vendor_city[each.key].result}"
  contact       = "Contact-${random_string.vendor_contact[each.key].result}"
}

resource "shop_service" "service" {
  for_each = local.services

  name  = "Service-${each.key}-${random_string.service_suffix[each.key].result}"
  price = random_integer.service_price[each.key].result
}

resource "shop_product" "product" {
  for_each = local.products

  name      = "Product-${each.key}-${random_string.product_suffix[each.key].result}"
  price     = random_integer.product_price[each.key].result
  vendor_id = local.vendor_ids[tonumber(each.key) % length(local.vendor_ids)]
}
