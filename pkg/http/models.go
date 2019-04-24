package http

type Info struct {
	Version string `json:"version"`
	Network string `json:"network"`
	Chain   struct {
		Height   int    `json:"height"`
		Tip      string `json:"tip"`
		Progress int    `json:"progress"`
	} `json:"chain"`
	Pool struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Agent    string `json:"agent"`
		Services string `json:"services"`
		Outbound int    `json:"outbound"`
		Inbound  int    `json:"inbound"`
	} `json:"pool"`
	Mempool struct {
		Tx   int `json:"tx"`
		Size int `json:"size"`
	} `json:"mempool"`
	Time struct {
		Uptime   int `json:"uptime"`
		System   int `json:"system"`
		Adjusted int `json:"adjusted"`
		Offset   int `json:"offset"`
	} `json:"time"`
	Memory struct {
		Total       int `json:"total"`
		JsHeap      int `json:"jsHeap"`
		JsHeapTotal int `json:"jsHeapTotal"`
		NativeHeap  int `json:"nativeHeap"`
		External    int `json:"external"`
	} `json:"memory"`
}
