package port

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParsePortsFromJSONMap(suite *testing.T) {
	suite.Parallel()

	suite.Run("should return empty iterator for empty input", func(t *testing.T) {
		t.Parallel()

		portIterator, err := ParsePortsFromJSONMap(bytes.NewReader([]byte{}))

		require.NoError(t, err)
		require.NotNil(t, portIterator)

		port, ok, err := portIterator.Next()
		require.Nil(t, port)
		require.False(t, ok)
		require.NoError(t, err)
	})

	suite.Run("should return error for wrong JSON begin", func(t *testing.T) {
		t.Parallel()

		portIterator, err := ParsePortsFromJSONMap(bytes.NewReader([]byte("*")))

		require.ErrorIs(t, err, ErrNotJSONObject)
		require.Nil(t, portIterator)
	})

	suite.Run("should return empty iterator for empty JSON object", func(t *testing.T) {
		t.Parallel()

		portIterator, err := ParsePortsFromJSONMap(bytes.NewReader([]byte("{}")))

		require.NoError(t, err)
		require.NotNil(t, portIterator)

		port, ok, err := portIterator.Next()
		require.Nil(t, port)
		require.False(t, ok)
		require.NoError(t, err)
	})

	suite.Run("should return iterator error for wrong port code type", func(t *testing.T) {
		t.Parallel()

		testMap := `{ true : [] }`
		portIterator, err := ParsePortsFromJSONMap(bytes.NewReader([]byte(testMap)))

		require.NoError(t, err)
		require.NotNil(t, portIterator)

		port, ok, err := portIterator.Next()
		require.Nil(t, port)
		require.False(t, ok)
		require.ErrorIs(t, err, ErrPortCodeType)
	})

	suite.Run("should return iterator error for wrong map value", func(t *testing.T) {
		t.Parallel()

		testMap := `{ "ABCD": [] }`
		portIterator, err := ParsePortsFromJSONMap(bytes.NewReader([]byte(testMap)))

		require.NoError(t, err)
		require.NotNil(t, portIterator)

		port, ok, err := portIterator.Next()
		require.Nil(t, port)
		require.False(t, ok)
		require.Error(t, err)
	})

	suite.Run("should return port from JSON with one port", func(t *testing.T) {
		t.Parallel()

		testMap := `
			{
			   "AEAJM":{
				  "name":"Ajman",
				  "city":"Ajman",
				  "country":"United Arab Emirates",
				  "alias":[
					 
				  ],
				  "regions":[
					 
				  ],
				  "coordinates":[
					 55.5136433,
					 25.4052165
				  ],
				  "province":"Ajman",
				  "timezone":"Asia/Dubai",
				  "unlocs":[
					 "AEAJM"
				  ],
				  "code":"52000"
			   }
			}`
		expectedPort := &Port{
			Code: "AEAJM",
			Details: Details{
				Name:    "Ajman",
				City:    "Ajman",
				Country: "United Arab Emirates",
				Alias:   []interface{}{},
				Regions: []interface{}{},
				Coordinates: []float64{
					55.5136433,
					25.4052165,
				},
				Province: "Ajman",
				Timezone: "Asia/Dubai",
				Unlocs: []string{
					"AEAJM",
				},
				Code: "52000",
			},
		}

		portIterator, err := ParsePortsFromJSONMap(bytes.NewReader([]byte(testMap)))

		require.NoError(t, err)
		require.NotNil(t, portIterator)

		port, ok, err := portIterator.Next()
		require.NotNil(t, port)
		require.True(t, ok)
		require.NoError(t, err)
		require.Equal(t, expectedPort, port)

		port, ok, err = portIterator.Next()
		require.Nil(t, port)
		require.False(t, ok)
		require.NoError(t, err)
	})

	suite.Run("should return all ports from JSON with multiple ports with the same code", func(t *testing.T) {
		t.Parallel()

		testMap := `
			{
			   "AEAJM":{
				  "name":"Ajman",
				  "city":"Ajman",
				  "country":"United Arab Emirates",
				  "alias":[
					 
				  ],
				  "regions":[
					 
				  ],
				  "coordinates":[
					 55.5136433,
					 25.4052165
				  ],
				  "province":"Ajman",
				  "timezone":"Asia/Dubai",
				  "unlocs":[
					 "AEAJM"
				  ],
				  "code":"52000"
			   },
			   "AEAUH":{
				  "name":"Abu Dhabi",
				  "coordinates":[
					 54.37,
					 24.47
				  ],
				  "city":"Abu Dhabi",
				  "province":"Abu Z¸aby [Abu Dhabi]",
				  "country":"United Arab Emirates",
				  "alias":[
					 
				  ],
				  "regions":[
					 
				  ],
				  "timezone":"Asia/Dubai",
				  "unlocs":[
					 "AEAUH"
				  ],
				  "code":"52001"
			   },
			   "AEAJM":{
				  "name":"New City",
				  "city":"New City",
				  "country":"United Arab Emirates",
				  "alias":[
					 
				  ],
				  "regions":[
					 
				  ],
				  "coordinates":[
					 55.5136433,
					 25.4052165
				  ],
				  "province":"Ajman",
				  "timezone":"Asia/Dubai",
				  "unlocs":[
					 "AEAJM"
				  ],
				  "code":"52000"
			   }
			}`

		portIterator, err := ParsePortsFromJSONMap(bytes.NewReader([]byte(testMap)))
		require.NoError(t, err)
		require.NotNil(t, portIterator)

		port1, ok, err := portIterator.Next()
		require.NotNil(t, port1)
		require.True(t, ok)
		require.NoError(t, err)

		port2, ok, err := portIterator.Next()
		require.NotNil(t, port2)
		require.True(t, ok)
		require.NoError(t, err)

		port3, ok, err := portIterator.Next()
		require.NotNil(t, port3)
		require.True(t, ok)
		require.NoError(t, err)

		port4, ok, err := portIterator.Next()
		require.Nil(t, port4)
		require.False(t, ok)
		require.NoError(t, err)

		require.Equal(t, port1.Code, port3.Code)
		require.NotEqual(t, port1.City, port3.City)
	})
}
