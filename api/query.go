package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

func parse(url string) string {

	resp, err := http.Get("https://www.pathofexile.com/trade/search/Metamorph/" + url)

	if err != nil {
		fmt.Println(err)
	}

	msg := "{\"query\":{\"status\":{\"option\":\"online\"},"
	contents, err := ioutil.ReadAll(resp.Body)
	// s := strings.Split(string(contents), " require([\"main\"]")
	s2 := strings.Split(string(contents), "\"state\":")

	s3 := strings.Split(s2[1], "//-->")
	s4 := strings.Split(s3[0], ",\"loggedIn\"")
	msg += s4[0][1:]
	msg += ",\"sort\":{\"price\":\"asc\"}}"

	msg = strings.Replace(msg, ",\"status\":\"online\"", "", -1)
	return msg
}
