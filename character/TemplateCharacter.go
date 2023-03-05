package character

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	def "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	impl "github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"

	. "genius-invokation-simulator-mod-template/mod"
	skill "genius-invokation-simulator-mod-template/skill/TemplateCharacter"
)

type TemplateCharacterType def.Character

var (
	TemplateCharacterEntity TemplateCharacterType = impl.NewCharacterWithOpts(
		impl.WithCharacterID(10001),
		impl.WithCharacterName("TemplateCharacter"),
		impl.WithCharacterAffiliation(enum.AffiliationUndefined),
		impl.WithCharacterVision(enum.ElementNone),
		impl.WithCharacterWeapon(enum.WeaponOthers),
		impl.WithCharacterHP(10),
		impl.WithCharacterMP(3),
		impl.WithCharacterSkills(
			skill.GetTemplateCharacterNormalAttackEntity(),
		),
	)
)

func GetTemplateCharacterEntity() def.Character {
	return TemplateCharacterEntity
}

func init() {
	GetMod().RegisterCharacter(GetTemplateCharacterEntity())
}
