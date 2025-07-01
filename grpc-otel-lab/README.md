# 🏃‍♂️ Lab 04 – Go gRPC ストリーミング & OpenTelemetry 動作確認手順

> **対象ブランチ**: `lab/04-grpc-otel`
>
> この記事のハンズオン環境をワンクリックまたは 3 コマンドで再現します。

---

## 0. 前提ツール

| ツール                  | バージョン        | 用途                         |
| -------------------- | ------------ | -------------------------- |
| **Go**               | 1.24 以上      | buf の go‑install & ローカル実行  |
| **Docker / Compose** | 20.10+ / v2+ | Jaeger + gRPC サービス起動       |
| **buf CLI**          | 最新           | .proto 生成（未インストールなら手順①で導入） |

---

## 1. クローン & buf インストール

```bash
# ブランチごとクローン
git clone -b lab/04-grpc-otel https://github.com/your-org/go-blog-examples.git
cd go-blog-examples

# buf が無い場合は go install 一発
command -v buf >/dev/null 2>&1 || {
  echo "Installing buf ...";
  go install github.com/bufbuild/buf/cmd/buf@latest;
  export PATH="$PATH:$(go env GOPATH)/bin";
}

# Proto → Go 生成
buf generate
```

---

## 2. Docker Compose で一発起動

```bash
# バックグラウンド起動 (-d) でサーバ + Jaeger を常駐
docker compose -f grpc-otel-lab/docker-compose.yml up --build -d
```

| コンテナ       | ポート   | 役割                                      |
| ---------- | ----- | --------------------------------------- |
| **server** | 50051 | gRPC 双方向チャットサーバ                         |
| **client** | —     | `wait-for-it` で server:50051 を検知 → 自動実行 |
| **jaeger** | 16686 | トレース UI                                 |

> **client コンテナが終了している場合**は↓
>
> ```bash
> # 新しいシェルを開き、内部で go run する方法
> docker compose -f grpc-otel-lab/docker-compose.yml run --service-ports --entrypoint bash client
> go run ./grpc-otel-lab/cmd/client   # ← コンテナ内で実行
> ```

### 2.1 ローカル端末でクライアントを動かす場合

サーバ & Jaeger をコンテナ、クライアントだけローカルにしたい場合は接続先を `localhost:50051` に書き換えてから実行します。

```bash
# サーバ & Jaeger を起動
docker compose -f grpc-otel-lab/docker-compose.yml up -d server jaeger
# ローカルターミナルでクライアント
go run grpc-otel-lab/cmd/client   # 接続先=localhost:50051
```

## 3. Jaeger UI でトレース確認. Jaeger UI でトレース確認

1. ブラウザで [http://localhost:16686](http://localhost:16686) を開く。
2. **Search → Service** に `chat-server` または `chat-client` を選択。
3. トレース一覧から最新をクリック → **Span** にメッセージ交換が時系列で並ぶ。

---

## 4. ローカル実行派向けワンライナー

```bash
# サーバ
go run grpc-otel-lab/cmd/server &
# クライアント（別ターミナル）
go run grpc-otel-lab/cmd/client
```

OpenTelemetry エクスポータのデフォルト先 `localhost:14250` に Jaeger がない場合は環境変数で URL を上書きしてください。

---

## 5. よくあるエラーと対処

| 症状                                             | 原因                                 | 対処                                                                                                                                  |
| ---------------------------------------------- | ---------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `buf: command not found`                       | buf 未インストール                        | `go install github.com/bufbuild/buf/cmd/buf@latest`                                                                                 |
| `protoc-gen-go.* not found`                    | 生成プラグイン未導入                         | `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest` |
| `name resolver error: produced zero addresses` | ローカル実行時に `server:50051` へ接続        | 接続先を `localhost:50051` に変更 or コンテナ内で実行                                                                                              |
| `service "client" is not running`              | `docker compose exec` に `-f` を付け忘れ | `docker compose -f grpc-otel-lab/docker-compose.yml exec client bash`                                                               |
| `connection refused`                           | client が server 起動前に接続             | `wait-for-it.sh` を使用 (本ラボで導入済)                                                                                                      |

---

## 6. 次のステップ

* **Proto に reaction / emoji フィールドを追加** → `buf generate` → UI で確認
* **otel exporter を Tempo へ切り替え**し、Grafana で統合表示
* **k6** を使って 100 並列ストリーム負荷試験 → Jaeger ヒートマップで遅延を分析

Happy tracing 🚀
