package task19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type replacement struct {
	from string
	to string
}

func (r *replacement)swap() {
	r.from, r.to = r.to, r.from
}

func getReplacements(molecule string, replacements []replacement) map[string]bool {
	outcomes := map[string]bool{};

	for _, r := range replacements {
		prefix := ""
		postfix := molecule
		pos := strings.Index(postfix, r.from)
		for pos > -1 {
			// Get absolute position in the molecule string
			pos += len(prefix)
			// Add outcome to the list
			outcomes[prefix + strings.Replace(postfix, r.from, r.to, 1)] = true
			// Replace pre and postfix
			prefix = molecule[:pos + len(r.from)]
			postfix = molecule[pos + len(r.from):]

			// Get new relative position
			pos = strings.Index(postfix, r.from)
		}
	}

	return outcomes
}

func sliceFilter(s []string, length int) []string {

	result := []string{}
	for _, v := range s {
		if len(v) > length {
			continue
		}
		result = append(result, v)
	}

	return result
}

func inSlice(haystack []string, needle string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func getMapKeys(m map[string]bool) []string {
    keys := []string{}
    for k := range m {
        keys = append(keys, k)
    }
	return keys
}

func Solve() {
	file, err := os.Open("res/task19.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	molecule := "HOH"
	replacements := []replacement{{"e", "H"}, {"e", "O"}, {"H", "HO"}, {"H", "OH"}, {"O", "HH"}}

	// Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, " => ") {
			parts := strings.Split(line, " => ")
			replacements = append(replacements, replacement{parts[0], parts[1]})
		} else if len(line) > 0 {
			molecule = line
		}
	}

	// Part 1
	fmt.Println(len(getReplacements(molecule, replacements)))

	// Part 2
	for k, r := range replacements {
		r.swap()
		replacements[k] = r
	}

//	fmt.Println(replacements)

	step := 1
	molecules := []string{molecule}
	for ;;step++ {
		fmt.Println("step", step, "Length", len(molecules))
		newMolecules := []string{}
		for _, molecule := range molecules {
//			fmt.Println("Molecule", len(molecule))
			for m, _ := range getReplacements(molecule, replacements) {
				if m == "e" || strings.Index(m, "e") == -1 {
					newMolecules = append(newMolecules, m)
				}
			}
		}

		molecules = newMolecules

		if inSlice(molecules, "e") {
			break
		}
	}

	fmt.Println("Steps:", step)
}
