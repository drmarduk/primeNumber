package main

import (
  "fmt"
  "os"
  "strconv"
  "flag"
  "time"
  )

// Array of futur prime numbers
var primeNumber []int
var starting_time time.Time

func main() {
  flag.Parse()
  input, err := parseInput(flag.Arg(0))
  if err != nil {
    fmt.Println("Error: your input is not a valid integer")
    os.Exit(1)
  } else {
    fmt.Println("Getting the primary number at position of", input)
    output(work(input))
  }
}

func parseInput(s string) (int, error){
  // Get the position of the wanted number from CLI
  target, err := strconv.Atoi(s)
  if err != nil {
    return 0, err
  }
  return target, nil
}

func work(target int) int{
  // Set the number that will involve in the program
  working := 2
  // Get the current time to output the time of processing
  starting_time = time.Now()
  // Infinite loop
  for {
    // Launch the validation of primary number
    prim(working)
    // if the size of the array is the wanted target
    // the program should stop itself
    if len(primeNumber) >= target {
      break
    }else{
      // Increment the working number and loop
      working++  
    }
  }
  return primeNumber[target - 1]
}

func prim(number int) {
  // For each elements in the primary number array
  for _, e := range primeNumber {
    // If the result of the division is an integer
    // It's not a primary number so go away
    if (e != 1 && number % e == 0) {
      return
    }
  }
  // Validation finished and the number is
  // a prime number
  primeNumber = append(primeNumber, number)
  return
}

func output(target int){
  // Getting the time resulting
  duration := time.Since(starting_time)
  // Output the result
  fmt.Println("Got it in", duration.String(), ":", target)
  return
}