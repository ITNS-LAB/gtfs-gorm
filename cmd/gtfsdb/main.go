package main

import (
	"flag"
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/service"
	"log/slog"
)

func main() {
	gtfsUrl := flag.String("url", "", "GTFS URL")
	gtfsFile := flag.String("file", "", "GTFS File")
	shapesEx := flag.Bool("shapesEx", false, "'shapes_ex'テーブルを作成するか")
	recalculateDist := flag.Bool("recal", false, "'shape_dist_traveled'を再計算するか")
	dsn := flag.String("dsn", "", "postgresのdsn(必須)")
	schema := flag.String("schema", "public", "格納するスキーマの指定")

	// 引数読み込み
	flag.Parse()

	// 引数チェック
	if *dsn == "" {
		slog.Error("dsnは必須オプションです。")
		return
	}
	if *gtfsUrl == "" && *gtfsFile == "" {
		slog.Error("'url' または 'file' のどちらかの指定が必要です。")
		return
	}

	// ロゴ表示
	fmt.Println("          __  ____                                                  __  ____         ____  \n   ____ _/ /_/ __/____      ____ _____  _________ ___        ____ _/ /_/ __/________/ / /_ \n  / __ `/ __/ /_/ ___/_____/ __ `/ __ \\/ ___/ __ `__ \\      / __ `/ __/ /_/ ___/ __  / __ \\\n / /_/ / /_/ __(__  )_____/ /_/ / /_/ / /  / / / / / /     / /_/ / /_/ __(__  ) /_/ / /_/ /\n \\__, /\\__/_/ /____/      \\__, /\\____/_/  /_/ /_/ /_/      \\__, /\\__/_/ /____/\\__,_/_.___/ \n/____/                   /____/                           /____/                           ")
	fmt.Println("")

	// ロジック
	if *gtfsFile != "" {
		if err := service.GtfsDbFile(*dsn, *gtfsFile, *schema); err != nil {
			slog.Error("処理中にエラーが発生したため終了します。", err)
			return
		}
	} else {
		if err := service.GtfsDbUrl(*dsn, *gtfsUrl, *schema); err != nil {
			slog.Error("処理中にエラーが発生したため終了します。", err)
		}
	}

	if *recalculateDist {
		if err := service.UpdateShapes(*dsn, *schema); err != nil {
			slog.Error("処理中にエラーが発生したため終了します。", err)
		}
	}

	fmt.Println("Process finished !!", *gtfsUrl, *gtfsFile, *shapesEx, *recalculateDist, *dsn, *schema)
}
