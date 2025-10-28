resource "fabric_workspace" "test" {
  display_name = "test"
  # Missing capacity_id - should trigger rule
}
