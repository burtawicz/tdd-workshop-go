// package main

func Count(sequence string) int {
	return len(sequence)
}

func Reverse(sequence string) string {
	r := []rune(sequence)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Separate(sequence string) string {
	return strings.Join(strings.Split(sequence, ""), ",")
}

// TODO: implement EncodeRot13(sequence string) string
// TODO: implement DecodeRot13(sequence string) string
// TODO: implement EncodeBase64(sequence string) string
// TODO: implement DecodeBase64(sequence string) string
// TODO: implement EncodeShiftLeft(sequence string, key int) string
// TODO: implement DecodeShiftLeft(sequence string, key int) string
// TODO: implement EncodeShiftRight(sequence string, key int) string
// TODO: implement DecodeShiftRight(sequence string, key int) string
