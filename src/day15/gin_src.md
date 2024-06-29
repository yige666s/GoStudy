# Gin原理解析
## Gin路由
### 基数树
基数树（Radix Tree），也叫压缩前缀树（Compressed Prefix Tree）或 PAT 位树（Patricia Trie），是一种优化后的前缀树（Trie），用来存储字符串集合。与标准的前缀树不同，基数树通过合并单一路径上的节点来节省空间。具体来说，它有以下特点：  
 路径压缩：在基数树中，如果某个节点只有一个子节点，那么该节点和它的唯一子节点会合并。这意味着中间节点（只有一个子节点的节点）不会单独存在，路径上的这些节点会被压缩成一个节点，保存这个路径上所有的字符。  
 节省空间：由于路径压缩，基数树减少了冗余节点，从而比标准的前缀树更节省空间。标准前缀树会为每个字符分配一个节点，而基数树则将唯一路径上的字符合并到一个节点中。

举个例子，假设我们要在基数树中存储以下字符串集合：["apple", "app", "apricot"]。
```
标准前缀树表示如下：
      root
      /  
     a
     |
     p
    / \
   p   r
   |   |
   l   i
   |   |
   e   c
       |
       o
       |
       t
基数树表示如下：
      root
      /  
     a
     |
     p
    / \
  "ple" "pricot"
```
在基数树中，"apple" 和 "app" 的共同前缀 "ap" 之后，"apple" 的剩余部分 "ple" 合并成一个节点。"apricot" 在前缀 "ap" 之后独立出来。通过这种方式，基数树减少了中间节点的数量，节省了空间，并提高了查找效率。

### 路由树节点

```
// tree.go

type node struct {
   // 节点路径，比如上面的s，earch，和upport
	path      string
	// 和children字段对应, 保存的是分裂的分支的第一个字符
	// 例如search和support, 那么s节点的indices对应的"eu"
	// 代表有两个分支, 分支的首字母分别是e和u
	indices   string
	// 儿子节点
	children  []*node
	// 处理函数链条（切片）
	handlers  HandlersChain
	// 优先级，子节点、子子节点等注册的handler数量
	priority  uint32
	// 节点类型，包括static, root, param, catchAll
	// static: 静态节点（默认），比如上面的s，earch等节点
	// root: 树的根节点
	// catchAll: 有*匹配的节点
	// param: 参数节点
	nType     nodeType
	// 路径上最大参数个数
	maxParams uint8
	// 节点是否是参数节点，比如上面的:post
	wildChild bool
	// 完整路径
	fullPath  string
}
```

### 请求方法树
在gin的路由中，每一个HTTP Method(GET、POST、PUT、DELETE…)都对应了一棵 radix tree，我们注册路由的时候会调用下面的addRoute函数：  
```
// gin.go
func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
   // liwenzhou.com...
   
   // 获取请求方法对应的树
	root := engine.trees.get(method)
	if root == nil {
	
	   // 如果没有就创建一个
		root = new(node)
		root.fullPath = "/"
		engine.trees = append(engine.trees, methodTree{method: method, root: root})
	}
	root.addRoute(path, handlers)
}
```

