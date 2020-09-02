@echo off
go test -timeout 30s -coverprofile=test_coverage.txt gitlab.com/enda-mullally/codechallenge/logic gitlab.com/enda-mullally/codechallenge/models
go tool cover -html=test_coverage.txt -o test_coverage.html
echo.
echo Done!
pause>nul
