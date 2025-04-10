package db

import (
    "encoding/json"
    "strconv"
    "strings"
    
    bolt "go.etcd.io/bbolt"
    "humpback/api/handle/models"
    "humpback/types"
)

func StatisticalCountInsert(info *types.StatisticalCountInfo) error {
    return SaveData(BucketStatistics, info.Id, info)
}

func StatisticsCountQuery(info *models.StatisticsCountQueryReqInfo) ([]*types.StatisticalCountInfo, error) {
    var results = make([]*types.StatisticalCountInfo, 0)
    if err := TransactionGet(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte(BucketStatistics))
        if bucket == nil {
            return ErrBucketNotExist
        }
        c := bucket.Cursor()
        k, v := c.Last()
        for {
            if k == nil {
                break
            }
            valid, err := checkDataIsValid(info, string(k))
            if err != nil {
                return err
            }
            if valid {
                tempValue := new(types.StatisticalCountInfo)
                if err := json.Unmarshal(v, tempValue); err != nil {
                    return err
                }
                results = append(results, tempValue)
            }
            k, v = c.Prev()
        }
        return nil
    }); err != nil {
        return nil, err
    }
    return results, nil
}

func checkDataIsValid(info *models.StatisticsCountQueryReqInfo, key string) (bool, error) {
    keys := strings.Split(key, "-")
    timestamp, err := strconv.ParseInt(keys[0], 10, 64)
    if err != nil {
        return false, err
    }
    if timestamp >= info.StartAt && (info.EndAt == 0 || timestamp < info.EndAt) {
        return true, nil
    }
    return false, nil
}
