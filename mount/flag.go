// Copyright 2015 Google Inc. All Rights Reserved.
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

// Helper functions for dealing with mount(8)-style flags.
package mount

// Set up a flag with the given name and usage string that can be used to
// receive options in the format accepted by mount(8) and generated for its
// external mount helpers.
//
// It is assumed that option name and values do not contain commas, and that
// the first equals sign in an option is the name/value separator. There is no
// support for escaping.
//
// For example, if the flag is called "o" and the invocation looks like this:
//
//     -o ro -o user,foo=bar=baz,qux
//
// then the map will contain:
//
//     "ro": "",
//     "user": "",
//     "foo": "bar=baz",
//     "qux": "",
//
func OptionFlag(
	options map[string]string,
	name string,
	usage string) {
	panic("TODO")
}
