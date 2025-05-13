package infrastructure

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/jp/domain/model"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/jp/domain/repository"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
	"gorm.io/gorm"
	"log/slog"
	"path/filepath"
)

func createGtfsJp[T any](filePath string, parser func(string) ([]T, error), db *gorm.DB) error {
	// ファイルのパース
	data, err := parser(filePath)
	if err != nil {
		slog.Warn(fmt.Sprintf("%sの読み込みができませんでした。ファイルが存在しない可能性があります。", filepath.Base(filePath)), "error", err, "path", filePath)
		return nil
	}

	// データベースへ挿入
	if err := db.CreateInBatches(&data, 1000).Error; err != nil {
		return fmt.Errorf("データベースへの挿入に失敗しました。%s", err)
	}
	slog.Info(fmt.Sprintf("%s の挿入が完了しました。", filepath.Base(filePath)))
	return nil
}

type gtfsJpRepository struct {
	Db *gorm.DB
}

func (g gtfsJpRepository) MigrateGtfsJp() error {
	gtfsJp := model.GtfsJp{}
	if err := g.Db.AutoMigrate(gtfsJp.Agency); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.AgencyJp); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.Routes); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.Stops); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.Calendar); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.CalendarDates); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.Trips); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.StopTimes); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.Shapes); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.Frequencies); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.Transfers); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.FeedInfo); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.FareAttribute); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.FareRules); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.Translations); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.OfficeJp); err != nil {
		return err
	}
	if err := g.Db.AutoMigrate(gtfsJp.PatternJp); err != nil {
		return err
	}
	return nil
}

func (g gtfsJpRepository) CreateGtfsJp(gtfsPath string) error {
	if err := createGtfsJp(filepath.Join(gtfsPath, "agency.txt"), gtfsjp.ParseAgency, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "agency_jp.txt"), gtfsjp.ParseAgencyJp, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "calendar.txt"), gtfsjp.ParseCalendar, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "calendar_dates.txt"), gtfsjp.ParseCalendarDates, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "routes.txt"), gtfsjp.ParseRoutes, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "stops.txt"), gtfsjp.ParseStops, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "shapes.txt"), gtfsjp.ParseShapes, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "trips.txt"), gtfsjp.ParseTrips, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "stop_times.txt"), gtfsjp.ParseStopTimes, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "transfers.txt"), gtfsjp.ParseTransfers, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "frequencies.txt"), gtfsjp.ParseFrequencies, g.Db); err != nil {
		return err
	}
	slog.Info("createGtfsJpの渡されるパス")
	slog.Info(gtfsPath)
	if err := createGtfsJp(filepath.Join(gtfsPath, "fare_attributes.txt"), gtfsjp.ParseFareAttributes, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "fare_rules.txt"), gtfsjp.ParseFareRules, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "feed_info.txt"), gtfsjp.ParseFeedInfo, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "translations.txt"), gtfsjp.ParseTranslations, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "office_jp.txt"), gtfsjp.ParseOfficeJp, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "pattern_jp.txt"), gtfsjp.ParsePatternJp, g.Db); err != nil {
		return err
	}
	return nil
}

func NewGtfsJpRepository(db *gorm.DB) repository.GtfsJpRepository {
	return gtfsJpRepository{Db: db}
}

type gtfsJpGeomRepository struct {
	Db *gorm.DB
}

func (g gtfsJpGeomRepository) MigrateGtfsJpGeom() error {
	gtfsJp := model.GtfsJpGeom{}
	if err := g.Db.AutoMigrate(gtfsJp.Agency); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.AgencyJp); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.Routes); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.Stops); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.Calendar); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.CalendarDates); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.Trips); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.StopTimes); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.Shapes); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.Frequencies); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.Transfers); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.FeedInfo); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.FareAttribute); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.FareRules); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.Translations); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.OfficeJp); err != nil {
		return nil
	}
	if err := g.Db.AutoMigrate(gtfsJp.PatternJp); err != nil {
		return nil
	}

	return nil
}

