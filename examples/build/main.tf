variable "unique_id" {
  description = "A unique ID to be appended to resource names for testing."
  type        = string
}

output "hello_world" {
  value = "Hello ${var.unique_id}"
}
