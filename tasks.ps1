param(
    [Parameter(Mandatory = $true)]
    [ValidateSet('run', 'build', 'test')]
    [string]$Action
)

$ErrorActionPreference = 'Stop'
Set-Location -LiteralPath $PSScriptRoot

switch ($Action) {
    'run' {
        go run ./cmd/server
    }

    'build' {
        New-Item -ItemType Directory -Force -Path .\bin | Out-Null
        go build -o .\bin\server.exe .\cmd\server
    }

    'test' {
        go test ./...
    }
}

if ($LASTEXITCODE -ne 0) {
    exit $LASTEXITCODE
}