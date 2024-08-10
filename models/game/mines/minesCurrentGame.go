package mines

type CurrentGameData struct {
	Amount     float64          `json:"amount"`
	MinesCount int64            `json:"mines_count"`
	Rounds     []MinesRoundData `json:"rounds"`
}

type MinesRoundData struct {
	Amount           float64 `json:"amount"`
	BoxNo            int64   `json:"box_no"`
	BetNo            int64   `json:"bet_no"`
	Payout           float64 `json:"payout"`
	AmountMultiplier float64 `json:"amount_multiplier"`
	PayOutMultiplier float64 `json:"payout_multiplier"`
}
