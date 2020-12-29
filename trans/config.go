package trans

const (
	CL = "cl"
	XT = "xt"
	DT = "dt"
	XH = "xh"
	ZH = "zh"
)

//YouDaoAuthConfig youdao config
type YouDaoAuthConfig struct {
	AppKey    string
	AppSecret string
}

type YouDaoResult struct {
	ErrorCode   string   `json:"errorCode"`
	Translation []string `json:"Translation"`
}

type Result struct {
	Title string
	SubTitle string
	Arg string
}

