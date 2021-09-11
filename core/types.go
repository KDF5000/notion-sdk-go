package core

type idObject struct {
	ID string `json:"id"`
}

type internalFile struct {
	URL        string `json:"url"`
	ExpiryTime string `json:"expiry_time"`
}

type externalFile struct {
	URL string `json:"url"`
}

type fileObject struct {
	Type         string        `json:"type"`
	InternalFile *internalFile `json:"file,omitempty"`
	ExternalFile *externalFile `json:"external,omitempty"`
}

type emojiObject struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

type richTextArrary []richTextObject


type relationObject struct {
	Relations *[]idObject `json:"relation"`
}

type numberObject struct {
	Number int64 `json:"number"`
}

type selectOption struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// "default", "gray", "brown",
	// "red", "orange", "yellow",
	// "green", "blue", "purple", "pink".
	Color string `json:"color"`
}

type multiSelectObject []selectOption

type dateObject struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type formulaObject struct {
	Type    string      `json:"type"`
	// type
	String  *string     `json:"string,omitempty"`
	Number  *int64      `json:"number,omitempty"`
	Boolean *bool       `json:"bool,omitempty"`
	Date    *dateObject `json:"date,omitempty"`
}

// type titleObject struct {
// 	Title richTextArrary `json:"title"`
// }

type annotationObject struct {
	Bold          bool `json:"bold"`
	Italic        bool `json:"italic"`
	Strikethrough bool `json:"strikethrough"`
	Underline     bool `json:"underline"`
	Code          bool `json:"code"`
	// "default", "gray", "brown", "orange",
	// "yellow", "green", "blue", "purple",
	// "pink", "red", "gray_background",
	// "brown_background", "orange_background",
	// "yellow_background", "green_background",
	// "blue_background", "purple_background",
	// "pink_background", "red_background"
	Color string `json:"color"`
}

type richTextObject struct {
	Type        string           `json:"type"`
	Text        textObject       `json:"text"`
	PlainText   string           `json:"plain_text"`
	Href        string           `json:"href,omitempty"`
	Annotations annotationObject `json:"annotations"`
}

type textObject struct {
	Content string  `json:"content"`
	Link    string  `json:"link"`
}

type parentObject struct {
	Type        string `json:"type"`
	DatabaseID  string `json:"database_id,omitempty"`
	PageID      string `json:"page_id,omitempty"`
	IsWorkspace bool   `json:"workspace,omitempty"`
}

type rollupObject struct {
	Type string `json:"type"`
	// by type
	Number  *int64      `json:"number,omitempty"`
	Date    *dateObject `json:"date,omitempty"`
}

type Page struct {
	Object         string                   `json:"object"`
	ID             string                   `json:"id"`
	CreatedTime    string                   `json:"created_time"`
	LastEditedTime string                   `json:"last_edited_time"`
	Archived       bool                     `json:"archived"`
	Icon           interface{}              `json:"icon"`
	// FileObject or emojiObject
	Cover          map[string]interface{}   `json:"cover"`
	Properties     map[string]propertyValue `json:"properties"`
	Parent         parentObject             `json:"parent"`
	Url            string                   `json:"string"`
}

type propertyValue struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	// Value map[string]interface{} `json:",remain"`

	// one of the following objects will be populated
	// accroding to the type value
	TitleObject  *richTextArrary    `json:"title,omitempty"`
	RichText     *richTextArrary    `json:"rich_text,omitempty"`
	Number       *numberObject      `json:"number,omitempty"`
	SingleSelect *selectOption      `json:"select,omitempty"`
	MultiSelect  *multiSelectObject `json:"multi_select,omitempty"`
	Date         *dateObject        `json:"date,omitempty"`
	Formula      *formulaObject     `json:"formula,omitempty"`
	Relation     *relationObject    `json:"relation,omitempty"`
    Rollup       *rollupObject `json:"rollup,omitempty"`
}
