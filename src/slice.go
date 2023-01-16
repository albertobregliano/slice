package slice

func Copy(a []any) (b []any) {
	b = append(a[:0:0], a...)
	return b
}

func Delete(a []any, i int) []any {
	return append(a[:i], a[i+1:]...)
}