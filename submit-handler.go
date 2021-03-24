func submitHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type submitType struct {
		Barcode   string `json:"barcode"`
	}

	decoder := json.NewDecoder(r.Body)

	var t submitType
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	fmt.Fprint(w, t.Barcode)
}
