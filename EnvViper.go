package envviper

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

//EnvViper - wrapper for Viper with automatic bind environment variables
type EnvViper struct {
	*viper.Viper
}

//NewEnvViper creates EnvViper instance
func NewEnvViper() *EnvViper {
	ev := new(EnvViper)
	ev.Viper = viper.New()
	return ev
}

//DefEnvViper creates EnvViper instance with global Viper
// You can use:
// import "github.com/satmaelstorm/envviper"
// viper := envviper.DefEnvViper()
// for fast migrate to EnvViper
func DefEnvViper() *EnvViper {
	ev := new(EnvViper)
	ev.Viper = viper.GetViper()
	return ev
}

//SetEnvPrefix - must not be use, panics if called
func (ev *EnvViper) SetEnvPrefix(prefix string) {
	panic("use SetEnvParams instead of SetEnvPrefix")
}

//SetEnvKeyReplacer - must not be use, panics if called
func (ev *EnvViper) SetEnvKeyReplacer(replacer *strings.Replacer) {
	panic("use SetEnvParams instead of SetEnvKeyReplacer")
}

//AutomaticEnv - must not be use, panics if called
func (ev *EnvViper) AutomaticEnv() {
	panic("use SetEnvParams instead of AutomaticEnv")
}

//SetEnvParamsSimple - call SetEnvParams with replace "." in config
//variables with "_" in environment variables
func (ev *EnvViper) SetEnvParamsSimple(prefix string) {
	ev.SetEnvParams(prefix, Replacement{
		InVar: ".",
		InEnv: "_",
	})
}

//SetEnvParams do:
// 1. Call viper.SetEnvPrefix(prefix)
// 2. Call viper.AutomaticEnv()
// 3. Call viper.SetEnvKeyReplacer() with replacements
// 4. Iterate through environment variables and bind variables with "prefix"
func (ev *EnvViper) SetEnvParams(prefix string, replacement ...Replacement) {
	ev.Viper.SetEnvPrefix(prefix)
	ev.Viper.AutomaticEnv()
	prefix = strings.ToUpper(prefix)
	var registrationReplacer *strings.Replacer
	if len(replacement) > 0 {
		viperReplacerStrings := make([]string, len(replacement)*2)
		registrationReplacerStrings := make([]string, len(replacement)*2)
		for i := 0; i < len(replacement); i++ {
			viperReplacerStrings[2*i] = replacement[i].InVar
			viperReplacerStrings[2*i+1] = replacement[i].InEnv
			registrationReplacerStrings[2*i] = replacement[i].InEnv
			registrationReplacerStrings[2*i+1] = replacement[i].InVar
		}
		replacer := strings.NewReplacer(viperReplacerStrings...)
		ev.Viper.SetEnvKeyReplacer(replacer)
		registrationReplacer = strings.NewReplacer(registrationReplacerStrings...)
	}
	prefix += "_"
	for _, pair := range os.Environ() {
		variable := strings.Split(pair, "=")
		if len(variable) < 1 {
			continue
		}
		if strings.HasPrefix(variable[0], prefix) {
			v := variable[0][len(prefix):]
			if registrationReplacer != nil {
				v = registrationReplacer.Replace(v)
			}
			_ = ev.BindEnv(v)
		}
	}
}
