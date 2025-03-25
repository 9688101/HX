package relay

import (
	"github.com/9688101/HX/relay/adaptor"
	"github.com/9688101/HX/relay/adaptor/aiproxy"
	"github.com/9688101/HX/relay/adaptor/ali"
	"github.com/9688101/HX/relay/adaptor/anthropic"
	"github.com/9688101/HX/relay/adaptor/aws"
	"github.com/9688101/HX/relay/adaptor/baidu"
	"github.com/9688101/HX/relay/adaptor/cloudflare"
	"github.com/9688101/HX/relay/adaptor/cohere"
	"github.com/9688101/HX/relay/adaptor/coze"
	"github.com/9688101/HX/relay/adaptor/deepl"
	"github.com/9688101/HX/relay/adaptor/gemini"
	"github.com/9688101/HX/relay/adaptor/ollama"
	"github.com/9688101/HX/relay/adaptor/openai"
	"github.com/9688101/HX/relay/adaptor/palm"
	"github.com/9688101/HX/relay/adaptor/proxy"
	"github.com/9688101/HX/relay/adaptor/replicate"
	"github.com/9688101/HX/relay/adaptor/tencent"
	"github.com/9688101/HX/relay/adaptor/vertexai"
	"github.com/9688101/HX/relay/adaptor/xunfei"
	"github.com/9688101/HX/relay/adaptor/zhipu"
	"github.com/9688101/HX/relay/apitype"
)

func GetAdaptor(apiType int) adaptor.Adaptor {
	switch apiType {
	case apitype.AIProxyLibrary:
		return &aiproxy.Adaptor{}
	case apitype.Ali:
		return &ali.Adaptor{}
	case apitype.Anthropic:
		return &anthropic.Adaptor{}
	case apitype.AwsClaude:
		return &aws.Adaptor{}
	case apitype.Baidu:
		return &baidu.Adaptor{}
	case apitype.Gemini:
		return &gemini.Adaptor{}
	case apitype.OpenAI:
		return &openai.Adaptor{}
	case apitype.PaLM:
		return &palm.Adaptor{}
	case apitype.Tencent:
		return &tencent.Adaptor{}
	case apitype.Xunfei:
		return &xunfei.Adaptor{}
	case apitype.Zhipu:
		return &zhipu.Adaptor{}
	case apitype.Ollama:
		return &ollama.Adaptor{}
	case apitype.Coze:
		return &coze.Adaptor{}
	case apitype.Cohere:
		return &cohere.Adaptor{}
	case apitype.Cloudflare:
		return &cloudflare.Adaptor{}
	case apitype.DeepL:
		return &deepl.Adaptor{}
	case apitype.VertexAI:
		return &vertexai.Adaptor{}
	case apitype.Proxy:
		return &proxy.Adaptor{}
	case apitype.Replicate:
		return &replicate.Adaptor{}
	}
	return nil
}
