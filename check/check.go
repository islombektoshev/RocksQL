package check

func That(ok bool) {
	if !ok {
		panic("check failed")
	}
}
