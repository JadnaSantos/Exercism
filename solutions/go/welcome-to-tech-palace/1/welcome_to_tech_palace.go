package techpalace

import (
    "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    message := "Welcome to the Tech Palace, "
    nameUpper := strings.ToUpper(customer)

    return message + nameUpper
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    // numStarsPerline qnt de estrela int \n
    // mensagem
   return strings.Repeat("*", numStarsPerLine) + "\n"+welcomeMsg+"\n" + strings.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	//strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")
    return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
