// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	nameCheck := r.Name != ""
	street, prs := r.Address["street"]
	addressCheck := prs && street != ""
	return nameCheck && addressCheck
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Age = 0
	r.Name = ""
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var result int
	for _, r := range residents {
		if r.HasRequiredInfo() {
			result++
		}
	}
	return result
}
