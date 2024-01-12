package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)



func main(){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx,"GET", "http://localhost:8080", nil)	
	if err != nil{
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
		if err != nil{
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)

}

func handler(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-time.After(5* time.Second):
		log.Println("Request processada com sucesso")
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
	}
}
