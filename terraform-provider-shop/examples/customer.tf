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

resource "shop_customer" "jay" {
  first_name    = "Jay"
  last_name     = "Patil"
  email_address = "jay.patil@example.com"
  phone_number  = "9876543210"
  address       = "Mumbai, India"
}
