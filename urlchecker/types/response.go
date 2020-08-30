package types

// Response is type of urlchecker response
type Response struct {
	url    string
	status string
}

// GenResponse is response generator of urlchecker
func GenResponse(url string, status string) Response {
	return Response{url: url, status: status}
}
