package order

import (
	"strings"
	"time"
)

// Activate struct contains response from API ActivateOrder.
type Activate struct {
	Status string `json:"status"`
	Data   struct {
		Message    string `json:"message"`
		StatusCode int    `json:"statusCode"`
	} `json:"data"`
}

// ActivateParams struct contains request parameter for API ActivateOrder.
type ActivateParams struct {
	Active  int `json:"active" validate:"required"`
	AgentID int `json:"agentId"`
}

// CancelledOrder struct contains response from API CancelOrder.
type CancelledOrder struct {
	Status string `json:"status"`
	Data   struct {
		Title      string `json:"title"`
		Message    string `json:"message"`
		StatusCode int    `json:"statusCode"`
	} `json:"data"`
}

// DomesticOrder struct contains response from API CreateDomesticOrder.
type DomesticOrder struct {
	Status string `json:"status"`
	Data   struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		StatusCode int    `json:"statusCode"`
		ID         string `json:"id"`
	} `json:"data"`
}

type DomesticOrderV3 struct {
	Metadata struct {
		Path           string `json:"path"`
		HttpStatusCode int    `json:"http_status_code"`
		HttpStatus     string `json:"http_status"`
		Timestamp      int    `json:"timestamp"`
	} `json:"metadata"`
	Data struct {
		Coverage    string `json:"coverage"`
		ExternalId  string `json:"external_id"`
		OrderId     string `json:"order_id"`
		PaymentType string `json:"payment_type"`
		Courier     struct {
			RateId          int  `json:"rate_id"`
			Amount          int  `json:"amount"`
			UseInsurance    bool `json:"use_insurance"`
			InsuranceAmount int  `json:"insurance_amount"`
			Cod             bool `json:"cod"`
		} `json:"courier"`
		Consignee struct {
			Name        string `json:"name"`
			Email       string `json:"email"`
			PhoneNumber string `json:"phone_number"`
		} `json:"consignee"`
		Consigner struct {
			Name        string `json:"name"`
			Email       string `json:"email"`
			PhoneNumber string `json:"phone_number"`
		} `json:"consigner"`
		Destination struct {
			Address      string `json:"address"`
			AreaId       int    `json:"area_id"`
			AreaName     string `json:"area_name"`
			CityId       int    `json:"city_id"`
			CityName     string `json:"city_name"`
			CountryId    int    `json:"country_id"`
			CountryName  string `json:"country_name"`
			Lat          string `json:"lat"`
			Lng          string `json:"lng"`
			Postcode     string `json:"postcode"`
			ProvinceId   int    `json:"province_id"`
			ProvinceName string `json:"province_name"`
			SuburbId     int    `json:"suburb_id"`
			SuburbName   string `json:"suburb_name"`
			EmailAddress string `json:"email_address"`
			CompanyName  string `json:"company_name"`
		} `json:"destination"`
		Origin struct {
			Address      string `json:"address"`
			AreaId       int    `json:"area_id"`
			AreaName     string `json:"area_name"`
			CityId       int    `json:"city_id"`
			CityName     string `json:"city_name"`
			CountryId    int    `json:"country_id"`
			CountryName  string `json:"country_name"`
			Lat          string `json:"lat"`
			Lng          string `json:"lng"`
			Postcode     string `json:"postcode"`
			ProvinceId   int    `json:"province_id"`
			ProvinceName string `json:"province_name"`
			SuburbId     int    `json:"suburb_id"`
			SuburbName   string `json:"suburb_name"`
			EmailAddress string `json:"email_address"`
			CompanyName  string `json:"company_name"`
		} `json:"origin"`
		Package struct {
			PackageType int     `json:"package_type"`
			Weight      float64 `json:"weight"`
			Length      int     `json:"length"`
			Width       int     `json:"width"`
			Height      int     `json:"height"`
			Price       int     `json:"price"`
			Items       []struct {
				Id    int    `json:"id"`
				Name  string `json:"name"`
				Qty   int    `json:"qty"`
				Price int    `json:"price"`
			} `json:"items"`
			International struct {
				CustomDeclaration struct {
					AdditionalDocument interface{} `json:"additional_document"`
					DocumentNumber     string      `json:"document_number"`
					TaxDocument        string      `json:"tax_document"`
				} `json:"custom_declaration"`
				DescriptionItem   string `json:"description_item"`
				DestinationPacket string `json:"destination_packet"`
				ItemType          string `json:"item_type"`
				MadeIn            string `json:"made_in"`
				Quantity          int    `json:"quantity"`
				Reason            string `json:"reason"`
				Unit              string `json:"unit"`
			} `json:"international"`
			ItemCategories []interface{} `json:"item_categories"`
		} `json:"package"`
	} `json:"data"`
}

func (r DomesticOrderV3) ToDomesticOrder() DomesticOrder {
	return DomesticOrder{
		Status: r.Metadata.HttpStatus,
		Data: struct {
			Title      string `json:"title"`
			Content    string `json:"content"`
			StatusCode int    `json:"statusCode"`
			ID         string `json:"id"`
		}{
			Title:      "Order",
			Content:    "",
			StatusCode: r.Metadata.HttpStatusCode,
			ID:         r.Data.OrderId,
		},
	}
}

// DomesticOrderParams struct contains request parameter for API CreateDomesticOrder.
type DomesticOrderParams struct {
	Origin                int        `json:"o" validate:"required"`
	Destination           int        `json:"d" validate:"required"`
	Length                float64    `json:"l" validate:"required"`
	Width                 float64    `json:"w" validate:"required"`
	Height                float64    `json:"h" validate:"required"`
	WeightTotal           float64    `json:"wt" validate:"required"`
	Value                 float64    `json:"v" validate:"required"`
	RateID                int        `json:"rateID" validate:"required"`
	ConsigneeName         string     `json:"consigneeName" validate:"required"`
	ConsigneePhoneNumber  string     `json:"consigneePhoneNumber" validate:"required"`
	ConsignerName         string     `json:"consignerName"`
	ConsignerPhoneNumber  string     `json:"consignerPhoneNumber"`
	OriginAddress         string     `json:"originAddress" validate:"required"`
	OriginDirection       string     `json:"originDirection" validate:"required"`
	DestinationAddress    string     `json:"destinationAddress" validate:"required"`
	DestinationDirection  string     `json:"destinationDirection" validate:"required"`
	ItemName              []ItemName `json:"itemName" validate:"required"`
	Contents              string     `json:"contents" validate:"required"`
	UseInsurance          bool       `json:"useInsurance"`
	ExternalID            string     `json:"externalID"`
	PaymentType           string     `json:"paymentType"`
	PackageType           int        `json:"packageType" validate:"required"`
	COD                   bool       `json:"cod"`
	OriginCoordinate      *string    `json:"originCoord"`
	DestinationCoordinate *string    `json:"destinationCoord"`
}

