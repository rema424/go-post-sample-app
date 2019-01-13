.PHONY: all

# コマンドライン引数で渡ってきたらその値で上書きする
project_id := ${PROJECT_ID}
version := ${GAE_VERSION}

# シェルコマンドの実行結果を代入する時は$でエスケープする
OAUTH2_ACCESS_TOKEN=$$(gcloud auth print-access-token 2> /dev/null)

dlv:
	dlv debug --headless --listen=:2345 --api-version=2

# ローカルサーバーの起動
# e.g) make serve
serve:
	goapp serve --host 0.0.0.0 main

# デプロイ
# e.g) make update version=1
update:
	appcfg.py update -A ${project_id} -V ${version} --oauth2_access_token=${OAUTH2_ACCESS_TOKEN} main

# トラフィックを徐々に切り替える (500が発生しない 基本的にはこっち)
# e.g) make migrate version=1
migrate:
	appcfg.py migrate_traffic -A ${project_id} -V ${version} --oauth2_access_token=${OAUTH2_ACCESS_TOKEN} main

# トラフィックを一瞬で切り替える (一瞬500が発生する)
# e.g) make set version=1
set:
	appcfg.py set_default_version -A ${project_id} -V ${version} --oauth2_access_token=${OAUTH2_ACCESS_TOKEN} main

# e.g) make open version=1
open:
	gcloud app browse --version ${version}