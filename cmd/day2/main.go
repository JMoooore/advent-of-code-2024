package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
  final_answer := 0
  file, err := os.Open("cmd/day2/day2.txt") 
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    //fmt.Println(scanner.Text())
    l_slice := strings.Split(scanner.Text(), " ")
    i_slice := make([]int, len(l_slice))
    for i, val := range l_slice {
      int, _ := strconv.Atoi(val)
      i_slice[i] = int
    }
    fmt.Println("------------------------------")
    fmt.Println(l_slice)
    fmt.Println(calcSafe(i_slice))
    if calcSafe(i_slice) {
      final_answer++
    } else {
      for i := 0; i < len(i_slice); i++ {
        damp := append([]int(nil), i_slice[:i]...)
        damp = append(damp, i_slice[i+1:]...)
        if calcSafe(damp) {
          final_answer++
          break
        }
      }
    } 
  }

  fmt.Printf("Final answer: %d\n", final_answer)
}

// 51 52 55 58 60 61 62 61
// Take in a slice of integers and calc whether it is safe or not
func calcSafe(line []int) bool {
  if line[0] > line[1] {
    slices.Reverse(line)
  }

  l, r := 0, 1
  for r < len(line) {
    if line[l] >= line[r] {
      return false
    }
    if line[r] - line[l] > 3 {
      return false
    }
    l++
    r++
  }
  return true
}

