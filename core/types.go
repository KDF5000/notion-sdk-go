package core

var (
	OBJECT_BLOCK             = "block"
	BLOCK_BULLETED_LIST_ITEM = "bulleted_list_item"
	BLOCK_HEADING1           = "heading_1"
	BLOCK_HEADING2           = "heading_2"
	BLOCK_HEADING3           = "heading_3"
	TYPE_TEXT                = "text"
)

type IdObject struct {
	ID *string `json:"id,omitempty"`
}

type InternalFile struct {
	URL        *string `json:"url,omitempty"`
	ExpiryTime *string `json:"expiry_time,omitempty"`
}

type ExternalFile struct {
	URL *string `json:"url,omitempty"`
}

type FileObject struct {
	Type         *string       `json:"type,omitempty"`
	InternalFile *InternalFile `json:"file,omitempty"`
	ExternalFile *ExternalFile `json:"external,omitempty"`
}

type EmojiObject struct {
	Type  *string `json:"type,omitempty"`
	Emoji *string `json:"emoji,omitempty"`
}

type RichTextArrary []richTextObject

type RelationObject struct {
	Relations []IdObject `json:"relation,omitempty"`
}

type NumberObject struct {
	Number *int64 `json:"number,omitempty"`
}

type SelectOption struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// "default", "gray", "brown",
	// "red", "orange", "yellow",
	// "green", "blue", "purple", "pink".
	Color *string `json:"color,omitempty"`
}

type MultiSelectObject []SelectOption

type DateObject struct {
	Start *string `json:"start,omitempty"`
	End   *string `json:"end,omitempty"`
}

type FormulaObject struct {
	Type *string `json:"type,omitempty"`
	// type
	String  *string     `json:"string,omitempty"`
	Number  *int64      `json:"number,omitempty"`
	Boolean *bool       `json:"bool,omitempty"`
	Date    *DateObject `json:"date,omitempty"`
}

// type titleObject struct {
// 	Title RichTextArrary `json:"title"`
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

type RollupObject struct {
	Type *string `json:"type,omitempty"`
	// by type
	Number *int64      `json:"number,omitempty"`
	Date   *DateObject `json:"date,omitempty"`
}

type propertyValue struct {
	ID   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	// Value map[string]interface{} `json:",remain"`

	// one of the following objects will be populated
	// accroding to the type value
	TitleObject  *RichTextArrary    `json:"title,omitempty"`
	RichText     *RichTextArrary    `json:"rich_text,omitempty"`
	Number       *NumberObject      `json:"number,omitempty"`
	SingleSelect *SelectOption      `json:"select,omitempty"`
	MultiSelect  *MultiSelectObject `json:"multi_select,omitempty"`
	Date         *DateObject        `json:"date,omitempty"`
	Formula      *FormulaObject     `json:"formula,omitempty"`
	Relation     *RelationObject    `json:"relation,omitempty"`
	Rollup       *RollupObject      `json:"rollup,omitempty"`
}

type Page struct {
	Object         *string `json:"object,omitempty"`
	ID             *string `json:"id,omitempty"`
	CreatedTime    *string `json:"created_time,omitempty"`
	LastEditedTime *string `json:"last_edited_time,omitempty"`
	Archived       *bool   `json:"archived,omitempty"`
	// FileObject or EmojObject
	Icon interface{} `json:"icon,omitempty"`
	// FileObject or EmojiObject
	Cover      map[string]interface{}   `json:"cover,omitempty"`
	Properties map[string]propertyValue `json:"properties,omitempty"`
	Parent     *parentObject            `json:"parent,omitempty"`
	Url        *string                  `json:"string,omitempty"`
}

type ParagraphBlobck struct {
	Text     RichTextArrary `json:"text,omitempty"`
	Children []Block        `json:"children,omitempty"`
}

type HeadingBlobck struct {
	Text RichTextArrary `json:"text,omitempty"`
}

type ListItemBlock struct {
	Text     RichTextArrary `json:"text,omitempty"`
	Children []Block        `json:"children,omitempty"`
}

type TodoBlock struct {
	Text     RichTextArrary `json:"text,omitempty"`
	Checked  *bool          `json:"checked,omitempty"`
	Children []Block        `json:"children,omitempty"`
}

type ToggleBlock struct {
	Text     RichTextArrary `json:"text,omitempty"`
	Children []Block        `json:"children,omitempty"`
}

type ChildPageBlock struct {
	Title *string `json:"title,omitempty"`
}

type EmbedBlock struct {
	Url *string `json:"url,omitempty"`
}

type ImageBlock struct {
	FileObject
}

type VideoBlock struct {
	FileObject
}

type FileBlock struct {
	FileObject
}

type PdfBlock struct {
	FileObject
}

type BookmarkBlock struct {
	Url     *string        `json:"url,omitempty"`
	Cpation RichTextArrary `json:"caption,omitempty"`
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
	ParagraphBlock        *ParagraphBlobck `json:"paragraph,omitempty"`
	Heading1Block         *HeadingBlobck   `json:"heading_1,omitempty"`
	Heading2Block         *HeadingBlobck   `json:"heading_2,omitempty"`
	Heading3Block         *HeadingBlobck   `json:"heading_3,omitempty"`
	BulletedListItemBlock *ListItemBlock   `json:"bulleted_list_item,omitempty"`
	NumberedListItemBlock *ListItemBlock   `json:"numbered_list_item,omitempty"`
	TodoBlockBlock        *TodoBlock       `json:"to_do,omitempty"`
	ToggleBlock           *ToggleBlock     `json:"toggle,omitempty"`
	ChildPageBlock        *ChildPageBlock  `json:"child_page,omitempty"`
	EmbedBlock            *EmbedBlock      `json:"embed,omitempty"`
	ImageBlock            *ImageBlock      `json:"image,omitempty"`
	VideoBlock            *VideoBlock      `json:"video,omitempty"`
	FileBlock             *FileBlock       `json:"file,omitempty"`
	PdfBlock              *PdfBlock        `json:"pdf,omitempty"`
	BookmarkBlock         *BookmarkBlock   `json:"bookmark,omitempty"`
}
