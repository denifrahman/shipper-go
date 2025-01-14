package pickup

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/denifrahman/shipper-go"
	"github.com/go-playground/validator/v10"
)

// The validator instance.
var validation = validator.New()

// CancelPickup cancel Pickup request.
func CancelPickup(params *CancelPickupParams) (CancelledPickup, error) {
	return CancelPickupWithContext(context.Background(), params)
}

// CancelPickupWithContext cancel Pickup request with context.
func CancelPickupWithContext(ctx context.Context, params *CancelPickupParams) (CancelledPickup, error) {
	var errValidation = validation.Struct(params)

	if errValidation != nil {
		log.Fatalln(errValidation.Error())
	}

	var endpoint = shipper.Conf.BaseURL + "/pickup/cancel"
	var responseStruct = CancelledPickup{}
	var JSONParams, errEncode = json.Marshal(params)

	if errEncode != nil {
		log.Fatalln(errEncode.Error())
	}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:            ctx,
		HTTPMethod:     "PATCH",
		Endpoint:       endpoint,
		AdditionalBody: bytes.NewBuffer(JSONParams),
	}, &responseStruct)

	return responseStruct, err
}

// CreatePickup assigns agent and activate orders.
func CreatePickup(params *CreatePickupParams) (CreatePickupV3, error) {
	return CreatePickupWithContext(context.Background(), params)
}

// CreatePickupWithContext assigns agent and activate orders with context.
func CreatePickupWithContext(ctx context.Context, params *CreatePickupParams) (CreatePickupV3, error) {
	var errValidation = validation.Struct(params)

	if errValidation != nil {
		log.Fatalln(errValidation.Error())
	}

	var endpoint = shipper.Conf.BaseURL + "/pickup"
	var responseStruct = CreatePickupV3{}
	var JSONParams, errEncode = json.Marshal(params.ToCreatePickupParamsV3())
	fmt.Println(string(JSONParams))
	if errEncode != nil {
		log.Fatalln(errEncode.Error())
	}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:            ctx,
		HTTPMethod:     "POST",
		Endpoint:       endpoint,
		AdditionalBody: bytes.NewBuffer(JSONParams),
	}, &responseStruct)

	return responseStruct, err
}

// GetAgents gets agent by origin suburb ID.
func GetAgents(suburbID int) (Agents, error) {
	return GetAgentsWithContext(context.Background(), suburbID)
}

// GetAgentsWithContext gets agent by origin suburb ID with context.
func GetAgentsWithContext(ctx context.Context, suburbID int) (Agents, error) {
	var endpoint = shipper.Conf.BaseURL + "/agents"
	var responseStruct = Agents{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
		AdditionalQuery: map[string]interface{}{
			"suburbId": fmt.Sprint(suburbID),
		},
	}, &responseStruct)

	return responseStruct, err
}

// GetPickupTimeSlots gets pickup timeslot.
func GetPickupTimeSlots(timezone string) (TimeSlots, error) {
	return GetPickupTimeSlotsWithContext(context.Background(), timezone)
}

// GetPickupTimeSlotsWithContext gets pickup timeslot.
func GetPickupTimeSlotsWithContext(ctx context.Context, timezone string) (TimeSlots, error) {
	var endpoint = shipper.Conf.BaseURL + "/pickup/timeslot"
	var responseStruct = TimeSlotsV3{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
		AdditionalQuery: map[string]interface{}{
			"time_zone": timezone,
		},
	}, &responseStruct)
	return responseStruct.ToTimeSlot(), err
}
