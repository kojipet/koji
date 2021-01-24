package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func delay(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func logTime() string {
	whatTime := fmt.Sprintf("%s", time.Now())
	timeStr := fmt.Sprintf(brightblack+" - %s", whatTime[:24])
	return timeStr
}

// Debug is the most verbose log level
func Debug(msg string) {
	// log level 1
	if logLevel <= 1 {
		fmt.Println(brightwhite + "[debug]\t" + white + msg + logTime())
	}
}

func roll(min, max int) int {
	return rand.Intn(max-min) + min
}

// Info is a normal message
func Info(msg string) {
	// log level 2
	if logLevel <= 2 {
		fmt.Println(brightcyan + "[info ]\t" + cyan + msg + logTime())
	}
}

// Error is a message that is the least verbose logging level
func Error(msg string) {
	// log level 3
	if logLevel <= 3 {
		fmt.Println(brightred + "[error]\t" + red + msg + logTime())
	}
}

func chooseFace(faceSet []string) string {
	return faceSet[roll(0, len(faceSet))]
}

func kojiHappyFace(k *koji) {
	fmt.Println(brightcyan + k.name + " " + nc + chooseFace(happyFaces))
}

func kojiSadFace(k *koji) {
	fmt.Println(brightred + k.name + " " + nc + chooseFace(sadFaces))
}

func kojiStressFace(k *koji) {
	fmt.Println(brightyellow + k.name + " " + nc + chooseFace(stressFaces))
}

func kojiDeadFace(k *koji) {
	fmt.Println(brightwhite + k.name + " " + nc + chooseFace(deadFaces))
}

func kojiEggFace(k *koji) {
	fmt.Println(brightcyan + k.name + " " + nc + chooseFace(eggFaces))
}
func kojiSleepFace(k *koji) {
	fmt.Println(brightcyan + k.name + " " + nc + chooseFace(sleepFaces))
}

func createDirIfItDontExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		handle("Could not create directory: ", err)
	}
}

func writeFile(filename, textToWrite string) {
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	handle("", err)
	defer file.Close()
	_, err = file.WriteString(textToWrite)
	err = file.Sync()
	handle("", err)
}

func createFile(filename string) {
	var _, err = os.Stat(filename)
	if os.IsNotExist(err) {
		var file, err = os.Create(filename)
		handle("", err)
		defer file.Close()
	}
}

func writeStrings(filename string, textToWrite []string) {
	if _, err := os.Stat("/path/to/your-file"); os.IsNotExist(err) {
		createFile(filename)
	}
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	handle("", err)
	defer file.Close()
	joinedStr := fmt.Sprintf(strings.Join(textToWrite, ""))
	_, err = file.WriteString(joinedStr)
	err = file.Sync()
	handle("", err)
}

func handle(msg string, err error) {
	if err != nil {
		fmt.Printf(brightred+"\n%s: %s"+nc, msg, err)
	}
}

func readFile(filename string) string {
	text, err := ioutil.ReadFile(filename)
	handle("Couldnt read the file: ", err)
	return string(text)
}

func writeBinFile(filename string, filedata []byte) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		createFile(filename)
	}
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	handle("", err)
	defer file.Close()
	file.Truncate(0)

	err = binary.Write(file, binary.LittleEndian, filedata)

	if err != nil {
		handle("Binary write failed: ", err)
	}
}

func readBinFile(filename string, filesize int) []byte {
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	handle("", err)
	defer file.Close()

	data := make([]byte, filesize)
	_, err = file.Read(data)

	if err != nil {
		handle("Binary read failed: ", err)
	}

	return data
}
