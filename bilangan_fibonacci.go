/*
Setiap suku baru dalam deret Fibonacci dihasilkan dengan menjumlahkan dua suku sebelumnya. Dengan memulai dengan 1 dan 2, 10 suku pertama adalah:
1, 2, 3, 5, 8, 13, 21, 34, 55, 89, …
Dengan mempertimbangkan suku-suku dalam deret Fibonacci yang nilainya tidak melebihi empat juta, temukan jumlah suku-suku yang bernilai genap.
*/

package main

import (
   "fmt"
   "time"
)
func main() {
   before := time.Now()
   firstNumber := 0
   secondNumber := 1
   sumEven := 0
   sumOdd := 0
   const max = 4_000_000
   for secondNumber < max {
      buff := firstNumber + secondNumber
      if secondNumber%2 == 0 {
         fmt.Println("Bilangannya Genap dari FIB :", secondNumber)
         sumEven += secondNumber
      } else {
         fmt.Println("Bilangannya Ganjil dari FIB :", secondNumber)
         sumOdd += secondNumber
      }
      firstNumber = secondNumber
      secondNumber = buff
   }
   fmt.Println("==================================================================")
   fmt.Printf("Jumlah dari bilangan ganjil : [%d]\n", sumOdd)
   fmt.Printf("Jumlah dari bilangan genap : [%d]\n", sumEven)
   after := time.Now()
   fmt.Println("Waktu eksekusi ", (after.Nanosecond() - before.Nanosecond()), "nano second")
