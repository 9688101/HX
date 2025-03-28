package vertexai

import (
	"net/http"

	claude "github.com/9688101/HX/relay/adaptor/vertexai/claude"
	gemini "github.com/9688101/HX/relay/adaptor/vertexai/gemini"
	"github.com/9688101/HX/relay/meta"
	"github.com/9688101/HX/relay/model"
	"github.com/gin-gonic/gin"
)

type VertexAIModelType int

const (
	VerterAIClaude VertexAIModelType = iota + 1
	VerterAIGemini
)

var modelMapping = map[string]VertexAIModelType{}
var modelList = []string{}

func init() {
	modelList = append(modelList, claude.ModelList...)
	for _, model := range claude.ModelList {
		modelMapping[model] = VerterAIClaude
	}

	modelList = append(modelList, gemini.ModelList...)
	for _, model := range gemini.ModelList {
		modelMapping[model] = VerterAIGemini
	}
}

type innerAIAdapter interface {
	ConvertRequest(c *gin.Context, relayMode int, request *model.GeneralOpenAIRequest) (any, error)
	DoResponse(c *gin.Context, resp *http.Response, meta *meta.Meta) (usage *model.Usage, err *model.ErrorWithStatusCode)
}

func GetAdaptor(model string) innerAIAdapter {
	adaptorType := modelMapping[model]
	switch adaptorType {
	case VerterAIClaude:
		return &claude.Adaptor{}
	case VerterAIGemini:
		return &gemini.Adaptor{}
	default:
		return nil
	}
}
