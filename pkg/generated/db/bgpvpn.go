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

const insertBGPVPNQuery = "insert into `bgpvpn` (`bgpvpn_type`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`key_value_pair`,`uuid`,`route_target`,`export_route_target_list_route_target`,`display_name`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`route_target_list_route_target`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBGPVPNQuery = "update `bgpvpn` set `bgpvpn_type` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`key_value_pair` = ?,`uuid` = ?,`route_target` = ?,`export_route_target_list_route_target` = ?,`display_name` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`route_target_list_route_target` = ?,`fq_name` = ?;"
const deleteBGPVPNQuery = "delete from `bgpvpn` where uuid = ?"

// BGPVPNFields is db columns for BGPVPN
var BGPVPNFields = []string{
	"bgpvpn_type",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"owner",
	"enable",
	"description",
	"key_value_pair",
	"uuid",
	"route_target",
	"export_route_target_list_route_target",
	"display_name",
	"global_access",
	"share",
	"perms2_owner",
	"perms2_owner_access",
	"route_target_list_route_target",
	"fq_name",
}

// BGPVPNRefFields is db reference fields for BGPVPN
var BGPVPNRefFields = map[string][]string{}

// CreateBGPVPN inserts BGPVPN to DB
func CreateBGPVPN(tx *sql.Tx, model *models.BGPVPN) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPVPNQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertBGPVPNQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.BGPVPNType),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		utils.MustJSON(model.ImportRouteTargetList.RouteTarget),
		utils.MustJSON(model.ExportRouteTargetList.RouteTarget),
		string(model.DisplayName),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		utils.MustJSON(model.RouteTargetList.RouteTarget),
		utils.MustJSON(model.FQName))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanBGPVPN(values map[string]interface{}) (*models.BGPVPN, error) {
	m := models.MakeBGPVPN()

	if value, ok := values["bgpvpn_type"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.BGPVPNType = models.VpnType(castedValue)

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

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.ImportRouteTargetList.RouteTarget)

	}

	if value, ok := values["export_route_target_list_route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.ExportRouteTargetList.RouteTarget)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

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

	if value, ok := values["route_target_list_route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.RouteTargetList.RouteTarget)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	return m, nil
}

// ListBGPVPN lists BGPVPN with list spec.
func ListBGPVPN(tx *sql.Tx, spec *db.ListSpec) ([]*models.BGPVPN, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "bgpvpn"
	spec.Fields = BGPVPNFields
	spec.RefFields = BGPVPNRefFields
	result := models.MakeBGPVPNSlice()
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
		m, err := scanBGPVPN(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowBGPVPN shows BGPVPN resource
func ShowBGPVPN(tx *sql.Tx, uuid string) (*models.BGPVPN, error) {
	list, err := ListBGPVPN(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateBGPVPN updates a resource
func UpdateBGPVPN(tx *sql.Tx, uuid string, model *models.BGPVPN) error {
	//TODO(nati) support update
	return nil
}

// DeleteBGPVPN deletes a resource
func DeleteBGPVPN(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteBGPVPNQuery)
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