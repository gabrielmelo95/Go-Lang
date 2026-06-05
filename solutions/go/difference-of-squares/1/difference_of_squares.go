package differenceofsquares

func SquareOfSum(n int) int {
	sum := n * (1 + n) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	squareOfSum := SquareOfSum(n)
	sumOfSquare := SumOfSquares(n)
	if squareOfSum >= sumOfSquare {
		return squareOfSum - sumOfSquare
	}
	return sumOfSquare - squareOfSum
}
