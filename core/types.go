package core

var (
	OBJECT_BLOCK             = "block"
	BLOCK_BULLETED_LIST_ITEM = "bulleted_list_item"
	BLOCK_HEADING1           = "heading_1"
	BLOCK_HEADING2           = "heading_2"
	BLOCK_HEADING3           = "heading_3"
	TYPE_TEXT                = "text"
)

type idObject struct {
	ID *string `json:"id,omitempty"`
}

type internalFile struct {
	URL        *string `json:"url,omitempty"`
	ExpiryTime *string `json:"expiry_time,omitempty"`
}

type externalFile struct {
	URL *string `json:"url,omitempty"`
}

type fileObject struct {
	Type         *string       `json:"type,omitempty"`
	InternalFile *internalFile `json:"file,omitempty"`
	ExternalFile *externalFile `json:"external,omitempty"`
}

type emojiObject struct {
	Type  *string `json:"type,omitempty"`
	Emoji *string `json:"emoji,omitempty"`
}

type richTextArrary []richTextObject

type relationObject struct {
	Relations []idObject `json:"relation,omitempty"`
}

type numberObject struct {
	Number *int64 `json:"number,omitempty"`
}

type selectOption struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// "default", "gray", "brown",
	// "red", "orange", "yellow",
	// "green", "blue", "purple", "pink".
	Color *string `json:"color,omitempty"`
}

type multiSelectObject []selectOption

type dateObject struct {
	Start *string `json:"start,omitempty"`
	End   *string `json:"end,omitempty"`
}

type formulaObject struct {
	Type *string `json:"type,omitempty"`
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
	Bold          *bool `json:"bold,omitempty"`
	Italic        *bool `json:"italic,omitempty"`
	Strikethrough *bool `json:"strikethrough,omitempty"`
	Underline     *bool `json:"underline,omitempty"`
	Code          *bool `json:"code,omitempty"`
	// "default", "gray", "brown", "orange",
	// "yellow", "green", "blue", "purple",
	// "pink", "red", "gray_background",
	// "brown_background", "orange_background",
	// "yellow_background", "green_background",
	// "blue_background", "purple_background",
	// "pink_background", "red_background"
	Color *string `json:"color,omitempty"`
}

type richTextObject struct {
	Type        *string           `json:"type,omitempty"`
	Text        *textObject       `json:"text,omitempty"`
	PlainText   *string           `json:"plain_text,omitempty"`
	Href        *string           `json:"href,omitempty"`
	Annotations *annotationObject `json:"annotations,omitempty"`
}

type textObject struct {
	Content *string `json:"content,omitempty"`
	Link    *string `json:"link,omitempty"`
}

type parentObject struct {
	Type        *string `json:"type,omitempty"`
	DatabaseID  *string `json:"database_id,omitempty"`
	PageID      *string `json:"page_id,omitempty"`
	IsWorkspace *bool   `json:"workspace,omitempty"`
}

type rollupObject struct {
	Type *string `json:"type,omitempty"`
	// by type
	Number *int64      `json:"number,omitempty"`
	Date   *dateObject `json:"date,omitempty"`
}

type propertyValue struct {
	ID   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
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
	Object         *string `json:"object,omitempty"`
	ID             *string `json:"id,omitempty"`
	CreatedTime    *string `json:"created_time,omitempty"`
	LastEditedTime *string `json:"last_edited_time,omitempty"`
	Archived       *bool   `json:"archived,omitempty"`
	// fileObject or emojObject
	Icon interface{} `json:"icon,omitempty"`
	// FileObject or emojiObject
	Cover      map[string]interface{}   `json:"cover,omitempty"`
	Properties map[string]propertyValue `json:"properties,omitempty"`
	Parent     *parentObject            `json:"parent,omitempty"`
	Url        *string                  `json:"string,omitempty"`
}

type paragraphBlobck struct {
	Text     richTextArrary `json:"text,omitempty"`
	Children []Block        `json:"children,omitempty"`
}

type headingBlobck struct {
	Text richTextArrary `json:"text,omitempty"`
}

type listItemBlock struct {
	Text     richTextArrary `json:"text,omitempty"`
	Children []Block        `json:"children,omitempty"`
}

type todoBlock struct {
	Text     richTextArrary `json:"text,omitempty"`
	Checked  *bool          `json:"checked,omitempty"`
	Children []Block        `json:"children,omitempty"`
}

type toggleBlock struct {
	Text     richTextArrary `json:"text,omitempty"`
	Children []Block        `json:"children,omitempty"`
}

type childPageBlock struct {
	Title *string `json:"title,omitempty"`
}

type embedBlock struct {
	Url *string `json:"url,omitempty"`
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
	Url     *string        `json:"url,omitempty"`
	Cpation richTextArrary `json:"caption,omitempty"`
}

type Block struct {
	Object *string `json:"object,omitempty"`
	ID     *string `json:"id,omitempty"`
	// "paragraph", "heading_1", "heading_2",
	// "heading_3", "bulleted_list_item",
	// "numbered_list_item", "to_do", "toggle",
	// "child_page", "embed", "image", "video",
	// "file", "pdf", "bookmark" and "unsupported"
	Type           *string `json:"type,omitempty"`
	CreatedTime    *string `json:"created_time,omitempty"`
	LastEditedTime *string `json:"last_edited_time,omitempty"`
	Archived       *bool   `json:"archived,omitempty"`
	HasChildren    *bool   `json:"has_children,omitempty"`

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
