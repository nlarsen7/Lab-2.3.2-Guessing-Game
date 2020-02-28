//Nicholas Larsen
//February 27, 2020
//Simple guessing game where the user tries to guess the computer's random number
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  var user int
  rand.Seed(time.Now().UnixNano())
  computer:=(rand.Intn(1000)+1)
  //setting variables for user and computer's number

  fmt.Println("Try and guess what number i am thinking of.  It is a number between 1-1000.  Enter your number now.")
  
  fmt.Scanln(&user)
  //user's choice

  for user<computer{
    fmt.Println("Wrong number, try one higher than that")
    fmt.Scanln(&user)

  if user>computer{
    fmt.Println("Wrong number, try one lower than that.")
    fmt.Scan(&user)
    //tells user if number is too high or low and inputs again
  }
  }
  if user==computer{
    fmt.Println("You got it!!")
    //tells user they got the number right!
  }
  }
