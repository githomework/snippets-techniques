func dbConnectionLocal() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("user=data dbname=common password=%s sslmode=disable", global.config.PgCode))
	if err != nil {
		log.Panicln("Cannot connect to DB")
	}
	return db

}

func dbTableau() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("user=data dbname=common password=%s sslmode=disable port=5433", global.config.PgCode))
	if err != nil {
		log.Panicln("Cannot connect to DB")
	}
	return db

}
