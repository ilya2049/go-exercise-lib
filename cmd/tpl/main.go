package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

const (
	defaultSection  = "other"
	defaultExercise = "test"
)

var (
	sectionFlag  = flag.String("section", defaultSection, "an exercise section")
	exerciseFlag = flag.String("exercise", defaultExercise, "an exercise name")
)

func main() {
	var section, exercise string

	flag.Parse()

	if *sectionFlag != "" {
		section = *sectionFlag
	}

	if *exerciseFlag != "" {
		exercise = *exerciseFlag
	}

	err := createTemplate(section, exercise)
	if err != nil {
		fmt.Println("create an exercise template", err)

		return
	}

	fmt.Println("OK")
}

func getSection() string {
	defaultSection := "other"

	if *sectionFlag != "" {
		return *sectionFlag
	}

	return defaultSection
}

const exerciseLibDir = "internal/lib"

func createTemplate(section, exercise string) error {
	exerciseDir := "/" + section + "/" + exercise

	err := os.MkdirAll(exerciseLibDir+exerciseDir, os.ModeDir|0755)
	if err != nil {
		return fmt.Errorf("create an exercise dir: %w", err)
	}

	if err := createMainFile(exerciseLibDir + exerciseDir); err != nil {
		return fmt.Errorf("create a main file: %w", err)
	}

	if err := createTestFile(exerciseLibDir + exerciseDir); err != nil {
		return fmt.Errorf("create a test file: %w", err)
	}

	return nil
}

func createMainFile(dir string) error {
	fileName := dir + "/main.go"

	templateData := []byte(`package main

func main() {
}`)
	err := os.WriteFile(fileName, templateData, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

func createTestFile(dir string) error {
	fileName := dir + "/main_test.go"

	templateData := []byte(`package main

func Example_main() {
	main()
	// Output: 
}`)
	err := os.WriteFile(fileName, templateData, fs.ModePerm)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
