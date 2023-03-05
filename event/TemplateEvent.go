package event

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	def "github.com/sunist-c/genius-invokation-simulator-backend/mod/definition"
	impl "github.com/sunist-c/genius-invokation-simulator-backend/mod/implement"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"

	. "genius-invokation-simulator-mod-template/mod"
)

type TemplateEVentType def.Event

var (
	TemplateEventEntity TemplateEVentType = impl.NewEventWithOpts(
		impl.WithEventID(30001),
		impl.WithEventTriggerAt(enum.AfterNone),
		impl.WithEventTriggerNow(func(ctx context.CallbackContext) bool {
			// todo: implement me
			panic("not implemented yet")
		}),
		impl.WithEventClearNow(func() bool {
			// todo: implement me
			panic("not implemented yet")
		}),
		impl.WithEventCallback(func(ctx *context.CallbackContext) {
			// todo: implement me
			panic("not implemented yet")
		}),
	)
)

func GetTemplateEventEntity() def.Event {
	return TemplateEventEntity
}

func init() {
	GetMod().RegisterEvent(GetTemplateEventEntity())
}
