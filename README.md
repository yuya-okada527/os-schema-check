# os-schema-check

## 概要

OpenSearch のインデックス設定 (`schema.json`) とバルクリクエスト (`*.jsonl` / `*.ndjson`) を読み込み、既存の `mappings.properties` に存在しないフィールドを検出するシンプルな CLI ツールです。未知フィールドを含むドキュメントはエラーとして扱われ、プロセスは非ゼロ終了コードになります。

## 前提

- Go (モジュール対応、`go run` が利用できるバージョン)
- スキーマファイルは `.json`、バルクデータは `.jsonl` もしくは `.ndjson` 拡張子であること

## 使い方

### 実行方法

```bash
go run ./cmd/os-schema-check <schema.json> <bulk.jsonl>
```

- 引数1: OpenSearch インデックス設定を含む JSON ファイルパス
- 引数2: Bulk API 形式 (1 行 action + 1 行 document の繰り返し) の NDJSON ファイルパス


### 例

リポジトリには動作確認用のサンプルファイルが [sample/](sample) にあります。

サンプルに含まれるファイルで実行する場合:

```bash
go run ./cmd/os-schema-check sample/schema.json sample/bulk_ok.jsonl
```

出力例:

```
Parsed JSON: {Mappings:{Properties:map[date:{Type:keyword} director:{Type:keyword} title:{Type:keyword}]}}
Document 1 is valid
Document 2 is invalid:
  - actors: unexpected field
```

未知フィールドが検出されると終了コード 1 で終了します。すべてのドキュメントが許可されたプロパティのみで構成されている場合は 0 で終了します。

## 実装メモ

- スキーマは `mappings.properties` に定義されたトップレベルフィールドのみを許可対象として解析します。
  - ネストされたフィールドの対応は未実装です。
