package request

type RequestOptions struct {
	RetryTimes int
	Timeout    int64
	Cookie     string
	UserAgent  string
	Refer      string
}
