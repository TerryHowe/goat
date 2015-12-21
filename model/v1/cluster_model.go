package v1

import (
	"github.com/TerryHowe/goat/model"
)

func ClusterGet(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprint(res, "{\"clusters\": []}")
	return
}
