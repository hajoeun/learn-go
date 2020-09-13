package syntax

import "fmt"

// FormattingVerb should print formatting verbs (형식 동사)
func FormattingVerb() {
	fmt.Printf("[%#v] A float: %f\n", "%f", 3.141592)
	fmt.Printf("[%#v] An integer: %d\n", "%d", 15)
	fmt.Printf("[%#v] A string: %s\n", "%s", "hello")
	fmt.Printf("[%#v] A Boolean: %t\n", "%t", false)
	fmt.Printf("[%#v] Values: %v %v %v\n", "%v", 1.2, "\t", true)
	fmt.Printf("[%#v] Values: %#v %#v %#v\n", "%#v", 1.2, "\t", true)
	fmt.Printf("[%#v] Types: %T %T %T\n", "%T", 1.2, "\t", true)
	fmt.Printf("[%#v] Percent sign: %%\n", "%%")
}
