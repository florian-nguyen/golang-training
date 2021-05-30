package flatten

// Flatten transforms a nested list with nil items into a single flattened list
func Flatten(list interface{}) []interface{} {
	switch list.(type) {
	case []interface{}:
		output := []interface{}{}
		// For each item of the interface array
		for _, item := range list.([]interface{}) {
			// Recursively call the Flatten function to deal with simple lists only
			for _, subitem := range Flatten(item) {
				output = append(output, subitem)
			}
		}
		return output

	// ignore nil type
	case nil:
		return nil
	}
	// Exit case: current argument is a simple list
	return []interface{}{list}

}
