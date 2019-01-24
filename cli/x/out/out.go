//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package out

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidwall/pretty"
	"gopkg.in/yaml.v2"
)

type OutFlags struct {
	Format string
}

func Failf(msg string, args ...interface{}) {
	Exitf(1, msg, args...)
}

func Warnf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}

func Infof(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, msg+"\n", args...)
}

func Exitf(code int, msg string, args ...interface{}) {
	if code == 0 {
		fmt.Fprintf(os.Stdout, msg+"\n", args...)
	} else {
		fmt.Fprintf(os.Stderr, msg+"\n", args...)
	}
	os.Exit(code)
}

func PrintObject(object interface{}, format string) {
	if format == "yaml" {
		yamlString, err := yaml.Marshal(object)
		if err != nil {
			Failf("unable to format yaml: %v", err)
		}
		fmt.Println(string(yamlString))
	} else if format == "json" {
		jsonString, err := json.Marshal(object)
		if err != nil {
			Failf("unable to format json: %v", err)
		}
		fmt.Printf("%s", pretty.Pretty(jsonString))
	} else {
		Failf("unknown format: %v", format)
	}
}

func PrintStrings(args ...string) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}