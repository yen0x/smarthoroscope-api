package horoscope

import (
	"encoding/json"
	"fmt"
	"net/http"
	"smarthoroscope_api/core"
)

func GetStats(w http.ResponseWriter, r *http.Request) {
	var s core.Sign
	name := r.URL.Query().Get("sign")
	date := r.URL.Query().Get("date")

	//TODO check date and name values
	if len(name) == 0 {
		w.WriteHeader(500)
		w.Write(([]byte)("Missing sign parameter"))
		return
	} else if len(date) == 0 {
		w.WriteHeader(500)
		w.Write(([]byte)("Missing date parameter"))
		return
	}

	m := core.GetConnection()
	defer m.Close()

	signDao := core.NewSignDao(m)

	sign, err := signDao.FindSignByDate(name, date)

	//Sign not found
	if err != nil && err.Error() == "not found" {
		sign = s.Generate(name, date)
		err = signDao.Save(sign)

		fmt.Println("generated sign " + name)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			w.Write(([]byte)("Internal error"))
			return
		}
	} else if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write(([]byte)("Internal error"))
		return
	} else {
		fmt.Println("found sign")
	}

	json.NewEncoder(w).Encode(sign)
}
