package infrastructure

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/domain/repository"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"gorm.io/gorm"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNewGtfsStaticRepository(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name string
		args args
		want repository.GtfsScheduleRepository
	}{
		{
			name: "Default",
			args: args{dsn: "host=localhost user=hoge password=hoge dbname=hoge port=5432 sslmode=disable"},
			want: &gtfsScheduleRepository{Dsn: "host=localhost user=hoge password=hoge dbname=hoge port=5432 sslmode=disable"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGtfsStaticRepository(tt.args.dsn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGtfsStaticRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gtfsScheduleRepository_ConnectDatabase(t *testing.T) {
	type fields struct {
		Db  *gorm.DB
		Dsn string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid connection",
			fields: fields{
				Dsn: "host=localhost user=hoge password=hoge dbname=hoge port=5432 sslmode=disable",
			},
			wantErr: false,
		},
		{
			name: "Invalid connection",
			fields: fields{
				Dsn: "host=invalid user=hoge password=hoge dbname=hoge port=5432 sslmode=disable",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gtfsScheduleRepository{
				Db:  tt.fields.Db,
				Dsn: tt.fields.Dsn,
			}
			if err := g.ConnectDatabase(); (err != nil) != tt.wantErr {
				t.Errorf("ConnectDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gtfsScheduleRepository_Create(t *testing.T) {
	type fields struct {
		Db  *gorm.DB
		Dsn string
	}
	type args struct {
		gtfsPath string
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_kitami_gtfs",
			fields: fields{
				Dsn: "host=localhost user=hoge password=hoge dbname=hoge port=5432 sslmode=disable",
			},
			args:    args{gtfsPath: filepath.Join("test", "kitami_gtfs")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gtfsScheduleRepository{
				Db:  tt.fields.Db,
				Dsn: tt.fields.Dsn,
			}
			// DBに接続
			if err := g.ConnectDatabase(); err != nil {
				t.Errorf("ConnectDatabse() error = %v", err)
			}

			// テーブルをマイグレーション
			if err := g.Migrate(); (err != nil) != tt.wantErr {
				t.Errorf("Migrate() error = %v", err)
			}

			// テーブルにデータをインサート
			if err := g.Create(filepath.Join("test", "kitami_gtfs")); err != nil {
				t.Errorf("Migrate() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 作成したテーブルをDrop
			var tables []string
			query := "SELECT tablename FROM pg_tables WHERE schemaname = 'public'"
			if err := g.Db.Raw(query).Scan(&tables).Error; err != nil {
				t.Errorf("Failed to retrieve table names: %v", err)
			}
			for _, table := range tables {
				if table == "spatial_ref_sys" {
					continue
				}
				err := g.Db.Migrator().DropTable(table)
				if err != nil {
					t.Error(err)
				}
			}

			// DBから切断
			if err := g.DisConnectDatabase(); err != nil {
				t.Errorf("DisConnectDatabse() error = %v", err)
			}
		})
	}
}

func Test_gtfsScheduleRepository_CreateSchema(t *testing.T) {
	type fields struct {
		Db  *gorm.DB
		Dsn string
	}
	type args struct {
		schema string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_kitami_gtfs",
			fields: fields{
				Dsn: "host=localhost user=hoge password=hoge dbname=hoge port=5432 sslmode=disable",
			},
			args:    args{schema: "test_schema"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gtfsScheduleRepository{
				Db:  tt.fields.Db,
				Dsn: tt.fields.Dsn,
			}
			// DBに接続
			if err := g.ConnectDatabase(); err != nil {
				t.Errorf("ConnectDatabse() error = %v", err)
			}

			// 新規スキーマ作成
			if err := g.CreateSchema(tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("CreateSchema() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 作成したスキーマを削除
			if err := g.Db.Exec("DROP SCHEMA test_schema CASCADE").Error; err != nil {
				t.Errorf("DROP SCHEMA test_schema CASCADE error = %v", err)
			}

			// DBから切断
			if err := g.DisConnectDatabase(); (err != nil) != tt.wantErr {
				t.Errorf("DisConnectDatabase() error = %v", err)
			}
		})
	}
}

func Test_gtfsScheduleRepository_DisConnectDatabase(t *testing.T) {
	type fields struct {
		Db  *gorm.DB
		Dsn string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test_kitami_gtfs",
			fields: fields{
				Dsn: "host=localhost user=hoge password=hoge dbname=hoge port=5432 sslmode=disable",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gtfsScheduleRepository{
				Db:  tt.fields.Db,
				Dsn: tt.fields.Dsn,
			}
			// DBに接続
			if err := g.ConnectDatabase(); err != nil {
				t.Errorf("ConnectDatabse() error = %v", err)
			}

			// DBから切断
			if err := g.DisConnectDatabase(); (err != nil) != tt.wantErr {
				t.Errorf("DisConnectDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gtfsScheduleRepository_Migrate(t *testing.T) {
	type fields struct {
		Db  *gorm.DB
		Dsn string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test_kitami_gtfs",
			fields: fields{
				Dsn: "host=localhost user=hoge password=hoge dbname=hoge port=5432 sslmode=disable",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gtfsScheduleRepository{
				Db:  tt.fields.Db,
				Dsn: tt.fields.Dsn,
			}
			// DBに接続
			if err := g.ConnectDatabase(); err != nil {
				t.Errorf("ConnectDatabse() error = %v", err)
			}

			// テーブルをマイグレーション
			if err := g.Migrate(); (err != nil) != tt.wantErr {
				t.Errorf("Migrate() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 作成したテーブルをDrop
			var tables []string
			query := "SELECT tablename FROM pg_tables WHERE schemaname = 'public'"
			if err := g.Db.Raw(query).Scan(&tables).Error; err != nil {
				t.Errorf("Failed to retrieve table names: %v", err)
			}
			for _, table := range tables {
				if table == "spatial_ref_sys" {
					continue
				}
				err := g.Db.Migrator().DropTable(table)
				if err != nil {
					t.Error(err)
				}
			}

			// DBから切断
			if err := g.DisConnectDatabase(); err != nil {
				t.Errorf("DisConnectDatabse() error = %v", err)
			}
		})
	}
}

func Test_gtfsScheduleRepository_SetSchema(t *testing.T) {
	type fields struct {
		Db  *gorm.DB
		Dsn string
	}
	type args struct {
		schema string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_kitami_gtfs",
			fields: fields{
				Dsn: "host=localhost user=hoge password=hoge dbname=hoge port=5432 sslmode=disable",
			},
			args:    args{schema: "test_schema"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gtfsScheduleRepository{
				Db:  tt.fields.Db,
				Dsn: tt.fields.Dsn,
			}
			// DBに接続
			if err := g.ConnectDatabase(); err != nil {
				t.Errorf("ConnectDatabse() error = %v", err)
			}

			// 新規スキーマ作成
			if err := g.CreateSchema(tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("CreateSchema() error = %v", err)
			}

			// 作成したスキーマに移動
			if err := g.SetSchema(tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("SetSchema() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 現在のスキーマを検索
			var currentSchema string
			if err := g.Db.Raw("SELECT current_schema()").Scan(&currentSchema).Error; err != nil {
				t.Errorf("SELECT current_schema error = %v", err)
			}
			t.Log(currentSchema)

			// 作成したスキーマを削除
			if err := g.Db.Exec("DROP SCHEMA test_schema CASCADE").Error; err != nil {
				t.Errorf("DROP SCHEMA test_schema CASCADE error = %v", err)
			}

			// DBから切断
			if err := g.DisConnectDatabase(); (err != nil) != tt.wantErr {
				t.Errorf("DisConnectDatabase() error = %v", err)
			}

		})
	}
}

func Test_gtfsScheduleRepository_ReadShapeIds(t *testing.T) {
	type fields struct {
		Db  *gorm.DB
		Dsn string
	}
	tests := []struct {
		name         string
		fields       fields
		wantShapeIds []string
		wantErr      bool
	}{
		{
			name: "test_kitami_gtfs",
			fields: fields{
				Dsn: "host=localhost user=hoge password=hoge dbname=hoge port=5432 sslmode=disable",
			},
			wantShapeIds: []string{
				"106", "107", "114", "115", "118", "120", "123", "124", "127",
				"131", "136", "140", "145", "148", "152", "153", "154", "155",
				"160", "165", "206", "207", "208", "209", "222", "223", "224",
				"225", "232", "233", "234", "235", "236", "237", "2491", "2492",
				"258", "259", "270", "271", "272", "273", "274", "291", "292",
				"293", "295", "300", "301", "302", "305", "306", "307", "308",
				"317", "318", "319", "341", "342", "343", "487", "488",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gtfsScheduleRepository{
				Db:  tt.fields.Db,
				Dsn: tt.fields.Dsn,
			}
			// DBに接続
			if err := g.ConnectDatabase(); err != nil {
				t.Errorf("ConnectDatabse() error = %v", err)
			}

			// テーブルをマイグレーション
			if err := g.Migrate(); (err != nil) != tt.wantErr {
				t.Errorf("Migrate() error = %v", err)
			}

			// テーブルにデータをインサート
			if err := g.Create(filepath.Join("test", "kitami_gtfs")); err != nil {
				t.Errorf("Migrate() error = %v, wantErr %v", err, tt.wantErr)
			}

			gotShapeIds, err := g.FindShapeIds()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindShapeIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotShapeIds, tt.wantShapeIds) {
				t.Errorf("FindShapeIds() gotShapeIds = %v, want %v", gotShapeIds, tt.wantShapeIds)
			}

			// 作成したテーブルをDrop
			var tables []string
			query := "SELECT tablename FROM pg_tables WHERE schemaname = 'public'"
			if err := g.Db.Raw(query).Scan(&tables).Error; err != nil {
				t.Errorf("Failed to retrieve table names: %v", err)
			}
			for _, table := range tables {
				if table == "spatial_ref_sys" {
					continue
				}
				err := g.Db.Migrator().DropTable(table)
				if err != nil {
					t.Error(err)
				}
			}

			// DBから切断
			if err := g.DisConnectDatabase(); (err != nil) != tt.wantErr {
				t.Errorf("DisConnectDatabase() error = %v", err)
			}
		})
	}
}

func Test_gtfsScheduleRepository_ReadShapes(t *testing.T) {
	type fields struct {
		Db  *gorm.DB
		Dsn string
	}
	type args struct {
		shapeId string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantShapes []ormstatic.Shape
		wantErr    bool
	}{
		{
			name: "test_kitami_gtfs",
			fields: fields{
				Dsn: "host=localhost user=hoge password=hoge dbname=hoge port=5432 sslmode=disable",
			},
			args:    args{shapeId: "106"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gtfsScheduleRepository{
				Db:  tt.fields.Db,
				Dsn: tt.fields.Dsn,
			}
			// DBに接続
			if err := g.ConnectDatabase(); err != nil {
				t.Errorf("ConnectDatabse() error = %v", err)
			}

			// テーブルをマイグレーション
			if err := g.Migrate(); (err != nil) != tt.wantErr {
				t.Errorf("Migrate() error = %v", err)
			}

			// テーブルにデータをインサート
			if err := g.Create(filepath.Join("test", "kitami_gtfs")); err != nil {
				t.Errorf("Migrate() error = %v, wantErr %v", err, tt.wantErr)
			}

			gotShapes, err := g.FindShapes(tt.args.shapeId)
			if err != nil {
				return
			}

			for _, s := range gotShapes {
				t.Log(*s.ShapeId, *s.ShapePtLat, *s.ShapePtLon, *s.ShapePtSequence)
			}

			// 作成したテーブルをDrop
			var tables []string
			query := "SELECT tablename FROM pg_tables WHERE schemaname = 'public'"
			if err := g.Db.Raw(query).Scan(&tables).Error; err != nil {
				t.Errorf("Failed to retrieve table names: %v", err)
			}
			for _, table := range tables {
				if table == "spatial_ref_sys" {
					continue
				}
				err := g.Db.Migrator().DropTable(table)
				if err != nil {
					t.Error(err)
				}
			}

			// DBから切断
			if err := g.DisConnectDatabase(); err != nil {
				t.Errorf("DisConnectDatabse() error = %v", err)
			}

		})
	}
}

func Test_gtfsScheduleRepository_UpdateShapes(t *testing.T) {
	type fields struct {
		Db  *gorm.DB
		Dsn string
	}
	type args struct {
		shapes []ormstatic.Shape
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_kitami_gtfs",
			fields: fields{
				Dsn: "host=localhost user=hoge password=hoge dbname=hoge port=5432 sslmode=disable",
			},
			args: args{shapes: []ormstatic.Shape{
				{ShapeId: dataframe.StrPtr("106"), ShapePtLat: dataframe.FloatPtr(43.7980950000001), ShapePtLon: dataframe.FloatPtr(143.857629523811), ShapePtSequence: dataframe.IntPtr(1), ShapeDistTraveled: dataframe.FloatPtr(0.0)},
				{ShapeId: dataframe.StrPtr("106"), ShapePtLat: dataframe.FloatPtr(43.7980760983335), ShapePtLon: dataframe.FloatPtr(143.858459586498), ShapePtSequence: dataframe.IntPtr(2), ShapeDistTraveled: dataframe.FloatPtr(0.1)},
				{ShapeId: dataframe.StrPtr("106"), ShapePtLat: dataframe.FloatPtr(43.7983846356787), ShapePtLon: dataframe.FloatPtr(143.862844298172), ShapePtSequence: dataframe.IntPtr(3), ShapeDistTraveled: dataframe.FloatPtr(0.5)},
				{ShapeId: dataframe.StrPtr("106"), ShapePtLat: dataframe.FloatPtr(43.7984666666666), ShapePtLon: dataframe.FloatPtr(143.863157142857), ShapePtSequence: dataframe.IntPtr(4), ShapeDistTraveled: dataframe.FloatPtr(0.6)},
				{ShapeId: dataframe.StrPtr("106"), ShapePtLat: dataframe.FloatPtr(43.7984430506168), ShapePtLon: dataframe.FloatPtr(143.863792516436), ShapePtSequence: dataframe.IntPtr(5), ShapeDistTraveled: dataframe.FloatPtr(0.7)},
				{ShapeId: dataframe.StrPtr("106"), ShapePtLat: dataframe.FloatPtr(43.7987545758103), ShapePtLon: dataframe.FloatPtr(143.868738335255), ShapePtSequence: dataframe.IntPtr(6), ShapeDistTraveled: dataframe.FloatPtr(1.0)},
				{ShapeId: dataframe.StrPtr("106"), ShapePtLat: dataframe.FloatPtr(43.7988766666667), ShapePtLon: dataframe.FloatPtr(143.869138095238), ShapePtSequence: dataframe.IntPtr(7), ShapeDistTraveled: dataframe.FloatPtr(1.1)},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gtfsScheduleRepository{
				Db:  tt.fields.Db,
				Dsn: tt.fields.Dsn,
			}

			// DBに接続
			if err := g.ConnectDatabase(); err != nil {
				t.Errorf("ConnectDatabse() error = %v", err)
			}

			// テーブルをマイグレーション
			if err := g.Migrate(); (err != nil) != tt.wantErr {
				t.Errorf("Migrate() error = %v", err)
			}

			// テーブルにデータをインサート
			if err := g.Create(filepath.Join("test", "kitami_gtfs")); err != nil {
				t.Errorf("Migrate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := g.UpdateShapes(tt.args.shapes); (err != nil) != tt.wantErr {
				t.Errorf("UpdateShapes() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 作成したテーブルをDrop
			var tables []string
			query := "SELECT tablename FROM pg_tables WHERE schemaname = 'public'"
			if err := g.Db.Raw(query).Scan(&tables).Error; err != nil {
				t.Errorf("Failed to retrieve table names: %v", err)
			}
			for _, table := range tables {
				if table == "spatial_ref_sys" {
					continue
				}
				err := g.Db.Migrator().DropTable(table)
				if err != nil {
					t.Error(err)
				}
			}

			// DBから切断
			if err := g.DisConnectDatabase(); err != nil {
				t.Errorf("DisConnectDatabse() error = %v", err)
			}
		})
	}
}
