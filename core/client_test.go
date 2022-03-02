package core

import "testing"

func TestPageBasic(t *testing.T) {
	client, err := NewClient(&Option{SecretKey: "secret_Jbt9BN8bFQ1UvYFtvoUFtkCZBsEdR9XBG45UI6kYmI6"})
	if err != nil {
		t.Fatal(err)
	}

	page, err := client.RetrivePage("bf07281b3df645f58c708661b9182bc3")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", page)
	for k, v := range page.Properties {
		t.Logf("key: %s", k)
		switch v.Type {
		case "title":
			if v.TitleObject == nil {
				t.Fatal("title object is nil")
			}
			t.Logf("TitleObject => %+v", *v.TitleObject)
		default:
			t.Log("unknown type")
		}
	}
}

func TestBlockBasic(t *testing.T) {
	client, err := NewClient(&Option{SecretKey: "secret_Jbt9BN8bFQ1UvYFtvoUFtkCZBsEdR9XBG45UI6kYmI6"})
	if err != nil {
		t.Fatal(err)
	}

	var nextCursor string
	var blocks []Block
	hasMore := true
	for hasMore {
		blocks, nextCursor, hasMore, err = client.RetriveBlockChildren("bf07281b3df645f58c708661b9182bc3", nextCursor, 3)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("%+v, next: %s, hasMore: %v", blocks, nextCursor, hasMore)
		for _, v := range blocks {
			t.Logf("block: %+v", v)
		}
	}
}

func TestBlockAppend(t *testing.T) {
	client, err := NewClient(&Option{SecretKey: "secret_Jbt9BN8bFQ1UvYFtvoUFtkCZBsEdR9XBG45UI6kYmI6"})
	if err != nil {
		t.Fatal(err)
	}

	var blocks []*Block
	var block1 Block
	block1.Object = OBJECT_BLOCK
	block1.Type = BLOCK_HEADING3
	var heading3Block HeadingBlobck
	date := "2021-09-12"
	heading3Block.Text = append(heading3Block.Text, RichTextObject{
		Type: TYPE_TEXT,
		Text: &TextObject{
			Content: date,
		},
	})
	block1.Heading3Block = &heading3Block
	blocks = append(blocks, &block1)

	var bulletedItem ListItemBlock
	item := "这是一个测试的BulletedItem"
	bulletedItem.Text = append(bulletedItem.Text,
		RichTextObject{
			Type: TYPE_TEXT,
			Text: &TextObject{
				Content: item,
			},
		})
	var block2 Block
	block2.Object = OBJECT_BLOCK
	block2.Type = BLOCK_BULLETED_LIST_ITEM
	block2.BulletedListItemBlock = &bulletedItem
	blocks = append(blocks, &block2)

	err = client.AppendBlock("bf07281b3df645f58c708661b9182bc3", blocks)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddPage(t *testing.T) {
	client, err := NewClient(&Option{SecretKey: "secret_Jbt9BN8bFQ1UvYFtvoUFtkCZBsEdR9XBG45UI6kYmI6"})
	if err != nil {
		t.Fatal(err)
	}

	var page Page
	// pageId := "f5577535d9ff4dae8912bc1a5bc83bca"
	// dbid := "f5577535d9ff4dae8912bc1a5bc83bca"
	dbid := "f5577535d9ff4dae8912bc1a5bc83bca"
	// pageT := "page_id"
	// dbT := "database_id"
	page.Parent = ParentObject{
		// Type:       &dbT,
		DatabaseID: dbid,
		// PageID:     &pageId,
	}

	True := true
	str := "这是一个测试的数据"
	var textBlock Block
	textBlock.Object = OBJECT_BLOCK
	textBlock.Type = BLOCK_PARAGRAPH
	var content ParagraphBlock
	content.Text = append(content.Text, RichTextObject{
		Type: TYPE_TEXT,
		Text: &TextObject{
			Content: str,
		},
		Annotations: &AnnotationObject{
			Bold: True,
		},
	})
	textBlock.ParagraphBlock = &content
	page.Children = append(page.Children, textBlock)
	page.Properties = make(map[string]PropertyValue)
	page.Properties["Name"] = PropertyValue{
		Type:        TYPE_TITLE,
		TitleObject: &RichTextArrary{},
	}

	var tags MultiSelectObject
	tags = append(tags, SelectOption{Name: "美食"}, SelectOption{Name: "科技"})
	page.Properties["Tags"] = PropertyValue{
		Type:        TYPE_MULTI_SELECT,
		MultiSelect: &tags,
	}

	err = client.CreatePage(&page)
	if err != nil {
		t.Fatal(err)
	}
}
