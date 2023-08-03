package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	env "git.com/searchEngineTumail"
)

func SearchMatch(term string, fromDate string, toDate string, maxResults int, fields []string) (interface{}, error) {
	/* term: word or text we are looking for
	   fromDate and toDate: timestamp range
	   maxResults: how many results it will look for
	   fields: if we want all of them to be analyzed, the param is an empty list. In other case: ["field1", "field2"]
	*/
	query := map[string]interface{}{
		"search_type": "match",
		"query": map[string]interface{}{
			"term":       term,
			"start_time": fromDate,
			"end_time":   toDate,
		},
		"from":        0,
		"max_results": maxResults,
		"_source":     fields,
	}
	queryStr, err := json.Marshal(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// se construye el request a la API search de zincsearch
	reqType, url, bodyJsonStr := "POST", "http://localhost:4080/api/mails/_search", string(queryStr)
	req, err := http.NewRequest(reqType, url, strings.NewReader(bodyJsonStr))
	if err != nil {
		log.Fatal("Error API request _search", err)
		return nil, err
	}
	req.SetBasicAuth(env.USERDB, env.PASSWORDDB)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	// se hace el request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	log.Println("Status code", resp.StatusCode)
	// se lee la respuesta entregada por la API
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// se decodifica la respuesta para acceder a sus atributos en un map
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data["hits"], nil
}
