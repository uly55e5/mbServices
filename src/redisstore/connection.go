package redisstore

import (
	"context"
)

func StoreConnection(connection string) error {
	if rdb == nil {
		err := Connect()
		if err != nil {
			return err
		}
	}
	isM, err := rdb.SIsMember(context.Background(), "connections", connection).Result()
	if err != nil {
		return err
	}
	if !isM {
		rdb.SAdd(context.Background(), "connections", connection)
	}
	return nil
}

func GetConnections() ([]string, error) {
	if rdb == nil {
		err := Connect()
		if err != nil {
			return nil, err
		}
	}
	js, err := rdb.SMembers(context.Background(), "connections").Result()
	if err != nil {
		return nil, err
	}

	return js, nil
}
