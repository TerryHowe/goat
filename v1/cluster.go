package v1

import (
	"fmt"
	"net/http"
)

func ClusterGet(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprint(res, "{\"clusters\": []}")
	return
}
