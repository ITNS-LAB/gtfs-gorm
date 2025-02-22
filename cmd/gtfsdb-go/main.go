package main

//名前変える
import (
	"flag"
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/interfaces"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/usecase"
	interfacesS "github.com/ITNS-LAB/gtfs-gorm/gtfsscheduledb/interfacesschedule"
	usecaseS "github.com/ITNS-LAB/gtfs-gorm/gtfsscheduledb/usecaseschedule"
	"github.com/m-mizutani/clog"
	"log/slog"
	"os"
)

func main() {
	slogInit()

	gtfsType := flag.String("type", "", "GTFSのタイプ")
	gtfsUrl := flag.String("url", "", "GTFSのURL")
	gtfsFile := flag.String("file", "", "GTFSのファイルパス")
	shapesEx := flag.Bool("shapesex", false, "shapes_exテーブルの作成")
	shapesDetail := flag.Bool("shapesdetail", false, "shapes_detailテーブルの作成")
	geom := flag.Bool("geom", false, "geomカラムにgeomを格納")
	recalculateDist := flag.Bool("recal", false, "'shape_dist_traveled'を再計算")
	dsn := flag.String("dsn", "", "Required: postgresのdsn 例)postgres://hoge:hoge@localhost:5432/hoge")
	schema := flag.String("schema", "public", "格納先のスキーマ")

	// 引数読み込み
	flag.Parse()

	options := usecase.CmdOptions{
		GtfsUrl:         *gtfsUrl,
		GtfsFile:        *gtfsFile,
		ShapesEx:        *shapesEx,
		ShapesDetail:    *shapesDetail,
		Geom:            *geom,
		RecalculateDist: *recalculateDist,
		Dsn:             *dsn,
		Schema:          *schema,
	}

	optionsSchedule := usecaseS.CmdOptions{
		GtfsUrl:         *gtfsUrl,
		GtfsFile:        *gtfsFile,
		ShapesEx:        *shapesEx,
		ShapesDetail:    *shapesDetail,
		Geom:            *geom,
		RecalculateDist: *recalculateDist,
		Dsn:             *dsn,
		Schema:          *schema,
	}

	// helpの出力
	if len(os.Args) < 2 || containsHelpFlag() {
		fmt.Println("GTFSをpostgresに格納するアプリケーションです")
		fmt.Println("Application to store GTFS into postgres")
		fmt.Println()
		fmt.Println("Usage: gtfsdb-go [OPTIONS]")
		flag.PrintDefaults()
		return
	}

	// 引数チェック
	if *dsn == "" {
		fmt.Println("dsnは必須オプションです。")
		return
	}
	if *gtfsUrl == "" && *gtfsFile == "" {
		fmt.Println("'url' または 'file' のどちらかのオプションが必要です。")
		return
	}
	if *gtfsType == "" {
		fmt.Println("gtfstypeは必須オプションです。")
	}

	// ロゴ表示
	fmt.Println("          __  ____         ____                    \n   ____ _/ /_/ __/________/ / /_        ____ _____ \n  / __ `/ __/ /_/ ___/ __  / __ \\______/ __ `/ __ \\\n / /_/ / /_/ __(__  ) /_/ / /_/ /_____/ /_/ / /_/ /\n \\__, /\\__/_/ /____/\\__,_/_.___/      \\__, /\\____/ \n/____/                               /____/")
	fmt.Println("")

	// ロジック
	if *gtfsType == "jp" {
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
	} else if *gtfsType == "schedule" {
		if *gtfsFile != "" {
			if err := interfacesS.GtfsDbFile(optionsSchedule); err != nil {
				slog.Error("処理中にエラーが発生したため終了します。", err)
				return
			}
		} else {
			if err := interfacesS.GtfsDbUrl(optionsSchedule); err != nil {
				slog.Error("処理中にエラーが発生したため終了します。", err)
			}
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
