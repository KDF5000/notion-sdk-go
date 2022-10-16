package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net"
	"time"
)

var (
	defaultHttpClient = http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   10  * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 10 * time.Second,
		},
		Timeout: 30 * time.Second,
	}
)

const (
	BASE_URI       = "https://api.notion.com/v1"
	NOTION_VERSION = "2021-08-16"
)

type Option struct {
	SecretKey string // Notion api secret key
}

type Client struct {
	option *Option
}

type RetriveBlockChildrenResp struct {
	Object     string  `json:"object"`
	NextCursor string  `json:"next_cursor"`
	HasMore    bool    `json:"has_more"`
	Result     []Block `json:"results"`
}

func NewClient(opt *Option) (*Client, error) {
	return &Client{option: opt}, nil
}

func (c *Client) CreatePage(page *Page) error {
	url := fmt.Sprintf("%s/pages", BASE_URI)
	payload, err := json.Marshal(page)
	if err != nil {
		return err
	}

	// fmt.Printf("%s, key:%s", string(payload), c.option.SecretKey)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Notion-Version", NOTION_VERSION)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.option.SecretKey))
	req.Header.Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	resp, err := defaultHttpClient.Do(req)
	if err != nil {
		return err
	}

	// r, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("%+v", string(r))
	if resp.StatusCode != 200 {
		return fmt.Errorf("code=%d, status=%s", resp.StatusCode, resp.Status)
	}

	return nil
}

func (c *Client) RetrivePage(id string) (*Page, error) {
	url := fmt.Sprintf("%s/pages/%s", BASE_URI, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Notion-Version", NOTION_VERSION)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.option.SecretKey))

	resp, err := defaultHttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("code=%d, status=%s", resp.StatusCode, resp.Status)
	}

	var page Page
	if err := json.NewDecoder(resp.Body).Decode(&page); err != nil {
		return nil, err
	}

	return &page, nil
}

func (c *Client) RetriveBlockChildren(blockId string, startCursor string, pageSize int) ([]Block, string, bool, error) {
	url := fmt.Sprintf("%s/blocks/%s/children?page_size=%d",
		BASE_URI, blockId, pageSize)
	if startCursor != "" {
		url = fmt.Sprintf("%s&start_cursor=%s", url, startCursor)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", false, err
	}

	req.Header.Set("Notion-Version", NOTION_VERSION)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.option.SecretKey))

	resp, err := defaultHttpClient.Do(req)
	if err != nil {
		return nil, "", false, err
	}

	if resp.StatusCode != 200 {
		return nil, "", false, fmt.Errorf("code=%d, status=%s", resp.StatusCode, resp.Status)
	}

	var blocksResp RetriveBlockChildrenResp
	if err := json.NewDecoder(resp.Body).Decode(&blocksResp); err != nil {
		return nil, "", false, err
	}

	return blocksResp.Result, blocksResp.NextCursor, blocksResp.HasMore, nil
}

func (c *Client) AppendBlock(blockId string, blocks []*Block) error {
	url := fmt.Sprintf("%s/blocks/%s/children", BASE_URI, blockId)

	payload := struct {
		Children []*Block `json:"children"`
	}{
		Children: blocks,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Notion-Version", NOTION_VERSION)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.option.SecretKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := defaultHttpClient.Do(req)
	if err != nil {
		return err
	}

	// r, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("%+v", string(r))
	if resp.StatusCode != 200 {
		return fmt.Errorf("code=%d, status=%s", resp.StatusCode, resp.Status)
	}

	return nil
}
