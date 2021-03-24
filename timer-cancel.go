	cmd1 := exec.Command("pdflatex", "-interaction=nonstopmode", "output.tex", "-output-format=pdf", "-quiet")
	timer1 := time.NewTimer(time.Second * time.Duration(timeout))
	go func(timer *time.Timer, cmd *exec.Cmd) {
		for _ = range timer.C {
			err := cmd.Process.Signal(os.Kill)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}(timer1, cmd1)
