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
