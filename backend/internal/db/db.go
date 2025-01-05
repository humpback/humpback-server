package db

/*

Buckets

Users: user_id, user_name, password, email, phone, create_time, update_time, is_admin, groups
Groups: group_id, group_name, create_time, update_time, users

Registries: registry_id, registry_name, url, user_name, password, isDefault, create_time, update_time, status

Configs: config_id, config_name, create_time, update_time, values(json)

Nodes: node_id, node_name, host_ip, host_port, create_time, update_time, status, groups, infos(json)
NodesGroups: group_id, group_name, create_time, update_time, nodes

templates: template_id, template_name, create_time, update_time, infos(json)

Services: service_id, service_name, create_time, update_time, type, status, containers, infos(json)

serviceId: {groupId}-{random-10}
containerId: humback-{serviceId}-{random-5}


*/

import (
	"bytes"
	"encoding/json"
	"fmt"

	bolt "go.etcd.io/bbolt"
)

type dbHelper struct {
	boltDB *bolt.DB
}

const BucketUsers = "Users"
const BucketGroups = "Groups"
const BucketRegistries = "Registries"
const BucketConfigs = "Configs"
const BucketNodes = "Nodes"
const BucketNodesGroups = "NodesGroups"
const BucketTemplates = "Templates"
const BucketServices = "Services"

var db *dbHelper

func InitDB() {
	db = &dbHelper{}
	boltDB, err := bolt.Open("humpback.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	db.boltDB = boltDB
}

func CloseDB() {
	db.boltDB.Close()
}

func GetDataById[T any](bucketName string, id string) (*T, error) {
	var result T
	err := db.boltDB.View(func(tx *bolt.Tx) error {
		bucket, bkErr := tx.CreateBucketIfNotExists([]byte(bucketName))
		if bkErr != nil {
			return fmt.Errorf("create bucket %s failed: %s", bucketName, bkErr)
		}
		data := bucket.Get([]byte(id))
		if data == nil {
			return fmt.Errorf("data with id %s not found in bucket %s", id, bucketName)
		}
		return json.Unmarshal(data, &result)
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetDataByIds[T any](bucketName string, ids []string) ([]T, error) {
	var results []T
	err := db.boltDB.View(func(tx *bolt.Tx) error {
		bucket, bkErr := tx.CreateBucketIfNotExists([]byte(bucketName))
		if bkErr != nil {
			return fmt.Errorf("create bucket %s failed: %s", bucketName, bkErr)
		}
		for _, id := range ids {
			data := bucket.Get([]byte(id))
			if data == nil {
				return fmt.Errorf("data with id %s not found in bucket %s", id, bucketName)
			}
			var result T
			if err := json.Unmarshal(data, &result); err != nil {
				return err
			}
			results = append(results, result)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}

func GetDataAll[T any](bucketName string) ([]T, error) {
	var results []T
	err := db.boltDB.View(func(tx *bolt.Tx) error {
		bucket, bkErr := tx.CreateBucketIfNotExists([]byte(bucketName))
		if bkErr != nil {
			return fmt.Errorf("create bucket %s failed: %s", bucketName, bkErr)
		}
		bucket.ForEach(func(k, v []byte) error {
			var result T
			if err := json.Unmarshal(v, &result); err != nil {
				return err
			}
			results = append(results, result)
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}

type CustomFilterData func(key string, value interface{}) bool

func GetDatabyQuery[T any](bucketName string, filter CustomFilterData) ([]T, error) {
	var results []T
	err := db.boltDB.View(func(tx *bolt.Tx) error {
		bucket, bkErr := tx.CreateBucketIfNotExists([]byte(bucketName))
		if bkErr != nil {
			return fmt.Errorf("create bucket %s failed: %s", bucketName, bkErr)
		}
		bucket.ForEach(func(k, v []byte) error {
			var result T
			if err := json.Unmarshal(v, &result); err != nil {
				return err
			}
			if filter(string(k), result) {
				results = append(results, result)
			}
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}

func GetDataByPrefix[T any](bucketName string, prefix string) ([]T, error) {
	var results []T
	err := db.boltDB.View(func(tx *bolt.Tx) error {
		bucket, bkErr := tx.CreateBucketIfNotExists([]byte(bucketName))
		if bkErr != nil {
			return fmt.Errorf("create bucket %s failed: %s", bucketName, bkErr)
		}
		c := bucket.Cursor()
		bp := []byte(prefix)
		for k, v := c.Seek([]byte(prefix)); k != nil && bytes.HasPrefix(k, bp); k, v = c.Next() {
			var result T
			if err := json.Unmarshal(v, &result); err != nil {
				return err
			}
			results = append(results, result)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}

func SaveData[T any](bucketName string, id string, data T) error {
	return db.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket %s failed: %s", bucketName, err)
		}
		encodedData, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("failed to encode data: %s", err)
		}
		return bucket.Put([]byte(id), encodedData)
	})
}
