package db

import (
	"encoding/json"
	"fmt"
	"strings"

	bolt "go.etcd.io/bbolt"
	"humpback/types"
)

func UserGetSupperAdmin() (*types.User, error) {
	users, err := UsersGetAll()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if types.IsSupperAdmin(user.Role) {
			return user, nil
		}
	}
	return nil, nil
}

func UsersGetAll() ([]*types.User, error) {
	return GetDataAll[types.User](BucketUsers)
}

func UserGetById(id string) (*types.User, error) {
	return GetDataById[types.User](BucketUsers, id)
}

func UsersGetByName(name string, isLower bool) ([]*types.User, error) {
	users, err := GetDataAll[types.User](BucketUsers)
	if err != nil {
		return nil, err
	}
	var result []*types.User
	for _, user := range users {
		if isLower && strings.ToLower(user.Username) == strings.ToLower(name) {
			result = append(result, user)
		}
		if !isLower && user.Username == name {
			result = append(result, user)
		}
	}
	return result, nil
}

func UsersGetByIds(ids []string, ignoreNotExist bool) ([]*types.User, error) {
	return GetDataByIds[types.User](BucketUsers, ids, ignoreNotExist)
}

func UserUpdate(id string, data *types.User) error {
	return SaveData[*types.User](BucketUsers, id, data)
}

func UserUpdateAndTeams(userInfo *types.User, teams []*types.Team) (string, error) {
	if err := TransactionUpdates(func(tx *bolt.Tx) error {
		var (
			teamBucket *bolt.Bucket
			userBucket *bolt.Bucket
		)
		userBucket = tx.Bucket([]byte(BucketUsers))
		if userBucket == nil {
			return ErrBucketNotExist
		}
		userData, err := json.Marshal(userInfo)
		if err != nil {
			return fmt.Errorf("failed to encode user data: %s", err)
		}
		if err = userBucket.Put([]byte(userInfo.UserId), userData); err != nil {
			return err
		}
		if len(teams) > 0 {
			teamBucket = tx.Bucket([]byte(BucketTeams))
			if teamBucket == nil {
				return ErrBucketNotExist
			}
			for _, team := range teams {
				teamData, err := json.Marshal(team)
				if err != nil {
					return fmt.Errorf("failed to encode team data: %s", err)
				}
				if err = teamBucket.Put([]byte(team.TeamId), teamData); err != nil {
					return err
				}
			}
		}
		return nil
	}); err != nil {
		return "", err
	}
	return userInfo.UserId, nil
}

func UserDelete(id string, teams []*types.Team) error {
	return TransactionUpdates(func(tx *bolt.Tx) error {
		var (
			teamBucket *bolt.Bucket
			userBucket *bolt.Bucket
		)
		userBucket = tx.Bucket([]byte(BucketUsers))
		if userBucket == nil {
			return ErrBucketNotExist
		}
		if err := userBucket.Delete([]byte(id)); err != nil {
			return err
		}
		if len(teams) > 0 {
			teamBucket = tx.Bucket([]byte(BucketTeams))
			if teamBucket == nil {
				return ErrBucketNotExist
			}
			for _, team := range teams {
				teamData, err := json.Marshal(team)
				if err != nil {
					return fmt.Errorf("failed to encode team data: %s", err)
				}
				if err = teamBucket.Put([]byte(team.TeamId), teamData); err != nil {
					return err
				}
			}
		}
		return nil
	})
}
