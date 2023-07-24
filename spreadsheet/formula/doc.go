// The unit tests for this package are unique in that we can take advantage of
// "cached" formula results that Excel/LibreOffice write to the sheet.  These
// are the computed results of a formula in string form.  By comparing these
// values to the value computed by the gooxml evaluation of the formula, adding
// a new test means just adding a new formula to one of the reference sheets
// with Excel. During the unit test, we evaluate the formula and compare it to
// the value that Excel computed.  If they're the same, the test passes.
package formula
