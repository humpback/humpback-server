package db

import (
	"encoding/json"
	"fmt"

	bolt "go.etcd.io/bbolt"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/types"
)

func TeamGetAll() ([]*types.Team, error) {
	return GetDataAll[types.Team](BucketTeams)
}

func TeamsQueryByIds(ids []string, ignoreNotExist bool) ([]*types.Team, error) {
	teams, err := GetDataByIds[types.Team](BucketTeams, ids, ignoreNotExist)
	if err != nil {
		if err == ErrKeyNotExist {
			return nil, response.NewBadRequestErr(locales.CodeTeamNotExist)
		}
		return nil, response.NewRespServerErr(err.Error())
	}
	return teams, nil
}

func TeamGetById(id string) (*types.Team, error) {
	info, err := GetDataById[types.Team](BucketTeams, id)
	if err != nil {
		if err == ErrKeyNotExist {
			return nil, response.NewBadRequestErr(locales.CodeTeamNotExist)
		}
		return nil, response.NewRespServerErr(err.Error())
	}
	return info, nil
}

func TeamGetByName(name string) (*types.Team, error) {
	teams, err := GetDataAll[types.Team](BucketTeams)
	if err != nil {
		return nil, response.NewRespServerErr(err.Error())
	}
	for _, team := range teams {
		if team.Name == name {
			return team, nil
		}
	}
	return nil, nil
}

func TeamUpdateAndUsers(teamInfo *types.Team, users []*types.User) (string, error) {
	if err := TransactionUpdates(func(tx *bolt.Tx) error {
		var (
			teamBucket *bolt.Bucket
			userBucket *bolt.Bucket
		)
		teamBucket = tx.Bucket([]byte(BucketTeams))
		if teamBucket == nil {
			return ErrBucketNotExist
		}
		teamData, err := json.Marshal(teamInfo)
		if err != nil {
			return fmt.Errorf("failed to encode team data: %s", err)
		}
		if err = teamBucket.Put([]byte(teamInfo.TeamId), teamData); err != nil {
			return err
		}
		if len(users) > 0 {
			userBucket = tx.Bucket([]byte(BucketUsers))
			if userBucket == nil {
				return ErrBucketNotExist
			}
			for _, user := range users {
				userData, err := json.Marshal(user)
				if err != nil {
					return fmt.Errorf("failed to encode user data: %s", err)
				}
				if err = userBucket.Put([]byte(user.UserId), userData); err != nil {
					return err
				}
			}
		}
		return nil
	}); err != nil {
		return "", response.NewRespServerErr(err.Error())
	}
	return teamInfo.TeamId, nil
}

func TeamDelete(id string, users []*types.User) error {
	if err := TransactionUpdates(func(tx *bolt.Tx) error {
		var (
			teamBucket *bolt.Bucket
			userBucket *bolt.Bucket
		)
		teamBucket = tx.Bucket([]byte(BucketTeams))
		if teamBucket == nil {
			return ErrBucketNotExist
		}

		if err := teamBucket.Delete([]byte(id)); err != nil {
			return err
		}

		if len(users) > 0 {
			userBucket = tx.Bucket([]byte(BucketUsers))
			if userBucket == nil {
				return ErrBucketNotExist
			}
			for _, user := range users {
				userData, err := json.Marshal(user)
				if err != nil {
					return fmt.Errorf("failed to encode user data: %s", err)
				}
				if err = userBucket.Put([]byte(user.UserId), userData); err != nil {
					return err
				}
			}
		}
		return nil
	}); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	return nil
}
