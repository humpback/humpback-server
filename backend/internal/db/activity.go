package db

import (
    "encoding/json"
    "fmt"
    
    bolt "go.etcd.io/bbolt"
    "humpback/types"
)

func ActivityUpdate(info *types.ActivityInfo, bucket string) error {
    return TransactionUpdates(func(tx *bolt.Tx) error {
        activityBucket := tx.Bucket([]byte(BucketActivities))
        if activityBucket == nil {
            return ErrBucketNotExist
        }
        childBucekt := activityBucket.Bucket([]byte(bucket))
        if childBucekt == nil {
            return ErrBucketNotExist
        }
        encodedData, err := json.Marshal(info)
        if err != nil {
            return fmt.Errorf("failed to encode data: %s", err)
        }
        return childBucekt.Put([]byte(info.ActivityId), encodedData)
    })
}

func ActivyQuery() {
    
}
