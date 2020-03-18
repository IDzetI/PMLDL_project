package main

import (
	"math"
)

//coordinate to degree
func xyzToDegrees(x [3]float64) [4]float64 {
	h := float64(5)
	r := float64(50)
	R := float64(48)
	//lengths of zero position

	lengths := [4]float64{
		5579.27704867277,
		5563.2328940038105,
		5577.328229187914,
		5593.331441243491}
	// coordinate of point C of every column
	C := [][]float64{
		{4450, -1930, 2650},
		{-4430, -1930, 2650},
		{-4430, 1970, 2650},
		{4450, 1970, 2650}}
	/*
		lengths := [4]float64{
			5680.6905447393065,
			5737.346609387781,
			5736.5756690723565,
			5683.231745467436}
		// coordinate of point C of every column
		C := [][] float64{
			{4366.35, -1943.798, 2970.303},
			{-4435.848, -1950.001, 2971.382},
			{-4433.995, 1948.413, 2973.691},
			{4368.262, 1942.21, 2973.331}}


	*/
	// calculate degrees from formulas
	for i := 0; i < 4; i++ {
		cosB := (x[0] - C[i][0]) /
			(math.Sqrt(math.Pow(x[0]-C[i][0], 2) + math.Pow(x[1]-C[i][1], 2)))
		sinB := (x[1] - C[i][1]) /
			(math.Sqrt(math.Pow(x[0]-C[i][0], 2) + math.Pow(x[1]-C[i][1], 2)))

		Cs := []float64{
			C[i][0] + r*cosB,
			C[i][1] + r*sinB,
			C[i][2]}

		cosE := r / math.Sqrt(math.Pow(x[0]-Cs[0], 2)+math.Pow(x[1]-Cs[1], 2)+math.Pow(x[2]-Cs[2], 2))

		cosD := math.Sqrt(math.Pow(x[0]-Cs[0], 2)+math.Pow(x[1]-Cs[1], 2)) /
			math.Sqrt(math.Pow(x[0]-Cs[0], 2)+math.Pow(x[1]-Cs[1], 2)+math.Pow(x[2]-Cs[2], 2))

		gamma := math.Acos(cosE) + math.Acos(cosD)

		B := []float64{
			Cs[0] + r*math.Cos(gamma)*cosB,
			Cs[1] + r*math.Cos(gamma)*sinB,
			Cs[2] + r*math.Sin(gamma)}

		L := r*(math.Pi-gamma) + math.Sqrt(math.Pow(x[0]-B[0], 2)+math.Pow(x[1]-B[1], 2)+math.Pow(x[2]-B[2], 2))
		//fmt.Println(L)
		H := h * (L - lengths[i]) / (2 * math.Pi * r)
		lengths[i] = (L - H - lengths[i]) / R / math.Pi * 180
	}
	//fmt.Println(lengths)
	return lengths
}
