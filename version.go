package main

import (
	"fmt"
	"net/http"
)

const versions = `{"versions": [{"status": "stable", "updated": "2015-12-11T00:00:00Z", "id": "v1"}]}`

func VersionGet(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprint(res, versions)
	return
}
