// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    googleBooksStruct, err := UnmarshalGoogleBooksStruct(bytes)
//    bytes, err = googleBooksStruct.Marshal()

package main

import "encoding/json"

func UnmarshalGoogleBooksStruct(data []byte) (GoogleBooksStruct, error) {
	var r GoogleBooksStruct
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GoogleBooksStruct) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GoogleBooksStruct struct {
	Kind       string `json:"kind"`
	TotalItems int64  `json:"totalItems"`
	Items      []Item `json:"items"`
}

type Item struct {
	Kind       Kind       `json:"kind"`
	ID         string     `json:"id"`
	Etag       string     `json:"etag"`
	SelfLink   string     `json:"selfLink"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
	SaleInfo   SaleInfo   `json:"saleInfo"`
	AccessInfo AccessInfo `json:"accessInfo"`
	SearchInfo SearchInfo `json:"searchInfo"`
}

type AccessInfo struct {
	Country                Country                `json:"country"`
	Viewability            Viewability            `json:"viewability"`
	Embeddable             bool                   `json:"embeddable"`
	PublicDomain           bool                   `json:"publicDomain"`
	TextToSpeechPermission TextToSpeechPermission `json:"textToSpeechPermission"`
	Epub                   Epub                   `json:"epub"`
	PDF                    Epub                   `json:"pdf"`
	WebReaderLink          string                 `json:"webReaderLink"`
	AccessViewStatus       AccessViewStatus       `json:"accessViewStatus"`
	QuoteSharingAllowed    bool                   `json:"quoteSharingAllowed"`
}

type Epub struct {
	IsAvailable  bool    `json:"isAvailable"`
	AcsTokenLink *string `json:"acsTokenLink,omitempty"`
}

type SaleInfo struct {
	Country     Country            `json:"country"`
	Saleability Saleability        `json:"saleability"`
	IsEbook     bool               `json:"isEbook"`
	ListPrice   *SaleInfoListPrice `json:"listPrice,omitempty"`
	RetailPrice *SaleInfoListPrice `json:"retailPrice,omitempty"`
	BuyLink     *string            `json:"buyLink,omitempty"`
	Offers      []Offer            `json:"offers,omitempty"`
}

type SaleInfoListPrice struct {
	Amount       float64      `json:"amount"`
	CurrencyCode CurrencyCode `json:"currencyCode"`
}

type Offer struct {
	FinskyOfferType int64          `json:"finskyOfferType"`
	ListPrice       OfferListPrice `json:"listPrice"`
	RetailPrice     OfferListPrice `json:"retailPrice"`
}

type OfferListPrice struct {
	AmountInMicros int64        `json:"amountInMicros"`
	CurrencyCode   CurrencyCode `json:"currencyCode"`
}

type SearchInfo struct {
	TextSnippet string `json:"textSnippet"`
}

type VolumeInfo struct {
	Title               string               `json:"title"`
	Subtitle            *string              `json:"subtitle,omitempty"`
	Authors             []string             `json:"authors"`
	Publisher           *string              `json:"publisher,omitempty"`
	PublishedDate       string               `json:"publishedDate"`
	Description         string               `json:"description"`
	IndustryIdentifiers []IndustryIdentifier `json:"industryIdentifiers"`
	ReadingModes        ReadingModes         `json:"readingModes"`
	PageCount           int64                `json:"pageCount"`
	PrintType           PrintType            `json:"printType"`
	Categories          []Category           `json:"categories"`
	MaturityRating      MaturityRating       `json:"maturityRating"`
	AllowAnonLogging    bool                 `json:"allowAnonLogging"`
	ContentVersion      string               `json:"contentVersion"`
	PanelizationSummary PanelizationSummary  `json:"panelizationSummary"`
	ImageLinks          ImageLinks           `json:"imageLinks"`
	Language            Language             `json:"language"`
	PreviewLink         string               `json:"previewLink"`
	InfoLink            string               `json:"infoLink"`
	CanonicalVolumeLink string               `json:"canonicalVolumeLink"`
	AverageRating       *int64               `json:"averageRating,omitempty"`
	RatingsCount        *int64               `json:"ratingsCount,omitempty"`
}

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

type IndustryIdentifier struct {
	Type       Type   `json:"type"`
	Identifier string `json:"identifier"`
}

type PanelizationSummary struct {
	ContainsEpubBubbles  bool `json:"containsEpubBubbles"`
	ContainsImageBubbles bool `json:"containsImageBubbles"`
}

type ReadingModes struct {
	Text  bool `json:"text"`
	Image bool `json:"image"`
}

type AccessViewStatus string

const (
	None   AccessViewStatus = "NONE"
	Sample AccessViewStatus = "SAMPLE"
)

type Country string

const (
	SE Country = "SE"
)

type TextToSpeechPermission string

const (
	Allowed                 TextToSpeechPermission = "ALLOWED"
	AllowedForAccessibility TextToSpeechPermission = "ALLOWED_FOR_ACCESSIBILITY"
)

type Viewability string

const (
	NoPages Viewability = "NO_PAGES"
	Partial Viewability = "PARTIAL"
)

type Kind string

const (
	BooksVolume Kind = "books#volume"
)

type CurrencyCode string

const (
	Sek CurrencyCode = "SEK"
)

type Saleability string

const (
	ForSale    Saleability = "FOR_SALE"
	NotForSale Saleability = "NOT_FOR_SALE"
)

type Category string

const (
	Computers Category = "Computers"
)

type Type string

const (
	Isbn10 Type = "ISBN_10"
	Isbn13 Type = "ISBN_13"
)

type Language string

const (
	En Language = "en"
)

type MaturityRating string

const (
	NotMature MaturityRating = "NOT_MATURE"
)

type PrintType string

const (
	Book PrintType = "BOOK"
)
