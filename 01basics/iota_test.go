package basics

import (
	"testing"
)

func TestCaroption(t *testing.T) {
	t.Parallel()

	var o CarOption //nolint:gosimple
	o = SunRoof | HeatedSeat

	if o&SunRoof == 0 {
		t.Errorf("except o&SunRoof == 0")
	}
}
