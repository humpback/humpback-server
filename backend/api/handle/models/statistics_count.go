package models

import (
    "humpback/common/locales"
    "humpback/common/response"
)

type StatisticsCountQueryReqInfo struct {
    StartAt int64 `json:"startAt"`
    EndAt   int64 `json:"endAt"`
}

func (s *StatisticsCountQueryReqInfo) Check() error {
    if s.StartAt > 0 && s.StartAt > s.EndAt {
        return response.NewBadRequestErr(locales.CodeRequestParamsInvalid)
    }
    return nil
}
