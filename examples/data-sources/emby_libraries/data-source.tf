# List all libraries on the Emby server.
data "emby_libraries" "all" {}

output "library_names" {
  value = keys(data.emby_libraries.all.libraries)
}
