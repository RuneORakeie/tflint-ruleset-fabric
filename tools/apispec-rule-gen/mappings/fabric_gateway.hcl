// Mapping for fabric_gateway resource
// Auto-generated from platform/definitions/gateways.json
// Properties merged from: CreateGatewayRequest, CreateVirtualNetworkGatewayRequest, CreateOnPremisesGatewayRequest
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_gateway" {
  import_path = "platform/definitions/gateways.json"

  // required, format: uuid
  attribute "capacity_id_virtual_network_gateway" {
    api_ref = "CreateVirtualNetworkGatewayRequest.capacityId"
  }

  // optional
  attribute "capacity_id_manual_manual" {
    api_ref = "manual.capacity_id_manual"
  }

  // optional
  attribute "capacity_id_virtual_network_gateway" {
    api_ref = "CreateVirtualNetworkGatewayRequest.capacityId"
  }

  // required, max 200 chars
  attribute "display_name_virtual_network_gateway" {
    api_ref = "CreateVirtualNetworkGatewayRequest.displayName"
    max_length = 200
  }

  // optional, max 200 chars
  attribute "display_name_gateway" {
    api_ref = "UpdateGatewayRequest.displayName"
    max_length = 200
  }

  // optional, max 200 chars
  attribute "display_name_gateway" {
    api_ref = "UpdateGatewayRequest.displayName"
    max_length = 200
  }

  // optional, max 200 chars
  attribute "display_name_gateway_manual" {
    api_ref = "manual.display_name_gateway"
    max_length = 200
  }

  // optional, max 200 chars
  attribute "display_name_virtual_network_gateway" {
    api_ref = "CreateVirtualNetworkGatewayRequest.displayName"
    max_length = 200
  }

  // required
  attribute "inactivity_minutes_before_sleep_virtual_network_gateway" {
    api_ref = "CreateVirtualNetworkGatewayRequest.inactivityMinutesBeforeSleep"
  }

  // optional, enum(10 values)
  attribute "inactivity_minutes_before_sleep_manual_manual" {
    api_ref = "manual.inactivity_minutes_before_sleep_manual"
    valid_values = ["30", "60", "90", "120", "150", "240", "360", "480", "720", "1440"]
  }

  // optional
  attribute "inactivity_minutes_before_sleep_virtual_network_gateway" {
    api_ref = "CreateVirtualNetworkGatewayRequest.inactivityMinutesBeforeSleep"
  }

  // required
  attribute "number_of_member_gateways_virtual_network_gateway" {
    api_ref = "CreateVirtualNetworkGatewayRequest.numberOfMemberGateways"
  }

  // optional
  attribute "number_of_member_gateways_manual_manual" {
    api_ref = "manual.number_of_member_gateways_manual"
  }

  // optional
  attribute "number_of_member_gateways_virtual_network_gateway" {
    api_ref = "CreateVirtualNetworkGatewayRequest.numberOfMemberGateways"
  }

  // required, enum(3 values)
  attribute "type" {
    api_ref = "CreateGatewayRequest.type"
    valid_values = ["OnPremises", "OnPremisesPersonal", "VirtualNetwork"]
  }

  // required
  attribute "virtual_network_azure_resource_virtual_network_gateway" {
    api_ref = "CreateVirtualNetworkGatewayRequest.virtualNetworkAzureResource"
  }

  // optional
  attribute "virtual_network_azure_resource_virtual_network_gateway" {
    api_ref = "CreateVirtualNetworkGatewayRequest.virtualNetworkAzureResource"
  }

  // Add manual customizations below with // MANUAL: comment
  // Example:
  // // MANUAL: custom constraint
  // attribute "display_name" {
  //   api_ref = "CreateXxxRequest.displayName"
  //   max_length = 256
  //   pattern = "^[a-zA-Z0-9_]+$"
  //   warn_on_exceed = true
  // }
}
