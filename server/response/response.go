package response

// Response generic
type Response struct {
	StatusCode int         `json:"statusCode"`
	Reason     int         `json:"reason"`
	Comment    string      `json:"comment"`
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
}

// BasicResponse from API
func BasicResponse(data interface{}, comment string, reason int) Response {
	success := false
	statusCode := 400
	if reason == 1 {
		success = true
		statusCode = 200
	}
	basicResponse := Response{StatusCode: statusCode, Reason: reason, Comment: comment, Success: success, Data: data}
	return basicResponse
}
