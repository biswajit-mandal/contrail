package models

// AliasIP

import "encoding/json"

// AliasIP
type AliasIP struct {
	AliasIPAddress       IpAddressType       `json:"alias_ip_address,omitempty"`
	Perms2               *PermType2          `json:"perms2,omitempty"`
	FQName               []string            `json:"fq_name,omitempty"`
	AliasIPAddressFamily IpAddressFamilyType `json:"alias_ip_address_family,omitempty"`
	DisplayName          string              `json:"display_name,omitempty"`
	Annotations          *KeyValuePairs      `json:"annotations,omitempty"`
	UUID                 string              `json:"uuid,omitempty"`
	ParentUUID           string              `json:"parent_uuid,omitempty"`
	ParentType           string              `json:"parent_type,omitempty"`
	IDPerms              *IdPermsType        `json:"id_perms,omitempty"`

	ProjectRefs                 []*AliasIPProjectRef                 `json:"project_refs,omitempty"`
	VirtualMachineInterfaceRefs []*AliasIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
}

// AliasIPProjectRef references each other
type AliasIPProjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// AliasIPVirtualMachineInterfaceRef references each other
type AliasIPVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *AliasIP) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAliasIP makes AliasIP
func MakeAliasIP() *AliasIP {
	return &AliasIP{
		//TODO(nati): Apply default
		IDPerms:              MakeIdPermsType(),
		AliasIPAddressFamily: MakeIpAddressFamilyType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		AliasIPAddress:       MakeIpAddressType(),
		Perms2:               MakePermType2(),
		FQName:               []string{},
	}
}

// MakeAliasIPSlice() makes a slice of AliasIP
func MakeAliasIPSlice() []*AliasIP {
	return []*AliasIP{}
}
