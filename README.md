# entry-exit-management
## What
- ICカードを使って人の入退出を管理するアプリケーション

## Will
Android App で入退出をチェックし Google Calendar に記録する
- felica のために Android App を使う
- Google API のために Web API 作って使いやすくする
- ブラウザからユーザ情報を管理したいので Webアプリ を作る
- 個人情報を扱うので TLS 必須，脆弱性スキャンもする
- プロメテウスを使って監視する ー＞ アラートは Slack へ
- 複数のコンテナを使うので Kubernetes 上で動かしたい
- プラットフォームは AKS か GKE （研究室で使うなら AKS の方がお金の処理が楽かも）
- タスク管理は jira を使う

## 構成
![[pdf/eem-full-map.pdf]]
- nginx・・・リバースプロキシ <- ingress がリバースプロキシのような機能を提供している
- node(Next.js)・・・Webアプリケーション -> サーバサイドレンダリングがしたいので React から Next.js を使うように変更
- Gin(golang)・・・Web API
- mariadb・・・データベース
- マニフェスト・・・https://github.com/munvei/eem-manifests で管理する
- Android App・・・別リポジトリで管理する予定
- CI/CD・・・CircleCI と AlgoCD を使う
