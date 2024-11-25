package main

import "example.com/web-service-gin/app"

func main() {
	router := app.GetRouter();
	router.Run("localhost:8080")
}