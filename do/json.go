package do

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func anyToJson(v any) (_ string, err error) {

	b := bytes.Buffer{}
	encoder := json.NewEncoder(&b)

	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")

	err = encoder.Encode(v)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func ginToJsonMap(c *gin.Context) (_ map[string]any, err error) {
	raw := map[string]any{}
	err = c.ShouldBindBodyWith(&raw, binding.JSON)
	return raw, err
}

func ginToJson(c *gin.Context) (_ string, err error) {
	raw, err := ginToJsonMap(c)
	if err != nil {
		return "", err
	}
	return anyToJson(raw)
}

// JsonMap: convert v, to json map,
func JsonMap(v *gin.Context) (_ map[string]any, err error) {
	return ginToJsonMap(v)
}

// Json: convert v, to json str,
func Json(v any) (_ string, err error) {
	switch v := v.(type) {
	case *gin.Context:
		return ginToJson(v)
	default:
		return anyToJson(v)
	}
}
