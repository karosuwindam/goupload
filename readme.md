# Web Uploader
Golang製のファイルアップローダ

# URLによるAPIについて

## /upload/

index.htmlの結果を表示するだけ

## /upload/key/keyword

|key|keyword|説明|
|:--|:--|:--|
|search|文字入力|文字入力対するファイル名が存在することを確認する<br>大文字、小文字は関係なく小文字共通で比較する。<br> 結果はJSON出力で、flagのKey名で1または0を返す|


## /uploadfile
指定フォルダ内のファイルリストををJSON出力する。 \
現在は、ファイル名とファイルサイズ

# ソースコード説明

|フォルダ名|package|説明|
|:--|:--|:--|
|dirread|dirread|指定フォルダ内のファイルリストを作る|
|html|-|HTMLファイル|
|logoutput|logoutput|ログ出力を作成する|
|upload|-|Upload先フォルダ|

