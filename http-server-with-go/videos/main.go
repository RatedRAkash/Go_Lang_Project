package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	http.HandleFunc("/", HandleGetVideos)

	http.HandleFunc("/post-videos", HandlePostVideos)

	// "nil" = Default Multiplexer
	http.ListenAndServe(":8080", nil)
}

func HandleGetVideos(w http.ResponseWriter, r *http.Request){
	for header, value := range r.Header {
		fmt.Printf("Key: %v \t Value: %v \n", header, value)
	}

	w.Header().Add("RamosCustomHeader", "TestValue")

	w.Write([]byte("Hola Ramos... Here is the Video Lists:\n"))

	// getting the Videos List
	videos := getVideos()

	videoBytes, err := json.Marshal(videos)

	if err!= nil {
		panic(err)
	}

	w.Write(videoBytes)
}

func HandlePostVideos(w http.ResponseWriter, r *http.Request){
	// Post Body (hit with a List of Videos with 3rd Bracket)
	//[
	// 	{   
	// 		"Id":"Cristiano Ronaldo",
	// 		"Title":"Portugese Superstar",
	// 		"Description":"He is the GOAT",
	// 		"Url":"https://cr7.com"
	// 	}
	// ]
	
	if r.Method == "POST" {
		// Converting Json Body Payload to "bytes"
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var videos []video
		err = json.Unmarshal(body, &videos)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Bad Request")
		}
		
		// add the Posted Videos to our LIST
		saveVideos(videos)

	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method Not Supported!")
	}
	
	// N.B.: "} else {" we have to give like this in the SAME LINE, otherwise Error will Occur
}