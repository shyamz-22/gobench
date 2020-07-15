package gobench

import "testing"

var val string
var ex error


func BenchmarkJson(b *testing.B) {
	var (
		value string
		err   error
	)
	config := Configuration{
		Patterns: []string{"info", "error"},
		Url:      "https://my.url.com",
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		value, err = JsonEncoder(config)
		//value, err = JsonMarshal(config)
	}

	val = value
	ex = err
}
