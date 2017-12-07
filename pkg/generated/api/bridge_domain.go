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

//BridgeDomainRESTAPI
type BridgeDomainRESTAPI struct {
	DB *sql.DB
}

type BridgeDomainCreateRequest struct {
	Data *models.BridgeDomain `json:"bridge-domain"`
}

//Path returns api path for collections.
func (api *BridgeDomainRESTAPI) Path() string {
	return "/bridge-domains"
}

//LongPath returns api path for elements.
func (api *BridgeDomainRESTAPI) LongPath() string {
	return "/bridge-domain/:id"
}

//SetDB sets db object
func (api *BridgeDomainRESTAPI) SetDB(db *sql.DB) {
	api.DB = db
}

//Create handle a Create REST API.
func (api *BridgeDomainRESTAPI) Create(c echo.Context) error {
	requestData := &BridgeDomainCreateRequest{
		Data: models.MakeBridgeDomain(),
	}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bridge_domain",
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
			return db.CreateBridgeDomain(tx, model)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "bridge_domain",
		}).Debug("db create failed on create")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusCreated, requestData)
}

//Update handles a REST Update request.
func (api *BridgeDomainRESTAPI) Update(c echo.Context) error {
	return nil
}

//Delete handles a REST Delete request.
func (api *BridgeDomainRESTAPI) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			return db.DeleteBridgeDomain(tx, id)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show handles a REST Show request.
func (api *BridgeDomainRESTAPI) Show(c echo.Context) error {
	id := c.Param("id")
	var result *models.BridgeDomain
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ShowBridgeDomain(tx, id)
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"bridge_domain": result,
	})
}

//List handles a List REST API Request.
func (api *BridgeDomainRESTAPI) List(c echo.Context) error {
	var result []*models.BridgeDomain
	var err error
	if err := common.DoInTransaction(
		api.DB,
		func(tx *sql.Tx) error {
			result, err = db.ListBridgeDomain(tx, &common.ListSpec{
				Limit: 1000,
			})
			return err
		}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"bridge-domains": result,
	})
}