#!/bin/bash

# Verify if the URL have been insert as a parameter
if [ -z "$1" ]; then
  echo "Uso: $0 <url>"
  exit 1
fi

URL="$1"

# Extract the fragment form the URL
FRAGMENT=$(echo "$URL" | grep -oP '(?<=#).+')

# Verify if the access_token exists
if [[ "$FRAGMENT" =~ access_token=([^&]+) ]]; then
  ACCESS_TOKEN="${BASH_REMATCH[1]}"
  echo "Access token encontrado: $ACCESS_TOKEN"

  # This copy the access token
  echo -n "$ACCESS_TOKEN" | clip

  echo "Access token copiado al portapapeles."
else
  echo "Access token no encontrado en la URL proporcionada."
  exit 1
fi
