# Look up a user by name.
data "emby_user" "example" {
  name = "example-user"
}

output "user_id" {
  value = data.emby_user.example.id
}

# You can also look up by ID.
data "emby_user" "by_id" {
  id = "user-id-here"
}
