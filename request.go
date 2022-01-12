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

	info := &Request{
		SearchRequest:  searchRequest,
		LengthOfOutput: lengthOfOutput,
	}

	if err := validateCredentials(info); err != nil {
		return nil, err
	}

	return info, nil
}

func validateCredentials(credentials *Request) error {
	if credentials.LengthOfOutput < 0 {
		credentials.LengthOfOutput = 1
	}

	if credentials.LengthOfOutput > 15 {
		credentials.LengthOfOutput = 15
	}

	if len(credentials.SearchRequest) < 1 {
		return errors.New("to short search request")
	}

	return nil
}
