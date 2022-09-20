package header

import http "github.com/saucesteals/fhttp"

func SecureTouchHeader(userAgent, sch, clientePoch, Instanceuuid, authorization, clientVersion, site string) http.Header {

	return http.Header{
	
		"sec-ch-ua":          {sch},
		"attempt":            {`0`},
		"sec-ch-ua-mobile":   {`?0`},
		"authorization":      {authorization},
		"clientepoch":        {clientePoch},
		"content-type":       {`application/json`},
		"accept":             {`application/json`},
		"instanceuuid":       {Instanceuuid},
		"user-agent":         {userAgent},
		"clientversion":      {clientVersion},
		"sec-ch-ua-platform": {`"Windows"`},
		"origin":             {site},
		"sec-fetch-site":     {`same-site`},
		"sec-fetch-mode":     {`cors`},
		"sec-fetch-dest":     {`empty`},
		"referer":            {site},
		"accept-encoding":    {`gzip, deflate, br`},
		"accept-language":    {`en-GB,en;q=0.9`},

		http.HeaderOrderKey: {
			"content-length",
			"sec-ch-ua",
			"content-encoding",
			"attempt",
			"sec-ch-ua-mobile",
			"authorization",
			"clientepoch",
			"content-type",
			"accept",
			"instanceuuid",
			"user-agent",
			"encrypted",
			"clientversion",
			"sec-ch-ua-platform",
			"origin",
			"sec-fetch-site",
			"sec-fetch-mode",
			"sec-fetch-dest",
			"referer",
			"accept-encoding",
			"accept-language",
			"cookie",
		},
		http.PHeaderOrderKey: {
			":method",
			":authority",
			":scheme",
			":path",
		},
	}

}
