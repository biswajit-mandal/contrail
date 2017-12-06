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

const insertServiceApplianceQuery = "insert into `service_appliance` (`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`key_value_pair`,`fq_name`,`uuid`,`display_name`,`annotations_key_value_pair`,`username`,`password`,`service_appliance_ip_address`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceApplianceQuery = "update `service_appliance` set `description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`key_value_pair` = ?,`fq_name` = ?,`uuid` = ?,`display_name` = ?,`annotations_key_value_pair` = ?,`username` = ?,`password` = ?,`service_appliance_ip_address` = ?;"
const deleteServiceApplianceQuery = "delete from `service_appliance` where uuid = ?"

// ServiceApplianceFields is db columns for ServiceAppliance
var ServiceApplianceFields = []string{
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"group_access",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"enable",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"share",
	"key_value_pair",
	"fq_name",
	"uuid",
	"display_name",
	"annotations_key_value_pair",
	"username",
	"password",
	"service_appliance_ip_address",
}

// ServiceApplianceRefFields is db reference fields for ServiceAppliance
var ServiceApplianceRefFields = map[string][]string{

	"physical_interface": {
		// <utils.Schema Value>
		"interface_type",
	},
}

const insertServiceAppliancePhysicalInterfaceQuery = "insert into `ref_service_appliance_physical_interface` (`from`, `to` ,`interface_type`) values (?, ?,?);"

// CreateServiceAppliance inserts ServiceAppliance to DB
func CreateServiceAppliance(tx *sql.Tx, model *models.ServiceAppliance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceApplianceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertServiceApplianceQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		utils.MustJSON(model.ServiceApplianceProperties.KeyValuePair),
		utils.MustJSON(model.FQName),
		string(model.UUID),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.ServiceApplianceUserCredentials.Username),
		string(model.ServiceApplianceUserCredentials.Password),
		string(model.ServiceApplianceIPAddress))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtPhysicalInterfaceRef, err := tx.Prepare(insertServiceAppliancePhysicalInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing PhysicalInterfaceRefs create statement failed")
	}
	defer stmtPhysicalInterfaceRef.Close()
	for _, ref := range model.PhysicalInterfaceRefs {
		_, err = stmtPhysicalInterfaceRef.Exec(model.UUID, ref.UUID, string(ref.Attr.InterfaceType))
		if err != nil {
			return errors.Wrap(err, "PhysicalInterfaceRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanServiceAppliance(values map[string]interface{}) (*models.ServiceAppliance, error) {
	m := models.MakeServiceAppliance()

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

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

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

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceApplianceProperties.KeyValuePair)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["annotations_key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["username"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceApplianceUserCredentials.Username = castedValue

	}

	if value, ok := values["password"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceApplianceUserCredentials.Password = castedValue

	}

	if value, ok := values["service_appliance_ip_address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceApplianceIPAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["ref_physical_interface"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.ServiceAppliancePhysicalInterfaceRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.PhysicalInterfaceRefs = append(m.PhysicalInterfaceRefs, referenceModel)

			attr := models.MakeServiceApplianceInterfaceType()
			referenceModel.Attr = attr

		}
	}

	return m, nil
}

// ListServiceAppliance lists ServiceAppliance with list spec.
func ListServiceAppliance(tx *sql.Tx, spec *db.ListSpec) ([]*models.ServiceAppliance, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "service_appliance"
	spec.Fields = ServiceApplianceFields
	spec.RefFields = ServiceApplianceRefFields
	result := models.MakeServiceApplianceSlice()
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
		m, err := scanServiceAppliance(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowServiceAppliance shows ServiceAppliance resource
func ShowServiceAppliance(tx *sql.Tx, uuid string) (*models.ServiceAppliance, error) {
	list, err := ListServiceAppliance(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateServiceAppliance updates a resource
func UpdateServiceAppliance(tx *sql.Tx, uuid string, model *models.ServiceAppliance) error {
	//TODO(nati) support update
	return nil
}

// DeleteServiceAppliance deletes a resource
func DeleteServiceAppliance(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceApplianceQuery)
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