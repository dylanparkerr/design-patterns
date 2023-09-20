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
    fmt.Println(myDice.rollStrat())
    fmt.Println(myDice2.rollStrat())
}
