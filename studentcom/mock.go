package studentcom

import (
	"net/url"
)

func mockHttpGetList(url url.URL) []byte {
	return []byte(`{
  "data": {
    "page_info": {
      "has_previous_page": true,
      "has_next_page": false,
      "num_pages": 10
    },
    "properties": [
      {
        "unit_types": [
          {
            "floorplan_area": 100.5,
            "bedroom_count_min": 1,
            "max_occupancy": 5,
            "kitchen_count": 1,
            "bedroom_count_max": 3,
            "bed_count": 3,
            "listings": [
              {
                "available_to": "2019-01-01",
                "min_duration": 100,
                "discount_value": "100.50",
                "id": "eyJ0eXBlIjoiUHJvcGVydHkiLCJpZCI6Mjc2fQ==",
                "move_out": "2019-01-01",
                "availability": "good",
                "type": "fixed",
                "move_in": "2018-01-01",
                "price_min": "400.50",
                "discount_type": "percentage",
                "available_from": "2018-01-01",
                "price_max": "400.50"
              }
            ],
            "room_arrangement": "cluster",
            "name": "deluxe suite",
            "id": "eyJ0eXBlIjoiUHJvcGVydHkiLCJpZCI6Mjc2fQ==",
            "category": "shared_room",
            "bathroom_count": 3,
            "floorplan_area_display_unit": "sqft"
          }
        ],
        "total_beds": 100,
        "country": {
          "name": "United Kingdom",
          "id": "eyJ0eXBlIjoiUHJvcGVydHkiLCJpZCI6Mjc2fQ==",
          "slug": "uk"
        },
        "images": [
          {
            "is_hero_image": true,
            "url": "https://cdn.student.com/media/cache/property_hero_banner_desktop/mstr/country/united-kingdom/city/london/property/urbanest-kings-cross/image-p5zimy.jpeg"
          }
        ],
        "payment_deposit_rules": [
          {
            "name": "Security Deposit",
            "type": "fixed_amount",
            "value": "100.23"
          }
        ],
        "address": "Canal Reach, Kings Cross, London",
        "city": {
          "name": "London",
          "id": "eyJ0eXBlIjoiUHJvcGVydHkiLCJpZCI6Mjc2fQ==",
          "slug": "london"
        },
        "description": "Grab a luxe pad in up-and-coming trendy hotspot Kings Cross a perfect entry point for experiencing London",
        "post_code": "N1C 4BE",
        "name": "Urbanest King's Cross",
        "currency": "gbp",
        "id": "eyJ0eXBlIjoiUHJvcGVydHkiLCJpZCI6Mjc2fQ==",
        "billing_cycle": "weekly",
        "facilities": [
          {
            "name": "Fully Furnished",
            "slug": "fully_furnished",
            "tags": [
              "amenity"
            ]
          }
        ]
      }
    ]
  }
}`)
}
