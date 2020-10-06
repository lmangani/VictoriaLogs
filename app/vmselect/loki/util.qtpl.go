// Code generated by qtc from "util.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/loki/util.qtpl:1
package loki

//line app/vmselect/loki/util.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/storage"
)

//line app/vmselect/loki/util.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/loki/util.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/loki/util.qtpl:7
func streammetricNameObject(qw422016 *qt422016.Writer, mn *storage.MetricName) {
//line app/vmselect/loki/util.qtpl:7
	qw422016.N().S(`{`)
//line app/vmselect/loki/util.qtpl:9
	if len(mn.MetricGroup) > 0 {
//line app/vmselect/loki/util.qtpl:9
		qw422016.N().S(`"__name__":`)
//line app/vmselect/loki/util.qtpl:10
		qw422016.N().QZ(mn.MetricGroup)
//line app/vmselect/loki/util.qtpl:10
		if len(mn.Tags) > 0 {
//line app/vmselect/loki/util.qtpl:10
			qw422016.N().S(`,`)
//line app/vmselect/loki/util.qtpl:10
		}
//line app/vmselect/loki/util.qtpl:11
	}
//line app/vmselect/loki/util.qtpl:12
	for j := range mn.Tags {
//line app/vmselect/loki/util.qtpl:13
		tag := &mn.Tags[j]

//line app/vmselect/loki/util.qtpl:14
		qw422016.N().QZ(tag.Key)
//line app/vmselect/loki/util.qtpl:14
		qw422016.N().S(`:`)
//line app/vmselect/loki/util.qtpl:14
		qw422016.N().QZ(tag.Value)
//line app/vmselect/loki/util.qtpl:14
		if j+1 < len(mn.Tags) {
//line app/vmselect/loki/util.qtpl:14
			qw422016.N().S(`,`)
//line app/vmselect/loki/util.qtpl:14
		}
//line app/vmselect/loki/util.qtpl:15
	}
//line app/vmselect/loki/util.qtpl:15
	qw422016.N().S(`}`)
//line app/vmselect/loki/util.qtpl:17
}

//line app/vmselect/loki/util.qtpl:17
func writemetricNameObject(qq422016 qtio422016.Writer, mn *storage.MetricName) {
//line app/vmselect/loki/util.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/util.qtpl:17
	streammetricNameObject(qw422016, mn)
//line app/vmselect/loki/util.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/util.qtpl:17
}

//line app/vmselect/loki/util.qtpl:17
func metricNameObject(mn *storage.MetricName) string {
//line app/vmselect/loki/util.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/util.qtpl:17
	writemetricNameObject(qb422016, mn)
//line app/vmselect/loki/util.qtpl:17
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/util.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/util.qtpl:17
	return qs422016
//line app/vmselect/loki/util.qtpl:17
}

//line app/vmselect/loki/util.qtpl:19
func streammetricRow(qw422016 *qt422016.Writer, timestamp int64, value float64) {
//line app/vmselect/loki/util.qtpl:19
	qw422016.N().S(`[`)
//line app/vmselect/loki/util.qtpl:20
	qw422016.N().F(float64(timestamp) / 1e3)
//line app/vmselect/loki/util.qtpl:20
	qw422016.N().S(`,"`)
//line app/vmselect/loki/util.qtpl:20
	qw422016.N().F(value)
//line app/vmselect/loki/util.qtpl:20
	qw422016.N().S(`"]`)
//line app/vmselect/loki/util.qtpl:21
}

//line app/vmselect/loki/util.qtpl:21
func writemetricRow(qq422016 qtio422016.Writer, timestamp int64, value float64) {
//line app/vmselect/loki/util.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/util.qtpl:21
	streammetricRow(qw422016, timestamp, value)
//line app/vmselect/loki/util.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/util.qtpl:21
}

//line app/vmselect/loki/util.qtpl:21
func metricRow(timestamp int64, value float64) string {
//line app/vmselect/loki/util.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/util.qtpl:21
	writemetricRow(qb422016, timestamp, value)
//line app/vmselect/loki/util.qtpl:21
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/util.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/util.qtpl:21
	return qs422016
//line app/vmselect/loki/util.qtpl:21
}

