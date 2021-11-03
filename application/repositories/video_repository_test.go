package repositories_test

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shnleeti/encodervideo/application/repositories"
	"github.com/shnleeti/encodervideo/domain"
	"github.com/shnleeti/encodervideo/framework/database"
	"github.com/stretchr/testify/require"
)

func TestNewVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
