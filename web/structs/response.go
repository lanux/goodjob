package structs

type ResponseBasic struct {
	Code    int         `json:"xml" xml`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}
