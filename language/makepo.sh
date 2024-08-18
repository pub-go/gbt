#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname $0)" && pwd)"
cd $SCRIPT_DIR

echo "usage: $0 [language-code]  用法：$0 [语言代码]"
echo "  extract messages from go source file       从 Go 源代码中抽取翻译字符串"
echo "  and make/update translate template file    并生成/更新给定语言的翻译文件"

# tree
#    -I 忽略指定的 pattern
#    -P 指定文件 pattern
#    -F 在文件夹后添加斜线结尾
#    -f 给每个文件添加路径前缀
#    -i 不要输入树形 直接每行一个文件
#    --noreport 不要输入统计信息: x directories, y files
# grep
#    -v 排除以斜线结尾的
tree -I "cmd" -P "*.go" -F -f -i --noreport .. | grep -v /$ | grep -v ^../$ |
    xargs xgettext -C --add-comments=TRANSLATORS: --force-po -kT -kN:1,2 -kX:2,1c -kXN:2,3,1c -o messages.pot

if [ $# -eq 0 ]; then
    echo "language-code not provide, exit               未提供语言代码，退出"
else
    if [ -f "$1.po" ]; then
        # 更新模板
        echo "update/更新 $1.po ..."
        msgmerge -U $1.po messages.pot
    else
        # 初始化翻译
        echo "create/创建 $1.po ..."
        GETTEXTCLDRDIR=./cldr msginit -i messages.pot -l $1 -o $1.po
    fi
fi
