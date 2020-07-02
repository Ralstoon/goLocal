// package main

// import (
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"
// )

// var logger = log.New(os.Stdout,"#",0)

// func hello(wr http.ResponseWriter,r *http.Request){
// 	wr.Write([]byte("hello!"))
// 	logger.Println("hello request finished!")
// }

// func timeMiddleWare(next http.Handler) http.Handler{
// 	return http.HandlerFunc(func (wr http.ResponseWriter,r *http.Request){
// 		timeStart := time.Now()
// 		// 业务代码Hello
// 		// next.ServeHTTP(wr,r)
// 		timeElapsed := time.Since(timeStart)
// 		logger.Println("timeMiddleWare:",timeElapsed)
// 	})
// }


// func main(){
// 	http.Handle("/",timeMiddleWare(http.HandlerFunc(hello)))

// 	err := http.ListenAndServe(":9000",nil)
// 	if err!=nil{
// 		logger.Panicln("serve listen err :",err)
// 	}
// }