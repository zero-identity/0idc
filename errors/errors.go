package errors

type RFCErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

var (
	INVALID_CLIENT = &RFCErrorResponse{
		Error:            "invalid_client",
		ErrorDescription: "Client authentication failed (e.g., unknown client, no client authentication included, or unsupported authentication method).",
	}

	INVALID_REQUEST = &RFCErrorResponse{
		Error:            "invalid_request",
		ErrorDescription: "The request is missing a required parameter, includes an unsupported parameter value (other than grant type), repeats a parameter, includes multiple credentials, utilizes more than one mechanism for authenticating the client, or is otherwise malformed.",
	}

	UNSUPPORTED_GRANT_TYPE = &RFCErrorResponse{
		Error:            "unsupported_grant_type",
		ErrorDescription: "The authorization grant type is not supported by the authorization server.",
	}
)
