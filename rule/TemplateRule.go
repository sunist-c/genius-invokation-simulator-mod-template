package rule

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/entity/model"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	def "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	impl "github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"

	. "genius-invokation-simulator-mod-template/mod"
)

type TemplateVictorCalculator struct{}

func (impl TemplateVictorCalculator) CalculateVictors(players []model.Player) (has bool, victors []model.Player) {
	// todo: implement me
	panic("implement me")
}

type TemplateReactionCalculator struct{}

func (impl TemplateReactionCalculator) ReactionCalculate(types []enum.ElementType) (reaction enum.Reaction, elementRemains []enum.ElementType) {
	// todo: implement me
	panic("implement me")
}

func (impl TemplateReactionCalculator) DamageCalculate(reaction enum.Reaction, targetCharacter model.Character, ctx *context.DamageContext) {
	// todo: implement me
	panic("implement me")
}

func (impl TemplateReactionCalculator) EffectCalculate(reaction enum.Reaction, targetPlayer model.Player) (ctx *context.CallbackContext) {
	// todo: implement me
	panic("implement me")
}

func (impl TemplateReactionCalculator) Attach(originalElements []enum.ElementType, newElement enum.ElementType) (resultElements []enum.ElementType) {
	// todo: implement me
	panic("implement me")
}

func (impl TemplateReactionCalculator) Relative(reaction enum.Reaction, relativeElement enum.ElementType) bool {
	// todo: implement me
	panic("implement me")
}

type TemplateRuleType def.Rule

var (
	TemplateRuleEntity TemplateRuleType = impl.NewRuleWithOpts(
		impl.WithRuleID(60001),
		impl.WithRuleImplement(enum.RuleTypeVictorCalculator, TemplateVictorCalculator{}),
		impl.WithRuleImplement(enum.RuleTypeReactionCalculator, TemplateReactionCalculator{}),
	)
)

func GetTemplateRuleEntity() def.Rule {
	return TemplateRuleEntity
}

func init() {
	GetMod().RegisterRule(GetTemplateRuleEntity())
}
