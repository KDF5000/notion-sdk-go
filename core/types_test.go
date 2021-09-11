package core

import (
	"encoding/json"
	"testing"
)

var (
    pageStr = `
{
    "object": "page",
    "id": "bf07281b-3df6-45f5-8c70-8661b9182bc3",
    "created_time": "2021-09-05T03:09:00.000Z",
    "last_edited_time": "2021-09-11T08:56:00.000Z",
    "cover": {
        "type": "external",
        "external": {
            "url": "https://www.notion.so/images/page-cover/nasa_wrights_first_flight.jpg"
        }
    },
    "icon": {
        "type": "file",
        "file": {
            "url": "https://s3.us-west-2.amazonaws.com/secure.notion-static.com/d7a14f04-dea6-436e-80b1-fe0cdfd2e4bc/default.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAT73L2G45O3KS52Y5%2F20210911%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20210911T085658Z&X-Amz-Expires=3600&X-Amz-Signature=6923c8aa7f70ba283bc8453b8a1cd164012dba6e71b5b0b7d316cab2c1232439&X-Amz-SignedHeaders=host",
            "expiry_time": "2021-09-11T09:56:57.996Z"
        }
    },
    "parent": {
        "type": "page_id",
        "page_id": "32a6e309-91af-466b-b957-cf470b3e2578"
    },
    "archived": false,
    "properties": {
        "title": {
            "id": "title",
            "type": "title",
            "title": [
                {
                    "type": "text",
                    "text": {
                        "content": "Nomo",
                        "link": null
                    },
                    "annotations": {
                        "bold": false,
                        "italic": false,
                        "strikethrough": false,
                        "underline": false,
                        "code": false,
                        "color": "default"
                    },
                    "plain_text": "Nomo",
                    "href": null
                }
            ]
        }
    },
    "url": "https://www.notion.so/Nomo-bf07281b3df645f58c708661b9182bc3"
}
`
)

func TestPage(t *testing.T) {
    var page Page
    err := json.Unmarshal([]byte(pageStr), &page)
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("page => %+v", page)
}
