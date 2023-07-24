**ooxml** is a library for creation of Office Open XML documents (.docx, .xlsx
and .pptx)

[![Build & Test Action
Status](https://github.com/devnw/ooxml/actions/workflows/build.yml/badge.svg)](https://github.com/devnw/ooxml/actions)
[![Go Report
Card](https://goreportcard.com/badge/go.devnw.com/ooxml)](https://goreportcard.com/report/go.devnw.com/ooxml)
[![codecov](https://codecov.io/gh/devnw/ooxml/branch/main/graph/badge.svg)](https://codecov.io/gh/devnw/ooxml)
[![Go
Reference](https://pkg.go.dev/badge/go.devnw.com/ooxml.svg)](https://pkg.go.dev/go.devnw.com/ooxml)
[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
[![PRs
Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)


## Status ##

- Documents (docx) [Word]
	- Read/Write/Edit
	- Formatting
	- Images
	- Tables
- Spreadsheets (xlsx) [Excel]
 	- Read/Write/Edit
 	- Cell formatting including conditional formatting
	- Cell validation (drop down combobox, rules, etc.)
    - Retrieve cell values as formatted by Excel (e.g. retrieve a date or number as displayed in Excel)
 	- Formula Evaluation (100+ functions supported currently, more will be added as required)
 	- Embedded Images
 	- All chart types
- PowerPoint (pptx) [PowerPoint]
	- Creation from templates
	- Textboxes/shapes


## Performance ##

There has been a great deal of interest in performance numbers for spreadsheet
creation/reading lately, so here are office numbers for this
[benchmark](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/lots-of-rows)
which creates a sheet with 30k rows, each with 100 columns.

    creating 30000 rows * 100 cells took 3.92506863s
    saving took 89ns
    reading took 9.522383048s

Creation is fairly fast, saving is very quick due to no reflection usage, and
reading is a bit slower. The downside is that the binary is large (33MB) as it
contains generated structs, serialization and deserialization code for all of
DOCX/XLSX/PPTX.

## Usage ##
    
    go get -u go.devnw.com/ooxml

## Document Examples ##

- [Simple Text Formatting](https://github.com/devnw/ooxml/tree/main/examples/document/simple) Text font colors, sizes, highlighting, etc.
- [Auto Generated Table of Contents](https://github.com/devnw/ooxml/tree/main/examples/document/toc) Creating document headings with an auto generated TOC based off of the headingds
- [Floating Image](https://github.com/devnw/ooxml/tree/main/examples/document/image) Placing an image somewhere on a page, absolutely positioned with different text wrapping.
- [Header & Footer](https://github.com/devnw/ooxml/tree/main/examples/document/header-footer) Creating headers and footers including page numbering.
- [Multiple Headers & Footers](https://github.com/devnw/ooxml/tree/main/examples/document/header-footer-multiple) Using different headers and footers depending on document section.
- [Inline Tables](https://github.com/devnw/ooxml/tree/main/examples/document/tables) Adding an table with and without borders.
- [Using Existing Word Document as a Template](https://github.com/devnw/ooxml/tree/main/examples/document/use-template) Opening a document as a template to re-use the styles created in the document.
- [Filling out Form Fields](https://github.com/devnw/ooxml/tree/main/examples/document/fill-out-form) Opening a document with embedded form fields, filling out the fields and saving the result as  a new filled form.
- [Editing an existing document](https://github.com/devnw/ooxml/tree/main/examples/document/edit-document) Open an existing document and replace/remove text without modifying formatting.

## Spreadsheet Examples ##
- [Simple](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/simple) A simple sheet with a few cells
- [Named Cells](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/named-cells) Different ways of referencing rows and cells
- [Cell Number/Date/Time Formats](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/number-date-time-formats) Creating cells with various number/date/time formats
- [Line Chart](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/line-chart)/[Line Chart 3D](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/line-chart-3d) Line Charts
- [Bar Chart](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/bar-chart) Bar Charts
- [Mutiple Charts](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/multiple-charts) Multiple charts on a single sheet
- [Named Cell Ranges](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/named-ranges) Naming cell ranges
- [Merged Cells](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/merged) Merge and unmerge cells
- [Conditional Formatting](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/conditional-formatting) Conditionally formatting cells, styling, gradients, icons, data bar
- [Complex](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/complex) Multiple charts, auto filtering and conditional formatting
- [Borders](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/borders) Individual cell borders and rectangular borders around a range of cells.
- [Validation](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/validation) Data validation including combo box dropdowns.
- [Frozen Rows/Cols](https://github.com/devnw/ooxml/tree/main/examples/spreadsheet/freeze-rows-cols) A sheet with a frozen header column and row

## Presentation Examples ##

- [Simple Text Boxes](https://github.com/devnw/ooxml/tree/main/examples/presentation/simple) Simple text boxes and shapes
- [Images](https://github.com/devnw/ooxml/tree/main/examples/presentation/image) Simple image insertion
- [Template](https://github.com/devnw/ooxml/tree/main/examples/presentation/use-template/simple) Creating a presentation from a template

## Raw Types ##

The OOXML specification is large and creating a friendly API to cover the entire
specification is a very time consuming endeavor.  This library attempts to
provide an easy to use API for common use cases in creating OOXML documents
while allowing users to fall back to raw document manipulation should the
library's API not cover a specific use case.

The raw XML based types reside in the ```schema/``` directory. These types are
accessible from the wrapper types via a ```X()``` method that returns the raw
type. 

For example, the library currently doesn't have an API for setting a document
background color. However it's easy to do manually via editing the
```CT_Background``` element of the document.

    dox := document.New()
    doc.X().Background = wordprocessingml.NewCT_Background()
	doc.X().Background.ColorAttr = &wordprocessingml.ST_HexColor{}
	doc.X().Background.ColorAttr.ST_HexColorRGB = color.RGB(50, 50, 50).AsRGBString()
