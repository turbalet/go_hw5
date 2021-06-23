package decode

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
	"testing"
)

func TestIntVal_UnmarshalJSON(t *testing.T) {
	testTable := []struct {
		data     []byte
		expected []byte
	}{
		{
			data: []byte(`[
					 {
					   "id": 1,
					   "address": {
						 "city_id": 5,
						 "street": "Satbayev"
					   },
					   "Age": 20
					 },
					 {
					   "id": 1,
					   "address": {
						 "city_id": "6",
						 "street": "Al-Farabi"
					   },
					   "Age": "32"
					 }
					]`),
			expected: []byte(`[
					 {"id": 1,
					   "address": {
						 "city_id": 5,
						 "street": "Satbayev"
					   },
					   "Age": 20
					 },
					 {
					   "id": 1,
					   "address": {
						 "city_id": 6,
						 "street": "Al-Farabi"
					   },
					   "Age": 32
					 }
					]`),
		},
	}


	for _, testCase := range testTable {
		var users []User
		var expectedUsers []User
		if err := json.Unmarshal(testCase.data, &users); err != nil {
			t.Error(err)
		}
		if err := json.Unmarshal(testCase.expected, &expectedUsers); err != nil {
			t.Error(err)
		}

		for _, u := range users {
			fmt.Println(u)
		}

		if !reflect.DeepEqual(users, expectedUsers) {
			t.Errorf("Incorrect result. Expect %v got %v", testCase.expected, testCase.data)
		}
	}

}

func TestIntVal_UnmarshalXML(t *testing.T) {
	testTable := []struct {
		data     []byte
		expected []byte
	}{
		{
			data: []byte(`
					<users>
						<user>
							<id>1</id>
							<address>
								<city_id>5</city_id>
								<street>"Satbayev"</street>
							</address>
							<age>20</age>
						</user>
						<user>
							<id>1</id>
							<address>
								<city_id>"6"</city_id>
								<street>"Al-Farabi"</street>
							</address>
							<age>"32"</age>
						</user>
					</users>
				`),
			expected: []byte(`
					<users>
						<user>
							<id>1</id>
							<address>
								<city_id>5</city_id>
								<street>"Satbayev"</street>
							</address>
							<age>20</age>
						</user>
						<user>
							<id>1</id>
							<address>
								<city_id>6</city_id>
								<street>"Al-Farabi"</street>
							</address>
							<age>32</age>
						</user>
					</users>
				`),
		},
	}
// problems with xml
	for _, testCase := range testTable {
		var users Users
		var expectedUsers Users
		if err := xml.Unmarshal(testCase.data, &users); err != nil {
			t.Error(err)
		}
		if err := xml.Unmarshal(testCase.expected, &expectedUsers); err != nil {
			t.Error(err)
		}

		fmt.Println("XML: ")
		for _, u := range users.Users {
			fmt.Println(u)
		}

		if !reflect.DeepEqual(users, expectedUsers) {
			t.Errorf("Incorrect result. Expect %v got %v", expectedUsers, users)
		}
	}


}
