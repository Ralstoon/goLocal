package main

import (
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func router1(wr http.ResponseWriter,r *http.Request){
	msg,err := ioutil.ReadAll(r.Body)
	if err!=nil{
		wr.Write([]byte("router1 error"))
	}
	writeLen,err := wr.Write(msg)
	if err!=nil || writeLen!=len(msg){
		log.Println(err,"write len:",writeLen)
	}
}

func main(){
	http.HandleFunc("/",router1)
	err := http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatal(err)
	}


	
}