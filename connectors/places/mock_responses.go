package places

import "encoding/json"

func getErrorMock() *AutocompleteResponse {
	var errorAnswerMock = &AutocompleteResponse{}
	json.Unmarshal([]byte(`{
     "error_message" : "You have exceeded your daily request quota for this API. If you did not set a custom daily request quota, verify your project has an active billing account: http://g.co/dev/maps-no-account",
     "predictions" : [],
     "status" : "OVER_QUERY_LIMIT"
  }`), errorAnswerMock)
	return errorAnswerMock
}

func getCitiesMock() *AutocompleteResponse {
	var citiesMock = &AutocompleteResponse{}
	json.Unmarshal([]byte(`{
     "predictions" : [
        {
           "description" : "Tel Aviv-Yafo, Israel",
           "id" : "78708c592bddf3a50b3a14a24024728d4fa8a31d",
           "matched_substrings" : [
              {
                 "length" : 3,
                 "offset" : 0
              }
           ],
           "place_id" : "ChIJH3w7GaZMHRURkD-WwKJy-8E",
           "reference" : "ChIJH3w7GaZMHRURkD-WwKJy-8E",
           "structured_formatting" : {
              "main_text" : "Tel Aviv-Yafo",
              "main_text_matched_substrings" : [
                 {
                    "length" : 3,
                    "offset" : 0
                 }
              ],
              "secondary_text" : "Israel"
           },
           "terms" : [
              {
                 "offset" : 0,
                 "value" : "Tel Aviv-Yafo"
              },
              {
                 "offset" : 15,
                 "value" : "Israel"
              }
           ],
           "types" : [ "locality", "political", "geocode" ]
        },
        {
           "description" : "Telford, UK",
           "id" : "3dc75343a58b2129a4b511abedb33ebbbbf222ac",
           "matched_substrings" : [
              {
                 "length" : 3,
                 "offset" : 0
              }
           ],
           "place_id" : "ChIJQ3JXDQt4cEgREYhv0ILtOZE",
           "reference" : "ChIJQ3JXDQt4cEgREYhv0ILtOZE",
           "structured_formatting" : {
              "main_text" : "Telford",
              "main_text_matched_substrings" : [
                 {
                    "length" : 3,
                    "offset" : 0
                 }
              ],
              "secondary_text" : "UK"
           },
           "terms" : [
              {
                 "offset" : 0,
                 "value" : "Telford"
              },
              {
                 "offset" : 9,
                 "value" : "UK"
              }
           ],
           "types" : [ "locality", "political", "geocode" ]
        },
        {
           "description" : "Telluride, CO, USA",
           "id" : "9e99af634dcf9072c6dc05a63961dfe418c4d70a",
           "matched_substrings" : [
              {
                 "length" : 3,
                 "offset" : 0
              }
           ],
           "place_id" : "ChIJc_TmcHvYPocR4eO6cSF37jg",
           "reference" : "ChIJc_TmcHvYPocR4eO6cSF37jg",
           "structured_formatting" : {
              "main_text" : "Telluride",
              "main_text_matched_substrings" : [
                 {
                    "length" : 3,
                    "offset" : 0
                 }
              ],
              "secondary_text" : "CO, USA"
           },
           "terms" : [
              {
                 "offset" : 0,
                 "value" : "Telluride"
              },
              {
                 "offset" : 11,
                 "value" : "CO"
              },
              {
                 "offset" : 15,
                 "value" : "USA"
              }
           ],
           "types" : [ "locality", "political", "geocode" ]
        },
        {
           "description" : "Teluk Intan, Perak, Malaysia",
           "id" : "237b85a204682b9ccf87a13c167a38ef9da284f0",
           "matched_substrings" : [
              {
                 "length" : 3,
                 "offset" : 0
              }
           ],
           "place_id" : "ChIJ5ZfdUiIVyzER6EvOk2YaQzM",
           "reference" : "ChIJ5ZfdUiIVyzER6EvOk2YaQzM",
           "structured_formatting" : {
              "main_text" : "Teluk Intan",
              "main_text_matched_substrings" : [
                 {
                    "length" : 3,
                    "offset" : 0
                 }
              ],
              "secondary_text" : "Perak, Malaysia"
           },
           "terms" : [
              {
                 "offset" : 0,
                 "value" : "Teluk Intan"
              },
              {
                 "offset" : 13,
                 "value" : "Perak"
              },
              {
                 "offset" : 20,
                 "value" : "Malaysia"
              }
           ],
           "types" : [ "locality", "political", "geocode" ]
        },
        {
           "description" : "Teltow, Germany",
           "id" : "30e2db089f2976c9deb289de3271ba06621b0c99",
           "matched_substrings" : [
              {
                 "length" : 3,
                 "offset" : 0
              }
           ],
           "place_id" : "ChIJ8bxbxZxbqEcRR85bEQV2GVE",
           "reference" : "ChIJ8bxbxZxbqEcRR85bEQV2GVE",
           "structured_formatting" : {
              "main_text" : "Teltow",
              "main_text_matched_substrings" : [
                 {
                    "length" : 3,
                    "offset" : 0
                 }
              ],
              "secondary_text" : "Germany"
           },
           "terms" : [
              {
                 "offset" : 0,
                 "value" : "Teltow"
              },
              {
                 "offset" : 8,
                 "value" : "Germany"
              }
           ],
           "types" : [ "locality", "political", "geocode" ]
        }
     ],
     "status" : "OK"
  }`), citiesMock)
	return citiesMock
}

func getPlaceDetailsMock() *PlaceDetails {
	details := &PlaceDetails{}
	json.Unmarshal([]byte(`{
   "html_attributions" : [],
   "result" : {
      "address_components" : [
         {
            "long_name" : "1",
            "short_name" : "1",
            "types" : [ "street_number" ]
         },
         {
            "long_name" : "Natana Rybaka Lane",
            "short_name" : "Natana Rybaka Ln",
            "types" : [ "route" ]
         },
         {
            "long_name" : "Shepetivka",
            "short_name" : "Shepetivka",
            "types" : [ "locality", "political" ]
         },
         {
            "long_name" : "Shepetivs'ka city council",
            "short_name" : "Shepetivs'ka city council",
            "types" : [ "administrative_area_level_3", "political" ]
         },
         {
            "long_name" : "Khmel'nyts'ka oblast",
            "short_name" : "Khmel'nyts'ka oblast",
            "types" : [ "administrative_area_level_1", "political" ]
         },
         {
            "long_name" : "Ukraine",
            "short_name" : "UA",
            "types" : [ "country", "political" ]
         },
         {
            "long_name" : "30400",
            "short_name" : "30400",
            "types" : [ "postal_code" ]
         }
      ],
      "formatted_address" : "Natana Rybaka Ln, 1, Shepetivka, Khmel'nyts'ka oblast, Ukraine, 30400",
      "geometry" : {
         "location" : {
            "lat" : 50.168899,
            "lng" : 27.0421056
         },
         "viewport" : {
            "northeast" : {
               "lat" : 50.1703102802915,
               "lng" : 27.0433398302915
            },
            "southwest" : {
               "lat" : 50.1676123197085,
               "lng" : 27.0406418697085
            }
         }
      }
   },
   "status" : "OK"
}`), details)
	return details
}
