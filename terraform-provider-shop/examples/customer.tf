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

resource "shop_customer" "yash" {
  first_name    = "Yash"
  last_name     = "Balasaheb Gaykar"
  email_address = "yash.gaykar@example.com"
  phone_number  = "9359687777"
  address       = "Pune, India"
}
