package main

import (
	"animeshop/internal/http"
	"animeshop/internal/store/inmemory"
	"context"
	"fmt"
)

func main() {
	store := inmemory.NewDB()

	srv := http.NewServer(context.Background(), ":8080", store)
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

	srv.WaitForGT()
}
