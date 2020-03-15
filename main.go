package main

import (
	"net/http"

	"./logoutput"
)

const Port = "8080"
const Ipdata = ""

var Logout logoutput.Data

func webstart() {
	http.HandleFunc("/uploadlist", uploadlist)
	http.HandleFunc("/upload/", upload)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./html"))))
	http.ListenAndServe(Ipdata+":"+Port, nil)
}

func main() {
	Logout.Setup("log.log")
	Logout.Out(1, "%v:%v Webserver start\n", Ipdata, Port)
	// fmt.Printf("%v:%v Webserver start\n", Ipdata, Port)
	webstart()
}
