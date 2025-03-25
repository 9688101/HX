package aiproxy

import "github.com/9688101/HX/relay/adaptor/openai"

var ModelList = []string{""}

func init() {
	ModelList = openai.ModelList
}
