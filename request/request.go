package request

var (
	retryTimes int
	timeout    int64
	rawCookie  string
	userAgent  string
	refer      string
)

func SetOptions(opt RequestOptions) {
	retryTimes = opt.RetryTimes
	timeout = opt.Timeout
	rawCookie = opt.Cookie
	userAgent = opt.UserAgent
	refer = opt.Refer
}
