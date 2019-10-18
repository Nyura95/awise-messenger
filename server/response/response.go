package response

// Response generic
type Response struct {
	StatusCode int
	Reason     int
	Comment    string
	Success    bool
	Data       interface{}
}

// BasicResponse from API
func BasicResponse(data interface{}, comment string, reason int) Response {
	var success = false
	var statusCode = 400
	if reason == 1 {
		success = true
		statusCode = 200
	}
	basicResponse := Response{statusCode, reason, comment, success, data}
	return basicResponse
}
