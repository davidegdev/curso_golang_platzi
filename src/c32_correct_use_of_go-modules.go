// ================================================
// 													MODIFY MODULES
// ================================================
// from .
// >git clone https://github.com/labstack/echo.git
// Modify the echo.go file to include a motivational message (Estoy Aprendiendo Go con Platzi!) at the end of the ASCII art and save changes.
// >cd ..
// >go mod edit -replace github.com/labstack/echo=./echo

// ================================================
// 													REVERSE CHANGES
// ================================================
// >go mod edit -dropreplace github.com/labstack/echo

// ================================================
// 													PACKAGE DEPENDENCIES
// ================================================
// >go mod vendor

// ================================================
// 													CLEAN UP UNUSED MODULES
// ================================================
// >go mod tidy