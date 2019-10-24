package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("cmd/sortstr/input")
	if err != nil {
		log.Fatalf("Reading file failed: %v", err)
		os.Exit(-1)
	}
	arr := strings.Split(strings.TrimSpace(string(dat)), " ")
	arr2 := make([]string, len(arr))
	copy(arr2, arr)

	fmt.Println("Puzirek sort:")
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println()

	fmt.Println("sort.Strings:")
	sort.Strings(arr2)
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("%v ", arr2[i])
	}
	fmt.Println()
}
