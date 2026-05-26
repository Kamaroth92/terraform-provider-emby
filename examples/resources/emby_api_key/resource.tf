# Create an API key for automation tools.
resource "emby_api_key" "automation" {
  app = "terraform-automation"
}

output "automation_token" {
  value     = emby_api_key.automation.access_token
  sensitive = true
}
