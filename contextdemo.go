package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func pingGoogle() error {
	res, err := http.Get("https://google.com")
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("got %s from Google", res.Status)
	}
	return nil
}

func pingWithContext(ctx context.Context) error {
	errc := make(chan error)
	go func() { errc <- pingGoogle() }()
	select {
	case err := <-errc:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	fmt.Println("main")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	if err := pingWithContext(ctx); err != nil {
		log.Printf("could not ping Google: %v", err)
	}
}
