package horoscope

import (
	"net/http"
	"encoding/json"
	"smarthoroscope_api/core"
)


func GetStats(w http.ResponseWriter, r *http.Request) {
	var s core.Sign
	name := r.URL.Query().Get("sign")
	if len(name) == 0 {
		w.WriteHeader(500)
		w.Write(([]byte)("Missing sign parameter"))
		return
	}
	sign := s.Generate(r.URL.Query().Get("sign"))
	json.NewEncoder(w).Encode(sign)
}
