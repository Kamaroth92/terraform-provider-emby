#!/usr/bin/env pwsh
# Usage:
#   .\generate-sdk.ps1 -EmbyHost "https://my-emby-server:8096"

param(
    [Parameter(Mandatory)]
    [string]$EmbyHost
)

$ErrorActionPreference = "Stop"

$specUrl = "$EmbyHost/emby/openapi"
$specFile = Join-Path $PSScriptRoot "openapi-spec.json"
$outputDir = Join-Path $PSScriptRoot "internal/client"

Write-Host "Downloading OpenAPI spec from $specUrl ..."
Invoke-WebRequest -Uri $specUrl -OutFile $specFile
Write-Host "Saved spec to $specFile"

Write-Host "Building codegen Docker image..."
docker build -f "$PSScriptRoot/Dockerfile.codegen" -t emby-go-codegen "$PSScriptRoot"

if (Test-Path $outputDir) {
    Write-Host "Clearing old generated files in $outputDir ..."
    Remove-Item -Recurse -Force "$outputDir/*"
}
else {
    New-Item -ItemType Directory -Path $outputDir -Force | Out-Null
}

Write-Host "Generating Go SDK..."
docker run --rm `
    -v "${specFile}:/spec/openapi.json" `
    -v "${outputDir}:/out" `
    emby-go-codegen `
    -i /spec/openapi.json

# This is a workaround for a type name collision 
# between the generated ServerConfiguration and the provider's ServerConfiguration struct.
Write-Host "Fixing ServerConfiguration type collision in configuration.go ..."
$configFile = Join-Path $outputDir "configuration.go"
(Get-Content $configFile -Raw) `
    -replace 'type ServerConfiguration struct', 'type OAPIServerConfig struct' `
    -replace 'type ServerConfigurations \[\]ServerConfiguration', 'type OAPIServerConfigs []OAPIServerConfig' `
    -replace 'ServerConfigurations', 'OAPIServerConfigs' `
    | Set-Content $configFile

Write-Host ""
Write-Host "Done! Generated SDK is in $outputDir"
