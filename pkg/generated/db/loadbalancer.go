package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertLoadbalancerQuery = "insert into `loadbalancer` (`status`,`provisioning_status`,`admin_state`,`vip_address`,`vip_subnet_id`,`operating_status`,`loadbalancer_provider`,`uuid`,`fq_name`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerQuery = "update `loadbalancer` set `status` = ?,`provisioning_status` = ?,`admin_state` = ?,`vip_address` = ?,`vip_subnet_id` = ?,`operating_status` = ?,`loadbalancer_provider` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?;"
const deleteLoadbalancerQuery = "delete from `loadbalancer` where uuid = ?"

// LoadbalancerFields is db columns for Loadbalancer
var LoadbalancerFields = []string{
	"status",
	"provisioning_status",
	"admin_state",
	"vip_address",
	"vip_subnet_id",
	"operating_status",
	"loadbalancer_provider",
	"uuid",
	"fq_name",
	"last_modified",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"display_name",
	"key_value_pair",
	"global_access",
	"share",
	"perms2_owner",
	"perms2_owner_access",
}

// LoadbalancerRefFields is db reference fields for Loadbalancer
var LoadbalancerRefFields = map[string][]string{

	"service_appliance_set": {
	// <utils.Schema Value>

	},

	"virtual_machine_interface": {
	// <utils.Schema Value>

	},

	"service_instance": {
	// <utils.Schema Value>

	},
}

const insertLoadbalancerServiceInstanceQuery = "insert into `ref_loadbalancer_service_instance` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerServiceApplianceSetQuery = "insert into `ref_loadbalancer_service_appliance_set` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerVirtualMachineInterfaceQuery = "insert into `ref_loadbalancer_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

// CreateLoadbalancer inserts Loadbalancer to DB
func CreateLoadbalancer(tx *sql.Tx, model *models.Loadbalancer) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertLoadbalancerQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.LoadbalancerProperties.Status),
		string(model.LoadbalancerProperties.ProvisioningStatus),
		bool(model.LoadbalancerProperties.AdminState),
		string(model.LoadbalancerProperties.VipAddress),
		string(model.LoadbalancerProperties.VipSubnetID),
		string(model.LoadbalancerProperties.OperatingStatus),
		string(model.LoadbalancerProvider),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtServiceApplianceSetRef, err := tx.Prepare(insertLoadbalancerServiceApplianceSetQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceApplianceSetRefs create statement failed")
	}
	defer stmtServiceApplianceSetRef.Close()
	for _, ref := range model.ServiceApplianceSetRefs {
		_, err = stmtServiceApplianceSetRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceApplianceSetRefs create failed")
		}
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertLoadbalancerVirtualMachineInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs create statement failed")
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	stmtServiceInstanceRef, err := tx.Prepare(insertLoadbalancerServiceInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceInstanceRefs create statement failed")
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {
		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanLoadbalancer(values map[string]interface{}) (*models.Loadbalancer, error) {
	m := models.MakeLoadbalancer()

	if value, ok := values["status"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerProperties.Status = castedValue

	}

	if value, ok := values["provisioning_status"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerProperties.ProvisioningStatus = castedValue

	}

	if value, ok := values["admin_state"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.LoadbalancerProperties.AdminState = castedValue

	}

	if value, ok := values["vip_address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerProperties.VipAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["vip_subnet_id"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerProperties.VipSubnetID = models.UuidStringType(castedValue)

	}

	if value, ok := values["operating_status"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerProperties.OperatingStatus = castedValue

	}

	if value, ok := values["loadbalancer_provider"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerProvider = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["ref_service_appliance_set"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.LoadbalancerServiceApplianceSetRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.ServiceApplianceSetRefs = append(m.ServiceApplianceSetRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.LoadbalancerVirtualMachineInterfaceRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_instance"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.LoadbalancerServiceInstanceRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

		}
	}

	return m, nil
}

// ListLoadbalancer lists Loadbalancer with list spec.
func ListLoadbalancer(tx *sql.Tx, spec *db.ListSpec) ([]*models.Loadbalancer, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "loadbalancer"
	spec.Fields = LoadbalancerFields
	spec.RefFields = LoadbalancerRefFields
	result := models.MakeLoadbalancerSlice()
	query, columns, values := db.BuildListQuery(spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.Query(query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "select query failed")
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "row error")
	}
	for rows.Next() {
		valuesMap := map[string]interface{}{}
		values := make([]interface{}, len(columns))
		valuesPointers := make([]interface{}, len(columns))
		for _, index := range columns {
			valuesPointers[index] = &values[index]
		}
		if err := rows.Scan(valuesPointers...); err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}
		for column, index := range columns {
			val := valuesPointers[index].(*interface{})
			valuesMap[column] = *val
		}
		log.WithFields(log.Fields{
			"valuesMap": valuesMap,
		}).Debug("valueMap")
		m, err := scanLoadbalancer(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowLoadbalancer shows Loadbalancer resource
func ShowLoadbalancer(tx *sql.Tx, uuid string) (*models.Loadbalancer, error) {
	list, err := ListLoadbalancer(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateLoadbalancer updates a resource
func UpdateLoadbalancer(tx *sql.Tx, uuid string, model *models.Loadbalancer) error {
	//TODO(nati) support update
	return nil
}

// DeleteLoadbalancer deletes a resource
func DeleteLoadbalancer(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerQuery)
	if err != nil {
		return errors.Wrap(err, "preparing delete query failed")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		return errors.Wrap(err, "delete failed")
	}
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}