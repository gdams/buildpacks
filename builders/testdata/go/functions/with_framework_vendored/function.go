// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package myfunc contains a test function.
package myfunc

import (
	"fmt"
	"net/http"

	// Blank-import the framework so that it doesn't get removed by go mod.
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"

	// Import from a require/replace directive in go.mod.
	"example.com/htmlreturn"
	// Import from a subdirectory.
	"example.com/myfunc/subdir"
)

// Func is a test function.
func Func(w http.ResponseWriter, r *http.Request) {
	if subdir.Pass == "PASS" && htmlreturn.GetReturn() == "PASS" {
		fmt.Fprintf(w, "PASS")
		return
	}
	fmt.Fprintf(w, "FAIL")
}
