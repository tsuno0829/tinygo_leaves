# tinygo_leaves
## 概要
TinyGo×leavesで
Qiita Link: 

## 実行手順
    1. LightGBMのモデルを学習して、モデルを保存
    ```
    $ docker compose build dev-py
    $ docker compose run --rm dev-py
    ```

    2. TinyGo×leavesでLightGBMのWASMを作成
    ```
    $ docker compose build dev-go
    $ docker compose run --rm dev-go bash
    $ make go_mod
    $ make build
    ```

    3. ローカルでlocalhostでWASMの実行確認
    ```
    $ go run server.go
    ```

## 参考
- https://github.com/tinygo-org/tinygo/tree/release/src/examples/wasm
