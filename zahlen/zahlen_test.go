package zahlen

import "fmt"

func ExampleDigits() {
	fmt.Println(Digits(143))
	fmt.Println(Digits(53))
	fmt.Println(Digits(7))
	fmt.Println(Digits(258))
	fmt.Println(Digits(258756))

	// Output:
	// [1 4 3]
	// [5 3]
	// [7]
	// [2 5 8]
	// [2 5 8 7 5 6]
}

func ExampleDigitStrings() {
	fmt.Println(DigitStrings([]int{1, 4, 3}))
	fmt.Println(DigitStrings([]int{5, 3}))
	fmt.Println(DigitStrings([]int{7}))
	fmt.Println(DigitStrings([]int{2, 5, 8}))
	fmt.Println(DigitStrings([]int{2, 5, 8, 7, 5, 6}))

	// Output:
	// ["ein" "vierz" "drei"]
	// ["fünf" "drei"]
	// ["sieben"]
	// ["zwei" "fünfz" "acht"]
	// ["zwei" "fünfz" "acht" "sieben" "fünfz" "sechs"]
}

func ExampleCompose() {
	fmt.Println(Compose([]string{"ein", "vierz", "drei"}))
	fmt.Println(Compose([]string{"fünf", "drei"}))
	fmt.Println(Compose([]string{"sieben"}))
	fmt.Println(Compose([]string{"zwei", "fünfz", "acht"}))
	fmt.Println(Compose([]string{"zwei", "fünfz", "acht", "sieben", "fünfz", "sechs"}))

	// Output:
	// einhundertdreiundvierzig
	// dreiundfünfzig
	// sieben
	// zweihundertachtundfünfzig
	// zweihundertachtundfünfzigtausendsiebenhundertfünfundfünfzig
}
