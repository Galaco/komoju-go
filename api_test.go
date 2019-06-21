package komoju

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestApiResponseError_Error(t *testing.T) {
	err := NewApiResponseError("401", "foo", "bar")
	if err == nil {
		t.Error("failed to create error response from parameters")
	}
	if err != nil && err.Error() != "Code: 401, Param: foo, Message: bar" {
		t.Error("incorrect error message returned")
	}
}

func TestNewApiResponseError(t *testing.T) {
	err := NewApiResponseError("401", "foo", "bar")
	if err == nil {
		t.Error("failed to create error response from parameters")
	}
}

func TestParseApiResponse(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte("MALFORMED{")))
	body := ApiResponseEvents{}
	err := ParseApiResponse(r, &body)
	if err == nil {
		t.Error("expected error from malformed response, but got none")
	}
}

func TestParseApiResponseError(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader(getResponseErrorMessage()))
	body := ApiResponseError{}
	err := ParseApiResponse(r, &body)
	if err != nil {
		t.Error("failed to parse error response message")
	}
}

func TestParseApiResponseGetEvents(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader(getEventsSampleResponse()))
	body := ApiResponseEvents{}
	err := ParseApiResponse(r, &body)
	if err != nil {
		t.Error(err)
	}
}

func TestParseApiResponseGetEvent(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader(getEventSampleResponse()))
	body := ApiResponseEvent{}
	err := ParseApiResponse(r, &body)
	if err != nil {
		t.Error(err)
	}
}

func getEventsSampleResponse() []byte {
	return []byte(`{
  "resource": "list",
  "total": 2,
  "page": 1,
  "per_page": 10,
  "last_page": 1,
  "data": [
    {
      "id": "4xwf1falwkeagr4mn0gg6n19d",
      "type": "payment.captured",
      "resource": "event",
      "data": {
        "id": "dvxisyithua801n3vyfb4q0oe",
        "resource": "payment",
        "status": "captured",
        "amount": 300,
        "tax": 30,
        "customer": null,
        "payment_deadline": "2018-11-20T14:59:59Z",
        "payment_details": {
          "type": "credit_card",
          "email": "gordon@example.com",
          "brand": "visa",
          "last_four_digits": "1111",
          "month": 3,
          "year": 2019
        },
        "payment_method_fee": 0,
        "total": 330,
        "currency": "JPY",
        "description": null,
        "captured_at": "2018-11-11T14:59:59Z",
        "external_order_num": "ORDER6",
        "metadata": {
        },
        "created_at": "2018-11-13T06:20:21Z",
        "amount_refunded": 0,
        "locale": "en",
        "refunds": [

        ],
        "refund_requests": [

        ]
      },
      "created_at": "2018-11-13T06:20:21Z"
    },
    {
      "id": "erwjx2l8u1ff38pxw2jvtgc0m",
      "type": "payment.updated",
      "resource": "event",
      "data": {
        "id": "dvxisyithua801n3vyfb4q0oe",
        "resource": "payment",
        "status": "captured",
        "amount": 300,
        "tax": 30,
        "customer": null,
        "payment_deadline": "2018-11-20T14:59:59Z",
        "payment_details": {
          "type": "credit_card",
          "email": "gordon@example.com",
          "brand": "visa",
          "last_four_digits": "1111",
          "month": 3,
          "year": 2019
        },
        "payment_method_fee": 0,
        "total": 330,
        "currency": "JPY",
        "description": null,
        "captured_at": "2018-11-11T14:59:59Z",
        "external_order_num": "ORDER6",
        "metadata": {
        },
        "created_at": "2018-11-13T06:20:21Z",
        "amount_refunded": 0,
        "locale": "en",
        "refunds": [

        ],
        "refund_requests": [

        ]
      },
      "created_at": "2018-11-13T06:20:21Z"
    }
  ]
}`)
}

func getEventSampleResponse() []byte {
	return []byte(`{
  "id": "bmon6b8awbd5wftxlpjxzpjte",
  "type": "payment.captured",
  "resource": "event",
  "data": {
    "id": "2g1ret84tfvgi7aii057prbeo",
    "resource": "payment",
    "status": "captured",
    "amount": 300,
    "tax": 30,
    "customer": null,
    "payment_deadline": "2018-11-20T14:59:59Z",
    "payment_details": {
      "type": "credit_card",
      "email": "gordon@example.com",
      "brand": "visa",
      "last_four_digits": "1111",
      "month": 3,
      "year": 2019
    },
    "payment_method_fee": 0,
    "total": 330,
    "currency": "JPY",
    "description": null,
    "captured_at": "2018-11-11T14:59:59Z",
    "external_order_num": "ORDER5",
    "metadata": {
    },
    "created_at": "2018-11-13T06:20:21Z",
    "amount_refunded": 0,
    "locale": "en",
    "refunds": [

    ],
    "refund_requests": [

    ]
  },
  "created_at": "2018-11-13T06:20:21Z"
}`)
}

func getResponseErrorMessage() []byte {
	return []byte(`{
  "error": {
     "message": "A required parameter (amount) is missing",
     "code": "missing_parameter",
     "param": "amount"
  }
}`)
}
