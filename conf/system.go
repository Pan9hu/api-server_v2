package conf

import "github.com/spf13/viper"

var (
	MeloVP *viper.Viper
)

func GetViper() *viper.Viper {
	return MeloVP
}