func (g gtfsJpGeomRepository) CreateGtfsJpGeom(gtfsPath string) error {
	if err := createGtfsJp(filepath.Join(gtfsPath, "agency.txt"), gtfsjp.ParseAgencyGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "agency_jp.txt"), gtfsjp.ParseAgencyJpGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "calendar.txt"), gtfsjp.ParseCalendarGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "calendar_dates.txt"), gtfsjp.ParseCalendarDatesGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "routes.txt"), gtfsjp.ParseRoutesGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "stops.txt"), gtfsjp.ParseStopsGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "shapes.txt"), gtfsjp.ParseShapesGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "trips.txt"), gtfsjp.ParseTripsGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "stop_times.txt"), gtfsjp.ParseStopTimesGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "transfers.txt"), gtfsjp.ParseTransfersGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "frequencies.txt"), gtfsjp.ParseFrequenciesGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "fare_attributes.txt"), gtfsjp.ParseFareAttributesGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "fare_rules.txt"), gtfsjp.ParseFareRulesGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "feed_info.txt"), gtfsjp.ParseFeedInfo, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "translations.txt"), gtfsjp.ParseTranslations, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "office_jp.txt"), gtfsjp.ParseOfficeJpGeom, g.Db); err != nil {
		return err
	}
	if err := createGtfsJp(filepath.Join(gtfsPath, "pattern_jp.txt"), gtfsjp.ParsePatternJpGeom, g.Db); err != nil {
		return err
	}
	return nil
}

func NewGtfsJpGeomRepository(db *gorm.DB) repository.GtfsJpGeomRepository {
	return gtfsJpGeomRepository{Db: db}
}

type tripRepository struct {
	Db *gorm.DB
}

func (t tripRepository) FindTripIds() (tripIds []string, err error) {
	t.Db.Table("trips").Select("trip_id").Distinct("trip_id").Order("trip_id asc").Find(&tripIds)
	return tripIds, nil
}

func (t tripRepository) FindShapeIdByTripId(tripId string) (string, error) {
	var shapeId string
	t.Db.Table("trips").Select("shape_id").Where("trip_id = ?", tripId).Find(&shapeId)
	return shapeId, nil
}

func NewTripRepository(db *gorm.DB) repository.TripRepository {
	return tripRepository{Db: db}
}

type tripGeomRepository struct {
	Db *gorm.DB
}

func (t tripGeomRepository) FindTripsGeomIds() (tripIds []string, err error) {
	t.Db.Table("trips_geom").Select("trip_id").Distinct("trip_id").Order("trip_id asc").Find(&tripIds)
	return tripIds, nil
}

