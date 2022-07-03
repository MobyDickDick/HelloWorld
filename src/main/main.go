package main

import (
	"fmt"
	"strings"
)

func Hello(name string, lang string) string {

	return evaluateGreeting(name, lang, func(name string, prefix string,
		suffix string) string {
		if name != "" {
			suffix = name
		}
		return prefix + suffix
	})
}

func evaluateGreeting(name string, lang string, f func(name string,
	prefix string, suffix string) string) string {

	prefix := ""
	switch strings.ToLower(lang) {
	case "spanish":
		prefix = f(name, "Hola, ", "Tierra")
		break
	case "french":
		prefix = f(name, "Bonjour, ", "monde")
	default:
		prefix = f(name, "Hello, ", "World!")
	}
	return prefix
}

func main() {
	fmt.Println(Hello("Max", ""))
}

/*
Olivias
package mainimport (
    "fmt"    "strings")
func Hello(name string, language string) string {
    prefix, postfix  := getPrefixPostfix(language)
    if name == "" {
        return prefix + postfix    } else {
        return prefix + name    }
}
func getPrefixPostfix(language string) (string,string){
    prefix := "Hello, "    postfix := "World!"    switch strings.ToLower(language) {
    case "spanish":
        prefix = "Hola, "        postfix = "el mundo"    case "french":
        prefix = "Bonjour, "        postfix = "le monde"    }
    return prefix, postfix}
func main() {
    fmt.Println(Hello("Olivia", ""))
}
*/
