#!/bin/bash

# Verificar si se ha pasado una URL como argumento
if [ -z "$1" ]; then
  echo "Uso: $0 <url>"
  exit 1
fi

URL="$1"

# Extraer el fragmento de la URL
FRAGMENT=$(echo "$URL" | grep -oP '(?<=#).+')

# Verificar si el fragmento contiene el access_token
if [[ "$FRAGMENT" =~ access_token=([^&]+) ]]; then
  ACCESS_TOKEN="${BASH_REMATCH[1]}"
  echo "Access token encontrado: $ACCESS_TOKEN"

  # Copiar el access token al portapapeles (usando clip para Windows)
  echo -n "$ACCESS_TOKEN" | clip

  echo "Access token copiado al portapapeles."
else
  echo "Access token no encontrado en la URL proporcionada."
  exit 1
fi
