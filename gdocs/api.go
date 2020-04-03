package gdocs

import (
	"net/http"

	"google.golang.org/api/docs/v1"
)

// CreateDocument TODO
func CreateDocument(client *http.Client, title string) (*docs.Document, error) {
	document := &docs.Document{
		Title: title,
	}

	service, err := docs.New(client)
	if err != nil {
		return document, err
	}

	return service.Documents.Create(document).Do()
}

// CreateFrontmatterRequests TODO
func CreateFrontmatterRequests(titles []string, insertIndex int) []*docs.Request {
	var (
		titlesNo = len(titles)
		requests = []*docs.Request{
			&docs.Request{InsertTable: &docs.InsertTableRequest{Columns: 2, Rows: int64(titlesNo), Location: &docs.Location{Index: 1}}},
		}
		rowIndex                          = int64(insertIndex * titlesNo)
		firstItemStartIndex, firstItemLen int64
	)

	titlesNo--

	for ; titlesNo > -1; titlesNo-- {
		firstItemStartIndex = rowIndex
		firstItemLen = firstItemStartIndex + int64(len(titles[titlesNo]))

		requests = append(
			requests,
			&docs.Request{InsertText: &docs.InsertTextRequest{Text: "<placeholder>", Location: &docs.Location{Index: rowIndex + 2}}},
			&docs.Request{InsertText: &docs.InsertTextRequest{Text: titles[titlesNo], Location: &docs.Location{Index: firstItemStartIndex}}},
			&docs.Request{
				UpdateTextStyle: &docs.UpdateTextStyleRequest{
					Fields: "*",
					Range: &docs.Range{
						StartIndex: firstItemLen + 2,
						EndIndex:   firstItemLen + 15,
					},
					TextStyle: &docs.TextStyle{
						BackgroundColor: &docs.OptionalColor{
							Color: &docs.Color{
								RgbColor: &docs.RgbColor{
									Blue:  0.922,
									Green: 0.62,
									Red:   0.427,
								},
							},
						},
						FontSize: &docs.Dimension{
							Unit:      "PT",
							Magnitude: 11,
						},
						WeightedFontFamily: &docs.WeightedFontFamily{
							FontFamily: "Montserrat",
							Weight:     300,
						},
					},
				},
			},
			&docs.Request{
				UpdateTextStyle: &docs.UpdateTextStyleRequest{
					Fields: "*",
					Range: &docs.Range{
						StartIndex: firstItemStartIndex,
						EndIndex:   firstItemLen,
					},
					TextStyle: &docs.TextStyle{
						Bold: true,
						FontSize: &docs.Dimension{
							Unit:      "PT",
							Magnitude: 12,
						},
						WeightedFontFamily: &docs.WeightedFontFamily{
							FontFamily: "Montserrat",
							Weight:     500,
						},
					},
				},
			},
		)

		rowIndex -= 5
	}

	zeroWidthBorder := &docs.TableCellBorder{DashStyle: "SOLID", Color: &docs.OptionalColor{Color: &docs.Color{RgbColor: &docs.RgbColor{Red: 1, Green: 1, Blue: 1}}}, Width: &docs.Dimension{Unit: "PT", Magnitude: 0}}

	requests = append(
		requests,
		&docs.Request{UpdateTableCellStyle: &docs.UpdateTableCellStyleRequest{
			Fields:             "*",
			TableStartLocation: &docs.Location{Index: 2},
			TableCellStyle: &docs.TableCellStyle{
				PaddingBottom: &docs.Dimension{Unit: "PT", Magnitude: 2},
				PaddingTop:    &docs.Dimension{Unit: "PT", Magnitude: 2},
				PaddingLeft:   &docs.Dimension{Unit: "PT", Magnitude: 2},
				PaddingRight:  &docs.Dimension{Unit: "PT", Magnitude: 2},
				BorderBottom:  zeroWidthBorder,
				BorderTop:     zeroWidthBorder,
				BorderLeft:    zeroWidthBorder,
				BorderRight:   zeroWidthBorder,
			},
		}},
	)

	return requests
}

