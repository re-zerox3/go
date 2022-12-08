package main

import(
	"net/http"
	"fmt"
	"log"
	"strconv"
)

func checkErr(err error){
				if err!= nil{
								log.Fatal(err)
				}
}


func Conversion(writer http.ResponseWriter, request *http.Request, conversion string){
	f:= request.URL.Query().Get("value")
	if f==""{
					msg :="ERROR: MISSING GET PARAMETER 'value'\nusage:http://localhost:5000/"+conversion+"?value=10" 
					_, err0 := writer.Write([]byte(msg))
					checkErr(err0)
	}else{
					x,err1:= strconv.ParseFloat(f,64)
					if err1 != nil{
									msg := "ERROR: BAD VALUE 'value'\nusage:http://localhost:5000/"+conversion+"?value=10" 
									_, err2 := writer.Write([]byte(msg))
									checkErr(err2)
					}else{
									switch conversion {
									case "FtoC":
												c := (5.0/9.0)*(x-32)
												_, err3 := writer.Write([]byte(fmt.Sprintf("%.3f F => %.3f C",x,c)))
												checkErr(err3)
				
									case "MtoK":  
												k := x *1.609344
												_, err3 := writer.Write([]byte(fmt.Sprintf("%.3f M => %.3f K",x,k)))
												checkErr(err3)

									case "GtoL": 
												l := x*3.7854
												_, err3 := writer.Write([]byte(fmt.Sprintf("%.3f G => %.3f L",x,l)))
												checkErr(err3)
								}
				}
		}
}


func FtoC(writer http.ResponseWriter, request *http.Request){
				Conversion(writer, request, "FtoC")
}


func MtoK(writer http.ResponseWriter, request *http.Request){
				Conversion(writer, request, "MtoK")
}



func GtoL(writer http.ResponseWriter, request *http.Request){
				Conversion(writer, request, "GtoL")
}


func Root(writer http.ResponseWriter,request *http.Request){
				_, err := writer.Write([]byte("NO ROUTE SPECIFIED\nusage:http://localhost:5000/routeName"))
				checkErr(err)
}


func main(){
	fmt.Println("Hello World.............")
	http.HandleFunc("/FtoC", FtoC)
	http.HandleFunc("/MtoK", MtoK)
	http.HandleFunc("/GtoL", GtoL)
	http.HandleFunc("/",Root)
	err := http.ListenAndServe("localhost:5000",nil)
	checkErr(err)
}



