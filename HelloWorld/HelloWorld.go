package main

import "fmt"

const (
	// languages available
	french  = "French"
	spanish = "Spanish"

	// Initial prefix for the result
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", "")) // Specify the language explicitly
}
