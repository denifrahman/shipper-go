package rates

import (
	"context"
	"encoding/json"
	"log"

	"github.com/denifrahman/shipper-go"
	"github.com/go-playground/validator/v10"
)

// The validator instance.
var validation = validator.New()

// GetDomesticRates gets available rates based on origin and destination area.
func GetDomesticRates(params *DomesticRatesParams) (DomesticRates, error) {
	return GetDomesticRatesWithContext(context.Background(), params)
}

// GetDomesticRatesWithContext gets available rates based on origin and destination area with context.
func GetDomesticRatesWithContext(ctx context.Context, params *DomesticRatesParams) (DomesticRates, error) {
	var errValidation = validation.Struct(params)

	if errValidation != nil {
		log.Fatalln(errValidation.Error())
	}

	var endpoint = shipper.Conf.BaseURL + "/domesticRates"
	var responseStruct = DomesticRates{}
	var additionalQueries map[string]interface{}

	tempJSON, _ := json.Marshal(params)
	json.Unmarshal(tempJSON, &additionalQueries)

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:             ctx,
		HTTPMethod:      "GET",
		Endpoint:        endpoint,
		AdditionalQuery: additionalQueries,
	}, &responseStruct)

	return responseStruct, err
}

// GetInternationalRates gets available rates based on origin and destination country.
func GetInternationalRates(params *InternationalRatesParams) (InternationalRates, error) {
	return GetInternationalRatesWithContext(context.Background(), params)
}

// GetInternationalRatesWithContext gets available rates based on origin and destination country with context.
func GetInternationalRatesWithContext(ctx context.Context, params *InternationalRatesParams) (InternationalRates, error) {
	var errValidation = validation.Struct(params)

	if errValidation != nil {
		log.Fatalln(errValidation.Error())
	}

	var endpoint = shipper.Conf.BaseURL + "/intlRates"
	var responseStruct = InternationalRates{}
	var additionalQueries map[string]interface{}

	tempJSON, _ := json.Marshal(params)
	json.Unmarshal(tempJSON, &additionalQueries)

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:             ctx,
		HTTPMethod:      "GET",
		Endpoint:        endpoint,
		AdditionalQuery: additionalQueries,
	}, &responseStruct)

	return responseStruct, err
}
