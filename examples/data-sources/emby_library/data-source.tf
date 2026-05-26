# Look up a single library by name.
data "emby_library" "movies" {
  name = "Movies"
}

output "movies_library_id" {
  value = data.emby_library.movies.id
}

output "movies_item_id" {
  value = data.emby_library.movies.item_id
}

# You can also look up by ID.
data "emby_library" "by_id" {
  id = "library-id-here"
}
