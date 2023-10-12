package zahlen

import "fmt"

func ExampleDigits() {
	fmt.Println(Digits(143))
	fmt.Println(Digits(53))
	fmt.Println(Digits(7))
	fmt.Println(Digits(258))

	// Output:
	// [1 4 3]
	// [5 3]
	// [7]
	// [2 5 8]
}

func ExampleDigitStrings() {
	fmt.Println(DigitStrings([]int{1, 4, 3}))
	fmt.Println(DigitStrings([]int{5, 3}))
	fmt.Println(DigitStrings([]int{7}))
	fmt.Println(DigitStrings([]int{2, 5, 8}))

	// Output:
	// ["ein" "vierz" "drei"]
	// ["fünf" "drei"]
	// ["sieben"]
	// ["zwei" "fünfz" "acht"]
}

func ExampleCompose() {
	fmt.Println(Compose([]string{"ein", "vierz", "drei"}))
	fmt.Println(Compose([]string{"fünf", "drei"}))
	fmt.Println(Compose([]string{"sieben"}))
	fmt.Println(Compose([]string{"zwei", "fünfz", "acht"}))

	// Output:
	// einhundertdreiundvierzig
	// dreiundfünfzig
	// sieben
	// zweihundertachtundfünfzig
}

func ExampleNumberString() {
	fmt.Println(NumberString(143))
	fmt.Println(NumberString(53))
	fmt.Println(NumberString(7))
	fmt.Println(NumberString(258))
	fmt.Println(NumberString(258756))

	// Output:
	// einhundertdreiundvierzig
	// dreiundfünfzig
	// sieben
	// zweihundertachtundfünfzig
	// zweihundertachtundfünfzigtausendsiebenhundertfünfundfünfzig
}
