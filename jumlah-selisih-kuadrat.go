/*
Jumlah kuadrat dari sepuluh bilangan asli pertama adalah,
1² + 2² + 3² + ... + 10² = 385
Kuadrat jumlah dari sepuluh bilangan asli pertama adalah,
(1 + 2 + 3 + ... + 10)² = 55² = 3025
Oleh karena itu, selisih antara jumlah kuadrat dari sepuluh bilangan asli pertama dan kuadrat dari jumlah tersebut adalah 3025–385=2640.
Temukan selisih antara jumlah kuadrat dari seratus bilangan asli pertama dan kuadrat dari jumlah tersebut!
*/


package main

import (
	"fmt"
	"time"
)
// solusi-1
func main() {
	before := time.Now()
	sumSquares := 0
	squaresOfTheSum := 0
	numberOfData := 100_000
	for i := 1; i <= numberOfData; i++ {
		sumSquares += i * i
		squaresOfTheSum += i
	}
	diffSums := (squaresOfTheSum * squaresOfTheSum) - sumSquares
	fmt.Println("Selisis Jumlah Kuadrat Adalah : ",diffSums)
	after := time.Now()
	fmt.Println("Waktu eksekusi", after.Nanosecond()-before.Nanosecond(), "nano second")

	sumSquares=(numberOfData*(numberOfData+1)*(2*numberOfData+1))/6
	squaresOfTheSum=(numberOfData*(numberOfData+1))/2

  
  // solusi-2
  package main

import (
	"fmt"
	"time"
)

func main() {
	before := time.Now()
	sumSquares := 0
	numberOfData := 100_000
	squaresOfTheSum:=0
	sumSquares=(numberOfData*(numberOfData+1)*(2*numberOfData+1))/6
	sumData:=(numberOfData*(numberOfData+1))/2
	squaresOfTheSum = sumData*sumData
	diffSums :=  squaresOfTheSum - sumSquares
	fmt.Println("Selisis Jumlah Kuadrat Adalah : ",diffSums)
	after := time.Now()
	fmt.Println("Waktu eksekusi", after.Nanosecond()-before.Nanosecond(), "nano second")
}
}
