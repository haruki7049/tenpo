package playing

type tickCounter struct {
	count tick
}

func newTickCounter() *tickCounter {
	tc := &tickCounter{}
	tc.count = tick{}

	return tc
}

func (tc *tickCounter) Count() int {
	return tc.count.inner
}

func (tc *tickCounter) Update() {
	tc.count.Update()
}

func (tc *tickCounter) UpdateWith(c int) {
	tc.count.UpdateWith(c)
}

type tick struct {
	inner int
}

func (t *tick) Update() {
	t.inner += 1
}

func (t *tick) UpdateWith(c int) {
	t.inner += c
}
