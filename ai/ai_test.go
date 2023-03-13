package ai_test

import (
	"testing"

	"github.com/aiomni/aish/ai"
)

func TestAskChatGPT(t *testing.T) {
	content, err := ai.AskChatGPT("删除当前文件夹和下面的所有文件")
	if err != nil {
		t.Error(err)
	}
	t.Errorf(content)
}
