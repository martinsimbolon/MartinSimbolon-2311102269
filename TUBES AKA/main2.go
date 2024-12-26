package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID        string
	UserID    string
	Amount    float64
	Timestamp time.Time
}

var operationCount int

func detectFraudIterative(transactions []Transaction, thresholdAmount float64, thresholdTime time.Duration) []string {
	suspicious := []string{}
	operationCount = 0 // Reset hitungan operasi

	for i := 0; i < len(transactions)-1; i++ {
		operationCount++

		current := transactions[i]
		next := transactions[i+1]

		operationCount++
		if current.UserID == next.UserID { // Transaksi oleh user yang sama
			operationCount++
			if next.Amount > thresholdAmount { // Jumlah transaksi terlalu besar
				suspicious = append(suspicious, next.ID)
			} else if next.Timestamp.Sub(current.Timestamp) < thresholdTime { // Waktu antar transaksi terlalu dekat
				operationCount++
				suspicious = append(suspicious, next.ID)
			}
		}
	}

	return suspicious
}

func main() {
	transactions := []Transaction{
		{"T1", "User1", 5000, time.Now().Add(-time.Hour)},
		{"T2", "User1", 12000, time.Now().Add(-30 * time.Minute)},
		{"T3", "User1", 1500, time.Now().Add(-10 * time.Minute)},
		{"T4", "User2", 3000, time.Now().Add(-2 * time.Hour)},
		{"T5", "User2", 9000, time.Now().Add(-1 * time.Minute)},
	}

	// Menambahkan transaksi dummy untuk memperbesar dataset
	for i := 6; i <= 10000; i++ {
		transactions = append(transactions, Transaction{
			ID:        fmt.Sprintf("T%d", i),
			UserID:    "User1",
			Amount:    9000 + float64(i), // Variasi jumlah
			Timestamp: time.Now().Add(time.Duration(-i) * time.Minute),
		})
	}

	thresholdAmount := 8000.0
	thresholdTime := 15 * time.Minute

	startTime := time.Now()

	// Gunakan fungsi iteratif
	suspiciousTransactions := detectFraudIterative(transactions, thresholdAmount, thresholdTime)

	durasi := time.Since(startTime)

	fmt.Println("Transaksi mencurigakan:")
	for _, id := range suspiciousTransactions {
		fmt.Println(" - ID Transaksi:", id)
	}

	fmt.Printf("Jumlah operasi (kompleksitas): %d\n", operationCount)
	fmt.Printf("Durasi eksekusi fungsi: %s (%d ns)\n", durasi, durasi.Nanoseconds())
}
