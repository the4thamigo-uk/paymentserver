#!/bin/bash
curl -X GET http://127.0.0.1:8080/payments/$1/$2 | jq