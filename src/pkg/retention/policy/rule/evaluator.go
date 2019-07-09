// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rule

import "github.com/goharbor/harbor/src/pkg/retention/res"

// Evaluator defines method of executing rule
type Evaluator interface {
	// Filter the inputs and return the filtered outputs
	//
	//  Arguments:
	//    artifacts []*res.Candidate : candidates for processing
	//
	//  Returns:
	//    []*res.Candidate : matched candidates for next stage
	//    error            : common error object if any errors occurred
	Process(artifacts []*res.Candidate) ([]*res.Candidate, error)

	// Specify what action is performed to the candidates processed by this evaluator
	Action() string
}

//  RuleFactory defines a factory method for creating rule evaluator
type RuleFactory func(parameters Parameters) Evaluator