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


func saveVideos(newVideos []video){
	videosCurrentList := getVideos()

	// ellipsis (...) operator, which unpacks the slice into individual elements
	allVideos := append(videosCurrentList, newVideos...)

	videoBytes, err := json.Marshal(allVideos)

	if err !=nil {
		panic(err)
	}
	
	// "0644" is the "Linux File Permission" for that particular File
	// updating the Existing "/videos.json", as we are Returning this for our "/" Home URL
	err = ioutil.WriteFile("./videos.json", videoBytes, 0644)

	if err !=nil {
		panic(err)
	}

}