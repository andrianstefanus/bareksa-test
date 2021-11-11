package function

import (
	st "bareksa-test/struck"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/subosito/gotenv"
)

func InitiateEnv(path string) {
	err := gotenv.Load(path)
	if err != nil {
		println(err.Error())
	}
}

func GetPORT() string {
	if os.Getenv("ENV") == "STAGING" {
		return os.Getenv("PORT_STAGING")
	} else {
		return os.Getenv("PORT_PRODUCTION")
	}
}

func ConvertValue(value interface{}, types string, def interface{}) interface{} {
	var result interface{}
	var ok bool

	if types == "string" {
		result, ok = value.(string)
	}
	if types == "int" {
		result, ok = value.(int)
	}

	if !ok {
		return def
	}

	return result
}

func ConvStrToInt(value interface{}, def int) int {
	result, ok := value.(int)

	if !ok {
		return def
	}

	return result
}

func ConvStr(value interface{}, def string) *string {
	result, ok := value.(string)

	if !ok {
		return &def
	}

	return &result
}

func GetContextValue(ctx context.Context, key string) *string {
	return ConvStr(ctx.Value(key), "")
}

func GetBody(body io.ReadCloser, dest interface{}) {
	data, _ := ioutil.ReadAll(body)
	json.Unmarshal([]byte(data), &dest)
}

func Marshal(target interface{}) string {
	result, err := json.Marshal(target)

	if err != nil {
		return fmt.Sprintf("%+v", target)
	}

	return string(result)
}

func UnMarshal(target string, dest *interface{}) {
	err := json.Unmarshal([]byte(target), dest)

	if err != nil {
		println(fmt.Sprintf("%s", err.Error()))
	}
}

func Logging(logs *[]st.Log, message string) {
	if os.Getenv("IS_LOGGING") == "true" {
		*logs = append(*logs, st.Log{Datetime: time.Now().String(), Message: message})
	}
}
