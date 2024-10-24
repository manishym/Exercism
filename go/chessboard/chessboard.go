package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
type File [8]bool
type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	ans := 0
	if f, ok := cb[file]; !ok {
		return 0
	} else {
		for _, elt := range f {
			if elt {
				ans++
			}
		}
	}
	return ans
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	ans := 0

	for _, val := range cb {
		if rank >= 1 && rank <= 8 && val[rank-1] {
			ans++
		}
	}
	return ans

}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	ans := 0
	for _, val := range cb {
		ans += len(val)
	}
	return ans
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	ans := 0
	for _, val := range cb {
		for _, x := range val {
			if x {
				ans++
			}
		}
	}
	return ans
}
