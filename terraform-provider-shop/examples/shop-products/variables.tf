variable "shop_endpoint" {
  type        = string
  description = "Base URL for the shop API."
  default     = "http://localhost:8080"
}

variable "customer_count" {
  type        = number
  description = "How many customers to create."
  default     = 2
  validation {
    condition     = var.customer_count >= 0
    error_message = "customer_count must be >= 0."
  }
}

variable "vendor_count" {
  type        = number
  description = "How many vendors to create."
  default     = 2
  validation {
    condition     = var.vendor_count >= 0
    error_message = "vendor_count must be >= 0."
  }
}

variable "service_count" {
  type        = number
  description = "How many services to create."
  default     = 1
  validation {
    condition     = var.service_count >= 0
    error_message = "service_count must be >= 0."
  }
}

variable "product_count" {
  type        = number
  description = "How many products to create."
  default     = 3
  validation {
    condition     = var.product_count >= 0
    error_message = "product_count must be >= 0."
  }
}

variable "product_requires_vendors" {
  type        = bool
  description = "Safety switch to require vendors when creating products."
  default     = true
  validation {
    condition     = !var.product_requires_vendors || var.product_count == 0 || var.vendor_count > 0
    error_message = "vendor_count must be > 0 when product_count > 0."
  }
}
