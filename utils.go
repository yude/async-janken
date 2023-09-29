package main

func Judge(One int, Two int) string {
	// Input:
	// - 0: ✊
	// - 1: ✌
	// - 2: ✋
	// Output:
	// - If winner exists: One or Two
	// - If draw: Draw

	r := (One - Two + 3) % 3

	if r == 0 {
		return "Draw"
	}

	if r == 2 {
		return "One"
	}

	return "Two"
}
