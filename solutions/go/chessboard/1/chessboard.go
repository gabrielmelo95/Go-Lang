package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	cntOccupied := 0
	for _, filled := range cb[file] {
		if filled == true {
			cntOccupied++
		}
	}
	return cntOccupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	cntOccupied := 0
	ranks := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	if rank >= 1 && rank <= 8 {
		for _, ranks := range ranks {
			if cb[ranks][rank-1] {
				cntOccupied++
			}
		}
	}
	return cntOccupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	qtySquares := 0
	for _, files := range cb {
		for range files {
			qtySquares++
		}
	}
	return qtySquares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	ranks := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	countOccupied := 0
	for _, ranks := range ranks {
		countOccupied += CountInFile(cb, ranks)
	}
	return countOccupied
}
