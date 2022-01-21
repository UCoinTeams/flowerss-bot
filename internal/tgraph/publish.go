package tgraph

import (
	"fmt"
	"html"
	"math/rand"
	"time"

	"go.uber.org/zap"
)

func PublishHtml(sourceTitle string, title string, rawLink string, htmlContent string) (string, error) {
	//html = fmt.Sprintf(
	//	"<p>本文章由 <a href=\"https://github.com/indes/flowerss-bot\">flowerss</a> 抓取自RSS，版权归<a href=\"\">源站点</a>所有。</p><hr>",
	//) + html + fmt.Sprintf(
	//	"<hr><p>本文章由 <a href=\"https://github.com/indes/flowerss-bot\">flowerss</a> 抓取自RSS，版权归<a href=\"\">源站点</a>所有。</p><p>查看原文：<a href=\"%s\">%s - %s</p>",
	//	rawLink,
	//	title,
	//	sourceTitle,
	//)

	htmlContent = html.UnescapeString(htmlContent) + fmt.Sprintf(
		"<hr><p>本页面是由 <a href=\"https://t.me/U2_Rss\">U2_Rss非官方频道</a> 抓取自种子介绍页</p><p>This page is crawled from the torrent introduction page by <a href=\"https://t.me/U2_Rss\">U2_Rss Unofficial Channel</a></p><p>このウェブページは <a href=\"https://t.me/U2_Rss\">U2_Rss非公式チャンネル</a> からトレントの紹介ウェブページを掴み取って</p><p>查看种子/View Torrent/トレントを見る：<a href=\"%s\">%s - %s</a></p>",
		rawLink,
		title,
		sourceTitle,
	)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	client := clientPool[rand.Intn(len(clientPool))]

	if page, err := client.CreatePageWithHTML(title+" - "+sourceTitle, sourceTitle, rawLink, htmlContent, true); err == nil {
		zap.S().Infof("Created telegraph page url: %s", page.URL)
		return page.URL, err
	} else {
		zap.S().Warnf("Create telegraph page failed, error: %s", err)
		return "", nil
	}
}
