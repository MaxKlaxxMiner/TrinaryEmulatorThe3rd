package tris

func (u Uint9) Add(add, carry Uint9) (sum, carryOut Uint9) {
	sum = u + add + carry
	if sum >= Uint9End {
		carryOut = 1
		sum -= Uint9End
	}
	return
}

func (u Uint27) Add(add, carry Uint27) (sum, carryOut Uint27) {
	sum = u + add + carry
	if sum >= Uint27End {
		carryOut = 1
		sum -= Uint27End
	}
	return
}

func (u Uint9) Sub(sub, borrow Uint9) (diff, borrowOut Uint9) {
	diff = u - sub - borrow
	if diff >= Uint9End {
		borrowOut = 1
		diff += Uint9End
	}
	return
}

func (u Uint27) Sub(sub, borrow Uint27) (diff, borrowOut Uint27) {
	diff = u - sub - borrow
	if diff >= Uint27End {
		borrowOut = 1
		diff += Uint27End
	}
	return
}

func (s Int9) Add(add, carry Int9) (sum, carryOut Int9) {
	su, car := Uint9(s).Add(Uint9(add), Uint9(carry))
	return Int9(su), Int9(car)
}

func (s Int27) Add(add, carry Int27) (sum, carryOut Int27) {
	su, car := Uint27(s).Add(Uint27(add), Uint27(carry))
	return Int27(su), Int27(car)
}

func (s Int9) Sub(sub, borrow Int9) (diff, borrowOut Int9) {
	d, b := Uint9(s).Sub(Uint9(sub), Uint9(borrow))
	return Int9(d), Int9(b)
}

func (s Int27) Sub(sub, borrow Int27) (diff, borrowOut Int27) {
	d, b := Uint27(s).Sub(Uint27(sub), Uint27(borrow))
	return Int27(d), Int27(b)
}
