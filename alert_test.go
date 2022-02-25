package zabbix_test

import (
	"fmt"
	"testing"

	zapi "github.com/fenbays/go-zabbix-api"
)

func TestAlertsGet(t *testing.T) {
	api := getAPI(t)

	alertParams := zapi.Params{"status": "2", "output": "alertid"}

	alerts, err := api.AlertsGet(alertParams)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(len(alerts))

}
