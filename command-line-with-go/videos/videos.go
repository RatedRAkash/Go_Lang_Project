package main

import (
	"io/ioutil"
	"encoding/json"
)

type video struct{
	Id string
	Title string
	Description string
	Url string
}

func getVideos()(videos []video){

	// As we are Declaring the variable for the First Time we are using ":="
	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err !=nil {
		panic(err)
	}
	
	// Re-using the same "err" Variable that's why we don't use ":=" instead we use "="
	err = json.Unmarshal(fileBytes, &videos)

	if err !=nil {
		panic(err)
	}

	return videos
}


func saveVideos(videos []video){

	videoBytes, err := json.Marshal(videos)

	if err !=nil {
		panic(err)
	}
	
	// "0644" is the "Linux File Permission" for that particular File
	err = ioutil.WriteFile("./new-videos.json", videoBytes, 0644)

	if err !=nil {
		panic(err)
	}

}