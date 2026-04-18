package hello

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix  = "Bonjour, "


func main() {
	// fmt.Println(greeting_recipirnt("Khalil"))
	fmt.Println(Hello("Khalil",""))
	fmt.Println(Hello("Khalil","Spanish"))
	fmt.Println(Hello("Khalil","French"))
}

func greeting() string {
	msg := englishHelloPrefix + "World!"
	return msg
}

func greeting_recipirnt(name string) string {
	msg := fmt.Sprintf("%s%s!", englishHelloPrefix, name)
	return msg
}

func Hello(name, language string) string{
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language{
	case "Spanish":
		prefix = spanishHelloPrefix

	case "French":
		prefix = frenchHelloPrefix

	default:
		prefix = englishHelloPrefix
	}
	return
}
