package main

import router "golang/internal/routers"

func main() {
	router := router.NewRouter()
	router.Run(":8081")
}
