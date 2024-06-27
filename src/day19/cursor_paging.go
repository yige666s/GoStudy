package main

// 基于偏移量分页
// GET liwenzhou.com/api/v1/books?page=1&size=10
// 第一页：Offset:0, Limit:10
// 第二页：Offset:10, Limit:10
// 第三页：Offset:20, Limit:10
// SELECT id, title FROM books ORDER BY id ASC LIMIT 10 OFFSET 10;
// 缺点：1. 基于偏移量的分页在数据量很大的场景下，查询效率会比较低；2.在并发场景下会出现元素重复或者跳过；3.显式的page参数在支持跳页的同时也会被爬虫并发请求

// 基于游标分页
// cursor——通常是一个不透明字符串。它表示的是这一页数据的最后那个元素，通过这个cursor API 就能准确的返回下一页的数据。
// 缺点：1. 不支持跳页（但现在流行无限滑动翻页）；2. 不太适合多检索条件的场景。
