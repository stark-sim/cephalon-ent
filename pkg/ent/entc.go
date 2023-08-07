//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	// 整合 versioned migrations
	if err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	}); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
