package poi

import "database/sql"

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) Create(poi *POI) error {
	_, err := repo.db.Exec("INSERT INTO pois (name, x_coord, y_coord) VALUES ($1, $2, $3);", poi.Name, poi.X, poi.Y)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) List() ([]POI, error) {
	rows, err := repo.db.Query("SELECT name, x_coord, y_coord FROM pois;")
	if err != nil {
		return nil, err
	}
	var pois []POI = make([]POI, 0, 10)
	var scanErr error
	for rows.Next() {
		actualPOI := NewDefaultPOI()
		scanErr = rows.Scan(&actualPOI.Name, &actualPOI.X, &actualPOI.Y)
		if scanErr != nil {
			return nil, scanErr
		}
		pois = append(pois, *actualPOI)
	}
	return pois, nil
}
