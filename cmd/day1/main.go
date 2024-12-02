package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
  "math"
)

func main() {
  file, err := os.Open("cmd/day1/day1a.txt")
  if err != nil {
    fmt.Println("Error opening file: ", err)
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  slice1 := []int{}
  slice2 := []int{}

  for scanner.Scan() {
    result := strings.Split(scanner.Text(), " ")
    num1 := strings.TrimSpace(result[0])
    num2 := strings.TrimSpace(result[3])
    i, err := strconv.Atoi(num1)
    if err != nil {
      fmt.Println("Error with Atoi")
      return
    }
    slice1 = append(slice1, i)
    j, err := strconv.Atoi(num2)
    if err != nil {
      fmt.Println("Error with Atoi")
      return
    }
    slice2 = append(slice2, j) 
  }
  sort.Ints(slice1)
  sort.Ints(slice2)

  sum := float64(0)
  for i, valA := range slice1 {
    diff := math.Abs(float64(valA - slice2[i]))
    sum += diff
  }
  fmt.Printf("Part 1: %.0f\n", sum)

  count := make(map[int]int)
  for _, valB := range slice2 {
    // 0 value of int is zero so no matter what just add 1?
    count[valB]++
  }
  similarityScore := 0
  for _, valA := range slice1 {
    similarityScore += valA * count[valA]
  }
  fmt.Printf("Part 2: %d\n", similarityScore)
}
