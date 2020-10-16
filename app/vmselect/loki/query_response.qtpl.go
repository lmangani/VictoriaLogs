// Code generated by qtc from "query_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/loki/query_response.qtpl:1
package loki

//line app/vmselect/loki/query_response.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmselect/netstorage"
)

// QueryResponse generates response for /api/v1/query.See https://prometheus.io/docs/prometheus/latest/querying/api/#instant-queries

//line app/vmselect/loki/query_response.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/loki/query_response.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/loki/query_response.qtpl:8
func StreamVectorQueryResponse(qw422016 *qt422016.Writer, rs []netstorage.Result) {
//line app/vmselect/loki/query_response.qtpl:8
	qw422016.N().S(`{"status":"success","data":{"resultType":"vector","result":[`)
//line app/vmselect/loki/query_response.qtpl:14
	if len(rs) > 0 {
//line app/vmselect/loki/query_response.qtpl:14
		qw422016.N().S(`{"metric":`)
//line app/vmselect/loki/query_response.qtpl:16
		streammetricNameObject(qw422016, &rs[0].MetricName)
//line app/vmselect/loki/query_response.qtpl:16
		qw422016.N().S(`,"value": [`)
//line app/vmselect/loki/query_response.qtpl:17
		qw422016.N().F(float64(rs[0].Timestamps[0]) / 1e3)
//line app/vmselect/loki/query_response.qtpl:17
		qw422016.N().S(`,"`)
//line app/vmselect/loki/query_response.qtpl:17
		qw422016.N().F(rs[0].Values[0])
//line app/vmselect/loki/query_response.qtpl:17
		qw422016.N().S(`"]}`)
//line app/vmselect/loki/query_response.qtpl:19
		rs = rs[1:]

//line app/vmselect/loki/query_response.qtpl:20
		for i := range rs {
//line app/vmselect/loki/query_response.qtpl:21
			r := &rs[i]

//line app/vmselect/loki/query_response.qtpl:21
			qw422016.N().S(`,{"metric":`)
//line app/vmselect/loki/query_response.qtpl:23
			streammetricNameObject(qw422016, &r.MetricName)
//line app/vmselect/loki/query_response.qtpl:23
			qw422016.N().S(`,"value": [`)
//line app/vmselect/loki/query_response.qtpl:24
			qw422016.N().F(float64(r.Timestamps[0]) / 1e3)
//line app/vmselect/loki/query_response.qtpl:24
			qw422016.N().S(`,"`)
//line app/vmselect/loki/query_response.qtpl:24
			qw422016.N().F(r.Values[0])
//line app/vmselect/loki/query_response.qtpl:24
			qw422016.N().S(`"]}`)
//line app/vmselect/loki/query_response.qtpl:26
		}
//line app/vmselect/loki/query_response.qtpl:27
	}
//line app/vmselect/loki/query_response.qtpl:27
	qw422016.N().S(`]}}`)
//line app/vmselect/loki/query_response.qtpl:31
}

//line app/vmselect/loki/query_response.qtpl:31
func WriteVectorQueryResponse(qq422016 qtio422016.Writer, rs []netstorage.Result) {
//line app/vmselect/loki/query_response.qtpl:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/query_response.qtpl:31
	StreamVectorQueryResponse(qw422016, rs)
//line app/vmselect/loki/query_response.qtpl:31
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/query_response.qtpl:31
}

//line app/vmselect/loki/query_response.qtpl:31
func VectorQueryResponse(rs []netstorage.Result) string {
//line app/vmselect/loki/query_response.qtpl:31
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/query_response.qtpl:31
	WriteVectorQueryResponse(qb422016, rs)
//line app/vmselect/loki/query_response.qtpl:31
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/query_response.qtpl:31
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/query_response.qtpl:31
	return qs422016
//line app/vmselect/loki/query_response.qtpl:31
}

//line app/vmselect/loki/query_response.qtpl:33
func StreamStreamsQueryResponse(qw422016 *qt422016.Writer, rs []netstorage.Result) {
//line app/vmselect/loki/query_response.qtpl:33
	qw422016.N().S(`{"status":"success","data":{"resultType":"streams","result":[`)
//line app/vmselect/loki/query_response.qtpl:39
	if len(rs) > 0 {
//line app/vmselect/loki/query_response.qtpl:39
		qw422016.N().S(`{"stream":`)
//line app/vmselect/loki/query_response.qtpl:41
		streammetricNameObject(qw422016, &rs[0].MetricName)
//line app/vmselect/loki/query_response.qtpl:41
		qw422016.N().S(`,"value": ["`)
//line app/vmselect/loki/query_response.qtpl:42
		qw422016.N().DL(rs[0].Timestamps[0] * 1e6)
//line app/vmselect/loki/query_response.qtpl:42
		qw422016.N().S(`",`)
//line app/vmselect/loki/query_response.qtpl:42
		qw422016.N().QZ(rs[0].Datas[0])
//line app/vmselect/loki/query_response.qtpl:42
		qw422016.N().S(`]}`)
//line app/vmselect/loki/query_response.qtpl:44
		rs = rs[1:]

//line app/vmselect/loki/query_response.qtpl:45
		for i := range rs {
//line app/vmselect/loki/query_response.qtpl:46
			r := &rs[i]

//line app/vmselect/loki/query_response.qtpl:46
			qw422016.N().S(`,{"stream":`)
//line app/vmselect/loki/query_response.qtpl:48
			streammetricNameObject(qw422016, &r.MetricName)
//line app/vmselect/loki/query_response.qtpl:48
			qw422016.N().S(`,"value": ["`)
//line app/vmselect/loki/query_response.qtpl:49
			qw422016.N().DL(r.Timestamps[0] * 1e6)
//line app/vmselect/loki/query_response.qtpl:49
			qw422016.N().S(`",`)
//line app/vmselect/loki/query_response.qtpl:49
			qw422016.N().QZ(r.Datas[0])
//line app/vmselect/loki/query_response.qtpl:49
			qw422016.N().S(`]}`)
//line app/vmselect/loki/query_response.qtpl:51
		}
//line app/vmselect/loki/query_response.qtpl:52
	}
//line app/vmselect/loki/query_response.qtpl:52
	qw422016.N().S(`]}}`)
//line app/vmselect/loki/query_response.qtpl:56
}

//line app/vmselect/loki/query_response.qtpl:56
func WriteStreamsQueryResponse(qq422016 qtio422016.Writer, rs []netstorage.Result) {
//line app/vmselect/loki/query_response.qtpl:56
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/query_response.qtpl:56
	StreamStreamsQueryResponse(qw422016, rs)
//line app/vmselect/loki/query_response.qtpl:56
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/query_response.qtpl:56
}

//line app/vmselect/loki/query_response.qtpl:56
func StreamsQueryResponse(rs []netstorage.Result) string {
//line app/vmselect/loki/query_response.qtpl:56
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/query_response.qtpl:56
	WriteStreamsQueryResponse(qb422016, rs)
//line app/vmselect/loki/query_response.qtpl:56
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/query_response.qtpl:56
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/query_response.qtpl:56
	return qs422016
//line app/vmselect/loki/query_response.qtpl:56
}