package zahlen

import (
	"slices"
)

// Digits erwartet eine Zahl und liefert eine Liste mit deren Ziffern.
func Digits(n int) []int {
	result := []int{}

	for n != 0 {
		result = append(result, n%10)
		n /= 10
	}

	slices.Reverse(result)

	return result
}

// DigitStrings erwartet eine Liste von Ziffern und liefert eine Liste mit deren Zeichenketten.
//
// Annahmen:
// * Die Zehnerstelle ist nicht 1. Dies wäre ein Sonderfall, den wir hier nicht behandeln.
// * Die Zahl hat höchstens 3 Stellen.
// * Sonderfälle und längere Zahlen müssen vorab abgefangen werden.
func DigitStrings(digits []int) []string {
	result := []string{}

	// Idee: Wir verwenden zwei Slices mit den Ziffern-Strings für die Einer-Hunderter- bzw. die Zehner-Stellen.
	//       Um die Ziffern-Strings zu finden, verwenden wir die Ziffern als Index in diese Slices.
	//       Dazu müssen wir auch die Stellen belegen, die wir gar nicht brauchen, also die Stellen 0 und 1 bei den Zehnern.
	dstrings_1_100 := []string{"null", "ein", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun"}
	dstrings_1_10 := []string{"", "", "zwanz", "dreiß", "vierz", "fünfz", "sechz", "siebenz", "achtz", "neunz"}

	// Position der Einer- Zehner- und Hunderterstelle bestimmen.
	// (Für den Fall, dass die Zahl weniger als 3 Stellen hat.)
	pos_1 := len(digits) - 1
	pos_10 := len(digits) - 2
	pos_100 := len(digits) - 3

	if pos_100 >= 0 { // Hunderterstelle vorhanden
		digit_100 := digits[pos_100]
		result = append(result, dstrings_1_100[digit_100])
	}

	if pos_10 >= 0 { // Zehnerstelle vorhanden
		digit_10 := digits[pos_10]
		result = append(result, dstrings_1_10[digit_10])
	}

	if pos_1 >= 0 { // Einerstelle vorhanden
		digit_1 := digits[pos_1]
		result = append(result, dstrings_1_100[digit_1])
	}

	return result
}

// Compose erwartet eine Liste von Ziffer-Strings und liefert einen String mit der zusammengesetzten Zahl.
//
// Annahmen:
// * Die Zehnerstelle ist nicht 1. Dies wäre ein Sonderfall, den wir hier nicht behandeln.
// * Die Zahl hat höchstens 3 Stellen.
// * Sonderfälle und längere Zahlen müssen vorab abgefangen werden.
func Compose(digits []string) string {
	result := ""

	// Position der Einer- Zehner- und Hunderterstelle bestimmen.
	// (Für den Fall, dass die Zahl weniger als 3 Stellen hat.)
	pos_1 := len(digits) - 1
	pos_10 := len(digits) - 2
	pos_100 := len(digits) - 3

	// Wenn die Hunderterstelle vorhanden ist, dann muss sie verwendet und mit "hundert" ergänzt werden.
	if len(digits) == 3 {
		result += digits[pos_100] + "hundert"
	}

	// Die Einerstelle ist immer vorhanden und kommt an zweiter Stelle.
	result += digits[pos_1]

	// Wenn die Zehnerstelle vorhanden ist, dann muss sie verwendet sowie mit "und...ig" ergänzt werden.
	if len(digits) >= 2 {
		result += "und" + digits[pos_10] + "ig"
	}

	return result
}

// NumberString erwartet eine Zahl und liefert einen String mit der Zahl in Worten.
// Annahme: Die Zahl hat höchstens 3 Stellen.
func NumberString(n int) string {

	// Sonderfälle abfangen:
	special_cases := []int{0, 1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	special_results := []string{"null", "eins", "zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechzehn", "siebzehn", "achtzehn", "neunzehn"}

	for i, special_case := range special_cases {
		if n == special_case {
			return special_results[i]
		}
	}

	digits := Digits(n)
	dstrings := DigitStrings(digits)
	result := Compose(dstrings)

	return result
}
