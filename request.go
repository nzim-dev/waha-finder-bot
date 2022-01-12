package info

import (
	"errors"
	"strconv"
	"strings"
)

type Request struct {
	SearchRequest  string
	LengthOfOutput int
}

func NewRequest(input string) (*Request, error) {
	credentials := strings.Split(input, ",")
	if len(credentials) != 2 {
		return nil, errors.New("wrong amount of arguments")
	}
	searchRequest := strings.TrimSpace(credentials[0])
	lengthOfOutput, err := strconv.Atoi(strings.TrimSpace(credentials[1]))
	if err != nil {
		return nil, err
	}

	return &Request{
		SearchRequest:  searchRequest,
		LengthOfOutput: lengthOfOutput,
	}, nil
}
