package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	array100 := rand.Perm(100)
	array1000 := rand.Perm(1000)
	array10000 := rand.Perm(10000)

	// insertionSort
	start := time.Now()
	insertionSort(array100)
	fmt.Println("[+] Tempo para 100 elementos usando insertionSort:", time.Since(start))

	start = time.Now()
	insertionSort(array1000)
	fmt.Println("[+] Tempo para 1000 elementos usando insertionSort:", time.Since(start))

	start = time.Now()
	insertionSort(array10000)
	fmt.Println("[+] Tempo para 10000 elementos usando insertionSort:]", time.Since(start))
	fmt.Println()

	// bubbleSort
	start = time.Now()
	bubbleSort(array100)
	fmt.Println("[+] Tempo para 100 elementos usando bubbleSort:", time.Since(start))

	start = time.Now()
	bubbleSort(array1000)
	fmt.Println("[+] Tempo para 1000 elementos usando bubbleSort:", time.Since(start))

	start = time.Now()
	bubbleSort(array10000)
	fmt.Println("[+] Tempo para 10000 elementos usando bubbleSort:", time.Since(start))
	fmt.Println()

	// selectionSort
	start = time.Now()
	selectionSort(array100)
	fmt.Println("[+] Tempo para 100 elementos usando selectionSort:", time.Since(start))

	start = time.Now()
	selectionSort(array1000)
	fmt.Println("[+] Tempo para 1000 elementos usando selectionSort:", time.Since(start))

	start = time.Now()
	selectionSort(array10000)
	fmt.Println("[+] Tempo para 10000 elementos usando selectionSort:]", time.Since(start))
}
