resource "payjp_plan" "example" {
  name        = "example_plan"
  amount      = 1000
  currency    = "jpy"
  interval    = "month"
  trial_days  = 10
  billing_day = 1
  metadata    = map("env", "example")
}