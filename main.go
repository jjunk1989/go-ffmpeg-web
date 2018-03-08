package main

func main() {
	err := logan.open()
	if err != nil {
		panic("open log file err" + err.Error())
	}
	err = ginLogan.open()
	if err != nil {
		panic("open log file err" + err.Error())
	}
	defer logan.close()
	defer ginLogan.close()

	r := engine()

	r.Run(":3000")
}
