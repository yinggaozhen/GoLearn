package yjson

import (
	"encoding/json"
	"strconv"
	"strings"
)

type PathType int

const (
	PATH_TYPE_STRING PathType = iota
	PATH_TYPE_NUMBER
)

type Result interface{
	Get(path interface{}) interface{}
}

type ParseResult struct {
	Result
	Data interface{}
}

type ParsePath struct {
	Raw string
	Type PathType
}

func newResult(data ...interface{}) ParseResult {
	if len(data) > 0 {
		return ParseResult{
			Data: data[0],
		}
	}

	return ParseResult{}
}

func Parse(parse string) (ParseResult, error) {
	result := newResult()
	e := json.Unmarshal([]byte(parse), &result.Data)
	return result, e
}

func (r ParseResult) clone(data ...interface{}) ParseResult {
	return newResult(data...)
}

func (r ParseResult) Get(path interface{}) interface{} {
	data := r.Data

	switch path.(type) {
	case int:
		return data.([]interface{})[path.(int)]
	case string:

		var cpData interface{}
		var ok bool
		var cr ParseResult

		firstPath, spiltPath := r.spiltFirstPath(path.(string))

		if _, ok = r.Data.(map[string]interface{}); ok {
			cpData = r.Data
			cr = r.clone(cpData.(map[string]interface{})[firstPath])
		} else if _, ok = r.Data.(string); ok {
			cpData = r.Data
		} else if _, ok := r.Data.([]interface{}); ok {
			cpData = r.Data

			if firstPath != "" {
				nFirstPath, _ := strconv.Atoi(firstPath)
				cr = r.clone(cpData.([]interface{})[nFirstPath])
			}

			if spiltPath != "" {
				spiltPath, _ = strconv.Atoi(spiltPath.(string))
			}
		}

		if spiltPath == "" && cr.Data != nil {
			return cr.Data
		} else if spiltPath != "" && cr.Data != nil {
			return cr.Get(spiltPath)
		} else {
			return nil
		}
	}

	return nil
}

func (r ParseResult) spiltFirstPath(path string) (string, interface{}) {
	paths := strings.SplitN(path, ".", 2)

	if len(paths) == 2 {
		return paths[0], paths[1]
	} else if len(paths) == 1 {
		return paths[0], ""
	} else {
		return "", ""
	}
}