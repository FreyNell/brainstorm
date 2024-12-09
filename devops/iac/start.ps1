#!/bin/bash

param (
    [Parameter(Mandatory = $true)]
    [string]$Environment # Argumento para especificar el ambiente (dev, prod, etc.)
)

$envFile = "..\..\secrets\$($Environment).env"

if (-Not (Test-Path $envFile)) {
    Write-Error "El archivo $envFile no existe."
    exit 1
}

Get-Content $envFile | ForEach-Object {

    # Omitir líneas en blanco o comentarios
    if ($_ -match '^\s*#' -or $_ -match '^\s*$') {
        return
    }

    # Dividir la línea en clave=valor
    $parts = $_ -split '=', 2

    if ($parts.Length -eq 2) {
        $key = $parts[0].Trim()
        $value = $parts[1].Trim()

        # Establecer la variable de entorno
        [System.Environment]::SetEnvironmentVariable($key, $value, "Process")
    }

}

docker-compose down
docker-compose up --build
