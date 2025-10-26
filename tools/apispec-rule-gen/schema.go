package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type schema struct {
	ProviderSchema providerSchema `json:"provider_schemas"`
}

type providerSchema struct {
	Fabric provider `json:"registry.terraform.io/microsoft/fabric"`
}

type provider struct {
	ResourceSchemas map[string]resourceSchema `json:"resource_schemas"`
}

type resourceSchema struct {
	Block block `json:"block"`
}

type block struct {
	Attributes map[string]attribute      `json:"attributes"`
	BlockTypes map[string]resourceSchema `json:"block_types"`
}

type attribute struct {
	Type        interface{} `json:"type"`
	Description string      `json:"description"`
	Sensitive   bool        `json:"sensitive"`
}

func loadProviderSchema() provider {
	schemaPath := getFullPath("schema/schema.json")
	fmt.Println("Loading provider schema from:", schemaPath)
	src, err := ioutil.ReadFile(schemaPath)

	if err != nil {
		panic(fmt.Sprintf("Failed to read schema file: %v\nPlease generate schema.json first:\n  cd %s/schema\n  terraform providers schema -json > schema.json", err, getFullPath("schema")))
	}

	var schema schema
	if err := json.Unmarshal(src, &schema); err != nil {
		panic(fmt.Sprintf("Failed to parse schema JSON: %v", err))
	}
	return schema.ProviderSchema.Fabric
}
