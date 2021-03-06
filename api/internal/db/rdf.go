package db

import "github.com/dgraph-io/dgo/v200/protos/api"

func nquadStr(subject, predicate, objectValue string) *api.NQuad {
	return &api.NQuad{
		Subject:   subject,
		Predicate: predicate,
		ObjectValue: &api.Value{
			Val: &api.Value_StrVal{StrVal: objectValue},
		},
	}
}

func nquadAll(subject string) *api.NQuad {
	return &api.NQuad{
		Subject:   subject,
		Predicate: "_STAR_ALL",
		ObjectValue: &api.Value{
			Val: &api.Value_DefaultVal{DefaultVal: "_STAR_ALL"},
		},
	}
}

func nquadBool(subject, predicate string, objectValue bool) *api.NQuad {
	return &api.NQuad{
		Subject:   subject,
		Predicate: predicate,
		ObjectValue: &api.Value{
			Val: &api.Value_BoolVal{BoolVal: objectValue},
		},
	}
}

func nquadRel(subject, predicate, objectID string) *api.NQuad {
	return &api.NQuad{
		Subject:   subject,
		Predicate: predicate,
		ObjectId:  objectID,
	}
}
