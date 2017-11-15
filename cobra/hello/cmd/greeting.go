package cmd

func greeting(lang string) string {
	switch lang {
	case "pl":
		return "Cześć"
	case "es":
		return "Hola"
	case "cn":
		return "你好"
	default: // "en"
		return "Hello"
	}
}
