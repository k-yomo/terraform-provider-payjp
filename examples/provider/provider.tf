variable "payjp_api_key" {}

provider "payjp" {
  api_key = var.payjp_api_key
}