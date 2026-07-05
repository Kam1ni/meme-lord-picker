package cache

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path"
)

func makeHash(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func createCachDir() string {
	cacheDir := path.Join(os.TempDir(), "meme-lord-picker")
	_, err := os.Stat(cacheDir)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(fmt.Sprintf("Failed to check if cache dir exists\n%s", err.Error()))
		}
		err := os.Mkdir(cacheDir, 0700)
		if err != nil {
			panic(fmt.Sprintf("Failed to create cache dir %s", cacheDir))
		}
	}
	return cacheDir
}

func (c CachingServer) getFileFromCache(url string) ([]byte, error) {
	hash := makeHash(url)
	return os.ReadFile(path.Join(c.cacheDir, hash))
}

func (c CachingServer) storeFileInCache(url string, data []byte) error {
	hash := makeHash(url)
	return os.WriteFile(path.Join(c.cacheDir, hash), data, 0500)
}
