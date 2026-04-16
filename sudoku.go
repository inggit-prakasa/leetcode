package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Global variable biar gampang (di production jangan gini ya, ini demo doang!)
// 0 artinya kosong
var board = [9][9]int{
	{5, 3, 0, 0, 7, 8, 9, 0, 2}, // 3 kosong
	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	{0, 9, 8, 3, 0, 2, 5, 6, 7}, // 2 kosong
	{8, 5, 0, 7, 6, 1, 4, 2, 3}, // 1 kosong
	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 6, 1, 5, 3, 7, 2, 8, 0}, // 1 kosong
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 0, 0, 2, 8, 0, 0, 7, 9}, // 4 kosong (Total sekitar 11-15 an di contoh ini, tambahin 0 sendiri kalau kurang pedes!)
}

func main() {
	clearScreen()
	fmt.Println("🔥 ALGOBRO SUDOKU SOLVER (GOLANG EDITION) 🔥")
	fmt.Println("Wait for it... compiling logic...")
	time.Sleep(2 * time.Second)

	if solve() {
		clearScreen()
		printBoard()
		fmt.Println("\n✅ SUKSES ABANGKUH! Menyala! 🔥")
	} else {
		fmt.Println("\n❌ Yah, board-nya unsolveable bro. Skill issue yang bikin soal.")
	}
}

// Ini fungsi DFS + Backtracking-nya
func solve() bool {
	// 1. Cari kotak kosong (yang angkanya 0)
	row, col, found := findEmpty()

	// Base Case: Kalau gak ada yang kosong, berarti KELAR! 🎉
	if !found {
		return true
	}

	// 2. Coba angka 1 sampai 9
	for num := 1; num <= 9; num++ {
		if isValid(row, col, num) {
			board[row][col] = num // ACTION: Taruh angka

			// --- VISUALISASI ---
			// Clear screen & print board biar keliatan animasinya
			// Warning: Ini bikin lambat karena I/O bound + Sleep
			clearScreen()
			printBoard()
			// time.Sleep(50 * time.Millisecond) // Atur speed di sini (makin kecil makin ngebut)
			// -------------------

			if solve() { // RECURSE
				return true
			}

			board[row][col] = 0 // BACKTRACK: Undo langkah tadi 🚩
		}
	}

	return false // Trigger backtrack ke level sebelumnya
}

// Helper buat nyari kotak kosong
func findEmpty() (int, int, bool) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == 0 {
				return r, c, true
			}
		}
	}
	return -1, -1, false
}

// Logic Validasi (The Police 👮‍♂️)
func isValid(row, col, num int) bool {
	// Cek Baris & Kolom
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		} // Udah ada di baris
		if board[i][col] == num {
			return false
		} // Udah ada di kolom
	}

	// Cek Kotak 3x3
	// Rumus sakti: (index / 3) * 3
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

// --- FUNGSI TAMPILAN (UI) ---

func printBoard() {
	fmt.Println("-------------------------")
	for i := 0; i < 9; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("|-------+-------+-------|")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				fmt.Print("| ")
			}
			if board[i][j] == 0 {
				fmt.Print(". ") // Tampilkan titik kalo kosong
			} else {
				// Pake warna ijo biar hacker vibes (ANSI Escape Code)
				fmt.Printf("\033[32m%d\033[0m ", board[i][j])
			}
		}
		fmt.Println("|")
	}
	fmt.Println("-------------------------")
}

func clearScreen() {
	c := exec.Command("cmd", "/c", "cls")
	// Kalau lo pake Windows, ganti jadi exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}
