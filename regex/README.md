# Regex CheatSheets

```sh
echo '
readme.md
document.pdf
image.png
 | grep -P '^\w+\.pdf$'
```

## Dot: `.` 

任意字符包括空格和特殊字符

```sh
echo 'abcABC123 .:!?' | grep -P '.'
# abcABC123 .:!?
```

## Character Sets: `[abc]`

```sh
echo 'bar ber bir bor bur' | grep -P 'b[aeio]r'
# bar ber bir bor
```

## Negated Character Sets: `[^abc]`

```sh
echo 'bar ber bir bor bur' | grep -P 'b[^eo]r'
# ber bor
```

## Letter Range: `[a-z]`

```sh
echo 'abcdefghijklmnopqrstuvwxyz' | grep -P '[e-o]'
# efghijklmno
```

## Number Range: `[0-9]`

```sh
echo '0123456789' | grep -P '[3-6]'
# 3456
```

---

## Asterisk: `*`

`{0,}`

```sh
echo 'br ber beer' | grep -P 'be*r'
# br ber beer
```

## Plus Sign `+`

`{1,}`

```sh
echo 'br ber beer' | grep -P 'be+r'
# ber beer
```

## Question Mark `?`

`{0,1}`

```sh
echo 'color colour' | grep -P 'colou?r'
# color colour
```

## Curly Braces `{m,n}`

```sh
echo 'ber beer beeer beeeer' | grep -P 'be{2}r'
# beer 

echo 'ber beer beeer beeeer' | grep -P 'be{3,}r'
# beeer beeeer

echo 'ber beer beeer beeeer' | grep -P 'be{1,3}r'
# ber beer beeer
```

```sh
echo 'Release 10/9/2021' | grep -P '[0-9]{4}'
# 2021

echo 'Release 10/9/2021' | grep -P '[0-9]{2,}'
# 10, 2021

echo 'Release 10/9/2021' | grep -P '[0-9]{1,4}'
# 10, 9, 2021
```

---

## Parentheses `()` Grouping

```sh
echo 'ha-ha,haa-haa' | grep -P '(haa)'
# haa, haa
```

## Group Reference

引用分组

```sh
echo 'ha-ha,haa-haa' | grep -P '(ha)-\1,(haa)-\2'
```

## `(?:)` Skip Group Reference

不捕获某个组

```sh
echo 'ha-ha,haa-haa' | grep -P '(?:ha)-ha,(haa)-\1'
```

## Pipe Character `|`

`|` 和 `[abc]` 的区别?

`[abc]` 单个字符级别, `|` 是表达式级别(支持多个字符!)

```sh
echo 'cat,Cat,rat' | grep -P '[c|C]at|rat'
# cat,Cat,rat
```

## Escape Character `\`

转义元字符

```sh
echo '(*) Asteriks.' | grep -P '(\*|\.)'
# *,.
```

## Caret Sign `^` Line Start

```sh
echo 'Basic Omellette Recipe

1. 3 eggs, beaten
2. 1 tsp sunflower oil
3. 1 tsp butter' | grep -P '^[0-9]'
```

## Dollar Sign `$` End of Line

```sh
echo 'https://domain.com/what-is-html.html
https://otherdomain.com/html-elements
https://website.com/html5-features.html' | grep -P 'html$'
```

---

## `\w` Word Character

Letter, Number, Underscore

```sh
echo 'abcABC123 _.:!?' | grep -P '\w'
# a,b,c,A,B,C,1,2,3,_
```

## `\W` Except Word Character

```sh
echo 'abcABC123 _.:!?' | grep -P '\w'
# . :!?
```

## Other Character

`\d` Number Character

`\D` Except Number Character

`\s` Space Character

`\S` Except Space Character

---

## `(?=)` Positive Lookahead

```sh
echo 'Date: 4 Aug 3PM' | grep -P '\d+(?=PM)'
# 3
```

## `(?!)` Negative Lookahead

```sh
echo 'Date: 4 Aug 3PM' | grep -P '\d+(?!PM)'
# 4
```

## `?<=` Positive Lookbehind

```sh
echo 'Product Code: 1064 Price: $5' | grep -P '(?<=\$)\d+'
# 5
```

## `?<!` Negative Lookbehind

```sh
echo 'Product Code: 1064 Price: $5' | grep -P '(?<!\$)\d+'
```

---

## Flags, modifiers

```sh
# Global Flag: select all matches
'domain.com, test.com, site.com'
/\w+\.com/g

# Multiline Flag
'domain.com
test.com
site.com' 
/\w+\.com/gm

# Case-insensitive Flag
'DOMAIN.COM
TEST.COM
SITE.COM'
/\w+\.com$/gmi
```

---

## Greedy Matching

越长越好

RegEx的默认行为

```sh
echo 'ber beer beeer beeeer' | grep -P '.*r'
# ber beer beeer beeeer
```

## Lazy Matching

`*?`

```sh
'ber beer beeer beeeer'
.*?r
/\w*?r/
# ber
```

# See Also

- [regex101](https://regex101.com/)
- [cheatsheet](https://regexlearn.com/cheatsheet)
