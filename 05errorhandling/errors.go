package errorhandling

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

type Book struct{}

var ErrNotFound = errors.New("not found")

func FindBook(isbn string) (*Book, error) {
	_ = isbn
	// 1. errors.New を使用するパターン
	return nil, ErrNotFound
}

func Validate(length int) error {
	if length <= 0 {
		// 2. fmt.Errorf を使用したパターン
		return fmt.Errorf("length must be greather than 0, length = %d", length) //nolint: err113
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

func ReadContents(rawURL string) ([]byte, error) {
	host, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return nil, err
	}

	ctx, f := context.WithDeadline(context.TODO(), time.Now().Add(5*time.Second))
	defer f()

	reader := bytes.NewReader(make([]byte, 0))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, host.String(), reader)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create context")
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// error interface として返却するときには *ポインタ* として返す
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			URL:        host.String(),
		}
	}

	return io.ReadAll(resp.Body)
}
