package db

import (
	"encoding/json"
	"fmt"

	bolt "go.etcd.io/bbolt"
	"humpback/api/handle/models"
	"humpback/common/response"
	"humpback/types"
)

func ActivityUpdate(info *types.ActivityInfo, bucket string) error {
	return TransactionUpdates(func(tx *bolt.Tx) error {
		activityBucket := tx.Bucket([]byte(BucketActivities))
		if activityBucket == nil {
			return ErrBucketNotExist
		}
		childBucket := activityBucket.Bucket([]byte(bucket))
		if childBucket == nil {
			return ErrBucketNotExist
		}
		encodedData, err := json.Marshal(info)
		if err != nil {
			return fmt.Errorf("failed to encode data: %s", err)
		}
		return childBucket.Put([]byte(info.ActivityId), encodedData)
	})
}

func ActivityQuery(info *models.ActivityQueryReqInfo, bucket string) (*response.QueryResult[types.ActivityInfo], error) {
	var (
		results    = &response.QueryResult[types.ActivityInfo]{Total: 0, List: make([]*types.ActivityInfo, 0)}
		startIndex = (info.PageInfo.Index-1)*info.PageInfo.Size + 1
		endIndex   = info.PageInfo.Index*info.PageInfo.Size + 1
	)
	if err := TransactionGet(func(tx *bolt.Tx) error {
		activityBucket := tx.Bucket([]byte(BucketActivities))
		if activityBucket == nil {
			return ErrBucketNotExist
		}
		childBucket := activityBucket.Bucket([]byte(bucket))
		if childBucket == nil {
			return ErrBucketNotExist
		}
		c := childBucket.Cursor()
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
				results.Total++
				if results.Total >= startIndex && results.Total < endIndex {
					tempValue := new(types.ActivityInfo)
					if err := json.Unmarshal(v, tempValue); err != nil {
						return err
					}
					results.List = append(results.List, tempValue)
				}
			}
			k, v = c.Prev()
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return results, nil
}
