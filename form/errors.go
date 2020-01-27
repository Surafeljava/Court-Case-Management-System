package form

// ValidationErrors represents input validation errors
type ValidationErrors map[string][]string

// Add method to add error messages for a given field to the map
func (ve ValidationErrors) Add(field, message string) {
	ve[field] = append(ve[field], message)
}

// Get method to retrieve the first error message for a given// field from the map.
func (ve ValidationErrors) Get(field string) string {
	ves := ve[field]
	if len(ves) == 0 {
		return ""
	}
	return ves[0]
}
