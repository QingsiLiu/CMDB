package response

type JSONResponse struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func NewJSONResponse(code int, msg string, result interface{}) *JSONResponse {
	return &JSONResponse{code, msg, result}
}

var (
	Unauthorization = NewJSONResponse(401, "unauthorization", nil)
	OK              = NewJSONResponse(200, "ok", nil)
	BadRequest      = NewJSONResponse(400, "bad request", nil)
)
