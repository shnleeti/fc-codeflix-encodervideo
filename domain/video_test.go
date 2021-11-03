package domain_test

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shnleeti/encodervideo/domain"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)

}

func TestValidateIdIsnotUUID(t *testing.T) {

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)

}
