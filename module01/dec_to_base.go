package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"
func DecToBase(dec, base int) string {
	converted_base := ""
	// version1
	//for dec != 0 {
	//	rem := dec % base
	//	if base >= 11 && rem >= 10 {
	//		remAlpha := ""
	//		switch rem {
	//		case 10:
	//			remAlpha = "A"
	//		case 11:
	//			remAlpha = "B"
	//		case 12:
	//			remAlpha = "C"
	//		case 13:
	//			remAlpha = "D"
	//		case 14:
	//			remAlpha = "E"
	//		case 15:
	//			remAlpha = "F"
	//		}
	//		converted_base = remAlpha + converted_base
	//	} else {
	//		converted_base = strconv.Itoa(rem) + converted_base
	//	}
	//	dec = dec / base
	//}

	// version2
	//for dec != 0 {
	//	rem := dec % base
	//	converted_base = fmt.Sprintf("%X%s", rem, converted_base)
	//	dec = dec / base
	//}

	// version3
	charset := "0123456789ABCDEF"
	for dec != 0 {
		rem := dec % base
		converted_base = string(charset[rem]) + converted_base
		dec = dec / base
	}

	return converted_base
}
