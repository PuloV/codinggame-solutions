package main

import "fmt"
// import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    for {
        var spaceX, spaceY int
        fmt.Scan(&spaceX, &spaceY)
        action := "HOLD"

        foundMontainOneLevelAbove := false
        foundMontainMoreLevelsAbove := false

        for i := 0; i < 8; i++ {
            // mountainH: represents the height of one mountain, from 9 to 0. Mountain heights are provided from left to right.
            var mountainH int
            fmt.Scan(&mountainH)
            if mountainH == spaceY - 1 && spaceX == i {
                action = "FIRE"
            }

            if mountainH == spaceY - 1 {
                foundMontainOneLevelAbove = true
            }

            if mountainH != 0 && spaceX == i  {
                foundMontainMoreLevelsAbove = true
            }
            // fmt.Fprintln(os.Stderr, fmt.Sprintf("More: %b One: %b ", foundMontainMoreLevelsAbove, foundMontainOneLevelAbove))
        }
        if !foundMontainOneLevelAbove && foundMontainMoreLevelsAbove {
            action = "FIRE"
        }
        // fmt.Fprintln(os.Stderr, mountainH)

        fmt.Println(action) // either:  FIRE (ship is firing its phase cannons) or HOLD (ship is not firing).
    }
}