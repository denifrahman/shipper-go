package order

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/denifrahman/shipper-go"
	"github.com/go-playground/validator/v10"
)

// The validator instance.
var validation = validator.New()

// ActivateOrder activates (initiate Shipper's pickup process) or Deactivate an Order.
func ActivateOrder(orderID string, params *ActivateParams) (Activate, error) {
	return ActivateOrderWithContext(context.Background(), orderID, params)
}

// ActivateOrderWithContext activates (initiate Shipper's pickup process) or Deactivate an Order with context.
func ActivateOrderWithContext(ctx context.Context, orderID string, params *ActivateParams) (Activate, error) {
	var errValidation = validation.Struct(params)

	if errValidation != nil {
		log.Fatalln(errValidation.Error())
	}

	var endpoint = shipper.Conf.BaseURL + "/activations/" + orderID
	var responseStruct = Activate{}
	var JSONParams, errEncode = json.Marshal(params)

	if errEncode != nil {
		log.Fatalln(errEncode.Error())
	}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:            ctx,
		HTTPMethod:     "PUT",
		Endpoint:       endpoint,
		AdditionalBody: bytes.NewBuffer(JSONParams),
	}, &responseStruct)

	return responseStruct, err
}

// CancelOrder cancels specific created order.
func CancelOrder(orderID string) (CancelledOrder, error) {
	return CancelOrderWithContext(context.Background(), orderID)
}

// CancelOrderWithContext cancels specific created order with context.
func CancelOrderWithContext(ctx context.Context, orderID string) (CancelledOrder, error) {
	var endpoint = shipper.Conf.BaseURL + "/orders/" + orderID + "/cancel"
	var responseStruct = CancelledOrder{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "PUT",
		Endpoint:   endpoint,
	}, &responseStruct)

	return responseStruct, err
}

// CreateDomesticOrder creates Shipper domestic order.
func CreateDomesticOrder(params *DomesticOrderParams) (DomesticOrder, error) {
	return CreateDomesticOrderWithContext(context.Background(), params)
}

// CreateDomesticOrderWithContext creates Shipper domestic order with context.
func CreateDomesticOrderWithContext(ctx context.Context, params *DomesticOrderParams) (DomesticOrder, error) {
	var errValidation = validation.Struct(params)

	if errValidation != nil {
		log.Fatalln(errValidation.Error())
	}

	var endpoint = shipper.Conf.BaseURL + "/order"
	var responseStruct = DomesticOrderV3{}
	var JSONParams, errEncode = json.Marshal(params.ToDomesticOrderParams())

	if errEncode != nil {
		log.Fatalln(errEncode.Error())
	}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:            ctx,
		HTTPMethod:     "POST",
		Endpoint:       endpoint,
		AdditionalBody: bytes.NewBuffer(JSONParams),
	}, &responseStruct)

	return responseStruct.ToDomesticOrder(), err
}

// CreateInternationalOrder creates Shipper international order.
func CreateInternationalOrder(params *InternationalOrderParams) (InternationalOrder, error) {
	return CreateInternationalOrderWithContext(context.Background(), params)
}

// CreateInternationalOrderWithContext creates Shipper international order with context.
func CreateInternationalOrderWithContext(ctx context.Context, params *InternationalOrderParams) (InternationalOrder, error) {
	var errValidation = validation.Struct(params)

	if errValidation != nil {
		log.Fatalln(errValidation.Error())
	}

	var endpoint = shipper.Conf.BaseURL + "/orders/internationals"
	var responseStruct = InternationalOrder{}
	var JSONParams, errEncode = json.Marshal(params)

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

// GetOrderDetail gets created order's detail.
func GetOrderDetail(orderID string) (DetailOrder, error) {
	return GetOrderDetailWithContext(context.Background(), orderID)
}

// GetOrderDetailWithContext gets created order's detail with context.
func GetOrderDetailWithContext(ctx context.Context, orderID string) (DetailOrder, error) {
	var endpoint = shipper.Conf.BaseURL + "/order/" + orderID
	var responseStruct = DetailOrderV3{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
	}, &responseStruct)
	d, _ := json.Marshal(responseStruct)
	fmt.Println(string(d))
	return responseStruct.ToDetailOrder(), err
}

// GetTrackingID gets the Tracking ID based on submitted order ID.
func GetTrackingID(orderID string) (DomesticOrder, error) {
	return GetTrackingIDWithContext(context.Background(), orderID)
}

// GetTrackingIDWithContext gets the Tracking ID based on submitted order ID with context.
func GetTrackingIDWithContext(ctx context.Context, orderID string) (DomesticOrder, error) {
	var endpoint = shipper.Conf.BaseURL + "/orders"
	var responseStruct = DomesticOrder{}

	var err = shipper.SendRequest(&shipper.RequestParameters{
		Ctx:        ctx,
		HTTPMethod: "GET",
		Endpoint:   endpoint,
		AdditionalQuery: map[string]interface{}{
			"id": orderID,
		},
	}, &responseStruct)

	return responseStruct, err
}
