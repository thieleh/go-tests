package address

import "strings"

// TypeOfAddress verifies if an address has a valid type
func TypeOfAddress(address string) bool {

	validTypes := []string{"street", "avenue", "rod"}
	lowerAddress := strings.ToLower(address)

	firstWord := strings.Split(lowerAddress, " ")[0]

	validAddress := false
	for _, typeOfAddress := range validTypes {
		if typeOfAddress == firstWord {
			validAddress = true
		}
	}

	if validAddress {
		return true
	}

	return false
}
