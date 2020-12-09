package HttpService

import (
	"context"
	"fmt"
	"net/http"
)

func Service(ctx context.Context,addr string) error {
	mux := 	http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("http service")
	})

	var s = &http.Server{
		Handler: mux,
		Addr:    addr,
	}
	go func() {
		<-ctx.Done()
		fmt.Println("Server Shutdown")
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}
