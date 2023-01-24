package controler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"

	"github.com/1amkaizen/jwtweb/entities"
)

func Index(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/home/index.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var stud entities.Student
		stud.Encode = request.Form.Get("encode")
		stud.Decode = request.Form.Get("decode")
		var encoded = base64.StdEncoding.EncodeToString([]byte(stud.Decode))
		clean := strings.Split(stud.Decode, ".")
		Head, _ := dec(clean[0])
		res, _ := pretty(Head)
		payload, _ := dec(clean[1])
		res2, _ := pretty(payload)

		temp, _ := template.ParseFiles("views/home/index.html")
		data := map[string]interface{}{
			"create":  encoded,
			"deco":    stud.Decode,
			"header":  res,
			"payload": res2,
		}
		temp.Execute(response, data)
	}
}

func dec(name string) (string, error) {
	decodedBy, err := base64.StdEncoding.DecodeString(name)
	if err != nil {
		return "", err
	}
	decodedStr := string(decodedBy)
	return decodedStr, nil
}

func pretty(str string) (string, error) {
	var prettyJson bytes.Buffer
	if err := json.Indent(&prettyJson, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJson.String(), nil
}
