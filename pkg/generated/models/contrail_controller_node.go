package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeContrailControllerNode makes ContrailControllerNode
func MakeContrailControllerNode() *ContrailControllerNode {
	return &ContrailControllerNode{
		//TODO(nati): Apply default
		ProvisioningLog:           "",
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
		ProvisioningState:         "",
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
	}
}

// MakeContrailControllerNode makes ContrailControllerNode
func InterfaceToContrailControllerNode(i interface{}) *ContrailControllerNode {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ContrailControllerNode{
		//TODO(nati): Apply default
		ProvisioningLog:           schema.InterfaceToString(m["provisioning_log"]),
		ProvisioningProgress:      schema.InterfaceToInt64(m["provisioning_progress"]),
		ProvisioningProgressStage: schema.InterfaceToString(m["provisioning_progress_stage"]),
		ProvisioningStartTime:     schema.InterfaceToString(m["provisioning_start_time"]),
		ProvisioningState:         schema.InterfaceToString(m["provisioning_state"]),
		UUID:                      schema.InterfaceToString(m["uuid"]),
		ParentUUID:                schema.InterfaceToString(m["parent_uuid"]),
		ParentType:                schema.InterfaceToString(m["parent_type"]),
		FQName:                    schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:                   InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:               schema.InterfaceToString(m["display_name"]),
		Annotations:               InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                    InterfaceToPermType2(m["perms2"]),
	}
}

// MakeContrailControllerNodeSlice() makes a slice of ContrailControllerNode
func MakeContrailControllerNodeSlice() []*ContrailControllerNode {
	return []*ContrailControllerNode{}
}

// InterfaceToContrailControllerNodeSlice() makes a slice of ContrailControllerNode
func InterfaceToContrailControllerNodeSlice(i interface{}) []*ContrailControllerNode {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ContrailControllerNode{}
	for _, item := range list {
		result = append(result, InterfaceToContrailControllerNode(item))
	}
	return result
}
