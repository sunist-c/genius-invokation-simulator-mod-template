package mod

import (
	def "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	impl "github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
)

type TemplateModType def.Mod

var (
	TemplateModEntity TemplateModType = impl.NewMod()
)

func GetMod() def.Mod {
	return TemplateModEntity
}
