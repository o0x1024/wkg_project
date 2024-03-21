package global

import (
	types2 "wkg-agent/model/types"
)

type BACKEND_API_STRUCT struct {
	Method string
	URL    string
}

var (
	Target        = make(chan string, 500)
	WkgURL        = "http://172.31.36.33:7788"
	NewDomainPath = "./result.txt"

	Alarm   = make(chan types2.Alarm, 1000)
	Version = "v1.0"
	AgentId = ""
	V3Token = ""

	BACKEND_API = []BACKEND_API_STRUCT{
		{Method: "POST", URL: "/v3/task/uploadDomains"},
	}
)
