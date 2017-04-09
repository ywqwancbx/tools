package main

import (
	"os"
	"fmt"
	"net/http"
)

 /*
  * http static file server .current (run path)
  * go build static.go
  * cp static /usr/local/bin
  * cd somedir
  * static (test for mac)
  */
func main()  {
	var args = os.Args
	var port string
	var path string
	if len(args)>1{
		port = args[1]
	}else{
		port = ":8088"
	}
	if len(args)>2{
		path = args[2]
	}else{
		pa,err :=os.Getwd()
		if err !=nil{
			fmt.Println("not found path ereor:",err)
			os.Exit(10002)
		}else{
			path = pa
		}
	}
	fmt.Printf("listen port %s  dir %s \r\n",port,path)
	panic(http.ListenAndServe(port, http.FileServer(http.Dir(path))))
}
