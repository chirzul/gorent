package helpers

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, statusCode int, msg string, err error, args ...interface{}) {
	var result = make(map[string]interface{})
	var desc string

	switch statusCode {
	case 200:
		desc = "OK"
		w.WriteHeader(http.StatusOK)
	case 201:
		desc = "Created"
		w.WriteHeader(http.StatusCreated)
	case 202:
		desc = "Accepted"
		w.WriteHeader(http.StatusAccepted)
	case 304:
		desc = "Not Modified"
		http.Error(w, "", http.StatusNotModified)
	case 400:
		desc = "Bad Request"
		http.Error(w, "", http.StatusBadRequest)
	case 401:
		desc = "Unauthorized"
		http.Error(w, "", http.StatusUnauthorized)
	case 403:
		desc = "Forbidden"
		http.Error(w, "", http.StatusForbidden)
	case 404:
		desc = "Not Found"
		http.Error(w, "", http.StatusNotFound)
	case 500:
		desc = "Internal Server Error"
		http.Error(w, "", http.StatusInternalServerError)
	case 502:
		desc = "Bad Gateway"
		http.Error(w, "", http.StatusBadGateway)
	default:
		desc = ""
	}

	if err != nil {
		result["code"] = statusCode
		result["description"] = desc
		result["error"] = err.Error()
		result["message"] = msg
	} else {
		result["code"] = statusCode
		result["description"] = desc
		result["message"] = msg
	}
	if len(args) > 0 {
		result["resultData"] = args[0]
	}

	json.NewEncoder(w).Encode(result)
}
