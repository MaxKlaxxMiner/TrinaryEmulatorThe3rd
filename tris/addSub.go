package tris

func (u Uint6) Add(add, carry Uint6) (sum, carryOut Uint6) {
	sum = u + add + carry
	if sum >= Uint6End {
		carryOut = 1
		sum -= Uint6End
	}
	return
}

func (u Uint18) Add(add, carry Uint18) (sum, carryOut Uint18) {
	sum = u + add + carry
	if sum >= Uint18End {
		carryOut = 1
		sum -= Uint18End
	}
	return
}

func (u Uint36) Add(add, carry Uint36) (sum, carryOut Uint36) {
	sum = u + add + carry
	if sum >= Uint36End {
		carryOut = 1
		sum -= Uint36End
	}
	return
}

func (u Uint6) Sub(sub, borrow Uint6) (diff, borrowOut Uint6) {
	diff = u - sub - borrow
	if diff >= Uint6End {
		borrowOut = 1
		diff += Uint6End
	}
	return
}

func (u Uint18) Sub(sub, borrow Uint18) (diff, borrowOut Uint18) {
	diff = u - sub - borrow
	if diff >= Uint18End {
		borrowOut = 1
		diff += Uint18End
	}
	return
}

func (u Uint36) Sub(sub, borrow Uint36) (diff, borrowOut Uint36) {
	diff = u - sub - borrow
	if diff >= Uint36End {
		borrowOut = 1
		diff += Uint36End
	}
	return
}