//line app/vmselect/loki/util.qtpl:23
func streammetricDataRow(qw422016 *qt422016.Writer, timestamp int64, data []byte) {
//line app/vmselect/loki/util.qtpl:23
	qw422016.N().S(`["`)
//line app/vmselect/loki/util.qtpl:24
	qw422016.N().DL(timestamp * 1e6)
//line app/vmselect/loki/util.qtpl:24
	qw422016.N().S(`","`)
//line app/vmselect/loki/util.qtpl:24
	qw422016.N().Z(data)
//line app/vmselect/loki/util.qtpl:24
	qw422016.N().S(`"]`)
//line app/vmselect/loki/util.qtpl:25
}

//line app/vmselect/loki/util.qtpl:25
func writemetricDataRow(qq422016 qtio422016.Writer, timestamp int64, data []byte) {
//line app/vmselect/loki/util.qtpl:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/util.qtpl:25
	streammetricDataRow(qw422016, timestamp, data)
//line app/vmselect/loki/util.qtpl:25
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/util.qtpl:25
}

//line app/vmselect/loki/util.qtpl:25
func metricDataRow(timestamp int64, data []byte) string {
//line app/vmselect/loki/util.qtpl:25
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/util.qtpl:25
	writemetricDataRow(qb422016, timestamp, data)
//line app/vmselect/loki/util.qtpl:25
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/util.qtpl:25
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/util.qtpl:25
	return qs422016
//line app/vmselect/loki/util.qtpl:25
}

//line app/vmselect/loki/util.qtpl:27
func streamvaluesWithTimestamps(qw422016 *qt422016.Writer, values []float64, timestamps []int64) {
//line app/vmselect/loki/util.qtpl:28
	if len(values) == 0 {
//line app/vmselect/loki/util.qtpl:28
		qw422016.N().S(`[]`)
//line app/vmselect/loki/util.qtpl:30
		return
//line app/vmselect/loki/util.qtpl:31
	}
//line app/vmselect/loki/util.qtpl:31
	qw422016.N().S(`[`)
//line app/vmselect/loki/util.qtpl:33
	/* inline metricRow call here for the sake of performance optimization */

//line app/vmselect/loki/util.qtpl:33
	qw422016.N().S(`[`)
//line app/vmselect/loki/util.qtpl:34
	qw422016.N().F(float64(timestamps[0]) / 1e3)
//line app/vmselect/loki/util.qtpl:34
	qw422016.N().S(`,"`)
//line app/vmselect/loki/util.qtpl:34
	qw422016.N().F(values[0])
//line app/vmselect/loki/util.qtpl:34
	qw422016.N().S(`"]`)
//line app/vmselect/loki/util.qtpl:36
	timestamps = timestamps[1:]
	values = values[1:]

//line app/vmselect/loki/util.qtpl:39
	if len(values) > 0 {
//line app/vmselect/loki/util.qtpl:41
		// Remove bounds check inside the loop below
		_ = timestamps[len(values)-1]

//line app/vmselect/loki/util.qtpl:44
		for i, v := range values {
//line app/vmselect/loki/util.qtpl:45
			/* inline metricRow call here for the sake of performance optimization */

//line app/vmselect/loki/util.qtpl:45
			qw422016.N().S(`,[`)
//line app/vmselect/loki/util.qtpl:46
			qw422016.N().F(float64(timestamps[i]) / 1e3)
//line app/vmselect/loki/util.qtpl:46
			qw422016.N().S(`,"`)
//line app/vmselect/loki/util.qtpl:46
			qw422016.N().F(v)
//line app/vmselect/loki/util.qtpl:46
			qw422016.N().S(`"]`)
//line app/vmselect/loki/util.qtpl:47
		}
//line app/vmselect/loki/util.qtpl:48
	}
//line app/vmselect/loki/util.qtpl:48
	qw422016.N().S(`]`)
//line app/vmselect/loki/util.qtpl:50
}

