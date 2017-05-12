package uid_test

import (
	"testing"

	"github.com/matthewmueller/uid"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestUID(t *testing.T) {
	id, err := uuid.FromString("e11c11ae-e9e7-4805-a9a1-65c4f5c2cbe3")
	if err != nil {
		t.Fatal(err)
	}

	result := uid.Hash(id.String())
	assert.Equal(t, "2z5bltk9wx333", result)
}
