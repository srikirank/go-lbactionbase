package common

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Local simple file-based cache that supports GraceTime for
// stale data of a maximum of 3 days.

// GraceTime Worst case including long-weekend, longer than that
// we probably want to refresh
const (
	GraceTime        = 3 * 24 * time.Hour
	FallbackCacheDir = "/tmp"
)

func CacheFetch(key string, ttlDuration time.Duration) ([]byte, error) {
	loc := cacheLoc(key)
	stat, err := os.Stat(loc)
	if err != nil {
		return nil, &CacheFetchError{}
	}
	mt := stat.ModTime()
	now := time.Now()

	// To force agents to update the cache during gracetime
	// while letting normal clients to access stale data during gracetime
	graceTimeAllowed := true
	env := os.Environ()
	for _, e := range env {
		if e == "LB_LOCAL_CACHE_GRACE_TIME_ALLOWED=false" {
			graceTimeAllowed = false
		}
	}

	log.Printf("GraceTime allowed for fetching cache-key: %s? %t", key, graceTimeAllowed)
	timeToCheck := ttlDuration
	if graceTimeAllowed {
		timeToCheck += GraceTime
	}

	if mt.Add(timeToCheck).Before(now) {
		return nil, &CacheExpiredError{}
	}

	data, err := ioutil.ReadFile(loc)
	if err != nil {
		return nil, &CacheFetchError{}
	}
	return data, nil
}

func CacheUpdate(key string, data []byte) error {
	loc := cacheLoc(key)
	err := ioutil.WriteFile(loc, data, 0644)
	if err != nil {
		return &CacheUpdateError{}
	}
	return nil
}

func ShouldAsyncUpdateBeforeExpiry(key string, ttlDuration time.Duration) bool {
	t, e := LastModTimeForCache(key)
	expT := t.Add(ttlDuration)
	if e != nil {
		return true
	}
	now := time.Now()
	nowPlus3Days := now.Add(3 * 24 * time.Hour)
	return expT.Before(nowPlus3Days)
}

func LastModTimeForCache(key string) (time.Time, error) {
	loc := cacheLoc(key)
	stat, err := os.Stat(loc)
	if err != nil {
		return time.Now(), &CacheFetchError{}
	}
	return stat.ModTime(), nil
}

func PwdCacheKeyWithSuffix(suffix string) (string, error) {
	b64f := base64.StdEncoding.EncodeToString([]byte(suffix))
	log.Printf("Cache key used: %s\n", b64f)
	return b64f, nil
}

func cacheLoc(key string) string {
	dir := cacheDir()
	return fmt.Sprintf("%s/%s", dir, key)
}

// ClearCacheFilesWithPrefix Removes all files starting with prefix, if empty, removes all files
func ClearCacheFilesWithPrefix(p string) {
	d := cacheDir()
	files, _ := ioutil.ReadDir(d)
	for _, file := range files {
		if !file.IsDir() {
			if strings.Contains(file.Name(), p) || len(p) == 0 {
				_ = os.Remove(filepath.Join(d, file.Name()))
			}
		}
	}
}

func cacheDir() string {
	dir := LBCustomActionsCacheDir()
	env := os.Environ()
	for _, e := range env {
		if strings.Contains(e, "LB_LOCAL_CACHE_OVERRIDE_DIR") {
			dir = strings.TrimPrefix(e, "LB_LOCAL_CACHE_OVERRIDE_DIR=")
			break
		}
	}
	_, err := os.Stat(dir)
	if err != nil {
		log.Printf("Cache Dir doesn't exist: %s. Attempting to create", dir)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Printf("Failed to create this dir structure: %s. Falling back to default: %s", dir, FallbackCacheDir)
			dir = FallbackCacheDir
		}
	}
	return dir
}
