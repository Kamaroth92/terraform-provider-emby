# The Emby provider supports authenticating via the EMBY_HOSTNAME and EMBY_APIKEY
# environment variables, or you can set them directly in the provider block.
#
# Environment variables:
#   EMBY_HOSTNAME — the base URL of your Emby server (e.g. http://emby.local:8096)
#   EMBY_APIKEY   — the API key for authentication

provider "emby" {
  # hostname = "http://emby.example.com:8096"
  # api_key  = "your-api-key-here"
}
