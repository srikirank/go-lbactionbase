package common

type CacheUpdateError struct{}

func (e *CacheUpdateError) Error() string {
	return "Could not update local cache"
}

type CacheFetchError struct{}

func (e *CacheFetchError) Error() string {
	return "Could not fetch from local cache"
}

type CacheExpiredError struct{}

func (e *CacheExpiredError) Error() string {
	return "Local Cache expired"
}
