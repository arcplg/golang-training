package scalar

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ObjectID = bson.ObjectID

func MarshalObjectID(id ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		fmt.Fprintf(w, "\"%s\"", id.Hex())
	})
}

func UnmarshalObjectID(v interface{}) (ObjectID, error) {
	str, ok := v.(string)
	if !ok {
		return bson.NilObjectID, fmt.Errorf("ObjectID must be a string")
	}
	return bson.ObjectIDFromHex(str)
}
func MarshalObjectIDScalar(id bson.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		fmt.Fprintf(w, "\"%s\"", id.Hex())
	})
}

const DateTime = time.RFC3339

// MarshalDateTime converts a Go time.Time to a string for GraphQL.
func MarshalDateTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		fmt.Fprintf(w, "\"%s\"", t.Format(DateTime))
	})
}

// UnmarshalDateTime converts a GraphQL string to Go time.Time.
func UnmarshalDateTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return time.Parse(DateTime, v)
	case []byte:
		return time.Parse(DateTime, string(v))
	default:
		return time.Time{}, fmt.Errorf("invalid type %T for DateTime", v)
	}
}

func MarshalDateTimeScalar(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		fmt.Fprintf(w, "\"%s\"", t.Format(DateTime))
	})
}
