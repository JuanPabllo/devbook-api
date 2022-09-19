package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(res http.ResponseWriter, statusCode int, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)

	if data != nil {
		if erro := json.NewEncoder(res).Encode(data); erro != nil {
			log.Fatal(erro)
		}
	}

}

func Error(res http.ResponseWriter, statusCode int, erro error) {
	JSON(res, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: erro.Error(),
	})
}
