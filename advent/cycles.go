package advent

// Brent's Algorithm for cycle detection
//def brent(f, x0) -> (int, int):
//"""Brent's cycle detection algorithm."""
//
//return lam, mu

type Brentable interface {
	Eq(other Brentable) bool
}

// Brent implements "Brent's Algorithm for Cycle Detection". x0 is the start state of a sequence of states where f(x)
// calculates the following state. States much be comparable, and always deterministically produce the next state.
func Brent(f func(Brentable) Brentable, x0 Brentable) (int, int) {
	// Main phase: search successive powers of two:
	power := 1
	lam := 1

	tortoise := x0
	hare := f(x0) // f(x0) is the element/node next to x0.

	for !tortoise.Eq(hare) {
		if power == lam { // time to start a new power of two?
			tortoise = hare
			power *= 2
			lam = 0
		}

		hare = f(hare)
		lam += 1
	}

	// Find the position of the first repetition of length λ
	tortoise = x0
	hare = x0

	for i := 0; i < lam; i++ {
		hare = f(hare)
	}
	// The distance between the hare and tortoise is now λ.

	// Next, the hare and tortoise move at same speed until they agree
	mu := 0
	for !tortoise.Eq(hare) {
		tortoise = f(tortoise)
		hare = f(hare)
		mu++
	}

	return lam, mu
}
