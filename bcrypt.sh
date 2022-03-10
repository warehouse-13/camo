#!/bin/sh

export PASSWORD=${CLEAR_PASSWORD:-dev}
python3 -c 'import bcrypt, os; print(bcrypt.hashpw(os.getenv("PASSWORD").encode(), bcrypt.gensalt()))' | tr -d b | tr -d "'" | base64
