resource "fabric_workspace" "test" {
  display_name = "test"
  capacity_id  = "test-capacity"
}

resource "fabric_eventhouse" "test_valid" {
  display_name = "valid-lakehouse"
  workspace_id = fabric_workspace.test.id
}
