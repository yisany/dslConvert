package elasticsql

const (
	// params
	defaultQueryMapStr string = `{"bool" : {"must": [{"match_all" : {}}]}}`
	defaultFrom        string = "0"
	defaultSize        string = "10"
	// format
	topBoolMust             string = `{"bool" : {"must" : [%v]}}`
	dslRangeGreaterAndEqual string = `{"range" : {"%v" : {"from" : "%v"}}}`
	dslRangeLessAndEqual    string = `{"range" : {"%v" : {"to" : "%v"}}}`
	dslRangeEqual           string = `{"match_phrase" : {"%v" : {"query" : "%v"}}}`
	dslRangeEqualMissing    string = `{"missing":{"field":"%v"}}`
	dslRangeGreater         string = `{"range" : {"%v" : {"gt" : "%v"}}}`
	dslRangeLess            string = `{"range" : {"%v" : {"lt" : "%v"}}}`
	dslRangeNotEqualMissing string = `{"bool" : {"must_not" : [{"missing":{"field":"%v"}}]}}`
	dslRangeNotEqual        string = `{"bool" : {"must_not" : [{"match_phrase" : {"%v" : {"query" : "%v"}}}]}}`
	dslRangeIn              string = `{"terms" : {"%v" : [%v]}}`
	dslRangeLike            string = `{"match_phrase" : {"%v" : {"query" : "%v"}}}`
	dslRangeNotLike         string = `{"bool" : {"must_not" : {"match_phrase" : {"%v" : {"query" : "%v"}}}}}`
	dslRangeNotIn           string = `{"bool" : {"must_not" : {"terms" : {"%v" : [%v]}}}}`
	dslRange                string = `{"range" : {"%v" : {"from" : "%v", "to" : "%v"}}}`
	dslExists               string = `{"exists" : {"field" : "%v"}}`
)
