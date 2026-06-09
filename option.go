package notifykit

type Option func(*[]any)

// Option keys
//
// The following keys are the keys supported by the built-in options of the Sender interface
const (
	OPTION_KEY_HOST               string = "host"
	OPTION_KEY_PATH               string = "path"
	OPTION_KEY_COMPRESSED         string = "compressed"
	OPTION_KEY_HEADER             string = "header"
	OPTION_KEY_TIMEOUT            string = "timeout"
	OPTION_KEY_PROXY              string = "proxy"
	OPTION_KEY_MAX_CONCUR_UPLOADS string = "max_concurrent_uploads"
)

// Host sets the host.
//
// The host in the New() function is required,
// but it needs to be set by function options in the initialization
func Host(hostOrIp string) Option {
	return func(args *[]any) {
		nv := map[string]string{
			OPTION_KEY_HOST: hostOrIp,
		}
		*args = append(*args, nv)
	}
}

// Path specifies a path to append to the host
func Path(path string) Option {
	return func(args *[]any) {
		nv := map[string]string{
			OPTION_KEY_PATH: path,
		}
		*args = append(*args, nv)
	}
}

// Compressed enables the request to compress API response
func Compressed(compressed bool) Option {
	return func(args *[]any) {
		nv := map[string]bool{
			OPTION_KEY_COMPRESSED: compressed,
		}
		*args = append(*args, nv)
	}
}

// Header sets the header value
func Header(key, value string) Option {
	return func(args *[]any) {
		nv := map[string]map[string]string{
			OPTION_KEY_HEADER: {
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
			OPTION_KEY_TIMEOUT: timeOut,
		}
		*args = append(*args, nv)
	}
}

// Proxy sets the http connection proxy
func Proxy(proxy string) Option {
	return func(args *[]any) {
		nv := map[string]string{
			OPTION_KEY_PROXY: proxy,
		}
		*args = append(*args, nv)
	}
}

// MaxConcurrentUploads sets the limit of processing max concurrent uploads
func MaxConcurrentUploads(max int) Option {
	return func(args *[]any) {
		nv := map[string]int{
			OPTION_KEY_MAX_CONCUR_UPLOADS: max,
		}
		*args = append(*args, nv)
	}
}
