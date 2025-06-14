#!/usr/bin/env bash
set -e

echo "🧪 Running benchmark & generating profiles..."

# 1) プロファイルを “その場” に出力
go test ./bench-lab/benchmark \
  -bench . -benchmem \
  -cpuprofile cpu.out \
  -memprofile mem.out \
  -benchtime 5s

# 2) ルート直下に移動（後続コマンドの参照先を合わせる）
mv cpu.out bench-lab/cpu.out
mv mem.out bench-lab/mem.out

echo "🌐 Launching pprof web UI on :8080"
echo "👉 Open your browser at http://localhost:8080/ui/  (Ctrl-Click or copy-paste)"
PPROF_NO_BROWSER=1 go tool pprof -http 0.0.0.0:8080 bench-lab/cpu.out