type DomesticOrderParamsV3 struct {
	Consignee struct {
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	} `json:"consignee"`
	Consigner struct {
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	} `json:"consigner"`
	Courier struct {
		Cod          bool `json:"cod"`
		RateId       int  `json:"rate_id"`
		UseInsurance bool `json:"use_insurance"`
	} `json:"courier"`
	Coverage    string `json:"coverage"`
	Destination struct {
		Address string `json:"address"`
		AreaId  int    `json:"area_id"`
		Lat     string `json:"lat"`
		Lng     string `json:"lng"`
	} `json:"destination"`
	ExternalId string `json:"external_id"`
	Origin     struct {
		Address string `json:"address"`
		AreaId  int    `json:"area_id"`
		Lat     string `json:"lat"`
		Lng     string `json:"lng"`
	} `json:"origin"`
	Package struct {
		Height      float64  `json:"height"`
		Items       []ItemV3 `json:"items"`
		Length      float64  `json:"length"`
		PackageType int      `json:"package_type"`
		Price       float64  `json:"price"`
		Weight      float64  `json:"weight"`
		Width       float64  `json:"width"`
	} `json:"package"`
	PaymentType string `json:"payment_type"`
}

type ItemV3 struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Qty   int    `json:"qty"`
}

func (r DomesticOrderParams) ToDomesticOrderParams() DomesticOrderParamsV3 {

	var items []ItemV3

	for _, val := range r.ItemName {
		items = append(items, ItemV3{
			Name:  val.Name,
			Price: val.Value,
			Qty:   val.Qty,
		})
	}
	return DomesticOrderParamsV3{
		Consignee: struct {
			Name        string `json:"name"`
			PhoneNumber string `json:"phone_number"`
		}{
			Name:        r.ConsigneeName,
			PhoneNumber: r.ConsignerPhoneNumber,
		},
		Consigner: struct {
			Name        string `json:"name"`
			PhoneNumber string `json:"phone_number"`
		}{
			Name:        r.ConsignerName,
			PhoneNumber: r.ConsignerPhoneNumber,
		},
		Courier: struct {
			Cod          bool `json:"cod"`
			RateId       int  `json:"rate_id"`
			UseInsurance bool `json:"use_insurance"`
		}{
			Cod:          r.COD,
			RateId:       r.RateID,
			UseInsurance: r.UseInsurance,
		},
		Destination: struct {
			Address string `json:"address"`
			AreaId  int    `json:"area_id"`
			Lat     string `json:"lat"`
			Lng     string `json:"lng"`
		}{
			Address: r.DestinationAddress,
			AreaId:  r.Destination,
			Lat:     strings.Split(*r.DestinationCoordinate, ",")[0],
			Lng:     strings.Split(*r.DestinationCoordinate, ",")[1],
		},
		Origin: struct {
			Address string `json:"address"`
			AreaId  int    `json:"area_id"`
			Lat     string `json:"lat"`
			Lng     string `json:"lng"`
		}{
			Address: r.OriginAddress,
			AreaId:  r.Origin,
			Lat:     strings.Split(*r.OriginCoordinate, ",")[0],
			Lng:     strings.Split(*r.OriginCoordinate, ",")[1],
		},
		Package: struct {
			Height      float64  `json:"height"`
			Items       []ItemV3 `json:"items"`
			Length      float64  `json:"length"`
			PackageType int      `json:"package_type"`
			Price       float64  `json:"price"`
			Weight      float64  `json:"weight"`
			Width       float64  `json:"width"`
		}{
			Height:      r.Height,
			Items:       items,
			Length:      r.Length,
			PackageType: r.PackageType,
			Price:       r.Value,
			Weight:      r.WeightTotal,
			Width:       r.Width,
		},
		Coverage:    "domestic",
		ExternalId:  r.ExternalID,
		PaymentType: r.PaymentType,
	}
}

// InternationalOrder struct contains response from API CreateInternationalOrder.
type InternationalOrder struct {
	Status string `json:"status"`
	Data   struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		StatusCode int    `json:"statusCode"`
		ID         string `json:"id"`
	} `json:"data"`
}

// InternationalOrderParams struct contains request parameter for API CreateInternationalOrder.
type InternationalOrderParams struct {
	Origin               int        `json:"o" validate:"required"`
	Destination          int        `json:"d" validate:"required"`
	Length               float64    `json:"l" validate:"required"`
	Width                float64    `json:"w" validate:"required"`
	Height               float64    `json:"h" validate:"required"`
	WeightTotal          float64    `json:"wt" validate:"required"`
	Value                float64    `json:"v" validate:"required"`
	RateID               int        `json:"rateID" validate:"required"`
	ConsigneeName        string     `json:"consigneeName" validate:"required"`
	ConsigneePhoneNumber string     `json:"consigneePhoneNumber" validate:"required"`
	ConsignerName        string     `json:"consignerName"`
	ConsignerPhoneNumber string     `json:"consignerPhoneNumber"`
	OriginAddress        string     `json:"originAddress" validate:"required"`
	OriginDirection      string     `json:"originDirection" validate:"required"`
	DestinationAddress   string     `json:"destinationAddress" validate:"required"`
	DestinationDirection string     `json:"destinationDirection" validate:"required"`
	DestinationArea      string     `json:"destinationArea"`
	DestinationSuburb    string     `json:"destinationSuburb"`
	DestinationCity      string     `json:"destinationCity"`
	DestinationProvince  string     `json:"destinationProvince"`
	DestinationPostCode  string     `json:"destinationPostcode"`
	ItemName             []ItemName `json:"itemName" validate:"required"`
	Contents             string     `json:"contents" validate:"required"`
	UseInsurance         bool       `json:"useInsurance"`
	ExternalID           string     `json:"externalID"`
	PaymentType          string     `json:"paymentType"`
	PackageType          int        `json:"packageType" validate:"required"`
}

