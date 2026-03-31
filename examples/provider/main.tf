provider "emby" {
  # hostname = "https://emby.local"
  # api_key  = "api-token"
}

data "emby_libraries" "all" {}

data "emby_library" "Movies" {
  name = "Movies"
}

data "emby_library" "TV" {
  name = "TV Shows"
  id = "3"
}

output "tv" {
  value = data.emby_library.TV.name
}

resource "emby_user" "example" {
  name                  = data.emby_libraries.all.libraries["Movies"].name
  is_administrator      = false
  enable_remote_access  = true
  enable_media_playback = true
}