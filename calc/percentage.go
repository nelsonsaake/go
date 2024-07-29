package calc

func Percentage(part, total int) float64 {
	if total > 0 {
		return (float64(part) / float64(total)) * 100
	} else {
		return 0
	}
}
