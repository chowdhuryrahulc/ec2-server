package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// docker build -t ec2dockerimage . : create docker image
// docker image ls: should show the image just creatred now
// docker run -p 8080:8081 -it ec2dockerimage: port maping.
//		8081 will be written in http.ListenAndServe
// 		8080 will be the port you access vi browser. Means "localhost:8080"
// go to localhost:8080
// localhost:8080/hi

/*
Old Approch to run AWS
	create AWS EC2 --> SSH into the VM --> Clone git repo --> Install docker --> Run docker project
New Approch
	Docker machine --> Direct to AWS

Process:
	1) go app 
	2) dockerize it 
	3) install aws cli 
	4) configure aws cli
	5) install docker machine
	6) create docker machine for your app, create EC2 instance and make the app live in one command
	7) switch to machine right from your local machine
	8) run the app
	9) docker-machine rm 

todo learn docker compose, docker swarm, html.EscapeString(r.URL.Path)
*/

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}