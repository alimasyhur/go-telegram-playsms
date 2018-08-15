package main

import (
	"reflect"
	"regexp"
)

//InArray function
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func empty(s string) bool {
	if len(s) > 0 {
		return false
	}
	return true
}

//FormatMessage function
//skip double space
func FormatMessage(input string) string {
	reLeadcloseWhtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	reInsideWhtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := reLeadcloseWhtsp.ReplaceAllString(input, "")
	final = reInsideWhtsp.ReplaceAllString(final, " ")
	return final
}
