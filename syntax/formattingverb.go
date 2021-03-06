package syntax

import "fmt"

// FormattingVerb should print formatting verbs (형식 동사)
func FormattingVerb() {
	fmt.Println("===== 형식 동사 기본 =====")
	fmt.Printf("[%#v] A float: %f\n", "%f", 3.141592)
	fmt.Printf("[%#v] An integer: %d\n", "%d", 15)
	fmt.Printf("[%#v] A string: %s\n", "%s", "hello")
	fmt.Printf("[%#v] A Boolean: %t\n", "%t", false)
	fmt.Printf("[%#v] Values: %v %v %v\n", "%v", 1.2, "\t", true)
	fmt.Printf("[%#v] Values: %#v %#v %#v\n", "%#v", 1.2, "\t", true)
	fmt.Printf("[%#v] Types: %T %T %T\n", "%T", 1.2, "\t", true)
	fmt.Printf("[%#v] Percent sign: %%\n", "%%")

	fmt.Println("===== 최소 너비 지정 =====")
	fmt.Printf("%12s | %s\n", "Product", "Cost in cents")
	fmt.Println("--------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)

	fmt.Println("===== 소수점 정밀도 지정 =====")
	value := 12.3456789
	fmt.Printf("%%7.3f: %7.3f\n", value)
	fmt.Printf("%%7.2f: %7.2f\n", value)
	fmt.Printf("%%7.1f: %7.1f\n", value)
	fmt.Printf("%%.1f: %.1f\n", value)
	fmt.Printf("%%.2f: %.2f\n", value)

}
