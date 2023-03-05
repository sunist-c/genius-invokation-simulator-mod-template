package TemplateCharacter

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	def "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	impl "github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"

	. "genius-invokation-simulator-mod-template/mod"
)

type TemplateCharacterNormalAttackType def.AttackSkill

var (
	TemplateCharacterNormalAttackEntity TemplateCharacterNormalAttackType = impl.NewAttackSkillWithOpts(
		impl.WithAttackSkillID(20001),
		impl.WithAttackSkillType(enum.SkillNormalAttack),
		impl.WithAttackSkillCost(map[enum.ElementType]uint{
			// todo: modify me
		}),
		impl.WithAttackSkillActiveDamageHandler(func(ctx def.Context) (elementType enum.ElementType, damageAmount uint) {
			// todo: implement me
			panic("not implemented yet")
		}),
		impl.WithAttackSkillBackgroundDamageHandler(func(ctx def.Context) (damageAmount uint) {
			// todo: implement me
			panic("not implemented yet")
		}),
	)
)

func GetTemplateCharacterNormalAttackEntity() def.AttackSkill {
	return TemplateCharacterNormalAttackEntity
}

func init() {
	GetMod().RegisterSkill(GetTemplateCharacterNormalAttackEntity())
}
