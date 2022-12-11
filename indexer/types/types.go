package types

type DataEnron struct {
	Index   string   `json:"index"`
	Records []Record `json:"records"`
}

type Record struct {
	Document Data `json:"document"`
}

type Data struct {
	MessageID               string
	Date                    string
	From                    string
	To                      string
	Subject                 string
	Cc                      string
	MimeVersion             string
	ContentType             string
	ContentTransferEncoding string
	XFrom                   string
	XTo                     string
	Xcc                     string
	Xbcc                    string
	XFolder                 string
	XOrigin                 string
	XFileName               string
	Body                    string
}
