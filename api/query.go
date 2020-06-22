package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Api(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()
	param := vars.Get("param")

	s := parse(param)
	w.Write([]byte(s))

}

func parse(url string) string {

	resp, err := http.Get("https://www.pathofexile.com/trade/search/Harvest/" + url)
	if err != nil {
		fmt.Println(err)
	}

	msg := "{\"query\":{\"status\":{\"option\":\"online\"},"
	contents, err := ioutil.ReadAll(resp.Body)
	s2 := strings.Split(string(contents), "\"state\":")
	if len(s2) <= 1 {
		return ""
	}

	s3 := strings.Split(s2[1], "//-->")

	if len(s3) < 1 {
		return ""
	}

	s4 := strings.Split(s3[0], ",\"loggedIn\"")
	msg += s4[0][1:]
	msg += ",\"sort\":{\"price\":\"asc\"}}"

	msg = strings.Replace(msg, ",\"status\":\"online\"", "", -1)
	return msg
}
