provider "restapi" {
    uri = "http://namwrncld800.nam.nsroot.net/cgi-bin/API_phonix.py"

  debug                = true
  write_returns_object = true
}

# This will make information about the user named "John Doe" available by finding him by first name
data "restapi_object" "John" {
  path = "/api/objects"
  search_key = "first"
  search_value = "John"
}
