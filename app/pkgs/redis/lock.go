package redis

import (
	"baozun.com/leak/app/pkgs/logger"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

// 分布式锁
type Lock struct {
	resource string
	token    string
	timeout  int
}

func (lock *Lock) key() string {
	return fmt.Sprintf("riskcontrol:redislock:%s", lock.resource)
}

func (lock *Lock) tryLock() (ok bool, err error) {
	_, err = redis.String(poll.Get().Do("SET", lock.key(), lock.token, "EX", int(lock.timeout), "NX"))
	if err == redis.ErrNil {
		// The lock was not successful, it already exists.
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (lock *Lock) UnlockDeferDefault() (err error) {
	time.Sleep(time.Duration(DefaultTimeout) * time.Second)
	_, err = poll.Get().Do("del", lock.key())
	return
}

func (lock *Lock) Unlock() (err error) {
	_, err = poll.Get().Do("del", lock.key())
	return
}

func (lock *Lock) AddTimeout(exTime int64) (ok bool, err error) {
	ttlTime, err := redis.Int64(poll.Get().Do("TTL", lock.key()))
	if err != nil {
		logger.Logger.Error("redis get failed:", err)
	}
	if ttlTime > 0 {
		_, err := redis.String(poll.Get().Do("SET", lock.key(), lock.token, "EX", int(ttlTime+exTime)))
		if err == redis.ErrNil {
			return false, nil
		}
		if err != nil {
			return false, err
		}
	}
	return false, nil
}

func TryLock(resource string, token string) (lock *Lock, ok bool, err error) {
	return TryLockWithTimeout(resource, token, DefaultTimeout)
}

func TryLockWithTimeout(resource string, token string, timeout int) (lock *Lock, ok bool, err error) {
	lock = &Lock{resource, token, timeout}

	ok, err = lock.tryLock()

	if !ok || err != nil {
		lock = nil
	}

	return
}
