package card

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	def "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	impl "github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"

	. "genius-invokation-simulator-mod-template/mod"
)

type TemplateCardType def.Card

var (
	TemplateCardEntity TemplateCardType = impl.NewCardWithOpts(
		impl.WithCardID(50001),
		impl.WithCardType(enum.CardUnknown),
		impl.WithCardCost(map[enum.ElementType]uint{
			// todo: modify me
		}),
	)
)

func GetTemplateCardEntity() def.Card {
	return TemplateCardEntity
}

func init() {
	GetMod().RegisterCard(GetTemplateCardEntity())
}
