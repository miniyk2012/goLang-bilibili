package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
)
import "fmt"

type ConfigParser struct {
	sections map[string]*Section
}

func newConfigParser() *ConfigParser {
	return &ConfigParser{sections: make(map[string]*Section)}
}

type Section struct {
	name    string
	options map[string]string
}

func newSection(name string) *Section {
	return &Section{name: name, options: make(map[string]string)}
}

func (section *Section) put(key, value string) {
	section.options[key] = value
}

func (section *Section) get(key string) (string, bool) {
	value, ok := section.options[key]
	return value, ok
}

func (section *Section) OptionNames() (options []string) {
	for option := range section.options {
		options = append(options, option)
	}
	return options
}

func (parser *ConfigParser) Read(path string) error {
	inputFile, inputError := os.Open(path)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return errors.New("can't open file")
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	// 解析为sections
	return parser.parserLines(inputReader)
}

func (parser *ConfigParser) Get(sectionName, optionName string) string {
	if section, ok := parser.sections[sectionName]; ok {
		if val, ok := section.get(optionName); ok {
			return val
		}
	}
	panic("")
}

func (parser *ConfigParser) GetSectionNames() (sections []string) {
	for section := range parser.sections {
		sections = append(sections, section)
	}
	return sections
}

func (parser *ConfigParser) GetSection(name string) *Section {
	return parser.sections[name]
}
func (parser *ConfigParser) parserLines(inputReader *bufio.Reader) error {
	var currentSection string
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return nil
		}
		val1, val2, lineType := parserLine(inputString)
		switch lineType {
		case 0:
			currentSection = val1
			parser.sections[currentSection] = newSection(val1)
		case 1:
			parser.sections[currentSection].put(val1, val2)
		}
	}
}

func parserLine(inputString string) (string, string, int) {
	var setcionRe = regexp.MustCompile(`\[(\w+)\]`)
	var optionRe = regexp.MustCompile(`(\w+)=(\S+)`)
	if setcionRe.MatchString(inputString) {
		return setcionRe.FindStringSubmatch(inputString)[1], "", 0
	} else if optionRe.MatchString(inputString) {
		values := optionRe.FindStringSubmatch(inputString)
		return values[1], values[2], 1
	} else {
		return "", "", 2
	}
}


func main() {
	path := "example.ini"
	configParser := newConfigParser()
	configParser.Read(path)
	for _, sectionName := range configParser.GetSectionNames() {
		fmt.Printf("[%s]\n", sectionName)
		for _, optionName := range configParser.GetSection(sectionName).OptionNames() {
			fmt.Printf("%s = %s\n", optionName, configParser.Get(sectionName, optionName))
		}
		println()
	}
}