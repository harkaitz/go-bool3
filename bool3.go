package bool3



type Bool string

func (b Bool) String() string {
	switch b {
	case "", "indeterminate": return "indeterminate"
	case "off", "false", "0": return ""
	default:                  return "checked"
	}
}

func (b *Bool) Set(value bool) {
	switch value {
	case true:  *b = "on"
	case false: *b = "off"
	}
}

func (b Bool) IsSet() bool {
	return b != ""
}

func (b Bool) Value() bool {
	switch b {
	case "", "indeterminate": return false
	case "off", "false", "0": return false
	default:                  return true
	}
}
