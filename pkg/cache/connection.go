package cache

import (
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func GetValkeyConnection() (*redis.Client, error) {
	address := fmt.Sprintf("%v:%v", viper.GetString("redis.host"), viper.GetString("redis.port"))

	valkeyClient := redis.NewClient(&redis.Options{
		Addr: address,
	})
	
	if valkeyClient == nil {
		return nil, errors.New("failed connect to valkey")
	}

	return valkeyClient, nil
}