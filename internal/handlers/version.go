package handlers

// Version returns template for version information
func Version() string {

	return `Client:
 Version: %s
 Go version: %s
Data sources: 	
 Template data:
  URL: %s
  Cache file: %s
  Cache TTL: %d mins
 Plugin data: 
  URL: %s
  Cache file: %s
  Cache TTL: %d mins
Defaults:
 Project name: %s
`
}
