package header

import http "github.com/saucesteals/fhttp"

func PostLoginHeader(userAgent, sch string) http.Header {
	return http.Header{
		"cache-control":             {`max-age=0`},
		"sec-ch-ua":                 {sch},
		"sec-ch-ua-mobile":          {`?0`},
		"sec-ch-ua-platform":        {`"Windows"`},
		"upgrade-insecure-requests": {`1`},
		"origin":                    {`https://my.asos.com`},
		"content-type":              {`application/x-www-form-urlencoded`},
		"user-agent":                {userAgent},
		"accept":                    {`text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`},
		"sec-fetch-site":            {`same-origin`},
		"sec-fetch-mode":            {`navigate`},
		"sec-fetch-user":            {`?1`},
		"sec-fetch-dest":            {`document`},
		"accept-encoding":           {`gzip, deflate, br`},
		"accept-language":           {`en-GB,en;q=0.9`},

		http.HeaderOrderKey: {
			"content-length",
			"cache-control",
			"sec-ch-ua",
			"sec-ch-ua-mobile",
			"sec-ch-ua-platform",
			"upgrade-insecure-requests",
			"origin",
			"content-type",
			"user-agent",
			"accept",
			"sec-fetch-site",
			"sec-fetch-mode",
			"sec-fetch-user",
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
