// Mapping for fabric_connection resource
// Auto-generated from platform/definitions/connections.json
// Properties merged from: CreateConnectionRequest, CreateCloudConnectionRequest, CreateVirtualNetworkGatewayConnectionRequest, CreateOnPremisesConnectionRequest
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_connection" {
  import_path = "platform/definitions/connections.json"

  // required
  attribute "connection_details" {
    api_ref = "CreateConnectionRequest.connectionDetails"
  }

  // required, enum(7 values)
  attribute "connectivity_type" {
    api_ref = "CreateConnectionRequest.connectivityType"
    valid_values = ["ShareableCloud", "PersonalCloud", "OnPremisesGateway", "OnPremisesGatewayPersonal", "VirtualNetworkGateway", "Automatic", "None"]
  }

  // required
  attribute "credential_details_cloud" {
    api_ref = "CreateCloudConnectionRequest.credentialDetails"
  }

  // required
  attribute "credential_details_onprem" {
    api_ref = "CreateOnPremisesConnectionRequest.credentialDetails"
  }

  // required
  attribute "credential_details_vnet" {
    api_ref = "CreateVirtualNetworkGatewayConnectionRequest.credentialDetails"
  }

  // optional
  attribute "credential_details_cloud" {
    api_ref = "CreateCloudConnectionRequest.credentialDetails"
  }

  // optional
  attribute "credential_details_cloud_manual" {
    api_ref = "manual.credential_details_cloud"
  }

  // optional
  attribute "credential_details_onprem" {
    api_ref = "CreateOnPremisesConnectionRequest.credentialDetails"
  }

  // optional
  attribute "credential_details_onprem_manual" {
    api_ref = "manual.credential_details_onprem"
  }

  // optional
  attribute "credential_details_vnet" {
    api_ref = "CreateVirtualNetworkGatewayConnectionRequest.credentialDetails"
  }

  // optional
  attribute "credential_details_vnet_manual" {
    api_ref = "manual.credential_details_vnet"
  }

  // required, max 200 chars
  attribute "display_name" {
    api_ref = "CreateConnectionRequest.displayName"
    max_length = 200
  }

  // required, format: uuid
  attribute "gateway_id_onprem" {
    api_ref = "CreateOnPremisesConnectionRequest.gatewayId"
  }

  // required, format: uuid
  attribute "gateway_id_vnet" {
    api_ref = "CreateVirtualNetworkGatewayConnectionRequest.gatewayId"
  }

  // optional
  attribute "gateway_id_onprem" {
    api_ref = "CreateOnPremisesConnectionRequest.gatewayId"
  }

  // optional
  attribute "gateway_id_onprem_manual" {
    api_ref = "manual.gateway_id_onprem"
  }

  // optional
  attribute "gateway_id_vnet" {
    api_ref = "CreateVirtualNetworkGatewayConnectionRequest.gatewayId"
  }

  // optional
  attribute "gateway_id_vnet_manual" {
    api_ref = "manual.gateway_id_vnet"
  }

  // optional, enum(4 values)
  attribute "privacy_level" {
    api_ref = "CreateConnectionRequest.privacyLevel"
    valid_values = ["None", "Private", "Organizational", "Public"]
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
