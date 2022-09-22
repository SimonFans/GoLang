package Method

type Thermostat struct {
	ID    string
	Value float64
}

func (t Thermostat) GetValue() float64 {
	return t.Value
}

// when assign value, need to use pointer variable
func (t *Thermostat) SetValue(val float64) {
	t.Value = val
}
