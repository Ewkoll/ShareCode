# Git Command 一级标题
## 二级
### 三级
#### 司机
##### 五级
##### 六级

- dsd
- dsd
- dsd
1. S1
2. S2
3. S3

*强调斜体*

**粗体效果**

* 居然和-一个效果

* Html块内的Makedown是不会被处理的。
* 出现横线是个没理解的结果，应该是<H1&></H1>的效果
* H2前面有/和没有居然还有差别，能够影响字体。

<H1><H3> 
第一个H1产生横线，后面跟进个H1影响字体大小。
<table>
    <tr>
        <td>Foo</td> 
        <td>Fo1</td>
    </tr>
    <tr>
        <td>Foo</td>
        <td>Fo1</td>
    </tr>
    <tr>
        <td>Foo</td>
        <td>Fo1</td>
    </tr>
</table>
<H1><H3> 

* S1

4 < 5

This is an H1
============

This is an H2
-------------

* Test
# 区块引用

> ## 这是一个标题。
> 
> 1.   这是第一行列表项。
> 2.   这是第二行列表项。
> 
> 给出一些例子代码：
> 
>     return shell_exec("echo $input | $markdown_script");
>> 这个是第二级区块。

*   Lorem ipsum dolor sit amet, consectetuer adipiscing elit.
    Aliquam hendrerit mi posuere lectus. Vestibulum enim wisi,
    viverra nec, fringilla in, laoreet vitae, risus.
*   Donec sit amet nisl. Aliquam semper ipsum sit amet velit.
    Suspendisse id sem consectetuer libero luctus adipiscing.


<ol>
    <li>Bird</li>
    <li>McHale</li>
    <li>Parish</li>
</ol>

<li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.
    Aliquam hendrerit mi posuere lectus. Vestibulum enim wisi,
    viverra nec, fringilla in, laoreet vitae, risus.
</li>
<li>
    Donec sit amet nisl. Aliquam semper ipsum sit amet velit.
    Suspendisse id sem consectetuer libero luctus adipiscing.
</li>

<H1></H1>

*   Bird
*   Magic

<ul>
<li>Bird</li>
<li>Magic</li>
</ul>

<H1></H1>

*   Bird

*   Magic

<ul>
<li><p>Bird</p></li>
<li><p>Magic</p></li>
</ul>

<H1></H1>

*   一列表项包含一个列表区块：

*   A list item with a blockquote:

    > This is a blockquote
    > inside a list item.


Use the `printf()` function.

<p>Use the <code>printf()</code> function.</p>

``There is a literal backtick (`) here.``
<p><code>There is a literal backtick (`) here.</code></p>

A single backtick in a code span: `` ` ``

A backtick-delimited string in a code span: `` `foo` ``

# 图片

很明显地，要在纯文字应用中设计一个「自然」的语法来插入图片是有一定难度的。

Markdown 使用一种和链接很相似的语法来标记图片，同样也允许两种样式： 行内式和参考式。

行内式的图片语法看起来像是：

![Alt text](E:/code/1.png "Just a image")

详细叙述如下：

一个惊叹号 !
接着一个方括号，里面放上图片的替代文字
接着一个普通括号，里面放上图片的网址，最后还可以用引号包住并加上 选择性的 'title' 文字。
参考式的图片语法则长得像这样：

![Alt text][id]
「id」是图片参考的名称，图片参考的定义方式则和连结参考一样：

[id]: url/to/image  "Optional title attribute"
到目前为止， Markdown 还没有办法指定图片的宽高，如果你需要的话，你可以使用普通的 <img> 标签。

其它

自动链接

Markdown 支持以比较简短的自动链接形式来处理网址和电子邮件信箱，只要是用尖括号包起来， Markdown 就会自动把它转成链接。一般网址的链接文字就和链接地址一样，例如：

<http://example.com/>
Markdown 会转为：
<a href="http://example.com/">http://example.com/</a>
邮址的自动链接也很类似，只是 Markdown 会先做一个编码转换的过程，把文字字符转成 16 进位码的 HTML 实体，这样的格式可以糊弄一些不好的邮址收集机器人，例如：

<address@example.com>Markdown 会转成：
<a href="&#x6D;&#x61;i&#x6C;&#x74;&#x6F;:&#x61;&#x64;&#x64;&#x72;&#x65;
&#115;&#115;&#64;&#101;&#120;&#x61;&#109;&#x70;&#x6C;e&#x2E;&#99;&#111;
&#109;">&#x61;&#x64;&#x64;&#x72;&#x65;&#115;&#115;&#64;&#101;&#120;&#x61;
&#109;&#x70;&#x6C;e&#x2E;&#99;&#111;&#109;</a>
在浏览器里面，这段字串（其实是 <a href="mailto:address@example.com">address@example.com</a>）会变成一个可以点击的「address@example.com」链接。

（这种作法虽然可以糊弄不少的机器人，但并不能全部挡下来，不过总比什么都不做好些。不管怎样，公开你的信箱终究会引来广告信件的。）

# 反斜杠

Markdown 可以利用反斜杠来插入一些在语法中有其它意义的符号，例如：如果你想要用星号加在文字旁边的方式来做出强调效果（但不用 <em> 标签），你可以在星号的前面加上反斜杠：

\*literal asterisks\*
Markdown 支持以下这些符号前面加上反斜杠来帮助插入普通的符号：

\   反斜线
`   反引号
*   星号
_   底线
{}  花括号
[]  方括号
()  括弧
#   井字号
+   加号
-   减号
.   英文句点
!   惊叹号
