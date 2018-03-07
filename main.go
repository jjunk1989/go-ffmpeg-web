package main

func main() {
	err := logan.open()
	if err != nil {
		panic("open log file err" + err.Error())
	}
	defer logan.close()
	// r := engine()

	// r.Run(":3000")
}
