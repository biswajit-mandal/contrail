package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeQosIdForwardingClassPair makes QosIdForwardingClassPair
func MakeQosIdForwardingClassPair() *QosIdForwardingClassPair {
	return &QosIdForwardingClassPair{
		//TODO(nati): Apply default
		Key:               0,
		ForwardingClassID: 0,
	}
}

// MakeQosIdForwardingClassPair makes QosIdForwardingClassPair
func InterfaceToQosIdForwardingClassPair(i interface{}) *QosIdForwardingClassPair {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &QosIdForwardingClassPair{
		//TODO(nati): Apply default
		Key:               schema.InterfaceToInt64(m["key"]),
		ForwardingClassID: schema.InterfaceToInt64(m["forwarding_class_id"]),
	}
}

// MakeQosIdForwardingClassPairSlice() makes a slice of QosIdForwardingClassPair
func MakeQosIdForwardingClassPairSlice() []*QosIdForwardingClassPair {
	return []*QosIdForwardingClassPair{}
}

// InterfaceToQosIdForwardingClassPairSlice() makes a slice of QosIdForwardingClassPair
func InterfaceToQosIdForwardingClassPairSlice(i interface{}) []*QosIdForwardingClassPair {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*QosIdForwardingClassPair{}
	for _, item := range list {
		result = append(result, InterfaceToQosIdForwardingClassPair(item))
	}
	return result
}
