package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 * ---
 * Hint: You can use the debug stream to print initialTX and initialTY, if Thor seems not follow your orders.
 **/

func getDir(dirX, dirY int) string {
	fmt.Fprintln(os.Stderr, fmt.Sprintf("dirX =%i dirY =%i ", dirX, dirY))
    if dirX * dirY == 0 {
 		switch {
 			case dirX == 0 && dirY > 0:
 				return "S"
  			case dirX == 0 && dirY < 0:
 				return "N"
 			case dirX > 0 && dirY == 0:
 				return "E"
 			case dirX < 0 && dirY == 0:
 				return "W"
 		}
    } else {
    	switch {
 			case dirX > 0 && dirY > 0:
 				return "SE"
  			case dirX > 0 && dirY < 0:
 				return "NW"
 			case dirX < 0 && dirY > 0:
 				return "SW"
 			case dirX < 0 && dirY < 0:
 				return "NE"
 		}
    }
    return "E"
}

func moveThor(direction string, initialTX, initialTY int ) (int, int){

	switch {
        case direction == "S":
        	initialTY ++

        case direction == "N":
        	initialTY --

        case direction == "E":
        	initialTX ++

        case direction == "W":
        	initialTX --

        case direction == "SE":
        	initialTY ++
        	initialTX ++

        case direction == "SW":
        	initialTY ++
        	initialTX --

        case direction == "NW":
        	initialTY --
        	initialTX ++

        case direction == "NE":
        	initialTY --
        	initialTX --
    }

    return initialTX, initialTY;
}

func main() {
    // lightX: the X position of the light of power
    // lightY: the Y position of the light of power
    // initialTX: Thor's starting X position
    // initialTY: Thor's starting Y position
    var lightX, lightY, initialTX, initialTY int
    fmt.Scan(&lightX, &lightY, &initialTX, &initialTY)

    for {
        // remainingTurns: The remaining amount of turns Thor can move. Do not remove this line.
        var remainingTurns int
        fmt.Scan(&remainingTurns)

        // fmt.Fprintln(os.Stderr, "Debug messages...")
        direction := getDir(lightX - initialTX, lightY - initialTY)
        fmt.Println(direction) // A single line providing the move to be made: N NE E SE S SW W or NW

        initialTX, initialTY = moveThor(direction,initialTX, initialTY)
    }
}