//line app/vmselect/loki/util.qtpl:50
func writevaluesWithTimestamps(qq422016 qtio422016.Writer, values []float64, timestamps []int64) {
//line app/vmselect/loki/util.qtpl:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/util.qtpl:50
	streamvaluesWithTimestamps(qw422016, values, timestamps)
//line app/vmselect/loki/util.qtpl:50
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/util.qtpl:50
}

//line app/vmselect/loki/util.qtpl:50
func valuesWithTimestamps(values []float64, timestamps []int64) string {
//line app/vmselect/loki/util.qtpl:50
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/util.qtpl:50
	writevaluesWithTimestamps(qb422016, values, timestamps)
//line app/vmselect/loki/util.qtpl:50
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/util.qtpl:50
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/util.qtpl:50
	return qs422016
//line app/vmselect/loki/util.qtpl:50
}

//line app/vmselect/loki/util.qtpl:52
func streamdatasWithTimestamps(qw422016 *qt422016.Writer, values [][]byte, timestamps []int64) {
//line app/vmselect/loki/util.qtpl:53
	if len(values) == 0 {
//line app/vmselect/loki/util.qtpl:53
		qw422016.N().S(`[]`)
//line app/vmselect/loki/util.qtpl:55
		return
//line app/vmselect/loki/util.qtpl:56
	}
//line app/vmselect/loki/util.qtpl:56
	qw422016.N().S(`[`)
//line app/vmselect/loki/util.qtpl:58
	/* inline metricRow call here for the sake of performance optimization */

//line app/vmselect/loki/util.qtpl:58
	qw422016.N().S(`["`)
//line app/vmselect/loki/util.qtpl:59
	qw422016.N().DL(timestamps[0] * 1e6)
//line app/vmselect/loki/util.qtpl:59
	qw422016.N().S(`","`)
//line app/vmselect/loki/util.qtpl:59
	qw422016.N().Z(values[0])
//line app/vmselect/loki/util.qtpl:59
	qw422016.N().S(`"]`)
//line app/vmselect/loki/util.qtpl:61
	timestamps = timestamps[1:]
	values = values[1:]

//line app/vmselect/loki/util.qtpl:64
	if len(values) > 0 {
//line app/vmselect/loki/util.qtpl:66
		// Remove bounds check inside the loop below
		_ = timestamps[len(values)-1]

//line app/vmselect/loki/util.qtpl:69
		for i, v := range values {
//line app/vmselect/loki/util.qtpl:70
			/* inline metricRow call here for the sake of performance optimization */

//line app/vmselect/loki/util.qtpl:70
			qw422016.N().S(`,["`)
//line app/vmselect/loki/util.qtpl:71
			qw422016.N().DL(timestamps[i] * 1e6)
//line app/vmselect/loki/util.qtpl:71
			qw422016.N().S(`","`)
//line app/vmselect/loki/util.qtpl:71
			qw422016.N().Z(v)
//line app/vmselect/loki/util.qtpl:71
			qw422016.N().S(`"]`)
//line app/vmselect/loki/util.qtpl:72
		}
//line app/vmselect/loki/util.qtpl:73
	}
//line app/vmselect/loki/util.qtpl:73
	qw422016.N().S(`]`)
//line app/vmselect/loki/util.qtpl:75
}

//line app/vmselect/loki/util.qtpl:75
func writedatasWithTimestamps(qq422016 qtio422016.Writer, values [][]byte, timestamps []int64) {
//line app/vmselect/loki/util.qtpl:75
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/util.qtpl:75
	streamdatasWithTimestamps(qw422016, values, timestamps)
//line app/vmselect/loki/util.qtpl:75
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/util.qtpl:75
}

//line app/vmselect/loki/util.qtpl:75
func datasWithTimestamps(values [][]byte, timestamps []int64) string {
//line app/vmselect/loki/util.qtpl:75
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/util.qtpl:75
	writedatasWithTimestamps(qb422016, values, timestamps)
//line app/vmselect/loki/util.qtpl:75
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/util.qtpl:75
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/util.qtpl:75
	return qs422016
//line app/vmselect/loki/util.qtpl:75
}
