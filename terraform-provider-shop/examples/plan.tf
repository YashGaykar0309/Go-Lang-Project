terraform {
  required_providers {
    shop = {
      source  = "yashgaykar/shop"
      version = "0.1.0"
    }
  }
}

provider "shop" {
  endpoint = "http://localhost:8080"
}

resource "shop_customer" "customer_1" {
  first_name    = "Customer 1"
  last_name     = "Grahak 1"
  email_address = "customer.grahak@example.com"
  phone_number  = "79698269842"
  address       = "Pune, India"
}

resource "shop_service" "service_1" {
  name = "Service 1"
  price = 3000
}

resource "shop_vendor" "vendor_1" {
  name = "Vendor 1"
  email_address = "vendor@example.com"
  phone_number  = "7658764876"
  address       = "Pune, India"
  contact       = "Person Name"
}

