package Model

type ResultItem struct {
	Code uint8  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