// CreateTitleRequests TODO
func CreateTitleRequests(insertIndex int, allTitles ...[]string) []*docs.Request {
	var (
		titles       = []string{}
		allTitlesLen = len(allTitles)
	)

	for i := 0; i < allTitlesLen; i++ {
		titles = append(titles, allTitles[i]...)
	}

	var (
		endIndex int64
		title    string
		titleLen int

		index    = int64(insertIndex)
		requests = []*docs.Request{
			&docs.Request{
				InsertText: &docs.InsertTextRequest{
					Text:     "\n\n\n",
					Location: &docs.Location{Index: index},
				},
			},
		}
		titleNo = len(titles)
	)

	index += 3

	for j := 0; j < titleNo; j++ {
		title = titles[j] + "\n"
		titleLen = len(title)
		endIndex = index + int64(titleLen)

		requests = append(requests, headingOneRequests(title, index, endIndex)...)

		index = endIndex
		endIndex = endIndex + int64(19)

		requests = append(requests, instructionTextRequests("insert text below ↓", index, endIndex)...)

		index = endIndex
		endIndex = endIndex + int64(3)

		requests = append(requests, normalTextRequests("\n\n\n", index, endIndex)...)

		index = endIndex
	}

	return requests
}

// AddText TODO
func AddText(client *http.Client, documentID string, requests []*docs.Request) error {
	service, err := docs.New(client)
	if err != nil {
		return err
	}

	_, err = service.Documents.BatchUpdate(
		documentID,
		&docs.BatchUpdateDocumentRequest{
			Requests: requests,
		}).Do()
	return err
}

func headingOneRequests(title string, index, endIndex int64) []*docs.Request {
	return []*docs.Request{
		&docs.Request{
			InsertText: &docs.InsertTextRequest{
				Text:     title,
				Location: &docs.Location{Index: index},
			},
		},
		&docs.Request{
			UpdateTextStyle: &docs.UpdateTextStyleRequest{
				Fields: "*",
				Range: &docs.Range{
					StartIndex: index,
					EndIndex:   endIndex,
				},
				TextStyle: &docs.TextStyle{
					FontSize: &docs.Dimension{
						Unit:      "PT",
						Magnitude: 20,
					},
					WeightedFontFamily: &docs.WeightedFontFamily{
						FontFamily: "Montserrat",
						Weight:     600,
					},
				},
			},
		},
	}
}

func normalTextRequests(text string, index, endIndex int64) []*docs.Request {
	return []*docs.Request{
		&docs.Request{
			InsertText: &docs.InsertTextRequest{
				Text:     text,
				Location: &docs.Location{Index: index},
			},
		},
		&docs.Request{
			UpdateTextStyle: &docs.UpdateTextStyleRequest{
				Fields: "*",
				Range: &docs.Range{
					StartIndex: index,
					EndIndex:   endIndex,
				},
				TextStyle: &docs.TextStyle{
					FontSize: &docs.Dimension{
						Unit:      "PT",
						Magnitude: 11,
					},
					WeightedFontFamily: &docs.WeightedFontFamily{
						FontFamily: "Montserrat",
						Weight:     300,
					},
				},
			},
		},
	}
}

func instructionTextRequests(text string, index, endIndex int64) []*docs.Request {
	return []*docs.Request{
		&docs.Request{
			InsertText: &docs.InsertTextRequest{
				Text:     text,
				Location: &docs.Location{Index: index},
			},
		},
		&docs.Request{
			UpdateTextStyle: &docs.UpdateTextStyleRequest{
				Fields: "*",
				Range: &docs.Range{
					StartIndex: index,
					EndIndex:   endIndex,
				},
				TextStyle: &docs.TextStyle{
					BackgroundColor: &docs.OptionalColor{
						Color: &docs.Color{
							RgbColor: &docs.RgbColor{
								Blue:  0.49,
								Red:   0.576,
								Green: 0.769,
							},
						},
					},
					FontSize: &docs.Dimension{
						Unit:      "PT",
						Magnitude: 11,
					},
					WeightedFontFamily: &docs.WeightedFontFamily{
						FontFamily: "Montserrat",
						Weight:     400,
					},
				},
			},
		},
	}
}
