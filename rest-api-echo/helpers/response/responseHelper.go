package responsehelper

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

/*
XMLJSONResponseHelper - Response Result Json or Xml
 */
func XMLJSONResponseHelper(outputData interface{}, typeResponse string, context echo.Context) error {

	if typeResponse == "json"{
		return context.JSON(http.StatusOK, outputData)
	}else if typeResponse == "xml"{
		return context.XMLPretty(http.StatusOK, outputData,  "  ")
	}else{
		var errMsg map[string]string = map[string]string{"Result":"Error", "Message":"No Data Response Select"}
		resultJSON, _ := json.Marshal(errMsg)
		return context.JSON(http.StatusOK, resultJSON)
	}
	return nil
}
