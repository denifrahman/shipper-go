package order

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/denifrahman/shipper-go"
	"github.com/joho/godotenv"
)

func init() {
	var err = godotenv.Load("../.env")

	if err != nil {
		panic("Error loading .env file.")
	}

	productionMode, errParse := strconv.ParseBool(os.Getenv("PRODUCTION_MODE"))

	if errParse != nil {
		productionMode = false
	}

	shipper.Conf.SetAPIKey(os.Getenv("API_KEY")).SetProductionMode(productionMode)
}

// CreateDomesticOrderParams is parameter for create domestic order in entire test.
var originCoordinate = "-6.308033944807303,106.73339847804874"
var destinationCoordinate = "49.020733179213,12.114381752908"
var CreateDomesticOrderParams = &DomesticOrderParams{
	Origin:               12921,
	Destination:          4645,
	Length:               10,
	Width:                10,
	Height:               10,
	WeightTotal:          0.4,
	Value:                100000,
	RateID:               59,
	ConsigneeName:        "Peoorang",
	ConsigneePhoneNumber: "089899878987",
	ConsignerName:        "Peorang",
	ConsignerPhoneNumber: "089891891818",
	OriginAddress:        "Mangga Dua Selatan",
	OriginDirection:      "Just follow the road",
	DestinationAddress:   "Pasar Baru",
	DestinationDirection: "Lurus terus",
	ItemName: []ItemName{
		{
			Name:  "Baju",
			Qty:   1,
			Value: 100000,
		},
	},
	Contents:              "Barang mudah pecah",
	UseInsurance:          true,
	ExternalID:            strconv.FormatInt(time.Now().UTC().UnixNano(), 10),
	PaymentType:           "cash",
	PackageType:           1,
	COD:                   false,
	OriginCoordinate:      &originCoordinate,
	DestinationCoordinate: &destinationCoordinate,
}

func TestCreateDomesticOrder(t *testing.T) {
	domesticOrder, err := CreateDomesticOrder(CreateDomesticOrderParams)

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(domesticOrder, "", "\t")

	fmt.Println("Domestic Order:")
	fmt.Println(string(s))
}

func TestCreateInternationalOrder(t *testing.T) {
	internationalOrder, err := CreateInternationalOrder(&InternationalOrderParams{
		Origin:               4819,
		Destination:          12,
		Length:               10,
		Width:                10,
		Height:               10,
		WeightTotal:          0.4,
		Value:                150000,
		RateID:               327,
		ConsigneeName:        "Peoorang",
		ConsigneePhoneNumber: "089899878987",
		ConsignerName:        "Peorang",
		ConsignerPhoneNumber: "089891891818",
		OriginAddress:        "Mangga Dua Selatan",
		OriginDirection:      "Just follow the road",
		DestinationAddress:   "Orchard Road 101",
		DestinationDirection: "Lurus terus",
		DestinationArea:      "Singapore",
		DestinationSuburb:    "Singapore",
		DestinationCity:      "Singapore",
		DestinationProvince:  "Singapore",
		DestinationPostCode:  "111111",
		ItemName: []ItemName{
			{
				Name:  "Baju",
				Qty:   1,
				Value: 150000,
			},
		},
		Contents:     "Barang mudah pecah",
		UseInsurance: false,
		PackageType:  2,
		PaymentType:  "cash",
		ExternalID:   strconv.FormatInt(time.Now().UTC().UnixNano(), 10),
	})

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(internationalOrder, "", "\t")

	fmt.Println("International Order:")
	fmt.Println(string(s))
}

func TestActivateOrder(t *testing.T) {
	domesticOrder, err := CreateDomesticOrder(CreateDomesticOrderParams)

	if err != nil {
		t.Error(err.Error())
	}

	activate, err := ActivateOrder(domesticOrder.Data.ID, &ActivateParams{
		Active: 1,
	})

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(activate, "", "\t")

	fmt.Println("Activate Order:")
	fmt.Println(string(s))
}

//func TestUpdateOrder(t *testing.T) {
//	domesticOrder, err := CreateDomesticOrder(CreateDomesticOrderParams)
//
//	if err != nil {
//		t.Error(err.Error())
//	}
//
//	updateOrder, err := UpdateOrder(domesticOrder.Data.ID, &UpdateOrderParams{
//		Length:      5,
//		Width:       5,
//		Height:      5,
//		WeightTotal: 1,
//	})
//
//	if err != nil {
//		t.Error(err.Error())
//	}
//
//	s, _ := json.MarshalIndent(updateOrder, "", "\t")
//
//	fmt.Println("Updated Order:")
//	fmt.Println(string(s))
//}

func TestCancelOrder(t *testing.T) {
	domesticOrder, err := CreateDomesticOrder(CreateDomesticOrderParams)

	if err != nil {
		t.Error(err.Error())
	}

	cancelOrder, err := CancelOrder(domesticOrder.Data.ID)

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(cancelOrder, "", "\t")

	fmt.Println("Cancelled Order:")
	fmt.Println(string(s))
}

func TestGetTrackingID(t *testing.T) {
	domesticOrder, err := CreateDomesticOrder(CreateDomesticOrderParams)

	if err != nil {
		t.Error(err.Error())
	}

	trackingID, err := GetTrackingID(domesticOrder.Data.ID)

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(trackingID, "", "\t")

	fmt.Println("Tracking ID:")
	fmt.Println(string(s))
}

func TestGetOrderDetail(t *testing.T) {
	//domesticOrder, err := CreateDomesticOrder(CreateDomesticOrderParams)
	//
	//if err != nil {
	//	t.Error(err.Error())
	//}

	detailOrder, err := GetOrderDetail("22BYDD7VGDDJK")

	if err != nil {
		t.Error(err.Error())
	}

	s, _ := json.MarshalIndent(detailOrder, "", "\t")
	fmt.Println(detailOrder.Data.Order.Tracking)
	fmt.Println("Order Detail:")
	fmt.Println(string(s))
}
