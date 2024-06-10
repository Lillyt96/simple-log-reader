package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_extractData(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *logData
	}{
		{
			name: "Success - one record of each found",
			args: args{
				input: "177.71.128.21 - - [10/Jul/2018:22:21:28 +0200] \"GET /intranet-analytics/ HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7\"\n",
			},
			want: &logData{
				urls: result{
					[]string{"/intranet-analytics/"},
				},
				ipAddresses: result{
					[]string{"177.71.128.21"},
				},
			},
		},
		{
			name: "Success - multiple records of each found",
			args: args{
				input: "177.71.128.21 - - [10/Jul/2018:22:21:28 +0200] \"GET /intranet-analytics/ HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7\"\n" +
					"168.41.191.40 - - [09/Jul/2018:10:10:38 +0200] \"GET http://example.net/blog/category/meta/ HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_7) AppleWebKit/534.24 (KHTML, like Gecko) RockMelt/0.9.58.494 Chrome/11.0.696.71 Safari/534.24\"\n" +
					"72.44.32.10 - - [09/Jul/2018:15:48:07 +0200] \"GET / HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (compatible; MSIE 10.6; Windows NT 6.1; Trident/5.0; InfoPath.2; SLCC1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET CLR 2.0.50727) 3gpp-gba UNTRUSTED/1.0\" junk extra\n",
			},
			want: &logData{
				urls: result{
					[]string{"/intranet-analytics/", "/blog/category/meta/", "/"},
				},
				ipAddresses: result{
					[]string{"177.71.128.21", "168.41.191.40", "72.44.32.10"},
				},
			},
		},
		{
			name: "Success - no records found",
			args: args{
				input: "this log does not contain an ip or url address",
			},
			want: &logData{
				urls:        result{},
				ipAddresses: result{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractLogData(tt.args.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_result_getUniqueValues(t *testing.T) {
	type fields struct {
		result []string
	}
	tests := []struct {
		name           string
		fields         fields
		wantUniqueList []string
	}{
		{
			name: "Success - IP Addresses",
			fields: fields{
				result: []string{
					"177.71.128.21",
					"177.71.128.21",
					"177.71.128.21",
					"200.71.128.21",
					"2.71.128.21",
					"2.71.128.21"},
			},
			wantUniqueList: []string{
				"177.71.128.21",
				"200.71.128.21",
				"2.71.128.21",
			},
		},
		{
			name: "Success - urls",
			fields: fields{
				result: []string{
					"/to-an-error",
					"http://example.net/faq/",
					"faq/how-to-install/",
					"faq/how-to-install/",
					"/blog/category/community/",
					"/blog/category/community/",
					"/blog/category/community/"},
			},
			wantUniqueList: []string{
				"/to-an-error",
				"http://example.net/faq/",
				"faq/how-to-install/",
				"/blog/category/community/",
			},
		},
		{
			name: "Success - no duplicates",
			fields: fields{
				result: []string{
					"/to-an-error",
					"http://example.net/faq/"},
			},
			wantUniqueList: []string{
				"/to-an-error",
				"http://example.net/faq/",
			},
		},
		{
			name: "Success - empty list",
			fields: fields{
				result: []string{},
			},
			wantUniqueList: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := result{
				result: tt.fields.result,
			}
			gotUniqueList := w.getUniqueValues()
			assert.ElementsMatchf(t, tt.wantUniqueList, gotUniqueList, "getUnique()")
		})
	}
}

func Test_result_getUniqueWithCount(t *testing.T) {
	type fields struct {
		result []string
	}
	tests := []struct {
		name                string
		fields              fields
		wantUniqueListCount map[string]int
	}{
		{
			name: "Success - IP Addresses",
			fields: fields{
				result: []string{
					"177.71.128.21",
					"177.71.128.21",
					"177.71.128.21",
					"200.71.128.21",
					"2.71.128.21",
					"2.71.128.21"},
			},
			wantUniqueListCount: map[string]int{
				"177.71.128.21": 3,
				"200.71.128.21": 1,
				"2.71.128.21":   2,
			},
		},
		{
			name: "Success - urls",
			fields: fields{
				result: []string{
					"/to-an-error",
					"http://example.net/faq/",
					"faq/how-to-install/",
					"faq/how-to-install/",
					"/blog/category/community/",
					"/blog/category/community/",
					"/blog/category/community/"},
			},
			wantUniqueListCount: map[string]int{
				"/to-an-error":              1,
				"http://example.net/faq/":   1,
				"faq/how-to-install/":       2,
				"/blog/category/community/": 3,
			},
		},
		{
			name: "Success - no duplicates",
			fields: fields{
				result: []string{
					"/to-an-error",
					"http://example.net/faq/"},
			},
			wantUniqueListCount: map[string]int{
				"/to-an-error":            1,
				"http://example.net/faq/": 1,
			},
		},
		{
			name: "Success - empty list",
			fields: fields{
				result: []string{},
			},
			wantUniqueListCount: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := result{
				result: tt.fields.result,
			}
			gotUniqueListCount := mapUniqueValuesToCount(w.result)
			assert.Equalf(t, tt.wantUniqueListCount, gotUniqueListCount, "getUniqueWithCount()")
		})
	}
}

func Test_result_getTop3Values(t *testing.T) {
	type fields struct {
		result []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []uniqueWithCount
	}{
		{
			name: "Success - get top 3 values IP Addresses",
			fields: fields{
				[]string{
					"168.41.191.43",
					"168.41.191.43",
					"168.41.191.43",
					"41.41.00.43",
					"41.41.00.43",
					"50.41.00.43",
					"50.41.00.43",
					"1.41.00.43",
				},
			},
			want: []uniqueWithCount{
				{
					value: "168.41.191.43",
					count: 3,
				},
				{
					value: "41.41.00.43",
					count: 2,
				},
				{
					value: "50.41.00.43",
					count: 2,
				},
			},
		},
		{
			name: "Success - get top 3 values URLs",
			fields: fields{
				[]string{
					"/translations/",
					"/docs/manage-websites/",
					"/temp-redirect",
					"/translations/",
					"/newsletter/",
					"/docs/manage-websites/",
					"/translations/",
					"/newsletter/",
				},
			},
			want: []uniqueWithCount{
				{
					value: "/translations/",
					count: 3,
				},
				{
					value: "/docs/manage-websites/",
					count: 2,
				},
				{
					value: "/newsletter/",
					count: 2,
				},
			},
		},
		{
			name: "Success - no top values (it will take the first they're read)",
			fields: fields{
				[]string{
					"/translations/",
					"/docs/manage-websites/",
					"/temp-redirect",
				},
			},
			want: []uniqueWithCount{
				{
					value: "/translations/",
					count: 1,
				},
				{
					value: "/docs/manage-websites/",
					count: 1,
				},
				{
					value: "/temp-redirect",
					count: 1,
				},
			},
		},
		{
			name: "Success - less than 3 values",
			fields: fields{
				[]string{
					"/translations/",
					"/translations/",
					"/docs/manage-websites/",
					"/docs/manage-websites/",
				},
			},
			want: []uniqueWithCount{
				{
					value: "/translations/",
					count: 2,
				},
				{
					value: "/docs/manage-websites/",
					count: 2,
				},
			},
		},
		{
			name: "Success - no values",
			fields: fields{
				[]string{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := result{
				result: tt.fields.result,
			}
			assert.ElementsMatchf(t, tt.want, w.getTop3Values(), "getTop3Values()")
		})
	}
}
