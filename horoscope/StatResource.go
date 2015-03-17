package horoscope

import (
	"net/http"
	"encoding/json"
)


func GetStats(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{0:{00:8,01:9,02:5,03:4}}")
}
