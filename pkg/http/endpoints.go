package http

func (hcc *HttpClientConfig) Info() (Info, error) {
	requestURL := hcc.Hostname + "/"

	var info Info

	err := call(requestURL, &info)

	return info, err

}
