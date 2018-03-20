package api

import (
	"github.com/vechain/thor/api/utils/types"
	"github.com/vechain/thor/logdb"
)

//LogInterface for query logs
type LogInterface struct {
	ldb *logdb.LogDB
}

//NewLogInterface new LogInterface
func NewLogInterface(ldb *logdb.LogDB) *LogInterface {
	return &LogInterface{
		ldb,
	}
}

//Filter query logs with option
func (li *LogInterface) Filter(option types.FilterOption) ([]types.Log, error) {
	op, err := option.ToLogFilter()
	if err != nil {
		return nil, err
	}
	logs, err := li.ldb.Filter(op)
	if err != nil {
		return nil, err
	}
	lgs := make([]types.Log, len(logs))
	for i, log := range logs {
		lgs[i] = types.ConvertLog(log)
	}
	return lgs, nil
}
