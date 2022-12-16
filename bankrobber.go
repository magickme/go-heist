package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)
  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }
  openedVault := rand.Intn(100)
  leftSafely := rand.Intn(5)
  if ((isHeistOn ==true) && (openedVault >= 70)) {
    fmt.Println("Grab and GO!")    
  } else if isHeistOn ==true {
    isHeistOn = false
    fmt.Println("The vault can't be opened.")
  } else {
    fmt.Print("isHeistOn is currently: ", isHeistOn)
  }
  if isHeistOn ==true {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("The heist failed.")
      case 1:
        isHeistOn = false
        fmt.Println("The heist failed badly.")
      case 2:
        isHeistOn = false
        fmt.Println("The heist failed really badly.")
      case 3:
        isHeistOn = false
        fmt.Println("The heist critically failed.")
      default:
        fmt.Println("Start the getaway car!")
    }
  }
  if isHeistOn ==true {
    amtStolen := 1000 + rand.Intn(1000000)
    fmt.Println("The amount stolen was: $", amtStolen)
  }
}
