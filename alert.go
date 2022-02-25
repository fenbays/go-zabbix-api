package zabbix

type Alert struct {
	AlertID   string     `json:"alertid"`
	ActionID  string     `json:"actionid"`
	AlertType string     `json:"alerttype"`
	Clock     string     `json:"clock"`
	Error     string     `json:"error"`
	Status    StatusType `json:"status"`
}

// Hosts is an array of Host
type Alerts []Alert

// HostsGet Wrapper for host.get
// https://www.zabbix.com/documentation/3.2/manual/api/reference/host/get
func (api *API) AlertsGet(params Params) (res Alerts, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("alert.get", params, &res)

	return
}
