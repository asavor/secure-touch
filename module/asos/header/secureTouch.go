package header

import http "github.com/saucesteals/fhttp"

func SecureTouchTokenHeader(userAgent, sch string) http.Header {

	return http.Header{
		"sec-ch-ua":          {sch},
		"sec-ch-ua-mobile":   {`?0`},
		"user-agent":         {userAgent},
		"sec-ch-ua-platform": {`"Windows"`},
		"accept":             {`*/*`},
		"sec-fetch-site":     {`same-site`},
		"sec-fetch-mode":     {`no-cors`},
		"sec-fetch-dest":     {`script`},
		"referer":            {`https://my.asos.com/`},
		"accept-encoding":    {`gzip, deflate, br`},
		"accept-language":    {`en-GB,en;q=0.9`},

		http.HeaderOrderKey: {
			"sec-ch-ua",
			"sec-ch-ua-mobile",
			"user-agent",
			"sec-ch-ua-platform",
			"accept",
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
