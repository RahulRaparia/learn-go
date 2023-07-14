package main 

func main () {
	input := "The quick brown fox jumps over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("double reversed: %q\n", doubleRev)
}


func Reverse (s string) string {
	b := []bytes(s)
	for i, j := 0, len(b)-1; i< len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

