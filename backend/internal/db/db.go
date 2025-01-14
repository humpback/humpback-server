package db

/*

Buckets

Users: user_id, user_name, password, email, phone, create_time, update_time, is_admin, groups
Groups: group_id, group_name, create_time, update_time, users

Registries: registry_id, registry_name, url, user_name, password, isDefault, create_time, update_time, status

Configs: config_id, config_name, create_time, update_time, values(json)

Nodes: node_id, node_name, host_ip, host_port, create_time, update_time, status, infos(json)
NodesGroups: group_id, group_name, create_time, update_time, nodes

templates: template_id, template_name, create_time, update_time, infos(json)

Services: service_id, service_name, create_time, update_time, type, status, containers, infos(json)

serviceId: {random-8}
containerName: humpback-{serviceId}-{version-5}-{random-5}


*/

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"humpback/config"

	bolt "go.etcd.io/bbolt"
)

type dbHelper struct {
	boltDB *bolt.DB
}

const BucketUsers = "Users"
const BucketGroups = "Groups"
const BucketSessions = "Sessions"
const BucketRegistries = "Registries"
const BucketConfigs = "Configs"
const BucketNodes = "Nodes"
const BucketNodesGroups = "NodesGroups"
const BucketTemplates = "Templates"
const BucketServices = "Services"

var (
	Buckets = []string{BucketUsers, BucketGroups, BucketSessions, BucketRegistries, BucketConfigs, BucketNodes, BucketNodesGroups, BucketTemplates, BucketServices}
)

var (
	ErrKeyNotExist     = errors.New("Key not found")
	ErrConnectAbnormal = errors.New("The database link is abnormal")
	ErrBucketNotExist  = errors.New("Bucket not exists")
)

var db *dbHelper

func InitDB() error {
	db = &dbHelper{}
	boltDB, err := bolt.Open(config.DBArgs().Root, 0600, nil)
	if err != nil {
		return fmt.Errorf("Open db failed: %s", err)
	}
	db.boltDB = boltDB
	if err = db.boltDB.Batch(func(tx *bolt.Tx) error {
		for _, bucket := range Buckets {
			if _, err := tx.CreateBucketIfNotExists([]byte(bucket)); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("Init Buckets failed: %s", err)
	}
	return nil
}

func CloseDB() {
	db.boltDB.Close()
}

func GetDataById[T any](bucketName string, id string) (*T, error) {
	var result T
	err := db.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrBucketNotExist
		}
		data := bucket.Get([]byte(id))
		if data == nil {
			return ErrKeyNotExist
		}
		return json.Unmarshal(data, &result)
	})
	if err != nil {
		return nil, checkErr(err)
	}
	return &result, nil
}

func GetDataByIds[T any](bucketName string, ids []string) ([]*T, error) {
	var results []*T
	err := db.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrBucketNotExist
		}
		for _, id := range ids {
			data := bucket.Get([]byte(id))
			if data == nil {
				return ErrKeyNotExist
			}
			result := new(T)
			if err := json.Unmarshal(data, result); err != nil {
				return err
			}
			results = append(results, result)
		}
		return nil
	})
	if err != nil {
		return nil, checkErr(err)
	}
	return results, nil
}

func GetDataAll[T any](bucketName string) ([]*T, error) {
	var results []*T
	err := db.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrBucketNotExist
		}
		bucket.ForEach(func(k, v []byte) error {
			result := new(T)
			if err := json.Unmarshal(v, &result); err != nil {
				return err
			}
			results = append(results, result)
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, checkErr(err)
	}
	return results, nil
}

type CustomFilterData func(key string, value interface{}) bool

func GetDataByQuery[T any](bucketName string, filter CustomFilterData) ([]*T, error) {
	var results []*T
	err := db.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrBucketNotExist
		}
		bucket.ForEach(func(k, v []byte) error {
			result := new(T)
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
		return nil, checkErr(err)
	}
	return results, nil
}

func GetDataByPrefix[T any](bucketName string, prefix string) ([]*T, error) {
	var results []*T
	err := db.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrBucketNotExist
		}
		c := bucket.Cursor()
		bp := []byte(prefix)
		for k, v := c.Seek([]byte(prefix)); k != nil && bytes.HasPrefix(k, bp); k, v = c.Next() {
			result := new(T)
			if err := json.Unmarshal(v, &result); err != nil {
				return err
			}
			results = append(results, result)
		}
		return nil
	})
	if err != nil {
		return nil, checkErr(err)
	}
	return results, nil
}

func SaveData[T any](bucketName string, id string, data T) error {
	if err := db.boltDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrBucketNotExist
		}
		encodedData, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("failed to encode data: %s", err)
		}
		return bucket.Put([]byte(id), encodedData)
	}); err != nil {
		return checkErr(err)
	}
	return nil
}

func DeleteData(bucketName string, id string) error {
	if err := db.boltDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrBucketNotExist
		}
		return bucket.Delete([]byte(id))
	}); err != nil {
		return checkErr(err)
	}
	return nil
}

func BatchDelete(bucketName string, ids []string) error {
	if err := db.boltDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrBucketNotExist
		}
		for _, id := range ids {
			if err := bucket.Delete([]byte(id)); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return checkErr(err)
	}
	return nil

}

func checkErr(err error) error {
	switch err {
	case bolt.ErrBucketNotFound:
		return ErrBucketNotExist
	case bolt.ErrInvalid,
		bolt.ErrDatabaseNotOpen,
		bolt.ErrDatabaseOpen,
		bolt.ErrInvalidMapping:
		return ErrConnectAbnormal
	default:
		return err
	}
}
