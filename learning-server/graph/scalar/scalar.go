package scalar

import (
	"fmt"
	"io"

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
