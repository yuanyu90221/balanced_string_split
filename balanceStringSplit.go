package balance_string

func balanceStringSplit(s string) int {
	count := 0
	amount := 0
	for _, r := range s {
		if r == 'R' {
			count += 1
		} else {
			count -= 1
		}
		if count == 0 {
			amount += 1
		}
	}
	return amount
}
