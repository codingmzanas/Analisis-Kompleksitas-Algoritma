package main

import (
	"fmt"
	"strings"
	"time"
)

func InsertionSortRecursive(names []string, n int) {
	if n <= 1 {
		return
	}

	InsertionSortRecursive(names, n-1)
	last := names[n-1]
	j := n - 2

	for j >= 0 && strings.Compare(names[j], last) > 0 {
		names[j+1] = names[j]
		j--
	}
	names[j+1] = last
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
	InsertionSortRecursive(subset, len(subset))
	duration := time.Since(start) // Hitung durasi

	fmt.Println("Setelah diurutkan:", subset)
	fmt.Printf("Waktu eksekusi: %.6f detik\n", duration.Seconds())
}
