package api

import "net/http"

func Api(w http.ResponseWriter, r *http.Request) {

	// b, err := json.Marshal(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// w.Write(b)

	vars, ok := r.URL.Query()["param"]

	if !ok || len(vars[0]) < 1 {
		w.Write([]byte("Url Param is missing"))
		return
	}
	s := parse(vars[0])
	w.Write([]byte(s))

}
