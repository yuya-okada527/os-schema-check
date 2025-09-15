# os-schema-check

## 概要

OpenSearch の マッピング定義（mapping） と Bulk NDJSONを受け取り、新たなマッピングが自動生成されてしまう可能性（= 未知フィールドのインデクシング） を事前に検出する CLI ツールです。

## 機能


バルク更新（NDJSON）に含まれるドキュメントと既存 mappingを突合し、以下を検出します：
- nknown fields：mapping 上に存在しないフィールド（パス）。
- dynamic: false/strict な領域では エラー になりうる対象。
- dynamic: true な領域では 新規マッピングが生成される候補 として 警告。

## 使い方

CLI

```bash
os-schema-check validate \
  --mapping ./mappings/product.json \
  --bulk ./bulk/products.ndjson \
```

オプション
- --mapping (required): mapping JSON ファイルパス
- --bulk (required): Bulk NDJSON ファイルパス（action/doc の 2 行ペア形式）

## 使用例

1) mapping（例）: mappings/product.json

```json
{
  "mappings": {
    "dynamic": "true",
    "properties": {
      "id": {"type": "keyword"},
      "title": {"type": "text"},
      "attrs": {
        "type": "object",
        "dynamic": false,
        "properties": {
          "color": {"type": "keyword"}
        }
      }
    }
  }
}
```

2) bulk（例）: bulk/products.ndjson

```json
{"index":{"_index":"products","_id":"SKU-001"}}
{"id":"SKU-001","title":"Sample","attrs":{"color":"black","size":"M"}}
{"update":{"_index":"products","_id":"SKU-002"}}
{"doc":{"attrs":{"color":"white"}}}
```

3) 実行

os-schema-check validate \
  --mapping ./mappings/product.json \
  --bulk ./bulk/products.ndjson

出力（text例）

```bash
[W] line 2  id=SKU-001  field=attrs.size  unknown in mapping (dynamic:true): example="M"
---
Summary: errors=0, warnings=1, docs=2
Would create new mappings under: attrs.size
⸻
```

## 技術要素

- 言語: Go
