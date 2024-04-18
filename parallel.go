package main

import (
    "fmt"
    "sync"
)

func calculateSum(matrix [][]int, row int, wg *sync.WaitGroup, resultChan chan int) {
    defer wg.Done() // Oznaczenie goroutine jako zakończonej po zakończeniu obliczeń
    sum := 0
    for _, num := range matrix[row] {
        sum += num
    }
    resultChan <- sum // Wysłanie wyniku do kanału
}

func main() {
    matrix := [][]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
    }

    numWorkers := len(matrix)
    resultChan := make(chan int, numWorkers)

    var wg sync.WaitGroup
    wg.Add(numWorkers)

    // Rozpoczęcie obliczeń w wielu goroutines
    for i := 0; i < numWorkers; i++ {
        go calculateSum(matrix, i, &wg, resultChan)
    }

    // Funkcja anonimowa do zamykania kanału po zakończeniu wszystkich obliczeń
    go func() {
        wg.Wait()
        close(resultChan)
    }()

    // Odbieranie wyników z kanału i sumowanie ich
    totalSum := 0
    for sum := range resultChan {
        totalSum += sum
    }

    fmt.Println("Suma elementów macierzy:", totalSum)
}