// DetailOrder struct contains response from API GetOrderDetail.
type DetailOrder struct {
	Status string `json:"status"`
	Data   struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Order   struct {
			Tracking []tracking `json:"tracking"`
			Detail   struct {
				ID            string `json:"_id"`
				TrackingID    string `json:"id"`
				GroupID       int    `json:"groupID"`
				SpecialID     string `json:"specialID"`
				ExternalID    string `json:"externalID"`
				LabelChecksum string `json:"labelChecksum"`
				Consigner     struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					PhoneNumber string `json:"phoneNumber"`
				}
				Consignee struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					PhoneNumber string `json:"phoneNumber"`
				}
				BatchID        int            `json:"batchID"`
				AWBNumber      string         `json:"awbNumber"`
				StickerNumber  string         `json:"stickerNumber"`
				ShipmentStatus shipmentStatus `json:"shipmentStatus"`
				InternalStatus shipmentStatus `json:"internalStatus"`
				ExternalStatus shipmentStatus `json:"externalStatus"`
				Origin         origin         `json:"origin"`
				Destination    origin         `json:"destination"`
				Package        struct {
					ItemType  string `json:"itemType"`
					Contents  string `json:"contents"`
					Price     item   `json:"price"`
					ItemName  int    `json:"itemName"`
					Dimension struct {
						Length item `json:"length"`
						Width  item `json:"width"`
						Height item `json:"height"`
					} `json:"dimension"`
					Weight        item            `json:"weight"`
					CubicalWeight item            `json:"cubicalWeight"`
					Fragile       int             `json:"fragile"`
					Type          int             `json:"type"`
					IsConfirmed   int             `json:"isConfirmed"`
					PictureURL    string          `json:"pictureURL"`
					Details       []PackageDetail `json:"details"`
				} `json:"package"`
				Driver struct {
					ID                 int    `json:"id"`
					Name               string `json:"name"`
					PhoneNumber        string `json:"phoneNumber"`
					VehicleType        string `json:"vehicleType"`
					VehicleNumber      string `json:"vehicleNumber"`
					IsPaymentCollected int    `json:"isPaymentCollected"`
					Feedback           struct {
						Score   int    `json:"score"`
						Comment string `json:"comment"`
					} `json:"feedback"`
					AgentID     int    `json:"agentID"`
					AgentName   string `json:"agentName"`
					AgentCityID int    `json:"agentCityID"`
				} `json:"driver"`
				Courier struct {
					ID           int    `json:"id"`
					Name         string `json:"name"`
					RateID       int    `json:"rate_id"`
					RateName     string `json:"rate_name"`
					ShipmentType int    `json:"shipmentType"`
					ActualID     int    `json:"actualID"`
					MinDay       int    `json:"min_day"`
					MaxDay       int    `json:"max_day"`
					Rate         item   `json:"rate"`
					ActualRate   struct {
						ID int `json:"id"`
						item
					} `json:"actualRate"`
				} `json:"courier"`
				Rates struct {
					Shipment        item `json:"shipment"`
					PaidShipment    item `json:"paidShipment"`
					ActualShipment  item `json:"actualShipment"`
					Insurance       item `json:"insurance"`
					PaidInsurance   item `json:"paidInsurance"`
					ActualInsurance item `json:"actualInsurance"`
					EscrowCost      item `json:"escrowCost"`
					FulfillmentCost item `json:"fulfillmentCost"`
					ItemPrice       item `json:"itemPrice"`
					Discount        item `json:"discount"`
				} `json:"rates"`
				UseInsurance    bool   `json:"useInsurance"`
				IsLabelPrinted  int    `json:"isLabelPrinted"`
				PaymentType     string `json:"paymentType"`
				Source          string `json:"source"`
				IsActive        bool   `json:"isActive"`
				IsAutoTrack     int    `json:"isAutoTrack"`
				IsCustomAWB     int    `json:"isCustomAWB"`
				ReadyTime       string `json:"readyTime"`
				PickUpTime      string `json:"pickUpTime"`
				PickUpCode      string `json:"PickUpCode"`
				CreationDate    string `json:"creationDate"`
				ActiveDate      string `json:"activeDate"`
				LastUpdatedDate string `json:"lastUpdatedDate"`
				CreatedBy       int    `json:"createdBy"`
				ShipmentArea    string `json:"shipmentArea"`
				COD             struct {
					Order        int `json:"order"`
					FreeDelivery int `json:"freeDelivery"`
				} `json:"cod"`
				Voucher          string `json:"voucher"`
				Whitelabel       string `json:"whitelabel"`
				JOBNumber        string `json:"JOBNumber"`
				Domain           string `json:"domain"`
				IsAutomateOrder  int    `json:"isAutomateOrder"`
				Channel          string `json:"channel"`
				IsHubless        int    `json:"isHubless"`
				ZonaID           int    `json:"zonaID"`
				IsJobMarketplace int    `json:"isJobMarketplace"`
				RequestID        string `json:"requestID"`
				PartitionKey     int    `json:"partitionKey"`
				HandlerIDs       []int  `json:"handler_ids"`
			} `json:"detail"`
		} `json:"order"`
		StatusCode int `json:"statusCode"`
	} `json:"data"`
}

