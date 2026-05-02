package main

import (
	"fmt"
	"os"
	"strings"
)

type local struct {
	country  string
	language string
}

func getLocales() []local {
	locales := make([]local, 0, 10)
	locales = append(locales, local{country: "CN", language: "en"})
	locales = append(locales, local{country: "CN", language: "fr"})
	locales = append(locales, local{country: "US", language: "en"})
	locales = append(locales, local{country: "UK", language: "en"})
	locales = append(locales, local{country: "GER", language: "de"})
	locales = append(locales, local{country: "FR", language: "fr"})
	locales = append(locales, local{country: "Ru", language: "ru"})

	return locales
}

func isLocalSupported(locales []local, language string, country string) bool {
	supported := false

	for _, value := range locales {
		if value.language == language && value.country == country {
			supported = true
			break
		}
	}

	return supported

}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide local informat of language_country")
		os.Exit(1)
	}

	s := os.Args[1]
	parts := strings.Split(s, "_")

	if len(parts) != 2 {
		fmt.Println("Invalid format. Use language_country e.g. en_US")
		os.Exit(1)
	}
	language := parts[0]
	country := parts[1]

	locales := getLocales()

	supported := isLocalSupported(locales, language, country)

	if supported {
		fmt.Println("Locale passed is supported")
	} else {
		fmt.Printf("Locale not supported: %v\n", s)
	}

}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type locale struct {
// 	language string
// 	territory string
// }

// func getLocales() map[locale]struct{} {
// 	supportedLocales := make(map[locale]struct{}, 5)
// 	supportedLocales[locale{"en", "US"}] = struct{}{}
// 	supportedLocales[locale{"en", "CN"}] = struct{}{}
// 	supportedLocales[locale{"fr", "CN"}] = struct{}{}
// 	supportedLocales[locale{"fr", "FR"}] = struct{}{}
// 	supportedLocales[locale{"ru", "RU"}] = struct{}{}
// 	return supportedLocales
// }

// func localeExists(l locale) bool {
// 	_, exists := getLocales()[l]
// 	return exists
// }

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("No locale passed")
// 		os.Exit(1)
// 	}

// 	localeParts := strings.Split(os.Args[1], "_")
// 	if len(localeParts) != 2 {
// 		fmt.Printf("Invalid locale passed: %v\n", os.Args[1])
// 		os.Exit(1)
// 	}

// 	passedLocale := locale{
// 		territory: localeParts[1],
// 		language:  localeParts[0],
// 	}

// 	if !localeExists(passedLocale) {
// 		fmt.Printf("Locale not supported: %v\n", os.Args[1])
// 		os.Exit(1)
// 	}
// 	fmt.Println("Locale passed is supported")
// }
