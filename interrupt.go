func waitForQuit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Printf("Quit")
			os.Exit(1)
		}
	}()

}
