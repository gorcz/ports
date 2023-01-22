package port

import (
	"encoding/json"
	"errors"
	"io"
)

// nolint: gocritic,lll
//go:generate mockgen -destination=./../../../mocks/pkg/model/port.go -package=mocks github.com/gorcz/ports/pkg/model/port Iterator

type Iterator interface {
	Next() (*Port, bool, error)
}

type iterator struct {
	jsonDecoder *json.Decoder
}

var (
	ErrNotJSONObject = errors.New("not JSON object")
	ErrPortCodeType  = errors.New("error port code type")
)

func ParsePortsFromJSONMap(ports io.Reader) (Iterator, error) {
	jsonDecoder := json.NewDecoder(ports)

	openingBracket, err := jsonDecoder.Token()
	if err == io.EOF {
		return newPortIterator(nil), nil
	}
	if err != nil {
		return nil, ErrNotJSONObject
	}

	if openingBracket.(json.Delim) != '{' {
		return nil, ErrNotJSONObject
	}

	return newPortIterator(jsonDecoder), nil
}

func newPortIterator(decoder *json.Decoder) *iterator {
	return &iterator{
		jsonDecoder: decoder,
	}
}

func (i *iterator) Next() (*Port, bool, error) {
	if i.jsonDecoder == nil {
		return nil, false, nil
	}

	if i.jsonDecoder.More() {
		key, err := i.jsonDecoder.Token()
		if err != nil {
			return nil, false, ErrPortCodeType
		}
		portCode, ok := key.(string)
		if !ok {
			return nil, false, ErrPortCodeType
		}

		port := &Port{
			Code: Code(portCode),
		}
		if err = i.jsonDecoder.Decode(&port.Details); err != nil {
			return nil, false, err
		}

		return port, true, nil
	}

	return nil, false, nil
}
