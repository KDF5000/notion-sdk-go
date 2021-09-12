package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/KDF5000/notion-sdk-go/core"
)

var (
	client    *core.Client
	pageId    string
	secretKey string
	content   string
)

func AppendContent(client *core.Client, pageId string, content string) error {
	page, err := client.RetrivePage(pageId)
	if err != nil {
		return err
	}

	// 2020-08-12T02:12:33.231Z
	lastEditTime, err := time.Parse(time.RFC3339, *page.LastEditedTime)
	if err != nil {
		return err
	}

	var blocks []*core.Block
	if lastEditTime.Local().Day() != time.Now().Day() {
		var block core.Block
		block.Object = &core.OBJECT_BLOCK
		block.Type = &core.BLOCK_HEADING3
		var heading3Block core.HeadingBlobck
		date := time.Now().Format("2006-01-02")
		heading3Block.Text = append(heading3Block.Text, core.RichTextObject{
			Type: &core.TYPE_TEXT,
			Text: &core.TextObject{
				Content: &date,
			},
		})
		block.Heading3Block = &heading3Block
		blocks = append(blocks, &block)
	}

	var bulletedItem core.ListItemBlock
	bulletedItem.Text = append(bulletedItem.Text,
		core.RichTextObject{
			Type: &core.TYPE_TEXT,
			Text: &core.TextObject{
				Content: &content,
			},
		})

	blocks = append(blocks, &core.Block{
		Object:                &core.OBJECT_BLOCK,
		Type:                  &core.BLOCK_BULLETED_LIST_ITEM,
		BulletedListItemBlock: &bulletedItem,
	})

	err = client.AppendBlock(pageId, blocks)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	flag.StringVar(&secretKey, "key", "", "notion secret key")
	flag.StringVar(&pageId, "page", "", "page id")
	flag.StringVar(&content, "content", "", "content to append into page")
	flag.Parse()

	if secretKey == "" || pageId == "" || content == "" {
		fmt.Println("secrey key,  pageid or content is empty")
		return
	}

	var err error
	client, err = core.NewClient(&core.Option{SecretKey: secretKey})
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := AppendContent(client, pageId, content); err != nil {
		fmt.Println(err)
		return
	}
}
