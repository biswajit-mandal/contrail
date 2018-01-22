package models

// ServiceGroup

import "encoding/json"

// ServiceGroup
type ServiceGroup struct {
	ServiceGroupFirewallServiceList *FirewallServiceGroupType `json:"service_group_firewall_service_list,omitempty"`
	UUID                            string                    `json:"uuid,omitempty"`
	IDPerms                         *IdPermsType              `json:"id_perms,omitempty"`
	FQName                          []string                  `json:"fq_name,omitempty"`
	DisplayName                     string                    `json:"display_name,omitempty"`
	Annotations                     *KeyValuePairs            `json:"annotations,omitempty"`
	Perms2                          *PermType2                `json:"perms2,omitempty"`
	ParentUUID                      string                    `json:"parent_uuid,omitempty"`
	ParentType                      string                    `json:"parent_type,omitempty"`
}

// String returns json representation of the object
func (model *ServiceGroup) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceGroup makes ServiceGroup
func MakeServiceGroup() *ServiceGroup {
	return &ServiceGroup{
		//TODO(nati): Apply default
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		ServiceGroupFirewallServiceList: MakeFirewallServiceGroupType(),
		UUID:    "",
		IDPerms: MakeIdPermsType(),
	}
}

// MakeServiceGroupSlice() makes a slice of ServiceGroup
func MakeServiceGroupSlice() []*ServiceGroup {
	return []*ServiceGroup{}
}
