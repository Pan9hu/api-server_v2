package api

import apiV2 "github.com/Pan9Hu/api-server_v2/api/v2"

type APIGroup struct {
	APIS apiV2.APIS
}

var APIGroupApp = new(APIGroup)
