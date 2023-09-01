package main

import (
	"fmt"
	"strconv"
)

func main() {
	existingVersions := []float64{0.2, 0.4, 0.7, 0.8, 1.2}
	newVersionStr := "1.2"
	newVersion, _ := strconv.ParseFloat(newVersionStr, 32)
	fmt.Println(newVersion)
	// return the nearest floating point numbers not the same
	isNewerVersion := true
	for _, version := range existingVersions {
		if newVersion <= version {
			isNewerVersion = false
			break
		}
	}
	fmt.Println(isNewerVersion)
}
