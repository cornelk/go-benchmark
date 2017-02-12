package benchmark

func generateIntSlice() [128]int {
	var m [128]int
	for i := 0; i < 128; i++ {
		m[i] = i
	}
	return m
}
