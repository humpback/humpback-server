package controller

import (
    "fmt"
    "log/slog"
    
    "humpback/api/handle/models"
    "humpback/common/response"
    "humpback/internal/db"
    "humpback/types"
)

var (
    StatisticsCh = make(chan *types.StatisticalCountInfo, 100)
)

type StatisticalCountEnvent struct {
    CreateAt int64           `json:"createAt"`
    Type     types.CountType `json:"type"`
    Num      int             `json:"num"`
    UserId   string          `json:"userId"`
}

func ReceiveStatisticsCount(stopCh <-chan struct{}) {
    defer close(StatisticsCh)
    slog.Info("[StatisticsCount] Startup wait channel.")
    for {
        select {
        case <-stopCh:
            return
        case info := <-StatisticsCh:
            if err := db.StatisticalCountInsert(info); err != nil {
                slog.Error("[StatisticsCount] Insert statistics count failed.", "Type", info.Type, "Id", info.Id, "Error", err)
            }
        }
    }
}

func InsertStatisticsCount(info *StatisticalCountEnvent) {
    StatisticsCh <- &types.StatisticalCountInfo{
        Id:       fmt.Sprintf("%d-%s-%s-%d", info.CreateAt, info.UserId, info.Type, info.Num),
        CreateAt: info.CreateAt,
        Type:     info.Type,
        Num:      info.Num,
        UserId:   info.UserId,
    }
}

func StatisticsCountQuery(info *models.StatisticsCountQueryReqInfo) ([]*types.StatisticalCountInfo, error) {
    list, err := db.StatisticsCountQuery(info)
    if err != nil {
        return nil, response.NewRespServerErr(err.Error())
    }
    return list, nil
}
