package office

// MinGoVersion is used to cause a compile time error if office is compiled with
// an older version of go.  Specifically it requires a feature in go 1.8
// regarding collecting all attributes from arbitrary xml used in decode
// office.XSDAny.
const MinGoVersion = requires_go_18
