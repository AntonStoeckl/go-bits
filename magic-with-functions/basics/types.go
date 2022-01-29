package basics

// IsBigger is a plain function that compares two integers
// and return true if the first is bigger than the second
func IsBigger(first, second int) bool {
	return first > second
}

// Compare is a higher order function that compares two integers via a function
// that is passed in as the third parameter using explicit annotation,
// so the full function signature is annotating the type of the parameter
func Compare(first, second int, compareFn func(first, second int) bool) bool {
	return compareFn(first, second)
}

// CompareFn is a function type
type CompareFn func(first, second int) bool

// OtherCompare is a higher order function that compares two integers via a
// function that is passed in as the third parameter using a function type
// annotation, so the function type is used to annotate the type of the
// parameter
func OtherCompare(first, second int, compareFn CompareFn) bool {
	return compareFn(first, second)
}
