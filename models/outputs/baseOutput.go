package outputs

type PaginationOutput struct {
	Page      int    `json:"page" default:"1"`
	Limit     int    `json:"limit" default:"10"`
	OrderBy   string `json:"order_by" default:"id"`
	OrderType string `json:"order_type" default:"asc"`
	TotalData int    `json:"total_data" default:"0"`
	TotalTake int    `json:"total_take" default:"0"`
	TotalPage int    `json:"total_page" default:"0"`
}

type BaseOutput struct {
	Code    int    `json:"code" default:"200"`
	Message string `json:"message" default:"Success: {message}"`
}

type BadRequestOutput struct {
	Code    int    `json:"code" default:"400"`
	Message string `json:"message" default:"Bad Request: {message}"`
}

type UnauthorizedOutput struct {
	Code    int    `json:"code" default:"401"`
	Message string `json:"message" default:"Unauthorized: {message}"`
}

type ForbiddenOutput struct {
	Code    int    `json:"code" default:"403"`
	Message string `json:"message" default:"Forbidden: {message}"`
}

type NotFoundOutput struct {
	Code    int    `json:"code" default:"404"`
	Message string `json:"message" default:"Not Found: {message}"`
}

type InternalServerErrorOutput struct {
	Code    int    `json:"code" default:"500"`
	Message string `json:"message" default:"Internal Server Error: {message}"`
}

type ServiceUnavailableOutput struct {
	Code    int    `json:"code" default:"503"`
	Message string `json:"message" default:"Service Unavailable: {message}"`
}
