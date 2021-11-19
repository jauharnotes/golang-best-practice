/*
2520 adalah angka terkecil yang dapat dibagi dengan masing-masing angka dari 1 sampai 10 tanpa sisa.
Berapa bilangan positif terkecil yang habis habis dibagi semua bilangan dari 1 sampai 20?
*/

package main

import (
   "fmt"
   "time"
)

func main() {
   before := time.Now()
   maxNumber:=20
   result:=1
   for i:=2;i<=maxNumber;i++ {
      result=kpk(result,i)
   }
   fmt.Println("Bilangan kelipatan terkecil dari 1 s.d. ",maxNumber,"adalah : ",result)
   after := time.Now()
   fmt.Println("Waktu eksekusi", after.Nanosecond()-before.Nanosecond(), "nano second")
}

func kpk(a int, b int) int {
   return (a * b) / ppb(a, b)
}
func ppb(a int, b int) int {
   for b != 0 {
      t := b
      b = a % b
      a = t
   }
   return a
}
