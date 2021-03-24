func quickQuery(statement string) string {

	var jsonStr = []byte(`{"statement":"` + statement + `"}`)

	req, err := http.NewRequest("POST", "http://localhost:"+global.config.FastConnPort+"/do", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}

	b, _ := ioutil.ReadAll(resp.Body)
	return string(b)

}
