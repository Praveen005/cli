package utility

import (
	"fmt"
	"os"

	"github.com/civo/cli/common"
)

const (
	_Green  = "\033[32m"
	_Yellow = "\033[33m"
	_Blue   = "\033[34m"
	_Magenta = "\033[35m"
	_Red    = "\033[31m"
	_Orange = "\033[38;5;214m"
	_Reset  = "\033[0m"
)

func Colorize(color, value string) string {
	return fmt.Sprintf("%s%s%s", color, value, _Reset)
}

// Green is the function to convert str to green in console
func Green(value string) string {
	return Colorize(_Green, value)
}

// Yellow is the function to convert str to yellow in console
func Yellow(value string) string {
	return Colorize(_Yellow, value)
}

// Orange is the function to convert str to orange in console
func Orange(value string) string{
	return Colorize(_Orange, value)
}

// Blue is the function to convert str to blue in console
func Blue(value string) string {
	return Colorize(_Blue, value)
}

// Magenta is the function to convert str to magenta in console
func Magenta(value string) string {
	return Colorize(_Magenta, value)
}

// Red is the function to convert str to red in console
func Red(value string) string {
	return Colorize(_Red, value)
}

// Error is the function to handler all error in the Cli
func Error(msg string, args ...interface{}) {
	common.IssueMessage()
	fmt.Fprintf(os.Stderr, "%s: %s\n", Red("Error"), fmt.Sprintf(msg, args...))
}

// Info is the function to handler all info messages in the Cli
func Info(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", Blue("Info"), fmt.Sprintf(msg, args...))
}

// Warning is the function to handler all warnings in the Cli
func Warning(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", Yellow("Warning"), fmt.Sprintf(msg, args...))
}

// YellowConfirm is the function to handler all delete confirm
func YellowConfirm(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s", Yellow("Warning"), fmt.Sprintf(msg, args...))
}

// RedConfirm is the function to handler the new version of the Cli
func RedConfirm(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s", Red("IMPORTANT"), fmt.Sprintf(msg, args...))
}

// ColorStatus is to print the status of the Instance or k8s Cluster
func ColorStatus(status string) string {

	var returnText string

	switch {
	case status == "ACTIVE":
		returnText = Green(status)
	case status == "SHUTOFF":
		returnText = Red(status)
	case status == "REBOOTING":
		returnText = Yellow(status)
	case status == "BUILDING":
		returnText = Yellow(status)
	case status == "INSTANCE-CREATE":
		returnText = Blue(status)
	case status == "INSTALLING":
		returnText = Blue(status)
	case status == "SCALING":
		returnText = Magenta(status)
	case status == "STOPPING":
		returnText = Yellow(status)
	default:
		returnText = Red("Unknown")
	}

	return returnText
}