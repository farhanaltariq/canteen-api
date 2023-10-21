package model

type SuccessResponse struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type UserAuthResponse struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

func NewSuccessResponse(id int, msg string) SuccessResponse {
	return SuccessResponse{
		UserID:  id,
		Message: msg,
	}
}

func NewErrorResponse(msg string) ErrorResponse {
	return ErrorResponse{
		Error: msg,
	}
}

func NewAccessTokenResponse(username string, accessToken string) UserAuthResponse {
	return UserAuthResponse{
		Username:    username,
		AccessToken: "Bearer " + accessToken,
	}
}
