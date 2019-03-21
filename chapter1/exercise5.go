package main

// One Away: There are three types of edits that can be performed on strings: insert a character,
// remove a character, or replace a character. Given two strings, write a function to check if they are
// one edit (or zero edits) away.

func isOnlyOneReplace(original, edit string) bool {
	for i := 0; i < len(original); i++ {
		if original[i] != edit[i] {
			// Replacement at the end.
			if i == len(original)-1 {
				return true
			}
			return original[i+1:] == edit[i+1:]
		}
	}
	return false
}

func isOnlyOneRemoval(original, edit string) bool {
	for i := 0; i < len(edit); i++ {
		if original[i] != edit[i] {
			return original[i+1:] == edit[i:]
		}
	}
	return true
}

func IsOnlyOneEdit(original, edit string) bool {
	// Or maybe, if same size: Attempt replace.
	// if size different by two
	diff := len(original) - len(edit)
	switch diff {
	case 0:
		return isOnlyOneReplace(original, edit)
	case 1:
		// attempt removal.
		return isOnlyOneRemoval(original, edit)
	case -1:
		// attempt insertion.
		return isOnlyOneRemoval(edit, original)
	default:
		return false
	}
}
