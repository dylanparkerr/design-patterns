package main

import (
    "fmt"
    "math/rand"
)

type RollStrategy func () int

func FairRollStrategy () int {
    return rand.Intn(5)+1
}

func RiggedRollStrategy () int {
    return 6
}

type Dice struct{
    rollStrat RollStrategy
}

func main(){
    myDice := Dice{rollStrat: FairRollStrategy}
    myDice2 := Dice{rollStrat: RiggedRollStrategy}
    fmt.Println("dice 1 rolled a: " + myDice.rollStrat())
    fmt.Println("dice 2 rolled a: " + myDice2.rollStrat())
}
