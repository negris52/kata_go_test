package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numRom = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
}

var numUint8 = map[int]uint8{
	0: 48,
	1: 49,
	2: 50,
	3: 51,
	4: 52,
	5: 53,
	6: 54,
	7: 55,
	8: 56,
	9: 57,
}

var mathUint8 = map[string]uint8{
	"+": 43,
	"-": 45,
	"*": 42,
	"/": 47,
}

var romUint8 = map[string]uint8{
	"I": 73,
	"V": 86,
	"X": 88,
}

func mathOperationsCheck(s string) uint8 {
	opCount := 0
	var operation uint8
	for i := 0; i < len(s); i++ {
		// count all math operations in string
		for k := range mathUint8 {
			if s[i] == mathUint8[k] {
				opCount++
				operation = mathUint8[k]
				if opCount > 1 {
					panic("Error: too many arithmetic operators, please check input")
				}
			}
		}
	}
	return operation
}

func mixedNumbersType(s string) bool {
	aNum := 0
	rNum := 0
	for i := 0; i < len(s); i++ {
		// count Arabic numbers
		for k := range numUint8 {
			if s[i] == numUint8[k] {
				aNum++
			}
		}
		// count Rome numbers
		for k := range romUint8 {
			if s[i] == romUint8[k] {
				rNum++
			}
		}
	}
	if aNum != 0 && rNum != 0 {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please enter data for calculation")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		delimit := mathOperationsCheck(text)
		if mixedNumbersType(text) {
			panic("Error: Only Rome or Arabic numbers available in one calculation, please check input")
		}
		parsed := strings.Split(text, string(delimit))
		if len(parsed) != 2 {
			panic("Error: 2 numbers should be entered for calculation, please check input")
		}
		firstNumStr := strings.TrimSpace(parsed[0])
		secondNumStr := strings.TrimSpace(parsed[1])
		firstNumInt, _ := strconv.Atoi(firstNumStr)
		secondNumInt, _ := strconv.Atoi(secondNumStr)
		calcMode := 0
		for k := range numRom {
			if firstNumStr == numRom[k] {
				firstNumInt = k
				calcMode = 1
			}
			if secondNumStr == numRom[k] {
				secondNumInt = k
			}
		}
		if firstNumInt > 10 || secondNumInt > 10 || firstNumInt == 0 || secondNumInt == 0 {
			panic("Error: Only 1 to 10 or I to X numbers available for calculation, please check input")
		}

		result := 0
		switch {
		case delimit == mathUint8["+"]:
			result = firstNumInt + secondNumInt
		case delimit == mathUint8["-"]:
			result = firstNumInt - secondNumInt
		case delimit == mathUint8["*"]:
			result = firstNumInt * secondNumInt
		case delimit == mathUint8["/"]:
			result = firstNumInt / secondNumInt
		default:
			panic("Error: no arithmetic operation found")
		}

		if calcMode == 0 {
			fmt.Printf("%v\n", result)
		} else {
			if result <= 0 {
				panic("Exception: Negative or 0 result with Roman numbers")
			}
			roman := NewRoman()
			fmt.Printf("%v\n", roman.ToRoman(result))
		}
	}
}
