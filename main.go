package main

import ()

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
}