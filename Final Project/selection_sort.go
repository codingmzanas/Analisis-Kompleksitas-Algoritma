package main 

import (
	"fmt"
	"strings"
	"time"
)

func SelectionSortIterative(names []string) {
	n := len(names)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if strings.Compare(names[j], names[minIndex]) < 0 {
				minIndex = j
			}
		}
		// Geser posisi
		names[i], names[minIndex] = names[minIndex], names[i]
	}
}

func main() {
	// Dataset nama
	names := []string{
		"Zaki", "Anisa", "Budi", "Citra", "Andi",
		"Dina", "Eka", "Fajar", "Gita", "Hendra",
		"Indah", "Joko", "Kirana", "Lutfi", "Mira",
		"Nanda", "Olivia", "Putra", "Qory", "Rizky",
		"Santi", "Teguh", "Umar", "Vina", "Winda",
	}

	fmt.Println("Dataset Lengkap:", names)

	var size int
	fmt.Print("Masukkan jumlah nama untuk diurutkan (5, 10, 15, 20, atau 25): ")
	fmt.Scan(&size)

	if size < 1 || size > len(names) {
		fmt.Println("Jumlah nama harus antara 1 dan 25.")
		return
	}

	// Ambil subset nama sesuai input
	subset := names[:size]
	fmt.Println("Sebelum diurutkan:", subset)

	start := time.Now() // Mulai pengukuran waktu
	SelectionSortIterative(subset)
	duration := time.Since(start) // Hitung durasi

	fmt.Println("Setelah diurutkan:", subset)
	fmt.Printf("Waktu eksekusi: %.6f detik\n", duration.Seconds())
}
