package test

import (
	"net/http/httptest"
	"testing"

	"github.com/liuvigongzuoshi/go-kriging-service/internal/app/schema"
	"github.com/liuvigongzuoshi/go-kriging-service/pkg/util/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDemo(t *testing.T) {
	const router = apiPrefix + "v1/demos"
	var err error

	w := httptest.NewRecorder()

	// post /demos
	addItem := &schema.Ordinary{
		Code:   uuid.MustUUID().String(),
		Name:   uuid.MustUUID().String(),
		Status: 1,
	}
	engine.ServeHTTP(w, newPostRequest(router, addItem))
	assert.Equal(t, 200, w.Code)
	var addItemRes ResID
	err = parseReader(w.Body, &addItemRes)
	assert.Nil(t, err)
}
