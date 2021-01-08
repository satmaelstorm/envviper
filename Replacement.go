package envviper

//Replacement represents strings for strings.Replacer in human readable format
type Replacement struct {
	InVar string
	InEnv string
}

//NewReplacement builds Replacement instance from two strings
func NewReplacement(inVar, inEnv string) Replacement {
	return Replacement{
		InVar: inVar,
		InEnv: inEnv,
	}
}
