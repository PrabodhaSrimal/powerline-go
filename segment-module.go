package main

import (
	pwl "github.com/justjanne/powerline-go/powerline"
	"os"
	"path"
)

func segmentModule(p *powerline) []pwl.Segment {
	var env string
	if env == "" {
		env, _ = os.LookupEnv("LOADEDMODULES")
	}
	if env != "" {
		envName := path.Base(env)
		return []pwl.Segment{{
			Name: "module",
			Content:    envName,
			Foreground: p.theme.ModuleFg,
			Background: p.theme.ModuleBg,
		}}
	}

	return []pwl.Segment{}
}
