# Look up user and library by name, then grant the user access to that library.
data "emby_user" "example" {
  name = "example-user"
}

data "emby_library" "movies" {
  name = "Movies"
}

resource "emby_user_library_access" "restricted" {
  user_id = data.emby_user.example.id
  library_ids = [
    data.emby_library.movies.item_id,
  ]
}
