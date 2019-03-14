package shipping

import "encoding/json"

//GetPricesMock is a mocked response for getting prices
func GetPricesMock(resp *GetPricesRes) error {
	return json.Unmarshal([]byte(`{
    "status": "success",
    "description": null,
    "rates": [
        {
            "id": "19024",
            "supplier_id": "1",
            "currency_symbol": "GBP",
            "currency_name": "United Kingdom Pounds",
            "product": {
                "id": "3",
                "title": "Pack+Unpack - Part Load",
                "description": "A comprehensive service that includes the packing of your items using suitable packing materials. We help to dismantle simpler items of furniture, although we leave the more complex items to you. We then load onto our vehicle and transport to your destination within 7 -10 days, depending on distance. We unload at your second location, unpack furniture and reassemble any items we dismantled at the beginning of the journey. We remove any packing debris.",
                "detailed_description": "Crew to professionally pack all your items using suitable packing materials.\r\nDismantle basic furniture items excluding IKEA and complex items.\r\nCrew collect and load items into vehicle.\r\nItems are taken to one of our consolidation warehouses for loading.\r\nTransport and delivery via our part load service within 7-10 working days of your items being collected.\r\nWe deliver to the ground floor of the property.\r\nCrew will unpack all boxes and place contents onto a flat surface.\r\nUnpack furniture items.\r\nAssemble basic furniture except IKEA and complex items.\r\nRemoval of packing debris.",
                "product_type": "pack_unpack",
                "pack_type": "loose",
                "transit_type": "road",
                "type": "part_load",
                "icon": null
            },
            "supplier": {
                "id": "1",
                "name": "ShippingForce",
                "icon": "https://direct.shippingforce.co.uk/uploads/user/Wo/Updated_SF_Logo.png"
            },
            "breakdown": {
                "inc_margin": 436.563,
                "inc_tax": 523.8756,
                "tax_eligable": 1
            }
        },
        {
            "id": "19344",
            "supplier_id": "1",
            "currency_symbol": "GBP",
            "currency_name": "United Kingdom Pounds",
            "product": {
                "id": "2",
                "title": "Pickup+Go - Part Load",
                "description": "We pick up the boxes you have packed from the ground floor of your first location and load them onto the vehicle. We then take them to the second location and unload again onto the ground floor.",
                "detailed_description": "You pack your belongings and bring them to the ground floor of your property.\r\nCrew load items into vehicle.\r\nWe deliver to the ground floor of the property.",
                "product_type": "pickup_go",
                "pack_type": "no_pack",
                "transit_type": "road",
                "type": "part_load",
                "icon": null
            },
            "supplier": {
                "id": "1",
                "name": "ShippingForce",
                "icon": "https://direct.shippingforce.co.uk/uploads/user/Wo/Updated_SF_Logo.png"
            },
            "breakdown": {
                "inc_margin": 439.6565,
                "inc_tax": 527.5878,
                "tax_eligable": 1
            }
        },
        {
            "id": "19025",
            "supplier_id": "1",
            "currency_symbol": "GBP",
            "currency_name": "United Kingdom Pounds",
            "product": {
                "id": "1",
                "title": "Pack+Go - Part Load",
                "description": "Let us take on some of the pressures of moving by helping to pack your possessions with suitable packing materials. We then load onto our vehicle and deliver to the ground floor of your second location.",
                "detailed_description": "Crew to professionally pack all your items using suitable packing materials.\r\nCrew collect and load items into vehicle.\r\nWe deliver to the ground floor of the property.",
                "product_type": "pack_go",
                "pack_type": "loose",
                "transit_type": "road",
                "type": "part_load",
                "icon": null
            },
            "supplier": {
                "id": "1",
                "name": "ShippingForce",
                "icon": "https://direct.shippingforce.co.uk/uploads/user/Wo/Updated_SF_Logo.png"
            },
            "breakdown": {
                "inc_margin": 501.5725,
                "inc_tax": 601.887,
                "tax_eligable": 1
            }
        },
        {
            "id": "19026",
            "supplier_id": "1",
            "currency_symbol": "GBP",
            "currency_name": "United Kingdom Pounds",
            "product": {
                "id": "5",
                "title": "Pickup+Go - Direct Load",
                "description": "We pick up your fully packed boxes from the ground floor of your first location and load them onto the vehicle. We then take them to the second location and unload again onto the ground floor.",
                "detailed_description": "You pack your belongings and bring them to the ground floor of your property.\r\nCrew load items into vehicle.\r\nWe deliver to the ground floor of the property.",
                "product_type": "pickup_go",
                "pack_type": "no_pack",
                "transit_type": "road",
                "type": "direct_load",
                "icon": null
            },
            "supplier": {
                "id": "1",
                "name": "ShippingForce",
                "icon": "https://direct.shippingforce.co.uk/uploads/user/Wo/Updated_SF_Logo.png"
            },
            "breakdown": {
                "inc_margin": 732.9065,
                "inc_tax": 879.4878,
                "tax_eligable": 1
            }
        },
        {
            "id": "19027",
            "supplier_id": "1",
            "currency_symbol": "GBP",
            "currency_name": "United Kingdom Pounds",
            "product": {
                "id": "4",
                "title": "Pack+Go - Direct Load",
                "description": "Let us take on some of the pressures of moving by helping to pack your possessions with suitable packing materials. We then load onto our vehicle and deliver to the ground floor of your second location.",
                "detailed_description": "Crew to professionally pack all your items using suitable packing materials.\r\nCrew collect and load items into vehicle.\r\nWe deliver to the ground floor of the property.",
                "product_type": "pack_go",
                "pack_type": "part_export_pack",
                "transit_type": "road",
                "type": "direct_load",
                "icon": null
            },
            "supplier": {
                "id": "1",
                "name": "ShippingForce",
                "icon": "https://direct.shippingforce.co.uk/uploads/user/Wo/Updated_SF_Logo.png"
            },
            "breakdown": {
                "inc_margin": 794.8225,
                "inc_tax": 953.787,
                "tax_eligable": 1
            }
        },
        {
            "id": "19028",
            "supplier_id": "1",
            "currency_symbol": "GBP",
            "currency_name": "United Kingdom Pounds",
            "product": {
                "id": "6",
                "title": "Pack+Unpack - Direct Load",
                "description": "A comprehensive service that includes the packing of your items using suitable packing materials. We help to dismantle simpler items of furniture, although we leave the more complex items to you. We then load onto our vehicle and transport to your destination within one to three days, depending on distance. We unload at your second location, unpack furniture and reassemble any items we dismantled at the beginning of the journey. We remove any packing debris",
                "detailed_description": "Crew to professionally pack all your items using suitable packing materials.\r\nDismantle basic furniture items excluding IKEA and complex items.\r\nCrew collect and load items into vehicle.\r\nItems are taken to one of our consolidation warehouses for loading.\r\nWithin 1 to 3 working days (depending on distance) your items reach the destination\r\nWe deliver to the ground floor of the property.\r\nCrew will unpack all boxes and place contents onto a flat surface.\r\nUnpack furniture items.\r\nAssemble basic furniture except IKEA and complex items.\r\nRemoval of packing debris.",
                "product_type": "pack_unpack",
                "pack_type": "part_export_pack",
                "transit_type": "road",
                "type": "direct_load",
                "icon": null
            },
            "supplier": {
                "id": "1",
                "name": "ShippingForce",
                "icon": "https://direct.shippingforce.co.uk/uploads/user/Wo/Updated_SF_Logo.png"
            },
            "breakdown": {
                "inc_margin": 850.563,
                "inc_tax": 1020.6756,
                "tax_eligable": 1
            }
        },
        {
            "id": "19029",
            "supplier_id": 2,
            "currency_symbol": "GBP",
            "currency_name": "United Kingdom Pounds",
            "product": {
                "id": "3",
                "title": "Pack+Unpack - Part Load",
                "description": "A comprehensive service that includes the packing of your items using suitable packing materials. We help to dismantle simpler items of furniture, although we leave the more complex items to you. We then load onto our vehicle and transport to your destination within 7 -10 days, depending on distance. We unload at your second location, unpack furniture and reassemble any items we dismantled at the beginning of the journey. We remove any packing debris.",
                "detailed_description": "Crew to professionally pack all your items using suitable packing materials.\r\nDismantle basic furniture items excluding IKEA and complex items.\r\nCrew collect and load items into vehicle.\r\nItems are taken to one of our consolidation warehouses for loading.\r\nTransport and delivery via our part load service within 7-10 working days of your items being collected.\r\nWe deliver to the ground floor of the property.\r\nCrew will unpack all boxes and place contents onto a flat surface.\r\nUnpack furniture items.\r\nAssemble basic furniture except IKEA and complex items.\r\nRemoval of packing debris.",
                "product_type": "pack_unpack",
                "pack_type": "loose",
                "transit_type": "road",
                "type": "part_load",
                "icon": null
            },
            "breakdown": {
                "inc_margin": 480.2193,
                "inc_tax": 576.26316,
                "tax_eligable": 1
            },
            "supplier": {
                "id": "2",
                "name": "BaggageHub",
                "icon": "https://www.baggagehub.com/wp-content/themes/absolute/images/misc/baggage-hub.png"
            }
        },
        {
            "id": "19030",
            "supplier_id": 2,
            "currency_symbol": "GBP",
            "currency_name": "United Kingdom Pounds",
            "product": {
                "id": "2",
                "title": "Pickup+Go - Part Load",
                "description": "We pick up the boxes you have packed from the ground floor of your first location and load them onto the vehicle. We then take them to the second location and unload again onto the ground floor.",
                "detailed_description": "You pack your belongings and bring them to the ground floor of your property.\r\nCrew load items into vehicle.\r\nWe deliver to the ground floor of the property.",
                "product_type": "pickup_go",
                "pack_type": "no_pack",
                "transit_type": "road",
                "type": "part_load",
                "icon": null
            },
            "breakdown": {
                "inc_margin": 483.62215,
                "inc_tax": 580.34658,
                "tax_eligable": 1
            },
            "supplier": {
                "id": "2",
                "name": "BaggageHub",
                "icon": "https://www.baggagehub.com/wp-content/themes/absolute/images/misc/baggage-hub.png"
            }
        },
        {
            "id": "19031",
            "supplier_id": 2,
            "currency_symbol": "GBP",
            "currency_name": "United Kingdom Pounds",
            "product": {
                "id": "1",
                "title": "Pack+Go - Part Load",
                "description": "Let us take on some of the pressures of moving by helping to pack your possessions with suitable packing materials. We then load onto our vehicle and deliver to the ground floor of your second location.",
                "detailed_description": "Crew to professionally pack all your items using suitable packing materials.\r\nCrew collect and load items into vehicle.\r\nWe deliver to the ground floor of the property.",
                "product_type": "pack_go",
                "pack_type": "loose",
                "transit_type": "road",
                "type": "part_load",
                "icon": null
            },
            "breakdown": {
                "inc_margin": 551.72975,
                "inc_tax": 662.0757,
                "tax_eligable": 1
            },
            "supplier": {
                "id": "2",
                "name": "BaggageHub",
                "icon": "https://www.baggagehub.com/wp-content/themes/absolute/images/misc/baggage-hub.png"
            }
        },
        {
            "id": "19032",
            "supplier_id": 2,
            "currency_symbol": "GBP",
            "currency_name": "United Kingdom Pounds",
            "product": {
                "id": "5",
                "title": "Pickup+Go - Direct Load",
                "description": "We pick up your fully packed boxes from the ground floor of your first location and load them onto the vehicle. We then take them to the second location and unload again onto the ground floor.",
                "detailed_description": "You pack your belongings and bring them to the ground floor of your property.\r\nCrew load items into vehicle.\r\nWe deliver to the ground floor of the property.",
                "product_type": "pickup_go",
                "pack_type": "no_pack",
                "transit_type": "road",
                "type": "direct_load",
                "icon": null
            },
            "breakdown": {
                "inc_margin": 806.19715,
                "inc_tax": 967.43658,
                "tax_eligable": 1
            },
            "supplier": {
                "id": "2",
                "name": "BaggageHub",
                "icon": "https://www.baggagehub.com/wp-content/themes/absolute/images/misc/baggage-hub.png"
            }
        },
        {
            "id": "19033",
            "supplier_id": 2,
            "currency_symbol": "GBP",
            "currency_name": "United Kingdom Pounds",
            "product": {
                "id": "4",
                "title": "Pack+Go - Direct Load",
                "description": "Let us take on some of the pressures of moving by helping to pack your possessions with suitable packing materials. We then load onto our vehicle and deliver to the ground floor of your second location.",
                "detailed_description": "Crew to professionally pack all your items using suitable packing materials.\r\nCrew collect and load items into vehicle.\r\nWe deliver to the ground floor of the property.",
                "product_type": "pack_go",
                "pack_type": "part_export_pack",
                "transit_type": "road",
                "type": "direct_load",
                "icon": null
            },
            "breakdown": {
                "inc_margin": 874.30475,
                "inc_tax": 1049.1657,
                "tax_eligable": 1
            },
            "supplier": {
                "id": "2",
                "name": "BaggageHub",
                "icon": "https://www.baggagehub.com/wp-content/themes/absolute/images/misc/baggage-hub.png"
            }
        },
        {
            "id": "19034",
            "supplier_id": 2,
            "currency_symbol": "GBP",
            "currency_name": "United Kingdom Pounds",
            "product": {
                "id": "6",
                "title": "Pack+Unpack - Direct Load",
                "description": "A comprehensive service that includes the packing of your items using suitable packing materials. We help to dismantle simpler items of furniture, although we leave the more complex items to you. We then load onto our vehicle and transport to your destination within one to three days, depending on distance. We unload at your second location, unpack furniture and reassemble any items we dismantled at the beginning of the journey. We remove any packing debris",
                "detailed_description": "Crew to professionally pack all your items using suitable packing materials.\r\nDismantle basic furniture items excluding IKEA and complex items.\r\nCrew collect and load items into vehicle.\r\nItems are taken to one of our consolidation warehouses for loading.\r\nWithin 1 to 3 working days (depending on distance) your items reach the destination\r\nWe deliver to the ground floor of the property.\r\nCrew will unpack all boxes and place contents onto a flat surface.\r\nUnpack furniture items.\r\nAssemble basic furniture except IKEA and complex items.\r\nRemoval of packing debris.",
                "product_type": "pack_unpack",
                "pack_type": "part_export_pack",
                "transit_type": "road",
                "type": "direct_load",
                "icon": null
            },
            "breakdown": {
                "inc_margin": 935.6193,
                "inc_tax": 1122.74316,
                "tax_eligable": 1
            },
            "supplier": {
                "id": "2",
                "name": "BaggageHub",
                "icon": "https://www.baggagehub.com/wp-content/themes/absolute/images/misc/baggage-hub.png"
            }
        },
        {
            "id": "19035",
            "product": {
                "id": "449508",
                "type": "shippingforce",
                "title": "Door to Door Courier Service",
                "description": "Service Includes Empty Box Delivery, Packed Box Collection, Delivery",
                "delivery_description": "Typical Delivery 7-10 days"
            },
            "currency_symbol": "GBP",
            "breakdown": {
                "inc_tax": 377.46,
                "inc_margin": 314.55
            }
        }
    ],
    "exceptions": [],
    "cant_quote": false,
    "id": "njz8SOzPgaDVdmPhKa7zyw",
    "volume": 37.4335465668,
    "volume_unit": "cuft",
    "converted_volume": 243.3180526842,
    "converted_volume_unit": "ibs",
    "origin": {
        "address": "London, UK",
        "lat": 51.5073509,
        "lng": -0.12775829999998,
        "country": "GB"
    },
    "destination": {
        "address": "Manchester, UK",
        "lat": 53.4807593,
        "lng": -2.2426305,
        "country": "GB"
    }
}`), resp)
}
