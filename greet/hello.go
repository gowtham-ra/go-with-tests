package greet

const (
	tamil   = "Tamil"
	french  = "French"
	spanish = "Spanish"

	helloPrefixEng     = "Hello, "
	helloPrefixSpanish = "Hola, "
	helloPrefixFrench  = "Bonjour, "
	helloPrefixTamil   = "வணக்கம், "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloPrefixSpanish
	case tamil:
		prefix = helloPrefixTamil
	case french:
		prefix = helloPrefixFrench
	default:
		prefix = helloPrefixEng
	}
	return
}
