package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func split(musicPath string, start string, end string, title string) {
}

func main() {
	// parse commandline options
	musicPath := flag.String("music", "", "a path to music")
	csvPath := flag.String("csv", "", "a path to csv")
	flag.Parse()

	if *musicPath == "" && *csvPath == "" {
		flag.Usage()
		os.Exit(0)
	}

	// read and parse
	file, error := os.Open(*csvPath)
	if error != nil {
		fmt.Println("Failed to open csv:", error)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, error := reader.ReadAll()
	if error != nil {
		fmt.Println("Error reading file:", error)
		os.Exit(1)
	}

	for _, record := range records {
		start := record[0]
		end := record[1]
		title := record[2]
		outputPath := fmt.Sprintf("%s.mp3", title)

		fmt.Print("Splitting: ", title, "...")
		command := exec.Command("ffmpeg", "-i", *musicPath, "-acodec", "mp3", "-to", end, "-ss", start, outputPath)
		_, error := command.Output()
		if error != nil {
			fmt.Println("Error executing ffmpeg:", error)
		}
		fmt.Println("Done")
	}
}
