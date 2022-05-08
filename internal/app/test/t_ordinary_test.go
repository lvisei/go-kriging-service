package test

import (
	"net/http/httptest"
	"testing"

	"github.com/lvisei/go-kriging-service/internal/app/schema"
	"github.com/stretchr/testify/assert"
)

func TestOrdinary(t *testing.T) {
	const router = apiPrefix + "v1/ordinary"
	var err error

	w := httptest.NewRecorder()

	// post /ordinary/grid
	grid := &schema.OrdinaryQueryGridParam{
		TrainParam: schema.OrdinaryTrainParam{
			Values: []float64{45.986076009952846, 46.223032113384235, 52.821454425024626, 89.19253247046487, 31.062802427638776},
			Lons:   []float64{31.995986076009952, 31.99622303211338, 32.002821454425025, 32.03919253247046, 31.981062802427637},
			Lats:   []float64{117.99598607600996, 117.99622303211338, 118.00282145442502, 118.03919253247047, 117.98106280242764},
			Sigma2: 0, Alpha: 100, Model: 1,
		},
		GridParam: schema.OrdinaryGridParam{
			Polygon: "{\"type\": \"Polygon\",\"coordinates\": [[[103.614373, 27.00541],[104.174357, 26.635252],[104.356163, 28.018448],[103.614373, 27.00541]]]}",
			Width:   0.01,
		},
	}
	engine.ServeHTTP(w, newPostRequest("%s/%s", grid, router, "grid"))
	assert.Equal(t, 200, w.Code)
	var gridRes schema.OrdinaryGridInfo
	err = parseReader(w.Body, &gridRes)
	assert.Nil(t, err)
	assert.Equal(t, 0.01, gridRes.GridMatrices.Width)

	// post /ordinary/grid-png
	gridPng := &schema.OrdinaryQueryGridPngParam{
		TrainParam: schema.OrdinaryTrainParam{
			Values: []float64{45.986076009952846, 46.223032113384235, 52.821454425024626, 89.19253247046487, 31.062802427638776},
			Lons:   []float64{31.995986076009952, 31.99622303211338, 32.002821454425025, 32.03919253247046, 31.981062802427637},
			Lats:   []float64{117.99598607600996, 117.99622303211338, 118.00282145442502, 118.03919253247047, 117.98106280242764},
			Sigma2: 0, Alpha: 100, Model: 1,
		},
		GridParam: schema.OrdinaryGridParam{
			Polygon: "{\"type\": \"Polygon\",\"coordinates\": [[[103.614373, 27.00541],[104.174357, 26.635252],[104.356163, 28.018448],[103.614373, 27.00541]]]}",
			Width:   0.01,
		},
		PlotPngParam: schema.OrdinaryPlotPngParam{
			Width: 100, Height: 100,
			Ylim: [2]float64{26.635252, 28.018448}, Xlim: [2]float64{103.614373, 104.356163},
			Colors: []schema.GridLevelColor{schema.GridLevelColor{Value: [2]float64{0, 15}, Color: [4]uint8{255, 128, 169, 255}}},
		},
	}
	engine.ServeHTTP(w, newPostRequest("%s/%s", gridPng, router, "grid-png"))
	assert.Equal(t, 200, w.Code)
}
