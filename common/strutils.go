// DBDeployer - The MySQL Sandbox
// Copyright © 2006-2018 Giuseppe Maxia
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

package common

import (
	"fmt"
	"os"
	"regexp"
)

// Given a path starting at the HOME directory
// returns a string where the literal value for $HOME
// is replaced by the string "$HOME"
func ReplaceLiteralHome( path string) string {
	// home := os.Getenv("HOME")
	// re := regexp.MustCompile(`^` + home)
	// return re.ReplaceAllString(path, "$$HOME")
	return ReplaceLiteralEnvVar(path, "HOME")
}

func ReplaceLiteralEnvVar(name string, env_var string) string {
	value := os.Getenv(env_var)
	re := regexp.MustCompile(value)
	return re.ReplaceAllString(name, "$$" + env_var)
}

func ReplaceEnvVar(name string, env_var string) string {
	value := os.Getenv(env_var)
	re := regexp.MustCompile(`\$` + env_var+ `\b`)
	return re.ReplaceAllString(name, value)
}

// Given a path with the variable "$HOME" at the start,
// returns a string with the value of HOME expanded
func ReplaceHomeVar(path string) string {
	// home := os.Getenv("HOME")
	// re := regexp.MustCompile(`^\$HOME\b`)
	//return re.ReplaceAllString(path, home)
	return ReplaceEnvVar(path, "HOME")
}

func MakeCustomizedUuid(port , node_num int ) string {
	re_digit := regexp.MustCompile(`\d`)
	group1 := fmt.Sprintf("%08d", port)
	group2 := fmt.Sprintf("%04d-%04d-%04d", node_num, node_num, node_num)
	group3 := fmt.Sprintf("%012d", port)
	//              12345678 1234 1234 1234 123456789012
	//    new_uuid="00000000-0000-0000-0000-000000000000"
	switch  {
	case node_num > 0 && node_num <= 9:
		group2 = re_digit.ReplaceAllString(group2, fmt.Sprintf("%d", node_num))
		group3 = re_digit.ReplaceAllString(group3, fmt.Sprintf("%d", node_num))
	// Number greater than 10 make little sense for this purpose.
	// But we keep the rule so that a valid UUID will be formatted in any case.
	case node_num >= 10000 && node_num <= 99999:
		group2 = fmt.Sprintf("%04d-%04d-%04d", 0, int(node_num / 10000), node_num - 10000 * int(node_num / 10000))
	case node_num >= 100000:
		group2 = fmt.Sprintf("%04d-%04d-%04d", int(node_num / 10000), 0, 0)
	case node_num >= 1000000:
		fmt.Printf("Node num out of boundaries: %d\n",node_num)
		os.Exit(1)
	}
	return fmt.Sprintf("%s-%s-%s", group1, group2, group3)
}

func Includes(main_string, contained string) bool {
	re := regexp.MustCompile(contained)
	return re.MatchString(main_string)

}
