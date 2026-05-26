# Create a restricted user with specific library access.
resource "emby_user" "kids" {
  name                    = "Kids"
  is_administrator        = false
  enable_remote_access    = false
  enable_content_deletion = false
  enable_all_folders      = false
  enable_all_channels     = false

  # Parental controls
  max_parental_rating = 12
  block_unrated_items = ["Movie", "Trailer", "Series"]

  # Subtitles always on
  subtitle_mode = "Always"
}

# Create an admin user with full access.
resource "emby_user" "admin" {
  name               = "admin-user"
  is_administrator   = true
  enable_all_folders = true
}
