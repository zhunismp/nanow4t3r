package main

import "github.com/zhunismp/nanow4t3r/services/product/infrastructures"

func main() {
	
	// Initialize the application
	// Running with graceful shutdown
	infrastructures.Start()
}
