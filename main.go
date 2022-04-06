package darkmode

import (
	"fmt"

	jxa "github.com/wobsoriano/go-jxa"
)

const property = "Application('System Events').appearancePreferences.darkMode"

func IsEnabled() bool {
	v, _ := jxa.RunJXA(fmt.Sprintf("%s()", property))
	return v == "true"
}

func Enable() bool {
	v, _ := jxa.RunJXA(fmt.Sprintf("%s = true", property))
	return v == "true"
}

func Disable() bool {
	v, _ := jxa.RunJXA(fmt.Sprintf("%s = false", property))
	return v == "true"
}

func Toggle() bool {
	v, _ := jxa.RunJXA(fmt.Sprintf("%s = !%s()", property, property))
	return v == "true"
}