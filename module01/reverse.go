package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {

	//version1
	//reverse_word := ""
	//for i := len(word) - 1; i >= 0; i-- {
	//	reverse_word += string(word[i])   // cast as string because word[i] returns a byte
	//}
	//return reverse_word

	//version2
	//reverse_word := ""
	//for i := 0; i < len(word); i++ {
	//	reverse_word = string(word[i]) + reverse_word
	//}
	//return reverse_word

	//version3 - allow chars with more than one byte(rune) ex: "語本日"
	reverse_word := ""
	for _, r := range word {
		reverse_word = string(r) + reverse_word
	}
	return reverse_word
}
