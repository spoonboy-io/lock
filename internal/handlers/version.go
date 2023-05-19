package handlers

// Version returns template for version information
func Version() string {

	return `Client:
 Version: %s
 Go version: %s
Datasources: 	
 Template data:
  URL: %s
  Cache file: %s
  Cache TTL: %d mins
 Plugin data: 
   URL: %s 
Defaults:
 Project name: %s
`
}
