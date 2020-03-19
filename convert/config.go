package elasticsql

const (
	DefaultQueryMapStr string = `{"bool" : {"must": [{"match_all" : {}}]}}`
	DefaultFrom        string = "0"
	DefaultSize        string = "10"
)
