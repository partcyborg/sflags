package viper

import (
	"log"
	"strings"

	"github.com/octago/sflags"
	"github.com/spf13/viper"
)

func GenerateTo(src []*sflags.Flag, prefix string, dst interface{}) error {
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	for _, sFlag := range src {
		log.Printf("binding env %s", sFlag.ViperName)
		v.BindEnv(sFlag.ViperName)
	}
	return v.Unmarshal(dst)
}

func ParseTo(cfg interface{}, prefix string, optFuncs ...sflags.OptFunc) error {
	flags, err := sflags.ParseStruct(cfg, optFuncs...)
	if err != nil {
		return err
	}
	return GenerateTo(flags, prefix, cfg)
}
