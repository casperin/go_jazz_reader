package str

func Or(left, right string) string {
	if left == "" {
		return right
	}
	return left
}
