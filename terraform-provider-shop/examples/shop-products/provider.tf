terraform {
  required_providers {
    shop = {
      source  = "yashgaykar/shop"
      version = "0.1.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.3"
    }
  }
}

provider "shop" {
  endpoint = var.shop_endpoint
}
