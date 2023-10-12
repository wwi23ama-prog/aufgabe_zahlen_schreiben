package zahlen

// Digits erwartet eine Zahl und liefert eine Liste mit deren Ziffern.
func Digits(n int) []int {
	result := []int{}

	// Die letzte Ziffer einer Zahl erhält man, wenn man die Zahl modulo 10 nimmt.
	// Die letzte Ziffer entfernt man, indem man die Zahl durch 10 teilt.
	// Schreiben Sie eine Schleife, die solange läuft, bis die Zahl 0 ist.
	// In jedem Schleifendurchlauf wird die letzte Ziffer der Zahl ermittelt und der Liste result angehängt.
	// Dann wird die Zahl durch 10 geteilt, um die letzte Ziffer zu entfernen.
	//
	// Anschließend muss die Liste noch umgedreht werden, damit die Ziffern in der richtigen Reihenfolge vorliegen.
	// Dazu können Sie die Funktion slices.Reverse verwenden.

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
	// Sie können die folgenden Slices verwenden. Kommentieren Sie sie ein, falls Sie sie verwenden:
	//dstrings_1_100 := []string{"null", "ein", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun"}
	//dstrings_1_10 := []string{"", "", "zwanz", "dreiß", "vierz", "fünfz", "sechz", "siebenz", "achtz", "neunz"}

	// Position der Einer- Zehner- und Hunderterstelle bestimmen.
	// (Für den Fall, dass die Zahl weniger als 3 Stellen hat.)
	//pos_1 := len(digits) - 1
	//pos_10 := len(digits) - 2
	//pos_100 := len(digits) - 3

	// Verwenden Sie nun mehrere if-Anweisungen, um die Ziffern-Strings in die Liste result einzufügen.
	// Prüfen Sie dabei jeweils zuerst, ob die entsprechende Stelle vorhanden ist.

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
	// Hinweis: Siehe auch die Funktion DigitStrings.

	// Verwenden Sie nun wieder if-Anweisungen, um die folgenden Bedingungen zu prüfen und den String result entsprechend zu ergänzen:
	// Wenn die Hunderterstelle vorhanden ist, dann muss sie verwendet und mit "hundert" ergänzt werden.
	// Die Einerstelle ist immer vorhanden und kommt an zweiter Stelle.
	// Wenn die Zehnerstelle vorhanden ist, dann muss sie verwendet sowie mit "und...ig" ergänzt werden.

	return result
}

// NumberString erwartet eine Zahl und liefert einen String mit der Zahl in Worten.
// Annahme: Die Zahl hat höchstens 3 Stellen.
func NumberString(n int) string {

	// 1. Sonderfälle abfangen:
	special_cases := []int{0, 1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	special_results := []string{"null", "eins", "zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechzehn", "siebzehn", "achtzehn", "neunzehn"}

	// Prüfen Sie mit einer Schleife, ob n einer der Sonderfälle ist und liefern Sie in diesem Fall das entsprechende Ergebnis zurück.
	for i, special_case := range special_cases {
		if n == special_case {
			return special_results[i]
		}
	}

	// Ab hier ist n keiner der Sonderfälle und kann mit den Funktionen Digits, DigitStrings und Compose behandelt werden.
	digits := Digits(n)
	dstrings := DigitStrings(digits)
	result := Compose(dstrings)

	return result
}
