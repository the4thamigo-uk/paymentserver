#!/bin/bash
curl -X POST -d @payment.json http://127.0.0.1:8080/payments | jq
