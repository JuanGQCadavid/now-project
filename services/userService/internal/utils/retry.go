package utils

import (
	"math/rand"
	"sync"
	"time"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
)

type backoffConfig struct {
	f         func() error
	onFailure func() error
	onSucess  func() error
}

type BackOffConfig func(*backoffConfig)

func WithFunction(f func() error) BackOffConfig {
	return func(cfg *backoffConfig) {
		cfg.f = f
	}
}

func WithOnFailure(onFailure func() error) BackOffConfig {
	return func(cfg *backoffConfig) {
		cfg.onFailure = onFailure
	}
}

func WithOnSucess(onSucess func() error) BackOffConfig {
	return func(cfg *backoffConfig) {
		cfg.onSucess = onSucess
	}
}

func ExponentialBackOffWithOpts(opts ...BackOffConfig) *sync.WaitGroup {
	var (
		wg         sync.WaitGroup = sync.WaitGroup{}
		attempts   int            = 5
		startTime  time.Duration  = 100 * time.Millisecond
		UpperLimit time.Duration  = 2 * time.Second
		cfg        *backoffConfig = &backoffConfig{}
	)

	for _, opt := range opts {
		opt(cfg)
	}

	wg.Add(1)

	go func() {
		defer wg.Done()
		for attempts > 0 {
			err := cfg.f()
			if err == nil {
				if cfg.onSucess != nil {
					go cfg.onSucess()
				}
				return
			}
			attempts -= 1
			toSleep := startTime.Milliseconds() + rand.Int63n(UpperLimit.Milliseconds())
			logs.Warning.Println("Sleeping ", toSleep)
		}

		if cfg.onFailure != nil {
			go cfg.onFailure()
		}
	}()

	return &wg
}

func ExponentialBackOff(f func() error, onFailure func() error, onSucess func() error) *sync.WaitGroup {
	var (
		wg         sync.WaitGroup = sync.WaitGroup{}
		attempts   int            = 5
		startTime  time.Duration  = 100 * time.Millisecond
		UpperLimit time.Duration  = 2 * time.Second
	)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for attempts > 0 {
			err := f()
			if err == nil {
				if onSucess != nil {
					go onSucess()
				}
				return
			}
			attempts -= 1
			toSleep := startTime.Milliseconds() + rand.Int63n(UpperLimit.Milliseconds())
			logs.Warning.Println("Sleeping ", toSleep)
		}

		go onFailure()
	}()

	return &wg
}
