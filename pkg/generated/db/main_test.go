package db

import (
	"os"
	"testing"

	"github.com/Juniper/contrail/pkg/utils"
	log "github.com/sirupsen/logrus"
)

var testServer *utils.TestServer

func TestMain(m *testing.M) {
	testServer = utils.NewTestServer()
	defer testServer.Close()
	log.Info("starting test")
	code := m.Run()
	log.Info("finished test")
	os.Exit(code)
}