package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

 // Abs returns the absolute value of x.
func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // n: the number of temperatures to analyse
    var n int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&n)

    scanner.Scan()
    temps := scanner.Text() // the n temperatures expressed as integers ranging from -273 to 5526

    closestTemp := 5526
    for _, v := range strings.Split(temps, " ") {
    	currentTemp,_ := strconv.Atoi(v)
    	switch  {
    		case Abs(currentTemp) == Abs(closestTemp) && currentTemp < closestTemp:
    			closestTemp = closestTemp
    		case Abs(currentTemp) == Abs(closestTemp) && currentTemp > closestTemp:
    			closestTemp = currentTemp
    		case Abs(closestTemp) > Abs(currentTemp):
    			closestTemp = currentTemp
    	}
    }
    // fmt.Fprintln(os.Stderr, "Debug messages...")

    fmt.Println(closestTemp)// Write answer to stdout
}