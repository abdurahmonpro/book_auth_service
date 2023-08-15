package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

func ConvertMapToStruct(inputMap map[string]interface{}) (*structpb.Struct, error) {
	marshledInputMap, err := json.Marshal(inputMap)
	outputStruct := &structpb.Struct{}
	if err != nil {
		return outputStruct, err
	}
	err = protojson.Unmarshal(marshledInputMap, outputStruct)

	return outputStruct, err
}

func ConvertStringToDate(text string) (date string, err error) {
	months := map[string]string{
		"январ":   "01",
		"феврал":  "02",
		"март":    "03",
		"апрел":   "04",
		"май":     "05",
		"июн":     "06",
		"июл":     "07",
		"август":  "08",
		"сентябр": "09",
		"октябр":  "10",
		"ноябр":   "11",
		"декабр":  "12",
	}

	texts := strings.Split(text, " ")

	for key, val := range months {
		if strings.Contains(text, key) {
			var (
				year int = 2023
				day  int
			)

			for _, word := range texts {
				if _, err := strconv.Atoi(word); err != nil {
					continue
				}

				if len(word) == 4 {
					year, err = strconv.Atoi(word)
					if err != nil {
						return "", err
					}
				} else if len(word) <= 2 {
					day, err = strconv.Atoi(word)
					if err != nil {
						return "", err
					}
				}
			}

			date = fmt.Sprintf("%d.%s.%d", day, val, year)
			if day < 10 {
				date = fmt.Sprintf("0%d.%s.%d", day, val, year)
			}

			break
		}
	}

	if date == "" {
		return date, errors.New("Not found months")
	}

	return date, nil
}

func StructToProto(p interface{}, s interface{}) error {
	body, err := json.Marshal(s)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &p)

	return err
}
