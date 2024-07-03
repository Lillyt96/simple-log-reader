package logParser

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// benchmark testing
func BenchmarkParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("programming-task-example-data.log")
	}
}

func BenchmarkParseConcurrently(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParseConcurrently("programming-task-example-data.log")
	}
}

// unit testing
func TestParse(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *Logs
	}{
		{
			name: "success - read and parse data",
			args: args{
				data: testLogExampleData(),
			},
			want: &Logs{
				logs: []Log{
					{
						Ip:     "50.112.00.11",
						Time:   "11/Jul/2018:17:31:56 +0200",
						Method: "GET",
						URL:    "/asset.js",
						Status: "200",
					},
					{
						Ip:     "168.41.191.9",
						Time:   "09/Jul/2018:22:56:45 +0200",
						Method: "GET",
						URL:    "/docs/",
						Status: "200",
					},
					{
						Ip:     "168.41.191.34",
						Time:   "10/Jul/2018:21:59:50 +0200",
						Method: "GET",
						URL:    "/faq/how-to/",
						Status: "200",
					},
					{
						Ip:     "50.112.00.11",
						Time:   "11/Jul/2018:17:33:01 +0200",
						Method: "GET",
						URL:    "/asset.css",
						Status: "200",
					},
				}},
		},
		{
			name: "success - incorrect log format is not parsed to Logs",
			args: args{
				data: "72.44.32.10 - - [09/Jul/2018:15:48:07 +0200] \"GET / HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (compatible; MSIE 10.6; Windows NT 6.1; Trident/5.0; InfoPath.2; SLCC1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET CLR 2.0.50727) 3gpp-gba UNTRUSTED/1.0\" junk extra\n" +
					"168.41.191.9 - - [09/Jul/2018:22:56:45 +0200] \"GET /docs/ HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (X11; Linux i686; rv:6.0) Gecko/20100101 Firefox/6.0\" 456 789\n",
			},
			want: &Logs{},
		},
		{
			name: "success - incorrect log with correct log - one log should be persisted",
			args: args{
				data: "72.44.32.10 - - [09/Jul/2018:15:48:07 +0200] \"GET / HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (compatible; MSIE 10.6; Windows NT 6.1; Trident/5.0; InfoPath.2; SLCC1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET CLR 2.0.50727) 3gpp-gba UNTRUSTED/1.0\"\n" +
					"168.41.191.9 - - [09/Jul/2018:22:56:45 +0200] \"GET /docs/ HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (X11; Linux i686; rv:6.0) Gecko/20100101 Firefox/6.0\" 456 789\n",
			},
			want: &Logs{
				logs: []Log{
					{
						Ip:     "72.44.32.10",
						Time:   "09/Jul/2018:15:48:07 +0200",
						Method: "GET",
						URL:    "/",
						Status: "200",
					},
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			err := os.WriteFile(dir+"test_data", []byte(tt.args.data), 0666)

			got, err := Parse(dir + "test_data")

			assert.NoError(t, err)

			assert.ElementsMatch(t, tt.want.logs, got.logs)
		})
	}
}

func TestParseConcurrently(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *Logs
	}{
		{
			name: "success - read and parse data",
			args: args{
				data: testLogExampleData(),
			},
			want: &Logs{
				logs: []Log{
					{
						Ip:     "50.112.00.11",
						Time:   "11/Jul/2018:17:31:56 +0200",
						Method: "GET",
						URL:    "/asset.js",
						Status: "200",
					},
					{
						Ip:     "168.41.191.9",
						Time:   "09/Jul/2018:22:56:45 +0200",
						Method: "GET",
						URL:    "/docs/",
						Status: "200",
					},
					{
						Ip:     "168.41.191.34",
						Time:   "10/Jul/2018:21:59:50 +0200",
						Method: "GET",
						URL:    "/faq/how-to/",
						Status: "200",
					},
					{
						Ip:     "50.112.00.11",
						Time:   "11/Jul/2018:17:33:01 +0200",
						Method: "GET",
						URL:    "/asset.css",
						Status: "200",
					},
				}},
		},
		{
			name: "success - incorrect log format is not parsed to Logs",
			args: args{
				data: "72.44.32.10 - - [09/Jul/2018:15:48:07 +0200] \"GET / HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (compatible; MSIE 10.6; Windows NT 6.1; Trident/5.0; InfoPath.2; SLCC1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET CLR 2.0.50727) 3gpp-gba UNTRUSTED/1.0\" junk extra\n" +
					"168.41.191.9 - - [09/Jul/2018:22:56:45 +0200] \"GET /docs/ HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (X11; Linux i686; rv:6.0) Gecko/20100101 Firefox/6.0\" 456 789\n",
			},
			want: &Logs{},
		},
		{
			name: "success - incorrect log with correct log - one log should be persisted",
			args: args{
				data: "72.44.32.10 - - [09/Jul/2018:15:48:07 +0200] \"GET / HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (compatible; MSIE 10.6; Windows NT 6.1; Trident/5.0; InfoPath.2; SLCC1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET CLR 2.0.50727) 3gpp-gba UNTRUSTED/1.0\"\n" +
					"168.41.191.9 - - [09/Jul/2018:22:56:45 +0200] \"GET /docs/ HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (X11; Linux i686; rv:6.0) Gecko/20100101 Firefox/6.0\" 456 789\n",
			},
			want: &Logs{
				logs: []Log{
					{
						Ip:     "72.44.32.10",
						Time:   "09/Jul/2018:15:48:07 +0200",
						Method: "GET",
						URL:    "/",
						Status: "200",
					},
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			err := os.WriteFile(dir+"test_data", []byte(tt.args.data), 0666)

			got, err := ParseConcurrently(dir + "test_data")

			assert.NoError(t, err)

			assert.ElementsMatch(t, tt.want.logs, got.logs)
		})
	}
}

func testLogExampleData() string {
	logs := "50.112.00.11 - admin [11/Jul/2018:17:31:56 +0200] \"GET /asset.js HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6\"\n" +
		"168.41.191.9 - - [09/Jul/2018:22:56:45 +0200] \"GET /docs/ HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (X11; Linux i686; rv:6.0) Gecko/20100101 Firefox/6.0\"\n" +
		"168.41.191.34 - - [10/Jul/2018:21:59:50 +0200] \"GET /faq/how-to/ HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/5.0)\"\n" +
		"50.112.00.11 - admin [11/Jul/2018:17:33:01 +0200] \"GET /asset.css HTTP/1.1\" 200 3574 \"-\" \"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6\""

	return logs
}
