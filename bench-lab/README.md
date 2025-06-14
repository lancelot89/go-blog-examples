# 🧪 Go Concurrency Benchmark Lab

ブログ第 3 回「Go 並行処理パターン完全ガイド」で解説した  
**Worker Pool × 画像パイプライン** をすぐに体験できる実行環境です。

---

## 0. 前提

| ツール | 推奨バージョン | 備考 |
|-------|---------------|------|
| **Git** | 2.x 以降 | `lab/03-concurrency-bench` ブランチを clone できれば OK |
| **Docker** | 24.x 以降 | `docker compose` が使えると楽（無くても動作可） |
| **Go** | 1.22 以降 | ローカル実行の場合のみ必須 |

> Docker や Go をローカルに入れたくない場合は **GitHub Codespaces** へどうぞ（手順⑤）。

---

## 1. ブランチの取得

```bash
git clone -b lab/03-concurrency-bench https://github.com/lancelot89/go-blog-examples.git
cd go-blog-examples
```

## 2. 最速ワンコマンド（Docker Compose）
```
docker compose -f bench-lab/docker/docker-compose.yml up --build
# ⏳ …ベンチ実行 → pprof Web UI が :8080 で待機
```
- 初回はイメージビルドのため 1〜2 分かかります。
- 終了したらブラウザで http://localhost:8080/ui/flamegraph を開くとCPU Flame Graph が確認できます。

## 3. ローカル Go で試す
```bash
bash bench-lab/scripts/run_bench.sh
#   1) go test -bench . -benchmem -cpuprofile cpu.out …
#   2) go tool pprof -http :8080 bench-lab/cpu.out
```
- ns/op・allocs/op の数値が表示 → 自動的に pprof UI が起動
- Ctrl-C で pprof UI を終了できます。

## 4. 自分で変更して再計測
1. bench-lab/pipeline/resize.go で draw.ApproxBiLinear を draw.CatmullRom に変更してみる
2. run_bench.sh を再実行
3. ns/op がどれくらい変わるか？ を確認

## 5. GitHub Codespaces（ブラウザだけで試す）
1. リポジトリ → Code ▾ → 「Create Codespace onlab/03-concurrency-bench」
2. 起動後 VS Code ターミナルで `bash bench-lab/scripts/run_bench.sh` を実行
3. 右下の Ports タブに 8080 が現れるので “Open in Browser”

## 6. CI で生成されたプロファイルをダウンロード
Push すると GitHub Actions が cpu.out / mem.out をアーティファクトとして保存します。
1. Actions → Bench & Pprof → 対象ジョブを開く
2. Artifacts セクションから pprof.zip をダウンロード
3. ローカルで go tool pprof cpu.out を開けば同じ解析が可能です

## 7. 目標値の目安
| 項目                  | 参考値 (Mac M1 8-core)                |
| ------------------- | ---------------------------------- |
| **ns/op**           | ≈ 10 000                           |
| **allocs/op**       | ≤ 3                                |
| **CPU Flame Graph** | `pipeline.Resize` が \~70 % 以上なら OK |

数値や Flame Graph を記事と見比べながら、
バッファサイズやアルゴリズムを変えてチューニング してみてください 🔥