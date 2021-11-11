package function

import (
	md "bareksa-test/model"
	st "bareksa-test/struck"
	"errors"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func DoRawQuery(connection *gorm.DB, sql string, params map[string]interface{}, destination interface{}, logs *[]st.Log) (int, interface{}, md.Err) {
	// do query
	Logging(logs, fmt.Sprintf("Do query [%s]", sql))
	result := connection.Raw(sql, params).Scan(destination)
	// err query
	if result.Error != nil {
		Logging(logs, fmt.Sprintf("Get error [%s]", result.Error.Error()))
		return 1, result, GetError(result.Error)
	}
	// success
	Logging(logs, "Query finished")
	return 0, result, md.Err{}
}

func DoBatchQuery(connection *gorm.DB, data interface{}, logs *[]st.Log) (int, interface{}, interface{}) {
	Logging(logs, "Do batch query")
	// do query
	if result := connection.CreateInBatches(data, 500); result.Error != nil {
		Logging(logs, fmt.Sprintf("Get error [%s]", result.Error.Error()))
		return 1, result, result.Error
	}
	// success
	Logging(logs, fmt.Sprintf("Get error [%s]", "Query finished"))
	return 0, nil, nil
}

func GetError(err error) md.Err {
	if err == nil {
		return md.Err{Status: ConvStrToInt(os.Getenv("RESPONSE_SUCCESS_CODE"), 553), Msg: os.Getenv("RESPONSE_SUCCESS")}
	}
	if errors.Is(err, gorm.ErrRegistered) {
		return md.Err{Status: ConvStrToInt(os.Getenv("ERR_DATA_DUPLICATE_CODE"), 553), Msg: os.Getenv("ERR_DATA_DUPLICATE")}
	}

	return md.Err{Status: ConvStrToInt(os.Getenv("ERR_DEFAULT_CODE"), 553), Msg: os.Getenv("ERR_DEFAULT")}
}
