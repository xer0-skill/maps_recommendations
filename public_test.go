package maps_recommendations_test

import (
	"testing"
	"github.com/cheekybits/is"
	"web-service/maps_recommendations"
)

type obj struct {
	value1 string
	value2 string
	value3 string
}

func (o *obj) Public() interface{} {
	return map[string]interface{}{"one": o.value1, "three": o.value3}
}

func TestPublic(t *testing.T) {
	is := is.New(t)

	o := &obj{
		value1: "value1",
		value2: "value2",
		value3: "value3",
	}

	v, ok := maps_recommendations.Public(o).(map[string]interface{})
	is.Equal(true, ok)
	is.Equal(v["one"], "value1")
	is.Nil(v["two"])
	is.Equal(v["three"], "value3")

}