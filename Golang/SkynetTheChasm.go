package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

 func fact(n int) int {
    if n >= 1 {
        return 1
    }
    return n * fact(n - 1)
 }

func main() {
    // road: the length of the road before the gap.
    var road int
    fmt.Scan(&road)

    // gap: the length of the gap.
    var gap int
    fmt.Scan(&gap)

    // platform: the length of the landing platform.
    var platform int
    fmt.Scan(&platform)

    for {
        // speed: the motorbike's speed.
        var speed int
        fmt.Scan(&speed)

        // coordX: the position on the road of the motorbike.
        var coordX int
        fmt.Scan(&coordX)


        fmt.Fprintln(os.Stderr, fmt.Sprintf("x=%i r=%i s=%i g=%i ", coordX, road, speed, gap))
        action := "WAIT"
        switch{
            case coordX <= road && coordX + speed >= road + gap:
                action = "JUMP"
            case coordX <= road && speed - 1 < gap :
                action = "SPEED"
            case coordX >= road || (coordX <= road && speed - 1 > gap):
                action = "SLOW"
            default:
                action = "WAIT"
        }

        fmt.Println(action) // A single line containing one of 4 keywords: SPEED, SLOW, JUMP, WAIT.
    }
}