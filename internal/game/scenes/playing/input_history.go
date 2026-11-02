package playing

type inputHistory struct {
	inner []tick
}

func (ih *inputHistory) Append(i int) {
	ih.inner = append(ih.inner, tick{inner: i})
}
