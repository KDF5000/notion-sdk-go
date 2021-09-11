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
	Type string `json:"type"`
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
	Content string `json:"content"`
	Link    string `json:"link"`
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
	Number *int64      `json:"number,omitempty"`
	Date   *dateObject `json:"date,omitempty"`
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
	Rollup       *rollupObject      `json:"rollup,omitempty"`
}

type Page struct {
	Object         string      `json:"object"`
	ID             string      `json:"id"`
	CreatedTime    string      `json:"created_time"`
	LastEditedTime string      `json:"last_edited_time"`
	Archived       bool        `json:"archived"`
	Icon           interface{} `json:"icon"`
	// FileObject or emojiObject
	Cover      map[string]interface{}   `json:"cover"`
	Properties map[string]propertyValue `json:"properties"`
	Parent     parentObject             `json:"parent"`
	Url        string                   `json:"string"`
}

type paragraphBlobck struct {
	Text     richTextArrary `json:"text"`
	Children []Block        `json:"children"`
}

type headingBlobck struct {
	Text richTextArrary `json:"text"`
}

type listItemBlock struct {
	Text     richTextArrary `json:"text"`
	Children []Block        `json:"children"`
}

type todoBlock struct {
	Text     richTextArrary `json:"text"`
	Checked  bool           `json:"checked"`
	Children []Block        `json:"children"`
}

type toggleBlock struct {
	Text     richTextArrary `json:"text"`
	Children []Block        `json:"children"`
}

type childPageBlock struct {
	Title string `json:"title"`
}

type embedBlock struct {
	Url string `json:"url"`
}

type imageBlock struct {
	fileObject
}

type videoBlock struct {
	fileObject
}

type fileBlock struct {
	fileObject
}

type pdfBlock struct {
	fileObject
}

type bookmarkBlock struct {
	Url     string         `json:"url"`
	Cpation richTextArrary `json:"caption"`
}

type Block struct {
	Object string `json:"object"`
	ID     string `json:"id"`
	// "paragraph", "heading_1", "heading_2",
	// "heading_3", "bulleted_list_item",
	// "numbered_list_item", "to_do", "toggle",
	// "child_page", "embed", "image", "video",
	// "file", "pdf", "bookmark" and "unsupported"
	Type           string `json:"type"`
	CreatedTime    string `json:"created_time"`
	LastEditedTime string `json:"last_edited_time"`
	Archived       bool   `json:"archived"`
	HasChildren    bool   `json:"has_children"`

	// one of the following blocks will be populated
	ParagraphBlock        *paragraphBlobck `json:"paragraph,omitempty"`
	Heading1Block         *headingBlobck   `json:"heading_1,omitempty"`
	Heading2Block         *headingBlobck   `json:"heading_2,omitempty"`
	Heading3Block         *headingBlobck   `json:"heading_3,omitempty"`
	BulletedListItemBlock *listItemBlock   `json:"bulleted_list_item,omitempty"`
	NumberedListItemBlock *listItemBlock   `json:"numbered_list_item,omitempty"`
	TodoBlockBlock        *todoBlock       `json:"to_do,omitempty"`
	ToggleBlock           *toggleBlock     `json:"toggle,omitempty"`
	ChildPageBlock        *childPageBlock  `json:"child_page,omitempty"`
	EmbedBlock            *embedBlock      `json:"embed,omitempty"`
	ImageBlock            *imageBlock      `json:"image,omitempty"`
	VideoBlock            *videoBlock      `json:"video,omitempty"`
	FileBlock             *fileBlock       `json:"file,omitempty"`
	PdfBlock              *pdfBlock        `json:"pdf,omitempty"`
	BookmarkBlock         *bookmarkBlock   `json:"bookmark,omitempty"`
}
