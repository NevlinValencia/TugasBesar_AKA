package main

import (
	"fmt"
	"math"
	"time"
)

// Iterative method to calculate total electricity bill
func calculateBillIterative(units int, ratePerUnit float64) float64 {
	total := 0.0
	for i := 1; i <= units; i++ {
		total += ratePerUnit
	}
	return total
}

// Recursive method to calculate total electricity bill
func calculateBillRecursive(units int, ratePerUnit float64) float64 {
	if units == 0 {
		return 0
	}
	return ratePerUnit + calculateBillRecursive(units-1, ratePerUnit)
}

func main() {
	fmt.Println("=== Aplikasi Pembayaran Tagihan PLN ===")

	var units int
	var ratePerUnit float64

	fmt.Print("Masukkan jumlah unit listrik yang digunakan: ")
	fmt.Scan(&units)
	fmt.Print("Masukkan tarif per unit listrik: ")
	fmt.Scan(&ratePerUnit)

	// Measure execution time for iterative calculation
	iterativeStart := time.Now()
	iterativeTotal := calculateBillIterative(units, ratePerUnit)
	iterativeDuration := time.Since(iterativeStart)
	fmt.Printf("Total tagihan listrik (metode iteratif): Rp %.2f\n", iterativeTotal)
	fmt.Printf("Waktu eksekusi metode iteratif: %v\n", iterativeDuration)

	// Measure execution time for recursive calculation
	recursiveStart := time.Now()
	recursiveTotal := calculateBillRecursive(units, ratePerUnit)
	recursiveDuration := time.Since(recursiveStart)
	fmt.Printf("Total tagihan listrik (metode rekursif): Rp %.2f\n", recursiveTotal)
	fmt.Printf("Waktu eksekusi metode rekursif: %v\n", recursiveDuration)

	// Compare execution times
	fmt.Println("\n=== Perbandingan Waktu Eksekusi ===")
	if iterativeDuration < recursiveDuration {
		fmt.Printf("Metode iteratif lebih cepat dibandingkan metode rekursif dengan selisih waktu: %s\n", recursiveDuration-iterativeDuration)
	} else if recursiveDuration < iterativeDuration {
		fmt.Printf("Metode rekursif lebih cepat dibandingkan metode iteratif dengan selisih waktu: %s\n", iterativeDuration-recursiveDuration)
	} else {
		fmt.Println("Kedua metode memiliki waktu eksekusi yang sama.")
	}

	if math.Abs(iterativeTotal-recursiveTotal) < 0.01 {
		fmt.Println("Hasil tagihan metode iteratif dan rekursif sesuai.")
	} else {
		fmt.Println("Terdapat perbedaan dalam hasil tagihan.")
	}
}
