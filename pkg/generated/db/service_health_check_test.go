package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"
)

//For skip import error.
var _ = errors.New("")

func TestServiceHealthCheck(t *testing.T) {
	// t.Parallel()
	db := testDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db, "metadata")
	mutexTable := common.UseTable(db, "service_health_check")
	// mutexProject := common.UseTable(db, "service_health_check")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeServiceHealthCheck()
	model.UUID = "service_health_check_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "service_health_check_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var ServiceInstancecreateref []*models.ServiceHealthCheckServiceInstanceRef
	var ServiceInstancerefModel *models.ServiceInstance
	ServiceInstancerefModel = models.MakeServiceInstance()
	ServiceInstancerefModel.UUID = "service_health_check_service_instance_ref_uuid"
	ServiceInstancerefModel.FQName = []string{"test", "service_health_check_service_instance_ref_uuid"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(ctx, tx, &models.CreateServiceInstanceRequest{
			ServiceInstance: ServiceInstancerefModel,
		})
	})
	ServiceInstancerefModel.UUID = "service_health_check_service_instance_ref_uuid1"
	ServiceInstancerefModel.FQName = []string{"test", "service_health_check_service_instance_ref_uuid1"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(ctx, tx, &models.CreateServiceInstanceRequest{
			ServiceInstance: ServiceInstancerefModel,
		})
	})
	ServiceInstancerefModel.UUID = "service_health_check_service_instance_ref_uuid2"
	ServiceInstancerefModel.FQName = []string{"test", "service_health_check_service_instance_ref_uuid2"}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceInstance(ctx, tx, &models.CreateServiceInstanceRequest{
			ServiceInstance: ServiceInstancerefModel,
		})
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	ServiceInstancecreateref = append(ServiceInstancecreateref, &models.ServiceHealthCheckServiceInstanceRef{UUID: "service_health_check_service_instance_ref_uuid", To: []string{"test", "service_health_check_service_instance_ref_uuid"}})
	ServiceInstancecreateref = append(ServiceInstancecreateref, &models.ServiceHealthCheckServiceInstanceRef{UUID: "service_health_check_service_instance_ref_uuid2", To: []string{"test", "service_health_check_service_instance_ref_uuid2"}})
	model.ServiceInstanceRefs = ServiceInstancecreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "service_health_check_admin_project_uuid"
	projectModel.FQName = []string{"default-domain-test", "admin-test"}
	projectModel.Perms2.Owner = "admin"
	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:admin-test", TenantAccess: 7})
	model.Perms2.Share = createShare
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateProject(ctx, tx, &models.CreateProjectRequest{
			Project: projectModel,
		})
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	//    //populate update map
	//    updateMap := map[string]interface{}{}
	//
	//
	//    common.SetValueByPath(updateMap, ".UUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceHealthCheckProperties.URLPath", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceHealthCheckProperties.TimeoutUsecs", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceHealthCheckProperties.Timeout", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceHealthCheckProperties.MonitorType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceHealthCheckProperties.MaxRetries", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceHealthCheckProperties.HTTPMethod", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceHealthCheckProperties.HealthCheckType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceHealthCheckProperties.ExpectedCodes", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceHealthCheckProperties.Enabled", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceHealthCheckProperties.DelayUsecs", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ServiceHealthCheckProperties.Delay", ".", 1.0)
	//
	//
	//
	//    if ".Perms2.Share" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.UserVisible", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OtherAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.GroupAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Group", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.LastModified", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Enable", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Description", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Creator", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Created", ".", "test")
	//
	//
	//
	//    if ".FQName" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".FQName", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".FQName", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DisplayName", ".", "test")
	//
	//
	//
	//    if ".Annotations.KeyValuePair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
	//    }
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "service_health_check_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var ServiceInstanceref []interface{}
	//    ServiceInstanceref = append(ServiceInstanceref, map[string]interface{}{"operation":"delete", "uuid":"service_health_check_service_instance_ref_uuid", "to": []string{"test", "service_health_check_service_instance_ref_uuid"}})
	//    ServiceInstanceref = append(ServiceInstanceref, map[string]interface{}{"operation":"add", "uuid":"service_health_check_service_instance_ref_uuid1", "to": []string{"test", "service_health_check_service_instance_ref_uuid1"}})
	//
	//    ServiceInstanceAttr := map[string]interface{}{}
	//
	//
	//
	//    common.SetValueByPath(ServiceInstanceAttr, ".InterfaceType", ".", "test")
	//
	//
	//
	//    ServiceInstanceref = append(ServiceInstanceref, map[string]interface{}{"operation":"update", "uuid":"service_health_check_service_instance_ref_uuid2", "to": []string{"test", "service_health_check_service_instance_ref_uuid2"}, "attr": ServiceInstanceAttr})
	//
	//    common.SetValueByPath(updateMap, "ServiceInstanceRefs", ".", ServiceInstanceref)
	//
	//
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceHealthCheck(ctx, tx,
			&models.CreateServiceHealthCheckRequest{
				ServiceHealthCheck: model,
			})
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateServiceHealthCheck(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("delete from `ref_service_health_check_service_instance` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing ServiceInstanceRefs delete statement failed")
		}
		_, err = stmt.Exec("service_health_check_dummy_uuid", "service_health_check_service_instance_ref_uuid")
		_, err = stmt.Exec("service_health_check_dummy_uuid", "service_health_check_service_instance_ref_uuid1")
		_, err = stmt.Exec("service_health_check_dummy_uuid", "service_health_check_service_instance_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs delete failed")
		}
		return nil
	})
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceInstance(ctx, tx,
			&models.DeleteServiceInstanceRequest{
				ID: "service_health_check_service_instance_ref_uuid"})
	})
	if err != nil {
		t.Fatal("delete ref service_health_check_service_instance_ref_uuid  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceInstance(ctx, tx,
			&models.DeleteServiceInstanceRequest{
				ID: "service_health_check_service_instance_ref_uuid1"})
	})
	if err != nil {
		t.Fatal("delete ref service_health_check_service_instance_ref_uuid1  failed", err)
	}
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceInstance(
			ctx,
			tx,
			&models.DeleteServiceInstanceRequest{
				ID: "service_health_check_service_instance_ref_uuid2",
			})
	})
	if err != nil {
		t.Fatal("delete ref service_health_check_service_instance_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteProject(ctx, tx, &models.DeleteProjectRequest{
			ID: projectModel.UUID})
	})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListServiceHealthCheck(ctx, tx, &models.ListServiceHealthCheckRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.ServiceHealthChecks) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceHealthCheck(ctxDemo, tx,
			&models.DeleteServiceHealthCheckRequest{
				ID: model.UUID},
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteServiceHealthCheck(ctx, tx,
			&models.DeleteServiceHealthCheckRequest{
				ID: model.UUID})
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateServiceHealthCheck(ctx, tx,
			&models.CreateServiceHealthCheckRequest{
				ServiceHealthCheck: model})
	})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		response, err := ListServiceHealthCheck(ctx, tx, &models.ListServiceHealthCheckRequest{
			Spec: &models.ListSpec{Limit: 1}})
		if err != nil {
			return err
		}
		if len(response.ServiceHealthChecks) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
