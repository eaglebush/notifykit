package notifykit

type Option func(*[]any)

// Host sets the host.
//
// The host in the New() function is required,
// but it needs to be set by function options in the initialization
func Host(hostOrIp string) Option {
	return func(args *[]any) {
		nv := map[string]string{
			"host": hostOrIp,
		}
		*args = append(*args, nv)
	}
}

// Path specifies a path to append to the host
func Path(path string) Option {
	return func(args *[]any) {
		nv := map[string]string{
			"path": path,
		}
		*args = append(*args, nv)
	}
}

// Compressed enables the request to compress API response
func Compressed(compressed bool) Option {
	return func(args *[]any) {
		nv := map[string]bool{
			"compressed": compressed,
		}
		*args = append(*args, nv)
	}
}

// Header sets the header value
func Header(key, value string) Option {
	return func(args *[]any) {
		nv := map[string]map[string]string{
			"header": {
				key: value,
			},
		}
		*args = append(*args, nv)
	}
}

// Timeout sets the request timeout
func Timeout(timeOut int) Option {
	return func(args *[]any) {
		nv := map[string]int{
			"timeout": timeOut,
		}
		*args = append(*args, nv)
	}
}

// Proxy sets the http connection proxy
func Proxy(proxy string) Option {
	return func(args *[]any) {
		nv := map[string]string{
			"proxy": proxy,
		}
		*args = append(*args, nv)
	}
}