func (t tripGeomRepository) UpdateTripsGeom(tripsGeom []model.TripGeomLine) error {
	tx := t.Db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, tripGeom := range tripsGeom {
		if result := tx.Model(&gtfsjp.TripGeom{}).
			Where("trip_id = ?", tripGeom.TripId).
			//Updates()部分を変更
			Update("geom", tripGeom.Geom); result.Error != nil {
			tx.Rollback() // エラーが発生したらロールバック
			return result.Error
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func NewTripGeomRepository(db *gorm.DB) repository.TripGeomRepository {
	return tripGeomRepository{Db: db}
}

type shapeRepository struct {
	Db *gorm.DB
}

func (s shapeRepository) FindShapeIds() (shapeIds []string, err error) {
	s.Db.Table("shapes").Select("shape_id").Distinct("shape_id").Order("shape_id asc").Find(&shapeIds)
	return shapeIds, nil
}

func (s shapeRepository) FindShapesByShapeId(shapeId string) (shapes []model.Shape, err error) {
	s.Db.Table("shapes").Where("shape_id = ?", shapeId).Order("shape_pt_sequence asc").Find(&shapes)
	return shapes, nil
}

func (s shapeRepository) UpdateShapes(shapes []model.Shape) error {
	tx := s.Db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, shape := range shapes {
		if result := tx.Model(&model.Shape{}).
			Where("shape_id = ? AND shape_pt_sequence = ?", shape.ShapeId, shape.ShapePtSequence).
			Updates(shape); result.Error != nil {
			tx.Rollback() // エラーが発生したらロールバック
			return result.Error
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (s shapeRepository) FindShapes() (shapes []model.Shape, err error) {
	s.Db.Find(&shapes)
	return shapes, nil
}

func NewShapeRepository(db *gorm.DB) repository.ShapeRepository {
	return shapeRepository{Db: db}
}

type shapeGeomRepository struct {
	Db *gorm.DB
}

func (s shapeGeomRepository) FindShapeGeomIds() (shapeIds []string, err error) {
	s.Db.Table("shapes").Select("shape_id").Distinct("shape_id").Order("shape_id").Find(&shapeIds)
	return shapeIds, nil
}

func (s shapeGeomRepository) FindShapesGeomByShapeId(shapeId string) (shapesGeom []model.ShapeGeom, err error) {
	if err := s.Db.Table("shapes").Where("shape_id = ?", shapeId).Order("shape_pt_sequence asc").Find(&shapesGeom).Error; err != nil {
		return shapesGeom, err
	}
	return shapesGeom, nil
}

func (s shapeGeomRepository) UpdateShapesGeom(shapesGeom []model.ShapeGeom) error {
	tx := s.Db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, shapeGeom := range shapesGeom {
		if result := tx.Model(&gtfsjp.Shape{}).
			Where("shape_id = ? AND shape_pt_sequence = ?", shapeGeom.ShapeId, shapeGeom.ShapePtSequence).
			Updates(shapeGeom); result.Error != nil {
			tx.Rollback() // エラーが発生したらロールバック
			return result.Error
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (s shapeGeomRepository) FindShapesGeom() (shapesGeom []model.ShapeGeom, err error) {
	s.Db.Find(&shapesGeom)
	return shapesGeom, nil
}

func NewShapeGeomRepository(db *gorm.DB) repository.ShapeGeomRepository {
	return shapeGeomRepository{Db: db}
}

type shapeExRepository struct {
	Db *gorm.DB
}

func (s shapeExRepository) MigrateShapesEx() error {
	if err := s.Db.AutoMigrate(&model.ShapeEx{}); err != nil {
		return err
	}
	return nil
}

func (s shapeExRepository) CreateShapesEx(shapesEx []model.ShapeEx) error {
	if err := s.Db.CreateInBatches(&shapesEx, 1000).Error; err != nil {
		return fmt.Errorf("データベースへの挿入に失敗しました。%s", err)
	}
	return nil
}

func (s shapeExRepository) UpdateShapesEx(shapeEx []model.ShapeEx) error {
	tx := s.Db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, shapeEx := range shapeEx {
		if result := tx.Model(&model.ShapeEx{}).
			Where("trip_id = ? AND shape_id = ? AND shape_pt_sequence = ?", shapeEx.TripId, shapeEx.ShapeId, shapeEx.ShapePtSequence).
			Updates(shapeEx); result.Error != nil {
			tx.Rollback() // エラーが発生したらロールバック
			return result.Error
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (s shapeExRepository) FindShapesExByTripsAndShapes() ([]model.ShapeEx, error) {
	var shapesEx []model.ShapeEx
	if err := s.Db.Table("shapes").
		Select("trips.trip_id, trips.shape_id, shapes.shape_pt_lat, shapes.shape_pt_lon, shapes.shape_pt_sequence, shapes.shape_dist_traveled, NULL AS stop_id").
		Joins("join trips on trips.shape_id = shapes.shape_id").
		Order("trips.trip_id").
		Order("shapes.shape_pt_sequence").
		Scan(&shapesEx).Error; err != nil {
		return shapesEx, err
	}
	return shapesEx, nil
}

func (s shapeExRepository) FindShapesExByTripId(tripId string) ([]model.ShapeEx, error) {
	var shapesEx []model.ShapeEx
	if err := s.Db.Table("shapes_ex").
		Select("trip_id, shape_id, shape_pt_lat, shape_pt_lon, shape_pt_sequence, shape_dist_traveled, stop_id").
		Where("trip_id = ?", tripId).
		Order("shape_pt_sequence").
		Scan(&shapesEx).Error; err != nil {
		return shapesEx, err
	}
	return shapesEx, nil
}

func (s shapeExRepository) FindTripWithStopLocationByTripId(tripId string) ([]model.TripWithStopLocation, error) {
	var tWSL []model.TripWithStopLocation
	if err := s.Db.Table("stop_times").
		Select("stop_times.trip_id, stop_times.stop_id, stop_times.stop_sequence, stops.stop_lat, stops.stop_lon").
		Joins("join stops on stop_times.stop_id = stops.stop_id").
		Where("trip_id = ?", tripId).
		Order("stop_sequence").
		Scan(&tWSL).Error; err != nil {
		return tWSL, err
	}
	return tWSL, nil
}

func NewShapeExRepository(db *gorm.DB) repository.ShapeExRepository {
	return shapeExRepository{Db: db}
}

type shapeExGeomRepository struct {
	Db *gorm.DB
}

func (s shapeExGeomRepository) MigrateShapesExGeom() error {
	if err := s.Db.AutoMigrate(&model.ShapeExGeom{}); err != nil {
		return err
	}
	return nil
}

func (s shapeExGeomRepository) CreateShapesExGeom(shapesExGeom []model.ShapeExGeom) error {
	if err := s.Db.CreateInBatches(&shapesExGeom, 1000).Error; err != nil {
		return fmt.Errorf("データベースへの挿入に失敗しました。%s", err)
	}
	return nil
}

func (s shapeExGeomRepository) UpdateShapesExGeom(shapesExGeom []model.ShapeExGeom) error {
	tx := s.Db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, shapeEx := range shapesExGeom {
		if result := tx.Model(&model.ShapeEx{}).
			Where("trip_id = ? AND shape_id = ? AND shape_pt_sequence = ?", shapeEx.TripId, shapeEx.ShapeId, shapeEx.ShapePtSequence).
			Updates(shapeEx); result.Error != nil {
			tx.Rollback() // エラーが発生したらロールバック
			return result.Error
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (s shapeExGeomRepository) FindShapesExGeomByTripsAndShapes() ([]model.ShapeExGeom, error) {
	var shapesExGeom []model.ShapeExGeom
	if err := s.Db.Table("shapes").
		Select("trips.trip_id, trips.shape_id, shapes.shape_pt_lat, shapes.shape_pt_lon, shapes.shape_pt_sequence, shapes.shape_dist_traveled, NULL AS stop_id, shapes.geom").
		Joins("join trips on trips.shape_id = shapes.shape_id").
		Order("trips.trip_id").
		Order("shapes.shape_pt_sequence").
		Scan(&shapesExGeom).Error; err != nil {
		return shapesExGeom, err
	}
	return shapesExGeom, nil
}

func (s shapeExGeomRepository) FindShapesExGeomByTripId(tripId string) ([]model.ShapeExGeom, error) {
	var shapesExGeom []model.ShapeExGeom
	if err := s.Db.Table("shapes_ex").
		Select("trip_id, shape_id, shape_pt_lat, shape_pt_lon, shape_pt_sequence, shape_dist_traveled, stop_id").
		Where("trip_id = ?", tripId).
		Order("shape_pt_sequence").
		Scan(&shapesExGeom).Error; err != nil {
		return shapesExGeom, err
	}
	return shapesExGeom, nil
}

func (s shapeExGeomRepository) FindTripWithStopLocationByTripId(tripId string) ([]model.TripWithStopLocation, error) {
	var tWSL []model.TripWithStopLocation
	if err := s.Db.Table("stop_times").
		Select("stop_times.trip_id, stop_times.stop_id, stop_times.stop_sequence, stops.stop_lat, stops.stop_lon").
		Joins("join stops on stop_times.stop_id = stops.stop_id").
		Where("trip_id = ?", tripId).
		Order("stop_sequence").
		Scan(&tWSL).Error; err != nil {
		return tWSL, err
	}
	return tWSL, nil
}

func NewShapeExGeomRepository(db *gorm.DB) repository.ShapeExGeomRepository {
	return shapeExGeomRepository{Db: db}
}

type shapeDetailRepository struct {
	Db *gorm.DB
}

func (s shapeDetailRepository) MigrateShapesDetail() error {
	if err := s.Db.AutoMigrate(&model.ShapeDetail{}); err != nil {
		return err
	}
	return nil
}

func (s shapeDetailRepository) CreateShapesDetail(shapesDetails []model.ShapeDetail) error {
	if err := s.Db.CreateInBatches(&shapesDetails, 1000).Error; err != nil {
		return fmt.Errorf("データベースへの挿入に失敗しました。%s", err)
	}
	return nil
}

func NewShapeDetailRepository(db *gorm.DB) repository.ShapeDetailRepository {
	return shapeDetailRepository{Db: db}
}

type shapeDetailGeomRepository struct {
	Db *gorm.DB
}

func (s shapeDetailGeomRepository) MigrateShapesDetailGeom() error {
	if err := s.Db.AutoMigrate(&model.ShapeDetailGeom{}); err != nil {
		return err
	}
	return nil
}

func (s shapeDetailGeomRepository) CreateShapesDetailGeom(shapesDetailGeom []model.ShapeDetailGeom) error {
	if err := s.Db.CreateInBatches(&shapesDetailGeom, 1000).Error; err != nil {
		return err
	}
	return nil
}

func NewShapeDetailGeomRepository(db *gorm.DB) repository.ShapeDetailGeomRepository {
	return shapeDetailGeomRepository{Db: db}
}

type stopTimesRepository struct {
	Db *gorm.DB
}

func (s stopTimesRepository) FindStopTimesByTripId(tripId string) ([]model.StopTime, error) {
	var stopTimes []model.StopTime
	if err := s.Db.Table("stop_times").
		Select("trip_id, arrival_time, departure_time, stop_id, stop_sequence, stop_headsign, pickup_type, drop_off_type, continuous_pickup, continuous_drop_off, shape_dist_traveled, timepoint").
		Where("trip_id = ?", tripId).
		Order("stop_sequence").
		Scan(&stopTimes).Error; err != nil {
		return stopTimes, err
	}
	return stopTimes, nil
}

func NewStopTimesRepository(db *gorm.DB) repository.StopTimeRepository {
	return stopTimesRepository{Db: db}
}
