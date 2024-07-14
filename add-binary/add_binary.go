package add_binary

//  Given two binary strings a and b, return their sum as a binary string.

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	res, carry := []byte(a), 0

	for ia, ib := len(a)-1, len(b)-1; ia >= 0; {
		if a[ia] == '1' {
			carry++
		}

		if ib >= 0 && b[ib] == '1' {
			carry++
		}

		switch carry {
		case 1:
			res[ia], carry = '1', 0
		case 2:
			res[ia], carry = '0', 1
		case 3:
			res[ia], carry = '1', 1
		}
		ia--
		ib--
	}

	if carry == 1 {
		temp := res
		res = make([]byte, len(temp)+1)
		copy(res[1:], temp)
		res[0] = '1'
	}

	return string(res)
}
