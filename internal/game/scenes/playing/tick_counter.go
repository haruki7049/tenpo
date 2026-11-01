package playing

type tickCounter struct {
	count int
}

func newTickCounter() *tickCounter {
	tc := &tickCounter{}
	tc.count = 0

	return tc
}

func (tc *tickCounter) Count() int {
	return tc.count
}

func (tc *tickCounter) Update() {
	tc.count += 1
}

func (tc *tickCounter) UpdateWith(c int) {
	tc.count += c
}
