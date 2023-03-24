package routes

import (
	"go-gql/services"
	"net/http"
)

func AddDummyData(w http.ResponseWriter, r *http.Request) {
	services.AddDummyData()
}
