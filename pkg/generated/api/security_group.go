package api

import (
	"database/sql"
	"net/http"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"

	log "github.com/sirupsen/logrus"
)

//SecurityGroupRESTAPI
type SecurityGroupRESTAPI struct {
	DB *sql.DB
}

type SecurityGroupCreateRequest struct {
	Data *models.SecurityGroup `json:"security-group"`
}

//Path returns api path for collections.
func (api *SecurityGroupRESTAPI) Path() string {
	return "/security-groups"
}

//LongPath returns api path for elements.
func (api *SecurityGroupRESTAPI) LongPath() string {
	return "/security-group/:id"
}

//SetDB sets db object
func (api *SecurityGroupRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *SecurityGroupRESTAPI) Create(c echo.Context) error {
	requestData := &SecurityGroupCreateRequest{
		Data: models.MakeSecurityGroup(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_group",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	model := requestData.Data
	if model == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.CreateSecurityGroup(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "security_group",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *SecurityGroupRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *SecurityGroupRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteSecurityGroup(tx, id)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *SecurityGroupRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	var result *models.SecurityGroup
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ShowSecurityGroup(tx, id)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"security_group": result,
	})
}

//List handles a List REST API Request.
func (api *SecurityGroupRESTAPI) List(c echo.Context) error {
	var result []*models.SecurityGroup
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListSecurityGroup(tx, &common.ListSpec{
				Limit: 1000,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"security-groups": result,
	})
}