### 注册路由
注册路由的逻辑主要有addRoute函数和insertChild方法。  
#### addRoute
```
// tree.go

// addRoute 将具有给定句柄的节点添加到路径中。
// 不是并发安全的
func (n *node) addRoute(path string, handlers HandlersChain) {
	fullPath := path
	n.priority++
	numParams := countParams(path)  // 数一下参数个数

	// 空树就直接插入当前节点
	if len(n.path) == 0 && len(n.children) == 0 {
		n.insertChild(numParams, path, fullPath, handlers)
		n.nType = root
		return
	}

	parentFullPathIndex := 0

walk:
	for {
		// 更新当前节点的最大参数个数
		if numParams > n.maxParams {
			n.maxParams = numParams
		}

		// 找到最长的通用前缀
		// 这也意味着公共前缀不包含“:”"或“*” /
		// 因为现有键不能包含这些字符。
		i := longestCommonPrefix(path, n.path)

		// 分裂边缘（此处分裂的是当前树节点）
		// 例如一开始path是search，新加入support，s是他们通用的最长前缀部分
		// 那么会将s拿出来作为parent节点，增加earch和upport作为child节点
		if i < len(n.path) {
			child := node{
				path:      n.path[i:],  // 公共前缀后的部分作为子节点
				wildChild: n.wildChild,
				indices:   n.indices,
				children:  n.children,
				handlers:  n.handlers,
				priority:  n.priority - 1, //子节点优先级-1
				fullPath:  n.fullPath,
			}

			// Update maxParams (max of all children)
			for _, v := range child.children {
				if v.maxParams > child.maxParams {
					child.maxParams = v.maxParams
				}
			}

			n.children = []*node{&child}
			// []byte for proper unicode char conversion, see #65
			n.indices = string([]byte{n.path[i]})
			n.path = path[:i]
			n.handlers = nil
			n.wildChild = false
			n.fullPath = fullPath[:parentFullPathIndex+i]
		}

		// 将新来的节点插入新的parent节点作为子节点
		if i < len(path) {
			path = path[i:]

			if n.wildChild {  // 如果是参数节点
				parentFullPathIndex += len(n.path)
				n = n.children[0]
				n.priority++

				// Update maxParams of the child node
				if numParams > n.maxParams {
					n.maxParams = numParams
				}
				numParams--

				// 检查通配符是否匹配
				if len(path) >= len(n.path) && n.path == path[:len(n.path)] {
					// 检查更长的通配符, 例如 :name and :names
					if len(n.path) >= len(path) || path[len(n.path)] == '/' {
						continue walk
					}
				}

				pathSeg := path
				if n.nType != catchAll {
					pathSeg = strings.SplitN(path, "/", 2)[0]
				}
				prefix := fullPath[:strings.Index(fullPath, pathSeg)] + n.path
				panic("'" + pathSeg +
					"' in new path '" + fullPath +
					"' conflicts with existing wildcard '" + n.path +
					"' in existing prefix '" + prefix +
					"'")
			}
			// 取path首字母，用来与indices做比较
			c := path[0]

			// 处理参数后加斜线情况
			if n.nType == param && c == '/' && len(n.children) == 1 {
				parentFullPathIndex += len(n.path)
				n = n.children[0]
				n.priority++
				continue walk
			}

			// 检查路path下一个字节的子节点是否存在
			// 比如s的子节点现在是earch和upport，indices为eu
			// 如果新加一个路由为super，那么就是和upport有匹配的部分u，将继续分列现在的upport节点
			for i, max := 0, len(n.indices); i < max; i++ {
				if c == n.indices[i] {
					parentFullPathIndex += len(n.path)
					i = n.incrementChildPrio(i)
					n = n.children[i]
					continue walk
				}
			}

			// 否则就插入
			if c != ':' && c != '*' {
				// []byte for proper unicode char conversion, see #65
				// 注意这里是直接拼接第一个字符到n.indices
				n.indices += string([]byte{c})
				child := &node{
					maxParams: numParams,
					fullPath:  fullPath,
				}
				// 追加子节点
				n.children = append(n.children, child)
				n.incrementChildPrio(len(n.indices) - 1)
				n = child
			}
			n.insertChild(numParams, path, fullPath, handlers)
			return
		}

		// 已经注册过的节点
		if n.handlers != nil {
			panic("handlers are already registered for path '" + fullPath + "'")
		}
		n.handlers = handlers
		return
	}
}
```
![路由构造](https://www.liwenzhou.com/images/Go/gin/addroute.gif)
#### insertChild

```
// tree.go
func (n *node) insertChild(numParams uint8, path string, fullPath string, handlers HandlersChain) {
  // 找到所有的参数
	for numParams > 0 {
		// 查找前缀直到第一个通配符
		wildcard, i, valid := findWildcard(path)
		if i < 0 { // 没有发现通配符
			break
		}

		// 通配符的名称必须包含':' 和 '*'
		if !valid {
			panic("only one wildcard per path segment is allowed, has: '" +
				wildcard + "' in path '" + fullPath + "'")
		}

		// 检查通配符是否有名称
		if len(wildcard) < 2 {
			panic("wildcards must be named with a non-empty name in path '" + fullPath + "'")
		}

		// 检查这个节点是否有已经存在的子节点
		// 如果我们在这里插入通配符，这些子节点将无法访问
		if len(n.children) > 0 {
			panic("wildcard segment '" + wildcard +
				"' conflicts with existing children in path '" + fullPath + "'")
		}

		if wildcard[0] == ':' { // param
			if i > 0 {
				// 在当前通配符之前插入前缀
				n.path = path[:i]
				path = path[i:]
			}

			n.wildChild = true
			child := &node{
				nType:     param,
				path:      wildcard,
				maxParams: numParams,
				fullPath:  fullPath,
			}
			n.children = []*node{child}
			n = child
			n.priority++
			numParams--

			// 如果路径没有以通配符结束
			// 那么将有另一个以'/'开始的非通配符子路径。
			if len(wildcard) < len(path) {
				path = path[len(wildcard):]

				child := &node{
					maxParams: numParams,
					priority:  1,
					fullPath:  fullPath,
				}
				n.children = []*node{child}
				n = child  // 继续下一轮循环
				continue
			}

			// 否则我们就完成了。将处理函数插入新叶子中
			n.handlers = handlers
			return
		}

		// catchAll
		if i+len(wildcard) != len(path) || numParams > 1 {
			panic("catch-all routes are only allowed at the end of the path in path '" + fullPath + "'")
		}

		if len(n.path) > 0 && n.path[len(n.path)-1] == '/' {
			panic("catch-all conflicts with existing handle for the path segment root in path '" + fullPath + "'")
		}

		// currently fixed width 1 for '/'
		i--
		if path[i] != '/' {
			panic("no / before catch-all in path '" + fullPath + "'")
		}

		n.path = path[:i]
		
		// 第一个节点:路径为空的catchAll节点
		child := &node{
			wildChild: true,
			nType:     catchAll,
			maxParams: 1,
			fullPath:  fullPath,
		}
		// 更新父节点的maxParams
		if n.maxParams < 1 {
			n.maxParams = 1
		}
		n.children = []*node{child}
		n.indices = string('/')
		n = child
		n.priority++

		// 第二个节点:保存变量的节点
		child = &node{
			path:      path[i:],
			nType:     catchAll,
			maxParams: 1,
			handlers:  handlers,
			priority:  1,
			fullPath:  fullPath,
		}
		n.children = []*node{child}

		return
	}

	// 如果没有找到通配符，只需插入路径和句柄
	n.path = path
	n.handlers = handlers
	n.fullPath = fullPath
}
```  
insertChild函数是根据path本身进行分割，将/分开的部分分别作为节点保存，形成一棵树结构。参数匹配中的:和*的区别是，前者是匹配一个字段而后者是匹配后面所有的路径。  
### 路由匹配
我们先来看gin框架处理请求的入口函数ServeHTTP：
```
// gin.go
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  // 这里使用了对象池
	c := engine.pool.Get().(*Context)
  // 这里有一个细节就是Get对象后做初始化
	c.writermem.reset(w)
	c.Request = req
	c.reset()

	engine.handleHTTPRequest(c)  // 我们要找的处理HTTP请求的函数

	engine.pool.Put(c)  // 处理完请求后将对象放回池子
}
```
精简版代码逻辑:
```
// gin.go
func (engine *Engine) handleHTTPRequest(c *Context) {
	// liwenzhou.com...

	// 根据请求方法找到对应的路由树
	t := engine.trees
	for i, tl := 0, len(t); i < tl; i++ {
		if t[i].method != httpMethod {
			continue
		}
		root := t[i].root
		// 在路由树中根据path查找
		value := root.getValue(rPath, c.Params, unescape)
		if value.handlers != nil {
			c.handlers = value.handlers
			c.Params = value.params
			c.fullPath = value.fullPath
			c.Next()  // 执行函数链条
			c.writermem.WriteHeaderNow()
			return
		}
	
	// liwenzhou.com...
	c.handlers = engine.allNoRoute
	serveError(c, http.StatusNotFound, default404Body)
}
```
路由匹配是由节点的 getValue方法实现的。getValue根据给定的路径(键)返回nodeValue值，保存注册的处理函数和匹配到的路径参数数据。  
如果找不到任何处理函数，则会尝试TSR(尾随斜杠重定向)。  
下面是getValue方法：
```
// tree.go

type nodeValue struct {
	handlers HandlersChain
	params   Params  // []Param
	tsr      bool
	fullPath string
}

// liwenzhou.com...

func (n *node) getValue(path string, po Params, unescape bool) (value nodeValue) {
	value.params = po
walk: // Outer loop for walking the tree
	for {
		prefix := n.path
		if path == prefix {
			// 我们应该已经到达包含处理函数的节点。
			// 检查该节点是否注册有处理函数
			if value.handlers = n.handlers; value.handlers != nil {
				value.fullPath = n.fullPath
				return
			}

			if path == "/" && n.wildChild && n.nType != root {
				value.tsr = true
				return
			}

			// 没有找到处理函数 检查这个路径末尾+/ 是否存在注册函数
			indices := n.indices
			for i, max := 0, len(indices); i < max; i++ {
				if indices[i] == '/' {
					n = n.children[i]
					value.tsr = (len(n.path) == 1 && n.handlers != nil) ||
						(n.nType == catchAll && n.children[0].handlers != nil)
					return
				}
			}

			return
		}

		if len(path) > len(prefix) && path[:len(prefix)] == prefix {
			path = path[len(prefix):]
			// 如果该节点没有通配符(param或catchAll)子节点
			// 我们可以继续查找下一个子节点
			if !n.wildChild {
				c := path[0]
				indices := n.indices
				for i, max := 0, len(indices); i < max; i++ {
					if c == indices[i] {
						n = n.children[i] // 遍历树
						continue walk
					}
				}

				// 没找到
				// 如果存在一个相同的URL但没有末尾/的叶子节点
				// 我们可以建议重定向到那里
				value.tsr = path == "/" && n.handlers != nil
				return
			}

			// 根据节点类型处理通配符子节点
			n = n.children[0]
			switch n.nType {
			case param:
				// find param end (either '/' or path end)
				end := 0
				for end < len(path) && path[end] != '/' {
					end++
				}

				// 保存通配符的值
				if cap(value.params) < int(n.maxParams) {
					value.params = make(Params, 0, n.maxParams)
				}
				i := len(value.params)
				value.params = value.params[:i+1] // 在预先分配的容量内扩展slice
				value.params[i].Key = n.path[1:]
				val := path[:end]
				if unescape {
					var err error
					if value.params[i].Value, err = url.QueryUnescape(val); err != nil {
						value.params[i].Value = val // fallback, in case of error
					}
				} else {
					value.params[i].Value = val
				}

				// 继续向下查询
				if end < len(path) {
					if len(n.children) > 0 {
						path = path[end:]
						n = n.children[0]
						continue walk
					}

					// ... but we can't
					value.tsr = len(path) == end+1
					return
				}

				if value.handlers = n.handlers; value.handlers != nil {
					value.fullPath = n.fullPath
					return
				}
				if len(n.children) == 1 {
					// 没有找到处理函数. 检查此路径末尾加/的路由是否存在注册函数
					// 用于 TSR 推荐
					n = n.children[0]
					value.tsr = n.path == "/" && n.handlers != nil
				}
				return

			case catchAll:
				// 保存通配符的值
				if cap(value.params) < int(n.maxParams) {
					value.params = make(Params, 0, n.maxParams)
				}
				i := len(value.params)
				value.params = value.params[:i+1] // 在预先分配的容量内扩展slice
				value.params[i].Key = n.path[2:]
				if unescape {
					var err error
					if value.params[i].Value, err = url.QueryUnescape(path); err != nil {
						value.params[i].Value = path // fallback, in case of error
					}
				} else {
					value.params[i].Value = path
				}

				value.handlers = n.handlers
				value.fullPath = n.fullPath
				return

			default:
				panic("invalid node type")
			}
		}

		// 找不到，如果存在一个在当前路径最后添加/的路由
		// 我们会建议重定向到那里
		value.tsr = (path == "/") ||
			(len(prefix) == len(path)+1 && prefix[len(path)] == '/' &&
				path == prefix[:len(prefix)-1] && n.handlers != nil)
		return
	}
}
```
## Gin中间件
gin框架涉及中间件相关有4个常用的方法，它们分别是c.Next()、c.Abort()、c.Set()、c.Get()。
### 中间件注册
我们注册路由时会将对应路由的函数和之前的中间件函数结合到一起：
```
func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)  // 将处理请求的函数与中间件函数结合
	group.engine.addRoute(httpMethod, absolutePath, handlers)
	return group.returnObj()
}
```
我们会将一个路由的中间件函数和处理函数结合到一起组成一条处理函数链条HandlersChain，而它本质上就是一个由HandlerFunc组成的切片  
```
const abortIndex int8 = math.MaxInt8 / 2

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	if finalSize >= int(abortIndex) {  // 这里有一个最大限制
		panic("too many handlers")
	}
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}
```
### 中间件执行
其中c.Next()就是很关键的一步，它的代码很简单：
```
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}
```
从上面的代码可以看到，这里通过索引遍历HandlersChain链条，从而实现依次调用该路由的每一个函数（中间件或处理请求的函数）。  
我们可以在中间件函数中通过再次调用c.Next()实现嵌套调用（func1中调用func2；func2中调用func3），  
或者通过调用c.Abort()中断整个调用链条，从当前函数返回。
```
func (c *Context) Abort() {
	c.index = abortIndex  // 直接将索引置为最大限制值，从而退出循环
}
```

### c.Set()/c.Get()
c.Set()和c.Get()这两个方法多用于在多个函数之间通过c传递数据的，比如我们可以在认证中间件中获取当前请求的相关信息（userID等）通过c.Set()存入c，然后在后续处理业务逻辑的函数中通过c.Get()来获取当前请求的用户。c就像是一根绳子，将该次请求相关的所有的函数都串起来了。

![调用链路](https://www.liwenzhou.com/images/Go/gin/gin_middleware3.png)

## 总结
1. gin框架路由使用前缀树，路由注册的过程是构造前缀树的过程，路由匹配的过程就是查找前缀2树的过程。
2. gin框架的中间件函数和处理函数是以切片形式的调用链条存在的，我们可以顺序调用也可以借助c.Next()方法实现嵌套调用。
3. 借助c.Set()和c.Get()方法我们能够在不同的中间件函数中传递数据。
