package main

import (
	"fmt"
	"os"
)

func main() {
	saveFile("firstFile", printThis())

}

func printThis() string {

	lyrics := "[Chorus] \n Shawty's like a melody in my head \n That I can't keep out, got me singing like \n Na-na-na-na, every day \n It's like my iPod stuck on replay, replay-ay-ay-ay \n Shawty's like a melody in my head \n That I can't keep out, got me singing like (Ayy) \n Na-na-na-na, every day \n It's like my iPod stuck on replay (J-J-J-JR), replay \n"

	return lyrics
}

func saveFile(fileName string, texts string) {
	fileData := []byte(texts)
	fileErr := os.MkdirAll("saves", os.ModePerm)
	if fileErr != nil {
		fmt.Println("cant create file")
	}

	owoError := os.WriteFile("saves/"+fileName+".txt", fileData, 0644)
	if owoError != nil {
		panic(owoError)
	}
	fmt.Println("file saved")
}
