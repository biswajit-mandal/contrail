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

const insertQosQueueQuery = "insert into `qos_queue` (`max_bandwidth`,`min_bandwidth`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`qos_queue_identifier`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateQosQueueQuery = "update `qos_queue` set `max_bandwidth` = ?,`min_bandwidth` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`qos_queue_identifier` = ?;"
const deleteQosQueueQuery = "delete from `qos_queue` where uuid = ?"

// QosQueueFields is db columns for QosQueue
var QosQueueFields = []string{
	"max_bandwidth",
	"min_bandwidth",
	"uuid",
	"fq_name",
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
	"description",
	"display_name",
	"key_value_pair",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"share",
	"qos_queue_identifier",
}

// QosQueueRefFields is db reference fields for QosQueue
var QosQueueRefFields = map[string][]string{}

// CreateQosQueue inserts QosQueue to DB
func CreateQosQueue(tx *sql.Tx, model *models.QosQueue) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertQosQueueQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertQosQueueQuery,
	}).Debug("create query")
	_, err = stmt.Exec(int(model.MaxBandwidth),
		int(model.MinBandwidth),
		string(model.UUID),
		utils.MustJSON(model.FQName),
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
		string(model.IDPerms.Description),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		int(model.QosQueueIdentifier))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanQosQueue(values map[string]interface{}) (*models.QosQueue, error) {
	m := models.MakeQosQueue()

	if value, ok := values["max_bandwidth"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.MaxBandwidth = castedValue

	}

	if value, ok := values["min_bandwidth"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.MinBandwidth = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

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

	if value, ok := values["description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

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

	if value, ok := values["qos_queue_identifier"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.QosQueueIdentifier = castedValue

	}

	return m, nil
}

// ListQosQueue lists QosQueue with list spec.
func ListQosQueue(tx *sql.Tx, spec *db.ListSpec) ([]*models.QosQueue, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "qos_queue"
	spec.Fields = QosQueueFields
	spec.RefFields = QosQueueRefFields
	result := models.MakeQosQueueSlice()
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
		m, err := scanQosQueue(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowQosQueue shows QosQueue resource
func ShowQosQueue(tx *sql.Tx, uuid string) (*models.QosQueue, error) {
	list, err := ListQosQueue(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateQosQueue updates a resource
func UpdateQosQueue(tx *sql.Tx, uuid string, model *models.QosQueue) error {
	//TODO(nati) support update
	return nil
}

// DeleteQosQueue deletes a resource
func DeleteQosQueue(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteQosQueueQuery)
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