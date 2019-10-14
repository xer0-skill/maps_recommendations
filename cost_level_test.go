package maps_recommendations_test

import (
	"github.com/cheekybits/is"
	"web-service/maps_recommendations"
	"testing"
)

func TestCostValues(t *testing.T) {
	is := is.New(t)

	is.Equal(int(maps_recommendations.Cost1), 1)
	is.Equal(int(maps_recommendations.Cost2), 2)
	is.Equal(int(maps_recommendations.Cost3), 3)
	is.Equal(int(maps_recommendations.Cost4), 4)
	is.Equal(int(maps_recommendations.Cost5), 5)

}

func TestCostString(t *testing.T) {
	is := is.New(t)
	is.Equal(maps_recommendations.Cost1.String(), "$")
	is.Equal(maps_recommendations.Cost2.String(), "$$")
	is.Equal(maps_recommendations.Cost3.String(), "$$$")
	is.Equal(maps_recommendations.Cost4.String(), "$$$$")
	is.Equal(maps_recommendations.Cost5.String(), "$$$$$")
}

func TestParseCost(t *testing.T) {
	is := is.New(t)
	is.Equal(maps_recommendations.Cost1, maps_recommendations.ParseCost("$"))
	is.Equal(maps_recommendations.Cost2, maps_recommendations.ParseCost("$$"))
	is.Equal(maps_recommendations.Cost3, maps_recommendations.ParseCost("$$$"))
	is.Equal(maps_recommendations.Cost4, maps_recommendations.ParseCost("$$$$"))
	is.Equal(maps_recommendations.Cost5, maps_recommendations.ParseCost("$$$$$"))
}

func TestParseCostRange(t *testing.T) {
	is := is.New(t)
	var l maps_recommendations.CostRange
	var err error
	l, err = maps_recommendations.ParseCostRange("$$...$$$")
	is.NoErr(err)
	is.Equal(l.From, maps_recommendations.Cost2)
	is.Equal(l.To, maps_recommendations.Cost3)
	l, err = maps_recommendations.ParseCostRange("$...$$$$$")
	is.NoErr(err)
	is.Equal(l.From, maps_recommendations.Cost1)
	is.Equal(l.To, maps_recommendations.Cost5)
}

func TestCostRangeString(t *testing.T) {
	is := is.New(t)
	r := maps_recommendations.CostRange{
		From: maps_recommendations.Cost2,
		To:   maps_recommendations.Cost4,
	}
	is.Equal("$$...$$$$", r.String())
}
