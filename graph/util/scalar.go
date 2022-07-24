package util

import (
	"time"
	"fmt"
	"io"
	"strconv"
    "encoding/json"
	
	"github.com/lib/pq"
	"github.com/99designs/gqlgen/graphql"
)


// Array string
func MarshalStringArray(a pq.StringArray) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		data, _ := json.Marshal(a)
		io.WriteString(w, string(data))
	})
}

func UnmarshalStringArray(v interface{}) (pq.StringArray, error) {
	a, ok := v.(pq.StringArray)
	if !ok {
		return nil, fmt.Errorf("failed to cast to pq.StringArray")
	}
	return a, nil
}


// Array integer
func MarshalInt32Array(a pq.Int32Array) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		data, _ := json.Marshal(a)
		io.WriteString(w, string(data))
	})
}

func UnmarshalInt32Array(v interface{}) (pq.Int32Array, error) {
	a, ok := v.(pq.Int32Array)
	if !ok {
		return nil, fmt.Errorf("failed to cast to pq.Int32Array")
	}
	return a, nil
}

// Time
func MarshalTimestamp(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(t.Unix(), 10))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int64); ok {
		return time.Unix(tmpStr, 0), nil
	}
	return time.Time{}, fmt.Errorf("time should be a unix timestamp")
}