FROM postgres:15.0-bullseye

# 日本語の文字列ソートが以下のようになるようLOCATEを指定
# あ
# あ(テスト)
# い
# い(テスト)
RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