type ShipperStatus struct {
	Code        int    `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type LogisticStatusV3 struct {
	Code        int    `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TrackingsV3 struct {
	ShipperStatus  ShipperStatus    `json:"shipper_status"`
	LogisticStatus LogisticStatusV3 `json:"logistic_status"`
	CreatedDate    time.Time        `json:"created_date"`
}

type PackageDetail struct {
	ItemID    int     `json:"itemID"`
	ItemName  string  `json:"itemName"`
	ItemPrice float64 `json:"itemPrice"`
	ItemQTY   int     `json:"itemQTY"`
}

type DetailOrderV3 struct {
	Metadata struct {
		Path           string `json:"path"`
		HttpStatusCode int    `json:"http_status_code"`
		HttpStatus     string `json:"http_status"`
		Timestamp      int    `json:"timestamp"`
	} `json:"metadata"`
	Data struct {
		Consignee struct {
			Name        string `json:"name"`
			PhoneNumber string `json:"phone_number"`
			Email       string `json:"email"`
		} `json:"consignee"`
		Consigner struct {
			Name        string `json:"name"`
			PhoneNumber string `json:"phone_number"`
			Email       string `json:"email"`
		} `json:"consigner"`
		Origin struct {
			Id           int    `json:"id"`
			StopId       int    `json:"stop_id"`
			Address      string `json:"address"`
			Direction    string `json:"direction"`
			Postcode     string `json:"postcode"`
			AreaId       int    `json:"area_id"`
			AreaName     string `json:"area_name"`
			SuburbId     int    `json:"suburb_id"`
			SuburbName   string `json:"suburb_name"`
			CityId       int    `json:"city_id"`
			CityName     string `json:"city_name"`
			ProvinceId   int    `json:"province_id"`
			ProvinceName string `json:"province_name"`
			CountryId    int    `json:"country_id"`
			CountryName  string `json:"country_name"`
			Lat          string `json:"lat"`
			Lng          string `json:"lng"`
			EmailAddress string `json:"email_address"`
			CompanyName  string `json:"company_name"`
		} `json:"origin"`
		Destination struct {
			Id           int    `json:"id"`
			StopId       int    `json:"stop_id"`
			Address      string `json:"address"`
			Direction    string `json:"direction"`
			Postcode     string `json:"postcode"`
			AreaId       int    `json:"area_id"`
			AreaName     string `json:"area_name"`
			SuburbId     int    `json:"suburb_id"`
			SuburbName   string `json:"suburb_name"`
			CityId       int    `json:"city_id"`
			CityName     string `json:"city_name"`
			ProvinceId   int    `json:"province_id"`
			ProvinceName string `json:"province_name"`
			CountryId    int    `json:"country_id"`
			CountryName  string `json:"country_name"`
			Lat          string `json:"lat"`
			Lng          string `json:"lng"`
			EmailAddress string `json:"email_address"`
			CompanyName  string `json:"company_name"`
		} `json:"destination"`
		ExternalId string `json:"external_id"`
		OrderId    string `json:"order_id"`
		Courier    struct {
			Name            string  `json:"name"`
			RateId          int     `json:"rate_id"`
			RateName        string  `json:"rate_name"`
			Amount          float64 `json:"amount"`
			UseInsurance    bool    `json:"use_insurance"`
			InsuranceAmount float64 `json:"insurance_amount"`
			Cod             bool    `json:"cod"`
			MinDay          int     `json:"min_day"`
			MaxDay          int     `json:"max_day"`
			PriceBreakdown  struct {
				SurchargeFee float64 `json:"surcharge_fee"`
				BasePrice    float64 `json:"base_price"`
				Discount     float64 `json:"discount"`
				InsuranceFee float64 `json:"insurance_fee"`
				FinalPrice   float64 `json:"final_price"`
			} `json:"price_breakdown"`
		} `json:"courier"`
		Package struct {
			Weight       float64 `json:"weight"`
			Length       float64 `json:"length"`
			Width        float64 `json:"width"`
			Height       float64 `json:"height"`
			VolumeWeight float64 `json:"volume_weight"`
			PackageType  int     `json:"package_type"`
			Items        []struct {
				Id    int     `json:"id"`
				Name  string  `json:"name"`
				Price float64 `json:"price"`
				Qty   int     `json:"qty"`
			} `json:"items"`
			International struct {
				CustomDeclaration struct {
					AdditionalDocument []interface{} `json:"additional_document"`
					DocumentNumber     string        `json:"document_number"`
					TaxDocument        string        `json:"tax_document"`
				} `json:"custom_declaration"`
				DescriptionItem   string `json:"description_item"`
				DestinationPacket string `json:"destination_packet"`
				ItemType          string `json:"item_type"`
				MadeIn            string `json:"made_in"`
				Quantity          int    `json:"quantity"`
				Reason            string `json:"reason"`
				Unit              string `json:"unit"`
			} `json:"international"`
			ItemCategories []interface{} `json:"item_categories"`
		} `json:"package"`
		PaymentType string `json:"payment_type"`
		Driver      struct {
			Name          string `json:"name"`
			Phone         string `json:"phone"`
			VehicleType   string `json:"vehicle_type"`
			VehicleNumber string `json:"vehicle_number"`
		} `json:"driver"`
		LabelCheckSum   string        `json:"label_check_sum"`
		CreationDate    time.Time     `json:"creation_date"`
		LastUpdatedDate time.Time     `json:"last_updated_date"`
		AwbNumber       string        `json:"awb_number"`
		Trackings       []TrackingsV3 `json:"trackings"`
		IsActive        bool          `json:"is_active"`
		IsHubless       bool          `json:"is_hubless"`
		PickupCode      string        `json:"pickup_code"`
		PickupTime      string        `json:"pickup_time"`
		ShipmentStatus  struct {
			Name        string    `json:"name"`
			Description string    `json:"description"`
			Code        int       `json:"code"`
			UpdatedBy   string    `json:"updated_by"`
			UpdatedDate time.Time `json:"updated_date"`
			TrackUrl    string    `json:"track_url"`
			Reason      string    `json:"reason"`
			CreatedDate time.Time `json:"created_date"`
		} `json:"shipment_status"`
		ProofOfDelivery struct {
			Photo     string `json:"photo"`
			Signature string `json:"signature"`
		} `json:"proof_of_delivery"`
		TimeSlotSelected struct {
			StartTime interface{} `json:"start_time"`
			EndTime   interface{} `json:"end_time"`
		} `json:"time_slot_selected"`
	} `json:"data"`
}

func (r DetailOrderV3) ToDetailOrder() DetailOrder {

	var trackingList []tracking
	var LogisticStatusList []trackStatus
	for _, val := range r.Data.Trackings {
		LogisticStatusList = append(LogisticStatusList, trackStatus{
			ID:          0,
			Name:        val.LogisticStatus.Name,
			Description: val.LogisticStatus.Description,
		})

		trackingList = append(trackingList, tracking{
			ID:       r.Data.OrderId,
			OrderID:  r.Data.OrderId,
			UniqueID: r.Data.OrderId,
			TrackStatus: trackStatus{
				ID:          val.LogisticStatus.Code,
				Name:        val.LogisticStatus.Name,
				Description: val.LogisticStatus.Description,
			},
			InternalStatus: trackStatus{
				ID:          val.ShipperStatus.Code,
				Name:        val.ShipperStatus.Name,
				Description: val.ShipperStatus.Description,
			},
			ExternalStatus: trackStatus{
				ID:          0,
				Name:        val.LogisticStatus.Name,
				Description: val.LogisticStatus.Description,
			},
			LogisticStatus: LogisticStatusList,
			CreatedBy:      "",
			CreatedDate:    val.CreatedDate.String(),
		})
	}

	var packageItem []item

	for _, val := range r.Data.Package.Items {
		packageItem = append(packageItem, item{
			Value: val.Price,
			UOM:   "",
		})
	}

	var packageDetail []PackageDetail
	for _, val := range r.Data.Package.Items {
		packageDetail = append(packageDetail, PackageDetail{
			ItemID:    val.Id,
			ItemName:  val.Name,
			ItemPrice: val.Price,
			ItemQTY:   val.Qty,
		})
	}

	return DetailOrder{
		Status: r.Metadata.HttpStatus,
		Data: struct {
			Title   string `json:"title"`
			Content string `json:"content"`
			Order   struct {
				Tracking []tracking `json:"tracking"`
				Detail   struct {
					ID            string `json:"_id"`
					TrackingID    string `json:"id"`
					GroupID       int    `json:"groupID"`
					SpecialID     string `json:"specialID"`
					ExternalID    string `json:"externalID"`
					LabelChecksum string `json:"labelChecksum"`
					Consigner     struct {
						ID          string `json:"id"`
						Name        string `json:"name"`
						PhoneNumber string `json:"phoneNumber"`
					}
					Consignee struct {
						ID          string `json:"id"`
						Name        string `json:"name"`
						PhoneNumber string `json:"phoneNumber"`
					}
					BatchID        int            `json:"batchID"`
					AWBNumber      string         `json:"awbNumber"`
					StickerNumber  string         `json:"stickerNumber"`
					ShipmentStatus shipmentStatus `json:"shipmentStatus"`
					InternalStatus shipmentStatus `json:"internalStatus"`
					ExternalStatus shipmentStatus `json:"externalStatus"`
					Origin         origin         `json:"origin"`
					Destination    origin         `json:"destination"`
					Package        struct {
						ItemType  string `json:"itemType"`
						Contents  string `json:"contents"`
						Price     item   `json:"price"`
						ItemName  int    `json:"itemName"`
						Dimension struct {
							Length item `json:"length"`
							Width  item `json:"width"`
							Height item `json:"height"`
						} `json:"dimension"`
						Weight        item            `json:"weight"`
						CubicalWeight item            `json:"cubicalWeight"`
						Fragile       int             `json:"fragile"`
						Type          int             `json:"type"`
						IsConfirmed   int             `json:"isConfirmed"`
						PictureURL    string          `json:"pictureURL"`
						Details       []PackageDetail `json:"details"`
					} `json:"package"`
					Driver struct {
						ID                 int    `json:"id"`
						Name               string `json:"name"`
						PhoneNumber        string `json:"phoneNumber"`
						VehicleType        string `json:"vehicleType"`
						VehicleNumber      string `json:"vehicleNumber"`
						IsPaymentCollected int    `json:"isPaymentCollected"`
						Feedback           struct {
							Score   int    `json:"score"`
							Comment string `json:"comment"`
						} `json:"feedback"`
						AgentID     int    `json:"agentID"`
						AgentName   string `json:"agentName"`
						AgentCityID int    `json:"agentCityID"`
					} `json:"driver"`
					Courier struct {
						ID           int    `json:"id"`
						Name         string `json:"name"`
						RateID       int    `json:"rate_id"`
						RateName     string `json:"rate_name"`
						ShipmentType int    `json:"shipmentType"`
						ActualID     int    `json:"actualID"`
						MinDay       int    `json:"min_day"`
						MaxDay       int    `json:"max_day"`
						Rate         item   `json:"rate"`
						ActualRate   struct {
							ID int `json:"id"`
							item
						} `json:"actualRate"`
					} `json:"courier"`
					Rates struct {
						Shipment        item `json:"shipment"`
						PaidShipment    item `json:"paidShipment"`
						ActualShipment  item `json:"actualShipment"`
						Insurance       item `json:"insurance"`
						PaidInsurance   item `json:"paidInsurance"`
						ActualInsurance item `json:"actualInsurance"`
						EscrowCost      item `json:"escrowCost"`
						FulfillmentCost item `json:"fulfillmentCost"`
						ItemPrice       item `json:"itemPrice"`
						Discount        item `json:"discount"`
					} `json:"rates"`
					UseInsurance    bool   `json:"useInsurance"`
					IsLabelPrinted  int    `json:"isLabelPrinted"`
					PaymentType     string `json:"paymentType"`
					Source          string `json:"source"`
					IsActive        bool   `json:"isActive"`
					IsAutoTrack     int    `json:"isAutoTrack"`
					IsCustomAWB     int    `json:"isCustomAWB"`
					ReadyTime       string `json:"readyTime"`
					PickUpTime      string `json:"pickUpTime"`
					PickUpCode      string `json:"PickUpCode"`
					CreationDate    string `json:"creationDate"`
					ActiveDate      string `json:"activeDate"`
					LastUpdatedDate string `json:"lastUpdatedDate"`
					CreatedBy       int    `json:"createdBy"`
					ShipmentArea    string `json:"shipmentArea"`
					COD             struct {
						Order        int `json:"order"`
						FreeDelivery int `json:"freeDelivery"`
					} `json:"cod"`
					Voucher          string `json:"voucher"`
					Whitelabel       string `json:"whitelabel"`
					JOBNumber        string `json:"JOBNumber"`
					Domain           string `json:"domain"`
					IsAutomateOrder  int    `json:"isAutomateOrder"`
					Channel          string `json:"channel"`
					IsHubless        int    `json:"isHubless"`
					ZonaID           int    `json:"zonaID"`
					IsJobMarketplace int    `json:"isJobMarketplace"`
					RequestID        string `json:"requestID"`
					PartitionKey     int    `json:"partitionKey"`
					HandlerIDs       []int  `json:"handler_ids"`
				} `json:"detail"`
			} `json:"order"`
			StatusCode int `json:"statusCode"`
		}{
			Title:   "",
			Content: "",
			Order: struct {
				Tracking []tracking `json:"tracking"`
				Detail   struct {
					ID            string `json:"_id"`
					TrackingID    string `json:"id"`
					GroupID       int    `json:"groupID"`
					SpecialID     string `json:"specialID"`
					ExternalID    string `json:"externalID"`
					LabelChecksum string `json:"labelChecksum"`
					Consigner     struct {
						ID          string `json:"id"`
						Name        string `json:"name"`
						PhoneNumber string `json:"phoneNumber"`
					}
					Consignee struct {
						ID          string `json:"id"`
						Name        string `json:"name"`
						PhoneNumber string `json:"phoneNumber"`
					}
					BatchID        int            `json:"batchID"`
					AWBNumber      string         `json:"awbNumber"`
					StickerNumber  string         `json:"stickerNumber"`
					ShipmentStatus shipmentStatus `json:"shipmentStatus"`
					InternalStatus shipmentStatus `json:"internalStatus"`
					ExternalStatus shipmentStatus `json:"externalStatus"`
					Origin         origin         `json:"origin"`
					Destination    origin         `json:"destination"`
					Package        struct {
						ItemType  string `json:"itemType"`
						Contents  string `json:"contents"`
						Price     item   `json:"price"`
						ItemName  int    `json:"itemName"`
						Dimension struct {
							Length item `json:"length"`
							Width  item `json:"width"`
							Height item `json:"height"`
						} `json:"dimension"`
						Weight        item            `json:"weight"`
						CubicalWeight item            `json:"cubicalWeight"`
						Fragile       int             `json:"fragile"`
						Type          int             `json:"type"`
						IsConfirmed   int             `json:"isConfirmed"`
						PictureURL    string          `json:"pictureURL"`
						Details       []PackageDetail `json:"details"`
					} `json:"package"`
					Driver struct {
						ID                 int    `json:"id"`
						Name               string `json:"name"`
						PhoneNumber        string `json:"phoneNumber"`
						VehicleType        string `json:"vehicleType"`
						VehicleNumber      string `json:"vehicleNumber"`
						IsPaymentCollected int    `json:"isPaymentCollected"`
						Feedback           struct {
							Score   int    `json:"score"`
							Comment string `json:"comment"`
						} `json:"feedback"`
						AgentID     int    `json:"agentID"`
						AgentName   string `json:"agentName"`
						AgentCityID int    `json:"agentCityID"`
					} `json:"driver"`
					Courier struct {
						ID           int    `json:"id"`
						Name         string `json:"name"`
						RateID       int    `json:"rate_id"`
						RateName     string `json:"rate_name"`
						ShipmentType int    `json:"shipmentType"`
						ActualID     int    `json:"actualID"`
						MinDay       int    `json:"min_day"`
						MaxDay       int    `json:"max_day"`
						Rate         item   `json:"rate"`
						ActualRate   struct {
							ID int `json:"id"`
							item
						} `json:"actualRate"`
					} `json:"courier"`
					Rates struct {
						Shipment        item `json:"shipment"`
						PaidShipment    item `json:"paidShipment"`
						ActualShipment  item `json:"actualShipment"`
						Insurance       item `json:"insurance"`
						PaidInsurance   item `json:"paidInsurance"`
						ActualInsurance item `json:"actualInsurance"`
						EscrowCost      item `json:"escrowCost"`
						FulfillmentCost item `json:"fulfillmentCost"`
						ItemPrice       item `json:"itemPrice"`
						Discount        item `json:"discount"`
					} `json:"rates"`
					UseInsurance    bool   `json:"useInsurance"`
					IsLabelPrinted  int    `json:"isLabelPrinted"`
					PaymentType     string `json:"paymentType"`
					Source          string `json:"source"`
					IsActive        bool   `json:"isActive"`
					IsAutoTrack     int    `json:"isAutoTrack"`
					IsCustomAWB     int    `json:"isCustomAWB"`
					ReadyTime       string `json:"readyTime"`
					PickUpTime      string `json:"pickUpTime"`
					PickUpCode      string `json:"PickUpCode"`
					CreationDate    string `json:"creationDate"`
					ActiveDate      string `json:"activeDate"`
					LastUpdatedDate string `json:"lastUpdatedDate"`
					CreatedBy       int    `json:"createdBy"`
					ShipmentArea    string `json:"shipmentArea"`
					COD             struct {
						Order        int `json:"order"`
						FreeDelivery int `json:"freeDelivery"`
					} `json:"cod"`
					Voucher          string `json:"voucher"`
					Whitelabel       string `json:"whitelabel"`
					JOBNumber        string `json:"JOBNumber"`
					Domain           string `json:"domain"`
					IsAutomateOrder  int    `json:"isAutomateOrder"`
					Channel          string `json:"channel"`
					IsHubless        int    `json:"isHubless"`
					ZonaID           int    `json:"zonaID"`
					IsJobMarketplace int    `json:"isJobMarketplace"`
					RequestID        string `json:"requestID"`
					PartitionKey     int    `json:"partitionKey"`
					HandlerIDs       []int  `json:"handler_ids"`
				} `json:"detail"`
			}{
				Tracking: trackingList,
				Detail: struct {
					ID            string `json:"_id"`
					TrackingID    string `json:"id"`
					GroupID       int    `json:"groupID"`
					SpecialID     string `json:"specialID"`
					ExternalID    string `json:"externalID"`
					LabelChecksum string `json:"labelChecksum"`
					Consigner     struct {
						ID          string `json:"id"`
						Name        string `json:"name"`
						PhoneNumber string `json:"phoneNumber"`
					}
					Consignee struct {
						ID          string `json:"id"`
						Name        string `json:"name"`
						PhoneNumber string `json:"phoneNumber"`
					}
					BatchID        int            `json:"batchID"`
					AWBNumber      string         `json:"awbNumber"`
					StickerNumber  string         `json:"stickerNumber"`
					ShipmentStatus shipmentStatus `json:"shipmentStatus"`
					InternalStatus shipmentStatus `json:"internalStatus"`
					ExternalStatus shipmentStatus `json:"externalStatus"`
					Origin         origin         `json:"origin"`
					Destination    origin         `json:"destination"`
					Package        struct {
						ItemType  string `json:"itemType"`
						Contents  string `json:"contents"`
						Price     item   `json:"price"`
						ItemName  int    `json:"itemName"`
						Dimension struct {
							Length item `json:"length"`
							Width  item `json:"width"`
							Height item `json:"height"`
						} `json:"dimension"`
						Weight        item            `json:"weight"`
						CubicalWeight item            `json:"cubicalWeight"`
						Fragile       int             `json:"fragile"`
						Type          int             `json:"type"`
						IsConfirmed   int             `json:"isConfirmed"`
						PictureURL    string          `json:"pictureURL"`
						Details       []PackageDetail `json:"details"`
					} `json:"package"`
					Driver struct {
						ID                 int    `json:"id"`
						Name               string `json:"name"`
						PhoneNumber        string `json:"phoneNumber"`
						VehicleType        string `json:"vehicleType"`
						VehicleNumber      string `json:"vehicleNumber"`
						IsPaymentCollected int    `json:"isPaymentCollected"`
						Feedback           struct {
							Score   int    `json:"score"`
							Comment string `json:"comment"`
						} `json:"feedback"`
						AgentID     int    `json:"agentID"`
						AgentName   string `json:"agentName"`
						AgentCityID int    `json:"agentCityID"`
					} `json:"driver"`
					Courier struct {
						ID           int    `json:"id"`
						Name         string `json:"name"`
						RateID       int    `json:"rate_id"`
						RateName     string `json:"rate_name"`
						ShipmentType int    `json:"shipmentType"`
						ActualID     int    `json:"actualID"`
						MinDay       int    `json:"min_day"`
						MaxDay       int    `json:"max_day"`
						Rate         item   `json:"rate"`
						ActualRate   struct {
							ID int `json:"id"`
							item
						} `json:"actualRate"`
					} `json:"courier"`
					Rates struct {
						Shipment        item `json:"shipment"`
						PaidShipment    item `json:"paidShipment"`
						ActualShipment  item `json:"actualShipment"`
						Insurance       item `json:"insurance"`
						PaidInsurance   item `json:"paidInsurance"`
						ActualInsurance item `json:"actualInsurance"`
						EscrowCost      item `json:"escrowCost"`
						FulfillmentCost item `json:"fulfillmentCost"`
						ItemPrice       item `json:"itemPrice"`
						Discount        item `json:"discount"`
					} `json:"rates"`
					UseInsurance    bool   `json:"useInsurance"`
					IsLabelPrinted  int    `json:"isLabelPrinted"`
					PaymentType     string `json:"paymentType"`
					Source          string `json:"source"`
					IsActive        bool   `json:"isActive"`
					IsAutoTrack     int    `json:"isAutoTrack"`
					IsCustomAWB     int    `json:"isCustomAWB"`
					ReadyTime       string `json:"readyTime"`
					PickUpTime      string `json:"pickUpTime"`
					PickUpCode      string `json:"PickUpCode"`
					CreationDate    string `json:"creationDate"`
					ActiveDate      string `json:"activeDate"`
					LastUpdatedDate string `json:"lastUpdatedDate"`
					CreatedBy       int    `json:"createdBy"`
					ShipmentArea    string `json:"shipmentArea"`
					COD             struct {
						Order        int `json:"order"`
						FreeDelivery int `json:"freeDelivery"`
					} `json:"cod"`
					Voucher          string `json:"voucher"`
					Whitelabel       string `json:"whitelabel"`
					JOBNumber        string `json:"JOBNumber"`
					Domain           string `json:"domain"`
					IsAutomateOrder  int    `json:"isAutomateOrder"`
					Channel          string `json:"channel"`
					IsHubless        int    `json:"isHubless"`
					ZonaID           int    `json:"zonaID"`
					IsJobMarketplace int    `json:"isJobMarketplace"`
					RequestID        string `json:"requestID"`
					PartitionKey     int    `json:"partitionKey"`
					HandlerIDs       []int  `json:"handler_ids"`
				}{
					ID:            r.Data.OrderId,
					ExternalID:    r.Data.ExternalId,
					LabelChecksum: r.Data.LabelCheckSum,
					Consigner: struct {
						ID          string `json:"id"`
						Name        string `json:"name"`
						PhoneNumber string `json:"phoneNumber"`
					}{
						ID:          "",
						Name:        r.Data.Consigner.Name,
						PhoneNumber: r.Data.Consigner.PhoneNumber,
					},
					Consignee: struct {
						ID          string `json:"id"`
						Name        string `json:"name"`
						PhoneNumber string `json:"phoneNumber"`
					}{
						ID:          "",
						Name:        r.Data.Consignee.Name,
						PhoneNumber: r.Data.Consignee.PhoneNumber,
					},
					AWBNumber: r.Data.AwbNumber,
					ShipmentStatus: shipmentStatus{
						Name:        r.Data.ShipmentStatus.Name,
						Description: r.Data.ShipmentStatus.Description,
						StatusCode:  r.Data.ShipmentStatus.Code,
						UpdatedBy:   r.Data.ShipmentStatus.UpdatedBy,
						UpdateDate:  r.Data.ShipmentStatus.UpdatedDate.String(),
						TrackUrl:    r.Data.ShipmentStatus.TrackUrl,
					},
					InternalStatus: shipmentStatus{
						Name:        r.Data.ShipmentStatus.Name,
						Description: r.Data.ShipmentStatus.Description,
						StatusCode:  r.Data.ShipmentStatus.Code,
						UpdatedBy:   r.Data.ShipmentStatus.UpdatedBy,
						UpdateDate:  r.Data.ShipmentStatus.UpdatedDate.String(),
						TrackUrl:    r.Data.ShipmentStatus.TrackUrl,
					},
					ExternalStatus: shipmentStatus{
						Name:        r.Data.ShipmentStatus.Name,
						Description: r.Data.ShipmentStatus.Description,
						StatusCode:  r.Data.ShipmentStatus.Code,
						UpdatedBy:   r.Data.ShipmentStatus.UpdatedBy,
						UpdateDate:  r.Data.ShipmentStatus.UpdatedDate.String(),
						TrackUrl:    r.Data.ShipmentStatus.TrackUrl,
					},
					Origin: origin{
						ID:           r.Data.Origin.Id,
						Address:      r.Data.Origin.Address,
						Direction:    r.Data.Origin.Direction,
						PostCode:     r.Data.Origin.Postcode,
						StopID:       r.Data.Origin.StopId,
						AreaID:       r.Data.Origin.AreaId,
						SuburbID:     r.Data.Origin.SuburbId,
						AreaName:     r.Data.Origin.AreaName,
						SuburbName:   r.Data.Origin.SuburbName,
						CityID:       r.Data.Origin.CityId,
						CityName:     r.Data.Origin.CityName,
						ProvinceID:   r.Data.Origin.ProvinceId,
						ProvinceName: r.Data.Origin.ProvinceName,
						CountryID:    r.Data.Origin.CountryId,
						CountryName:  r.Data.Origin.CountryName,
						Latitude:     r.Data.Origin.Lat,
						Longitude:    r.Data.Origin.Lng,
					},
					Destination: origin{
						ID:           r.Data.Destination.Id,
						Address:      r.Data.Destination.Address,
						Direction:    r.Data.Destination.Direction,
						PostCode:     r.Data.Destination.Postcode,
						StopID:       r.Data.Destination.StopId,
						AreaID:       r.Data.Destination.AreaId,
						SuburbID:     r.Data.Destination.SuburbId,
						AreaName:     r.Data.Destination.AreaName,
						SuburbName:   r.Data.Destination.SuburbName,
						CityID:       r.Data.Destination.CityId,
						CityName:     r.Data.Destination.CityName,
						ProvinceID:   r.Data.Destination.ProvinceId,
						ProvinceName: r.Data.Destination.ProvinceName,
						CountryID:    r.Data.Destination.CountryId,
						CountryName:  r.Data.Destination.CountryName,
						Latitude:     r.Data.Destination.Lat,
						Longitude:    r.Data.Destination.Lng,
					},
					Package: struct {
						ItemType  string `json:"itemType"`
						Contents  string `json:"contents"`
						Price     item   `json:"price"`
						ItemName  int    `json:"itemName"`
						Dimension struct {
							Length item `json:"length"`
							Width  item `json:"width"`
							Height item `json:"height"`
						} `json:"dimension"`
						Weight        item            `json:"weight"`
						CubicalWeight item            `json:"cubicalWeight"`
						Fragile       int             `json:"fragile"`
						Type          int             `json:"type"`
						IsConfirmed   int             `json:"isConfirmed"`
						PictureURL    string          `json:"pictureURL"`
						Details       []PackageDetail `json:"details"`
					}{
						Dimension: struct {
							Length item `json:"length"`
							Width  item `json:"width"`
							Height item `json:"height"`
						}{
							Length: item{
								Value: r.Data.Package.Length,
							},
							Width: item{
								Value: r.Data.Package.Width,
							},
							Height: item{
								Value: r.Data.Package.Height,
							},
						},
						Weight: item{
							Value: r.Data.Package.Weight,
						},
						CubicalWeight: item{
							Value: r.Data.Package.VolumeWeight,
						},
						PictureURL: "",
						Details:    packageDetail,
					},
					Driver: struct {
						ID                 int    `json:"id"`
						Name               string `json:"name"`
						PhoneNumber        string `json:"phoneNumber"`
						VehicleType        string `json:"vehicleType"`
						VehicleNumber      string `json:"vehicleNumber"`
						IsPaymentCollected int    `json:"isPaymentCollected"`
						Feedback           struct {
							Score   int    `json:"score"`
							Comment string `json:"comment"`
						} `json:"feedback"`
						AgentID     int    `json:"agentID"`
						AgentName   string `json:"agentName"`
						AgentCityID int    `json:"agentCityID"`
					}{
						ID:            0,
						Name:          r.Data.Driver.Name,
						PhoneNumber:   r.Data.Driver.Phone,
						VehicleType:   r.Data.Driver.VehicleType,
						VehicleNumber: r.Data.Driver.VehicleNumber,
					},
					Courier: struct {
						ID           int    `json:"id"`
						Name         string `json:"name"`
						RateID       int    `json:"rate_id"`
						RateName     string `json:"rate_name"`
						ShipmentType int    `json:"shipmentType"`
						ActualID     int    `json:"actualID"`
						MinDay       int    `json:"min_day"`
						MaxDay       int    `json:"max_day"`
						Rate         item   `json:"rate"`
						ActualRate   struct {
							ID int `json:"id"`
							item
						} `json:"actualRate"`
					}{
						ID:       0,
						Name:     r.Data.Courier.Name,
						RateID:   r.Data.Courier.RateId,
						RateName: r.Data.Courier.RateName,
						MinDay:   r.Data.Courier.MinDay,
						MaxDay:   r.Data.Courier.MaxDay,
						Rate: item{
							Value: r.Data.Courier.Amount,
						},
					},
					Rates: struct {
						Shipment        item `json:"shipment"`
						PaidShipment    item `json:"paidShipment"`
						ActualShipment  item `json:"actualShipment"`
						Insurance       item `json:"insurance"`
						PaidInsurance   item `json:"paidInsurance"`
						ActualInsurance item `json:"actualInsurance"`
						EscrowCost      item `json:"escrowCost"`
						FulfillmentCost item `json:"fulfillmentCost"`
						ItemPrice       item `json:"itemPrice"`
						Discount        item `json:"discount"`
					}{
						Shipment: item{Value: r.Data.Courier.Amount},
						Insurance: item{
							Value: r.Data.Courier.InsuranceAmount,
						},
						EscrowCost: item{
							Value: r.Data.Courier.PriceBreakdown.SurchargeFee,
						},
						ActualInsurance: item{
							Value: r.Data.Courier.PriceBreakdown.InsuranceFee,
						},
						Discount: item{
							Value: r.Data.Courier.PriceBreakdown.Discount,
						},
					},
					UseInsurance:    r.Data.Courier.UseInsurance,
					PaymentType:     r.Data.PaymentType,
					IsActive:        r.Data.IsActive,
					PickUpTime:      r.Data.PickupTime,
					PickUpCode:      r.Data.PickupCode,
					CreationDate:    r.Data.CreationDate.String(),
					LastUpdatedDate: r.Data.LastUpdatedDate.String(),
				},
			},
			StatusCode: r.Metadata.HttpStatusCode,
		},
	}
}

// UpdatedOrder struct contains response from API UpdateOrder.
type UpdatedOrder struct {
	Status string `json:"status"`
	Data   struct {
		Title           string `json:"title"`
		Content         string `json:"content"`
		CorrectedFields struct {
			Weight              float64 `json:"weight"`
			VolumeWeight        float64 `json:"volumeWeight"`
			Length              float64 `json:"length"`
			Height              float64 `json:"height"`
			Width               float64 `json:"width"`
			CompulsoryInsurance int     `json:"compulsoryInsurance"`
			Insurance           int     `json:"insurance"`
			FinalRate           int     `json:"finalRate"`
		} `json:"correctedFields"`
		StatusCode int `json:"statusCode"`
	} `json:"data"`
}

// UpdateOrderParams struct contains request parameter for API UpdateOrder.
type UpdateOrderParams struct {
	Length      float64 `json:"l" validate:"required"`
	Width       float64 `json:"w" validate:"required"`
	Height      float64 `json:"h" validate:"required"`
	WeightTotal float64 `json:"wt" validate:"required"`
}

// ItemName struct contains list of items for API CreateOrder.
type ItemName struct {
	Name  string `json:"name"`
	Qty   int    `json:"qty"`
	Value int    `json:"value"`
}

// item struct contains item description for API GetOrderDetail.
type item struct {
	Value float64 `json:"value"`
	UOM   string  `json:"UoM"`
}

// origin struct contains data of origin & destination for API GetOrderDetail.
type origin struct {
	ID           int    `json:"id"`
	Address      string `json:"address"`
	Direction    string `json:"direction"`
	PostCode     string `json:"postcode"`
	StopID       int    `json:"stopID"`
	AreaID       int    `json:"areaID"`
	SuburbID     int    `json:"suburbID"`
	AreaName     string `json:"areaName"`
	SuburbName   string `json:"suburbName"`
	CityID       int    `json:"cityID"`
	CityName     string `json:"cityName"`
	ProvinceID   int    `json:"provinceID"`
	ProvinceName string `json:"provinceName"`
	CountryID    int    `json:"countryID"`
	CountryName  string `json:"countryName"`
	Latitude     string `json:"lat"`
	Longitude    string `json:"lng"`
}

// shipmentStatus struct contains description of shipment status for API GetOrderDetail.
type shipmentStatus struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StatusCode  int    `json:"statusCode"`
	UpdatedBy   string `json:"updatedBy"`
	UpdateDate  string `json:"updateDate"`
	IsCanceled  int    `json:"isCanceled"`
	TrackUrl    string `json:"track_url"`
}

// tracking struct contains list of tracking data for API GetOrderDetail.
type tracking struct {
	ID             string        `json:"id"`
	OrderID        string        `json:"orderID"`
	UniqueID       string        `json:"uniqueID"`
	TrackStatus    trackStatus   `json:"trackStatus"`
	InternalStatus trackStatus   `json:"internalStatus"`
	ExternalStatus trackStatus   `json:"externalStatus"`
	LogisticStatus []trackStatus `json:"logisticStatus"`
	CreatedBy      string        `json:"createdBy"`
	CreatedDate    string        `json:"createdDate"`
}

// trackStatus struct contains data of tracking status for tracking struct.
type trackStatus struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
