package main

func Bool() {
	// !(a && b) == ? !a || !b

	a, b := true, true
	println("!(a && b) = ", !(a && b), " !a || !b = ", !a || !b)

	a = true
	b = false
	println("!(a && b) = ", !(a && b), " !a || !b = ", !a || !b)

	a = false
	b = true
	println("!(a && b) = ", !(a && b), " !a || !b = ", !a || !b)

	a = false
	b = false
	println("!(a && b) = ", !(a && b), " !a || !b = ", !a || !b)
}
