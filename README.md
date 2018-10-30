# paymentserver

[![Build Status](https://travis-ci.org/the4thamigo-uk/paymentserver.svg?branch=master)](https://travis-ci.org/the4thamigo-uk/paymentserver?branch=master)
[![Coverage Status](https://coveralls.io/repos/the4thamigo-uk/paymentserver/badge.svg?branch=master&service=github)](https://coveralls.io/github/the4thamigo-uk/paymentserver?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/the4thamigo-uk/paymentserver)](https://goreportcard.com/report/github.com/the4thamigo-uk/paymentserver)
[![Godoc](https://godoc.org/github.com/the4thamigo-uk/paymentserver?status.svg)](https://godoc.org/github.com/the4thamigo-uk/paymentserver)

## Description

This is an example project demonstrating a REST API for a server that allows CRUD operations on payment records. 

The following are the key limitations, but please see the issues list for known
problems and planned improvements :

- The server currently stores data in a simple in-memory datastore implementation, so the server is not stateless.
- There is currently no support for authentication or user identity, so all payments are public.
- BDD test suite is not comprehensive, and mainly for illustrative purposes.
- Missing unit tests for ./pkg/server, however the other packages are fairly well covered and some of the server
behaviour is covered by the BDD tests.

## Getting Started

To run unit tests :

    go test -v ./pkg/...

To build :

    go build ./cmd/paymentserver

To run BDD tests, using the specified server executable and listening address, do the following (requires build step):

    go test -v ./cmd/bdd -server ../../paymentserver -address :8080

Start the server with :

    ./paymentserver -l :8080

Then you can use the test scripts :

    cd ./cmd/paymentserver/scripts

To create a payment : 
  
    ./create.sh

or :

    curl -X POST -d @payment.json http://127.0.0.1:8080/payments
  
which both return json of the form :

    {
      "data": [
        {
          "type": "Payment",
          "id": "b56df16e-a310-4762-854e-e672bffac14d",
          "version": 1,
          "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
          "attributes": {
            "amount": "100.21",
            "beneficiary_party": {
              "account_name": "W Owens",
              "account_number": "31926819",
              "account_number_code": "BBAN",
              "account_type": 0,
              "address": "1 The Beneficiary Localtown SE2",
              "bank_id": "403000",
              "bank_id_code": "GBDSC",
              "name": "Wilfred Jeremiah Owens"
            },
            "charges_information": {
              "bearer_code": "SHAR",
              "sender_charges": [
                {
                  "amount": "5.00",
                  "currency": "GBP"
                },
                {
                  "amount": "10.00",
                  "currency": "USD"
                }
              ],
              "receiver_charges_amount": "1.00",
              "receiver_charges_currency": "USD"
            },
            "currency": "GBP",
            "debtor_party": {
              "account_name": "EJ Brown Black",
              "account_number": "BG18RZBB91550123456789",
              "account_number_code": "IBAN",
              "address": "10 Debtor Crescent Sourcetown NE1",
              "bank_id": "203301",
              "bank_id_code": "GBDSC",
              "name": "Emelia Jane Brown"
            },
            "end_to_end_reference": "Wil piano Jan",
            "fx": {
              "contract_reference": "FX123",
              "exchange_rate": "2.00000",
              "original_amount": "200.42",
              "original_currency": "USD"
            },
            "numeric_reference": "1002001",
            "payment_id": "123456789012345678",
            "payment_purpose": "Paying for goods/services",
            "payment_scheme": "FPS",
            "payment_type": "Credit",
            "processing_date": "2017-01-18",
            "reference": "Payment for Em's piano lessons",
            "scheme_payment_sub_type": "InternetBanking",
            "scheme_payment_type": "ImmediatePayment",
            "sponsor_party": {
              "account_number": "56781234",
              "bank_id": "123123",
              "bank_id_code": "GBDSC"
            }
          }
        }
      ],
      "_links": {
        "index": {
          "title": "Payments server",
          "href": "/",
          "method": "GET"
        },
        "payment:create": {
          "title": "Create payment",
          "href": "/payments",
          "method": "POST"
        },
        "payment:delete": {
          "title": "Delete a payment",
          "href": "/payments/b56df16e-a310-4762-854e-e672bffac14d/1",
          "method": "DELETE"
        },
        "payment:laod": {
          "title": "Load payment",
          "href": "/payments/b56df16e-a310-4762-854e-e672bffac14d/1",
          "method": "GET"
        },
        "payment:list": {
          "title": "List payments",
          "href": "/payments",
          "method": "GET"
        },
        "payment:save": {
          "title": "Save a payment",
          "href": "/payments/b56df16e-a310-4762-854e-e672bffac14d/1",
          "method": "PUT"
        },
        "payment:update": {
          "title": "Update a payment",
          "href": "/payments/b56df16e-a310-4762-854e-e672bffac14d/1",
          "method": "PATCH"
        },
        "self": {
          "title": "Create payment",
          "href": "/payments",
          "method": "POST"
        }
      }
    }

The embedded [HAL style](http://stateless.co/hal_specification.html) links help the client to navigate to other functions.

To run other scripts you need to pass the id for the returned resource, and the current version of the data : 

    ./update.sh b56df16e-a310-4762-854e-e672bffac14d 0

Note that passing version 0 will always update the latest version of the resource, but passing a version number that
does not match the version in the store will error, thereby implementing an optimistic offline-lock :

    ./update.sh b56df16e-a310-4762-854e-e672bffac14d 42
    {
      "error": {
        "message": "Error saving object 'b56df16e-a310-4762-854e-e672bffac14d'. Expected version 42, actual version 1",
        "code": 400
      }
    }
