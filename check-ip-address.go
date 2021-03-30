	


  const timeout = 500 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel() // important to avoid a resource leak
	var res net.Resolver

	addr, err := res.LookupIP(ctx, "ip", "name of pc")
	if err != nil {
		return false
	}
	//log.Println(r.Header.Get("X-FORWARDED-FOR"))

	return r.Header.Get("X-FORWARDED-FOR") == addr[0].String()
