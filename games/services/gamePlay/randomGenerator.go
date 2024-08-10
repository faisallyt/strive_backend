package gameplay

const (
	profitPercentage = 0.2
)

func ComputeResult(multiplier float64, amount float64, username string) bool {
	// logic to determine win or lose on the basis of multiplier and amount
	y := (1 - profitPercentage) / float64(multiplier)
	return statisticalResultGenerator(y)
}

func statisticalResultGenerator(probability float64) bool {
	// logic to generate result on the basis of probability
	return true
}
