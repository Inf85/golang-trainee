package helpers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func TypeResponse(inputData interface{}, typeResponse string, writer http.ResponseWriter) {
	switch typeResponse {
	case "json":
		resultJSON, _ := json.Marshal(inputData)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(resultJSON)
	case "xml":
		resultXML, _ := xml.MarshalIndent(inputData, "", "	")

		writer.Header().Set("Content-Type", "application/xml")
		writer.WriteHeader(http.StatusOK)
		writer.Write(resultXML)
	default:
		var errMsg map[string]string = map[string]string{"Result":"Error", "Message":"No Data Response Select"}
		resultJSON, _ := json.Marshal(errMsg)
		writer.Write(resultJSON)
	}

	//return writer
}
