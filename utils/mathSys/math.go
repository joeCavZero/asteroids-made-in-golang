package mathSys

func Lerp(start, end, t float32) float32 {
	return start + t*(end-start)
}
