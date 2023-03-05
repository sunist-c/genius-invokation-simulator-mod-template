package test

import (
	"testing"

	gist "github.com/sunist-c/genius-invokation-simulator-backend/mod/testing"

	. "genius-invokation-simulator-mod-template/mod"
)

func TestGeneration(t *testing.T) {
	t.Run("CheckModImplementations", gist.CheckModImplementTestingFunction(GetMod()))
	t.Run("CheckRepeatedEntityID", gist.CheckRepeatEntityIDTestingFunction(GetMod()))
	t.Run("CheckDescriptions", gist.CheckDescriptionsTestingFunction(GetMod()))
	t.Run("CheckCharacterImplementations", gist.RunCharacterImplementsTestingFunction(GetMod()))
	t.Run("CheckSkillImplementations", gist.RunSkillImplementsTestingFunction(GetMod()))
	t.Run("CheckEventImplementations", gist.RunEventImplementsTestingFunction(GetMod()))
	//t.Run("CheckSummonImplementations", gist.RunSummonImplementsTestingFunction(GetMod()))
	t.Run("CheckCardImplementations", gist.RunCardImplementsTestingFunction(GetMod()))
	t.Run("CheckRuleImplementations", gist.RunRuleImplementsTestingFunction(GetMod()))
}
