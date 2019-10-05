# How To Use

## 1. 必要なソフトウェアのインストール

### Docker

```
/* HomeBrew更新 */
$ brew update

/* Dockerインストール */
$ brew install docker
$ brew cask install docker
```

下のコマンドを実行してバージョン情報が出力されればOK  

```
/* インストール確認 */
$ docker --version
```

インストール確認後、Dockerを起動する
Launchpadから起動してもOK

```
/* Docker起動 */
$ open /Applications/Docker.app
```

メニューバーのDockerアイコンをクリックし、`Docker Desktop is running`が表示されていればOK

## 2. セットアップ

### マスタダウンロード

マスタ実行ファイル, 設定ファイルをダウンロードする

[Google Drive](https://drive.google.com/drive/folders/14YnVFhW6lRltP7_dhXOrQTG1WVVDBkUT?usp=sharing)  

ダウンロードしたzipファイルは任意のディレクトリに展開する  

### ソルバ

マスタを展開したディレクトリにソルバを配置する  

### Dockerイメージ

下のコマンドを実行する  
実行時間目安 : 3~4分

```
/* ベースイメージ */
$ make docker-build-base

/* ソルバイメージ */
$ make docekr-build-solver SOLVER_IMAGE=<image> SOURCE_PY=<solver>

// <image> = procon30-solver:***
// <solver> = ソルバのファイル名
// 例: SOLVER_IMAGE=procon30-solver:test SOURCE_PY=solver.py
```

実行終了後に下のコマンドを実行し `procon30-sovler` が存在するか確認する

```
/* Dockerイメージ表示 */
$ docker images
REPOSITORY            TAG                    IMAGE ID            CREATED             SIZE
procon30-solver       latest                 3f14358ade7b        56 minutes ago      491MB
alpine                procon30-solver-base   1f3e832f5e4b        42 hours ago        491MB
alpine                latest                 caf27325b298        8 months ago        5.53M
```

### 設定ファイル

`config.toml` をエディタで開き、下のようにサーバURLやトークンを入力する

```
[GameServer]
url = "ゲームサーバURL"
token = "トークン"

[Solver]
image = "procon30-solver" <- 変更しない
```

## 3. 起動

下のコマンドを実行する

```
/* マスタ実行(MacOS) */
$ ./procon30_yuge_kyogi_darwin
```

## 4. 試合

基本的に操作の必要はなし  
エラーが表示された場合は別紙「errorcode.md」を参照して対処する

### コマンド

マスタ実行中にコマンドを入力することで動作指示が可能

#### token \<Token\>

トークン変更

#### solver \<SolverImage\>

使用するソルバイメージを変更

#### viewer \<BattleID\>

ビューワ起動

#### refresh

試合情報再取得

#### exit, q

マスタ終了
