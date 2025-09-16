# ReadyGo

新しいGoプロジェクトを素早くセットアップするためのツールです。必要なファイルと設定を自動で生成します。

## 機能

ReadyGoは完全なGoプロジェクト構造を自動作成します：

- **Goモジュール**: 依存関係管理用の`go.mod`を初期化
- **Gitリポジトリ**: コミット準備が完了した新しいGitリポジトリをセットアップ
- **README.md**: プロジェクトドキュメントのテンプレート
- **Hello World**: "Hello, World!"プログラムの基本的な`main.go`
- **Gitignore**: 一般的な除外パターンを含む標準Go用`.gitignore`ファイル
- **Taskfile**: ビルド、テスト、開発に必要なタスクを含む`Taskfile.yml`

## インストール

```bash
go install github.com/HMasataka/readygo@latest
```

## 使用方法

新しいGoプロジェクトを作成：

```bash
readygo my-awesome-project
```

これにより、すべての必要なファイルと設定を含む`my-awesome-project`ディレクトリが作成されます。

### 作成されるファイル

```
my-awesome-project/
├── README.md          # プロジェクトドキュメント
├── main.go           # Hello Worldアプリケーション
├── go.mod            # Goモジュールファイル
├── .gitignore        # Git除外パターン
├── Taskfile.yml      # タスクランナー設定
└── .git/             # Gitリポジトリ
```

## 生成されるTaskfileコマンド

生成される`Taskfile.yml`には以下の便利なタスクが含まれます：

- `task` または `task default` - 利用可能なタスクを表示
- `task build` - バイナリをビルド
- `task build-all` - 複数プラットフォーム用にビルド（Linux、macOS、Windows）
- `task run` - アプリケーションを実行
- `task test` - テストを実行
- `task fmt` - Goコードをフォーマット
- `task vet` - go vetを実行
- `task clean` - ビルド成果物をクリーンアップ
- `task check` - すべてのチェックを実行（fmt、vet、test）

## 要件

- Go 1.16以上（埋め込みテンプレート用）
- Git（リポジトリ初期化用）
- [Task](https://taskfile.dev/)（生成されたTaskfileを使用する場合、オプション）

## ライセンス

MIT License

