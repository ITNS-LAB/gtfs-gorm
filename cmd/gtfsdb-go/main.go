package main

import (
	"flag"
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/interfaces"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/usecase"
	"github.com/m-mizutani/clog"
	"log/slog"
	"os"
)

func main() {
	slogInit()

	gtfsUrl := flag.String("url", "", "GTFSのURL")
	gtfsFile := flag.String("file", "", "GTFSのファイルパス")
	shapesEx := flag.Bool("shapesex", false, "shapes_exテーブルの作成")
	recalculateDist := flag.Bool("recal", false, "'shape_dist_traveled'を再計算")
	dsn := flag.String("dsn", "", "Required: postgresのdsn 例)postgres://hoge:hoge@localhost:5432/hoge")
	schema := flag.String("schema", "public", "格納先のスキーマ")

	// 引数読み込み
	flag.Parse()
	options := usecase.CmdOptions{
		GtfsUrl:         *gtfsUrl,
		GtfsFile:        *gtfsFile,
		ShapesEx:        *shapesEx,
		RecalculateDist: *recalculateDist,
		Dsn:             *dsn,
		Schema:          *schema,
	}

	if len(os.Args) < 2 || containsHelpFlag() {
		fmt.Println("GTFSをpostgresに格納するアプリケーションです")
		fmt.Println("Application to store GTFS into postgres")
		fmt.Println()
		fmt.Println("Usage: gtfsdb-go-go [OPTIONS]")
		flag.PrintDefaults()
		return
	}

	// 引数チェック
	if options.Dsn == "" {
		fmt.Println("dsnは必須オプションです。")
		return
	}
	if options.GtfsUrl == "" && options.GtfsFile == "" {
		fmt.Println("'url' または 'file' のどちらかのオプションが必要です。")
		return
	}

	// ロゴ表示
	//fmt.Println("          __  ____                                                  __  ____         ____  \n   ____ _/ /_/ __/____      ____ _____  _________ ___        ____ _/ /_/ __/________/ / /_ \n  / __ `/ __/ /_/ ___/_____/ __ `/ __ \\/ ___/ __ `__ \\      / __ `/ __/ /_/ ___/ __  / __ \\\n / /_/ / /_/ __(__  )_____/ /_/ / /_/ / /  / / / / / /     / /_/ / /_/ __(__  ) /_/ / /_/ /\n \\__, /\\__/_/ /____/      \\__, /\\____/_/  /_/ /_/ /_/      \\__, /\\__/_/ /____/\\__,_/_.___/ \n/____/                   /____/                           /____/                           ")
	fmt.Println("          __  ____         ____                    \n   ____ _/ /_/ __/________/ / /_        ____ _____ \n  / __ `/ __/ /_/ ___/ __  / __ \\______/ __ `/ __ \\\n / /_/ / /_/ __(__  ) /_/ / /_/ /_____/ /_/ / /_/ /\n \\__, /\\__/_/ /____/\\__,_/_.___/      \\__, /\\____/ \n/____/                               /____/")
	fmt.Println("")

	// ロジック
	if *gtfsFile != "" {
		if err := interfaces.GtfsDbFile(options); err != nil {
			slog.Error("処理中にエラーが発生したため終了します。", err)
			return
		}
	} else {
		if err := interfaces.GtfsDbUrl(options); err != nil {
			slog.Error("処理中にエラーが発生したため終了します。", err)
		}
	}

	slog.Info("Process finished !!")
}

func containsHelpFlag() bool {
	for _, arg := range os.Args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	return false
}

func slogInit() {
	slogLogger := slog.New(clog.New(
		clog.WithColor(true),
		clog.WithSource(true),
		clog.WithLevel(slog.LevelDebug),
	))
	slog.SetDefault(slogLogger)
}
