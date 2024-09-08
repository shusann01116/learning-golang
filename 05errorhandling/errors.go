package errorhandling

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Book struct{}

var ErrNotFound = errors.New("not found")

func findBook(isbn string) (*Book, error) {
	_ = isbn
	// 1. errors.New を使用するパターン
	return nil, ErrNotFound
}

func validate(length int) error {
	if length <= 0 {
		// 2. fmt.Errorf を使用したパターン
		return fmt.Errorf("length must be greather than 0, length = %d", length)
	}

	// ...
	return nil
}

// 3. 独自構造体を使用するパターン
type HTTPError struct {
	StatusCode int
	URL        string
}

func (he *HTTPError) Error() string {
	return fmt.Sprintf("http status code = %d, url = %s", he.StatusCode, he.URL)
}

func ReadContents(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// error interface として返却するときには *ポインタ* として返す
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			URL:        url,
		}
	}

	return io.ReadAll(resp.Body)
}
