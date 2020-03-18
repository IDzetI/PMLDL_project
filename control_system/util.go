package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

//read all file
func ReadAll(file string) string {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	return string(s)
}

//function to wait enter key
func waitEnter() {
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}

//string of 3 number to coordinate
func stringToCoordinates(str string) ([3]float64, bool) {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)

	//check number of inputs
	values := strings.Split(str, " ")
	if len(values) != 3 {
		fmt.Println("ERROR Please enter correct x y z")
		return [3]float64{0.0, 0.0, 0.0}, true
	}

	return stringArrayToCoordinates(values)
}

//array of 3 string number to coordinate
func stringArrayToCoordinates(values []string) ([3]float64, bool) {
	x := [3]float64{}
	for i := 0; i < 3; i++ {
		//convert strings to float64
		buff, err := strconv.ParseFloat(values[i], 64)
		if err != nil {
			fmt.Println(err)
			return x, true
		}
		x[i] = buff
	}
	return x, false
}

// compare two array of 3 float64
func compareArrayOfFloat64_3(a [3]float64, b [3]float64) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func lengthOfVector(v [3]float64) float64 {
	return math.Sqrt(math.Pow(v[0], 2) + math.Pow(v[1], 2) + math.Pow(v[2], 2))
}
func lengthOfVector4(v [4]float64) float64 {
	return math.Sqrt(math.Pow(v[0], 2) + math.Pow(v[1], 2) + math.Pow(v[2], 2) +  math.Pow(v[3], 2))
}


func distanceBeatenPoints(a [3]float64, b [3]float64) float64 {
	return lengthOfVector([3]float64{a[0] - b[0], a[1] - b[1], a[2] - b[2]})
}

func distanceBeatenPoints4(a [4]float64, b [4]float64) float64 {
	return lengthOfVector4([4]float64{a[0] - b[0], a[1] - b[1], a[2] - b[2], a[3] - b[3]})
}


func difOfVectors(a [4]float64, b [4]float64) [4]float64  {
	return [4]float64{a[0]-b[0],a[1]-b[1],a[2]-b[2],a[3]-b[3]}
}

func sumOfVectors(a [3]float64, b [3]float64) [3]float64 {
	return [3]float64{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}
func sumOfVectors4(a [4]float64, b [4]float64) [4]float64 {
	return [4]float64{a[0] + b[0], a[1] + b[1], a[2] + b[2],a[3] + b[3]}
}

func mulVectorByScalar(a [3]float64, b float64) [3]float64 {
	return [3]float64{a[0] * b, a[1] * b, a[2] * b}
}
func mulVectorByScalar4(a [4]float64, b float64) [4]float64 {
	return [4]float64{a[0] * b, a[1] * b, a[2] * b,a[3] * b}
}


func formatstring3(input [3]float64) string {
	// to convert a float number to a string
	str := ""
	for i := 0; i < len(input); i++ {
		str = str + "\t" + strconv.FormatFloat(input[i], 'f', 6, 64)
	}
	return str
}
func formatstring4(input [4]float64) string {
	// to convert a float number to a string
	str := ""
	for i := 0; i < len(input); i++ {
		str = str + "\t" + strconv.FormatFloat(input[i], 'f', 6, 64)
	}
	return str
}



func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func Float64Uint16(number float64) []uint16 {
	return (*[4]uint16)(unsafe.Pointer(&number))[:]
}


func Uint16Float64(u []uint16) float64  {
	return  math.Float64frombits(uint64(u[3])<<48 | uint64(u[2]) << 32 | uint64(u[1]) << 16 | uint64(u[0]))
}

func LineToAngles(s string) [4]float64 {
	values := strings.Split(s, " ")
	var ans [4]float64

	for i := 0; i < len(values); i++ {
		buff, err := strconv.ParseFloat(values[i], 64)
		ans[i] = buff
		if err != nil {
			fmt.Println(err)
		}
	}
	return ans
}

func checkCoordinate(cord [3]float64) bool {
	limX := [2]float64{-4000,4000}
	limY := [2]float64{-2000,2000}
	limZ := [2]float64{-500,1800}
	return checkValue(cord[0],limX[0],limX[1]) &&
			checkValue(cord[1],limY[0],limY[1]) &&
			checkValue(cord[2],limZ[0],limZ[1])
}

func checkValue(x float64, min float64, max float64) bool{
	return x >= min && x <= max
}


//getPosition
func getPosition(strPos []string, oldPos [3]float64, text string) [3]float64 {
	position, err := stringArrayToCoordinates(strPos)
	if err || !checkCoordinate(position){
		fmt.Println("The entered position is not correct")
		return oldPos
	}
	fmt.Print(text)
	fmt.Println(position)
	return position
}

//set value
func getValue(value string, oldValue float64, name string, min float64, max float64) float64 {
	newValue, err := strconv.ParseFloat(value, 64)
	if err != nil || !checkValue(newValue,min,max){
		fmt.Println("The entered " + name + " is not correct")
		return oldValue
	}
	fmt.Print("New " + name + " = ")
	fmt.Println(newValue)
	return newValue
}