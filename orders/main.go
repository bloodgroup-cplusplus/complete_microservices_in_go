package main

import "context"

func main() {
	// initiate the transparency

	store := NewStore()
	svc := NewService(store)

	svc.CreateOrder(context.Background())

}
