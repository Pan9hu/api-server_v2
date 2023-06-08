package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type PropertyUtil struct {
}

type AppConfigProperties map[string]string

/*
读取Property文件
Scanner将配置文件中的信息以key value形式存入map中
*/

func (pu *PropertyUtil) ReadPropertiesFile(filename string) (AppConfigProperties, error) {
	config := AppConfigProperties{}
	if len(filename) == 0 {
		return config, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}

	if scanErr := scanner.Err(); scanErr != nil {
		log.Fatal(scanErr)
		return nil, scanErr
	}

	return config, nil
}
